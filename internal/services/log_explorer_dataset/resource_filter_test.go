package log_explorer_dataset_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogExplorerDatasetUpdateFilterLifecycle(t *testing.T) {
	tests := map[string]struct {
		stateFilter    types.String
		planFilter     types.String
		wantPresent    bool
		wantFilter     string
		wantStateValue string
	}{
		"omitted means no filter": {stateFilter: types.StringNull(), planFilter: types.StringNull()},
		"configured value replaces": {
			stateFilter: types.StringValue("old"), planFilter: types.StringValue("new"),
			wantPresent: true, wantFilter: "new", wantStateValue: "new",
		},
		"explicit empty clears": {stateFilter: types.StringValue("old"), planFilter: types.StringValue(""), wantPresent: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			stateModel := knownDatasetModel(ctx)
			stateModel.Filter = test.stateFilter
			planModel := stateModel
			planModel.Filter = test.planFilter
			state := datasetState(t, ctx, stateModel)
			plan := datasetPlan(t, ctx, planModel)

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body := decodeRequestBody(t, r)
				value, present := body["filter"]
				if present != test.wantPresent || (present && value != test.wantFilter) {
					t.Errorf("got filter value %#v (present %t), want %q (present %t)", value, present, test.wantFilter, test.wantPresent)
				}
				responseFilter := test.wantStateValue
				writeDatasetResponse(t, w, datasetResponse{Enabled: true, DeletionProtection: true, Filter: &responseFilter, Fields: []map[string]any{}, ObjectType: "zone", ObjectID: testZoneID})
			}))
			defer ts.Close()

			resp := &resource.UpdateResponse{State: state}
			newDatasetResource(t, ts.URL).Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("Update returned diagnostics: %v", resp.Diagnostics)
			}
			var got log_explorer_dataset.LogExplorerDatasetModel
			if diags := resp.State.Get(ctx, &got); diags.HasError() {
				t.Fatalf("failed to read state: %v", diags)
			}
			if got.Filter.IsNull() || got.Filter.ValueString() != test.wantStateValue {
				t.Errorf("got state filter %v, want %q", got.Filter, test.wantStateValue)
			}
		})
	}
}

func TestLogExplorerDatasetFilterResponseState(t *testing.T) {
	for _, test := range []struct {
		name       string
		filterJSON string
		want       string
	}{
		{name: "missing"},
		{name: "null", filterJSON: `,"filter":null`},
		{name: "empty", filterJSON: `,"filter":""`},
		{name: "nonempty", filterJSON: `,"filter":"server-filter"`, want: "server-filter"},
	} {
		for _, operation := range []string{"create", "read", "update", "import"} {
			t.Run(operation+"/"+test.name, func(t *testing.T) {
				ctx := context.Background()
				model := knownDatasetModel(ctx)
				model.Filter = types.StringValue("previous-filter")
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintf(w, `{"success":true,"errors":[],"messages":[],"result":{
						"dataset_id":%q,"dataset":"dns_logs","object_type":"zone","object_id":%q,
						"enabled":true,"deletion_protection":true,"fields":[],
						"created_at":"2026-09-01T12:00:00Z","updated_at":"2026-09-02T12:00:00Z"%s
					}}`, testDatasetID, testZoneID, test.filterJSON)
				}))
				defer server.Close()
				r := newDatasetResource(t, server.URL)
				state := datasetState(t, ctx, model)
				switch operation {
				case "create":
					resp := &resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
					r.Create(ctx, resource.CreateRequest{Plan: datasetPlan(t, ctx, model)}, resp)
					if resp.Diagnostics.HasError() {
						t.Fatal(resp.Diagnostics)
					}
					state = resp.State
				case "read":
					resp := &resource.ReadResponse{State: state}
					r.Read(ctx, resource.ReadRequest{State: state}, resp)
					if resp.Diagnostics.HasError() {
						t.Fatal(resp.Diagnostics)
					}
					state = resp.State
				case "update":
					model.Filter = types.StringUnknown()
					resp := &resource.UpdateResponse{State: state}
					r.Update(ctx, resource.UpdateRequest{Plan: datasetPlan(t, ctx, model), State: state}, resp)
					if resp.Diagnostics.HasError() {
						t.Fatal(resp.Diagnostics)
					}
					state = resp.State
				case "import":
					resp := &resource.ImportStateResponse{State: tfsdk.State{Schema: state.Schema}}
					r.ImportState(ctx, resource.ImportStateRequest{ID: "zones/" + testZoneID + "/" + testDatasetID}, resp)
					if resp.Diagnostics.HasError() {
						t.Fatal(resp.Diagnostics)
					}
					state = resp.State
				}
				var got log_explorer_dataset.LogExplorerDatasetModel
				if diags := state.Get(ctx, &got); diags.HasError() {
					t.Fatal(diags)
				}
				if !got.Filter.Equal(types.StringValue(test.want)) {
					t.Fatalf("filter = %v, want %q", got.Filter, test.want)
				}
			})
		}
	}
}
