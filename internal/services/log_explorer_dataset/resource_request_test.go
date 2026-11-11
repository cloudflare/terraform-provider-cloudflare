package log_explorer_dataset_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogExplorerDatasetCreateSettings(t *testing.T) {
	for name, tc := range map[string]struct {
		enabled, protection types.Bool
		wantUpdate          bool
	}{
		"API defaults":       {types.BoolUnknown(), types.BoolUnknown(), false},
		"matching settings":  {types.BoolValue(true), types.BoolValue(true), false},
		"disable ingest":     {types.BoolValue(false), types.BoolUnknown(), true},
		"disable protection": {types.BoolUnknown(), types.BoolValue(false), true},
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			model := baseDatasetModel(ctx)
			model.Enabled, model.DeletionProtection = tc.enabled, tc.protection
			filter := `{"where":{"key":"QueryName","operator":"eq","value":"example.com"}}`
			model.Filter = types.StringValue(filter)
			model.Fields = customfield.NewObjectListMust(ctx, []log_explorer_dataset.LogExplorerDatasetFieldsModel{
				{Name: types.StringValue("timestamp"), Enabled: types.BoolValue(true)},
				{Name: types.StringValue("queryname"), Enabled: types.BoolValue(true)},
			})
			wantFields := []any{map[string]any{"name": "timestamp", "enabled": true}, map[string]any{"name": "queryname", "enabled": true}}
			enabled, protection := true, true
			updates := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				body := decodeRequestBody(t, req)
				switch req.Method {
				case http.MethodPost:
					if _, exists := body["enabled"]; exists {
						t.Error("POST contains update-only enabled")
					}
					if _, exists := body["deletion_protection"]; exists {
						t.Error("POST contains update-only deletion_protection")
					}
					if body["dataset"] != "dns_logs" {
						t.Errorf("dataset = %v", body["dataset"])
					}
				case http.MethodPut:
					updates++
					if _, exists := body["dataset"]; exists {
						t.Error("PUT contains create-only dataset")
					}
					value, present := body["enabled"].(bool)
					if !present {
						t.Error("PUT is missing required enabled")
					}
					enabled = value
					if value, present := body["deletion_protection"].(bool); present {
						protection = value
					}
				default:
					t.Errorf("unexpected method %s", req.Method)
				}
				if !reflect.DeepEqual(body["fields"], wantFields) {
					t.Errorf("%s lost field selection: %v", req.Method, body["fields"])
				}
				if body["filter"] != filter {
					t.Errorf("%s lost filter: %v", req.Method, body["filter"])
				}
				writeDatasetResponse(t, w, datasetResponse{Enabled: enabled, DeletionProtection: protection, Filter: &filter, ObjectType: "zone", ObjectID: testZoneID,
					Fields: []map[string]any{{"name": "timestamp", "enabled": true}, {"name": "queryname", "enabled": true}},
				})
			}))
			defer server.Close()
			plan := datasetPlan(t, ctx, model)
			resp := &resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
			newDatasetResource(t, server.URL).Create(ctx, resource.CreateRequest{Plan: plan}, resp)
			if resp.Diagnostics.HasError() {
				t.Fatal(resp.Diagnostics)
			}
			wantEnabled, wantProtection := true, true
			if !tc.enabled.IsUnknown() {
				wantEnabled = tc.enabled.ValueBool()
			}
			if !tc.protection.IsUnknown() {
				wantProtection = tc.protection.ValueBool()
			}
			if enabled != wantEnabled || protection != wantProtection {
				t.Fatalf("API settings = %t/%t, want %t/%t", enabled, protection, wantEnabled, wantProtection)
			}
			if (updates == 1) != tc.wantUpdate || updates > 1 {
				t.Fatalf("updates = %d, want update %t", updates, tc.wantUpdate)
			}
			var got log_explorer_dataset.LogExplorerDatasetModel
			if diags := resp.State.Get(ctx, &got); diags.HasError() {
				t.Fatal(diags)
			}
			if got.ID.ValueString() != testDatasetID || got.DatasetID.ValueString() != testDatasetID {
				t.Fatalf("server identity lost: %v", got.ID)
			}
			if got.Enabled.ValueBool() != enabled || got.DeletionProtection.ValueBool() != protection {
				t.Fatal("state does not match API settings")
			}
		})
	}
}

func TestLogExplorerDatasetUpdateRequest(t *testing.T) {
	ctx := context.Background()
	stateModel := knownDatasetModel(ctx)
	planModel := stateModel
	planModel.DeletionProtection = types.BoolValue(false)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body := decodeRequestBody(t, req)
		if _, exists := body["dataset"]; exists {
			t.Error("PUT contains create-only dataset")
		}
		if body["enabled"] != true {
			t.Errorf("required unchanged enabled = %v", body["enabled"])
		}
		if body["deletion_protection"] != false {
			t.Errorf("protection = %v", body["deletion_protection"])
		}
		writeDatasetResponse(t, w, datasetResponse{Enabled: true, DeletionProtection: false, ObjectType: "zone", ObjectID: testZoneID})
	}))
	defer server.Close()
	state := datasetState(t, ctx, stateModel)
	resp := &resource.UpdateResponse{State: state}
	newDatasetResource(t, server.URL).Update(ctx, resource.UpdateRequest{Plan: datasetPlan(t, ctx, planModel), State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
}

func TestLogExplorerDatasetCreateKeepsIdentityAfterUpdateFailure(t *testing.T) {
	ctx := context.Background()
	model := baseDatasetModel(ctx)
	model.DeletionProtection = types.BoolValue(false)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{"success": false, "errors": []any{map[string]any{"code": 1000, "message": "update rejected"}}})
			return
		}
		writeDatasetResponse(t, w, datasetResponse{Enabled: true, DeletionProtection: true, ObjectType: "zone", ObjectID: testZoneID})
	}))
	defer server.Close()
	plan := datasetPlan(t, ctx, model)
	resp := &resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
	newDatasetResource(t, server.URL).Create(ctx, resource.CreateRequest{Plan: plan}, resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("expected update failure")
	}
	var got log_explorer_dataset.LogExplorerDatasetModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatal(diags)
	}
	if got.ID.ValueString() != testDatasetID {
		t.Fatalf("created dataset lost from state: %v", got.ID)
	}
}
