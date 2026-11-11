package log_explorer_dataset_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestLogExplorerDatasetDeleteHonorsAPIProtection(t *testing.T) {
	tests := map[string]struct {
		protected bool
		status    int
		wantError bool
	}{
		"protected":   {protected: true, status: http.StatusBadRequest, wantError: true},
		"unprotected": {protected: false, status: http.StatusOK},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			model := knownDatasetModel(ctx)
			model.DeletionProtection = types.BoolValue(test.protected)
			state := datasetState(t, ctx, model)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodDelete {
					t.Errorf("got method %s, want DELETE", r.Method)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.status)
				if test.wantError {
					fmt.Fprint(w, `{"success":false,"errors":[{"code":1000,"message":"deletion is protected"}],"messages":[],"result":null}`)
					return
				}
				fmt.Fprint(w, `{"success":true,"errors":[],"messages":[],"result":{}}`)
			}))
			defer ts.Close()

			resp := &resource.DeleteResponse{State: state}
			newDatasetResource(t, ts.URL).Delete(ctx, resource.DeleteRequest{State: state}, resp)
			if resp.Diagnostics.HasError() != test.wantError {
				t.Errorf("Delete diagnostics error=%t, want %t: %v", resp.Diagnostics.HasError(), test.wantError, resp.Diagnostics)
			}
		})
	}
}

func TestLogExplorerDatasetCreateReturnsAPIError(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := decodeRequestBody(t, r)
		if body["dataset"] != "DNS_LOGS" {
			t.Errorf("got dataset %#v, want case-mismatched registered type", body["dataset"])
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"success":false,"errors":[{"code":1001,"message":"dataset is not registered"}],"messages":[],"result":null}`)
	}))
	defer ts.Close()

	model := baseDatasetModel(ctx)
	model.Dataset = types.StringValue("DNS_LOGS")
	plan := datasetPlan(t, ctx, model)
	resp := &resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
	newDatasetResource(t, ts.URL).Create(ctx, resource.CreateRequest{Plan: plan}, resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("Create did not report the API rejection")
	}
}
