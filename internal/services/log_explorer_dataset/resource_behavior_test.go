package log_explorer_dataset_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	testDatasetID = "0199179f7c7f7b37a3f4de2f8c634e1a"
	testZoneID    = "zone-123"
	testAccountID = "account-123"
)

type datasetResponse struct {
	Enabled            bool
	DeletionProtection bool
	Filter             *string
	Fields             []map[string]any
	ObjectType         string
	ObjectID           string
}

func newDatasetResource(t *testing.T, baseURL string) *log_explorer_dataset.LogExplorerDatasetResource {
	t.Helper()

	r := log_explorer_dataset.NewResource().(*log_explorer_dataset.LogExplorerDatasetResource)
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
	)
	resp := &resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure failed: %v", resp.Diagnostics)
	}
	return r
}

func baseDatasetModel(ctx context.Context) log_explorer_dataset.LogExplorerDatasetModel {
	return log_explorer_dataset.LogExplorerDatasetModel{
		ID:                 types.StringUnknown(),
		DatasetID:          types.StringUnknown(),
		AccountID:          types.StringNull(),
		ZoneID:             types.StringValue(testZoneID),
		Dataset:            types.StringValue("dns_logs"),
		DeletionProtection: types.BoolUnknown(),
		Enabled:            types.BoolUnknown(),
		Filter:             types.StringNull(),
		Fields:             customfield.NullObjectList[log_explorer_dataset.LogExplorerDatasetFieldsModel](ctx),
		ObjectType:         types.StringUnknown(),
	}
}

func datasetState(t *testing.T, ctx context.Context, model log_explorer_dataset.LogExplorerDatasetModel) tfsdk.State {
	t.Helper()

	state := tfsdk.State{Schema: log_explorer_dataset.ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to build dataset state: %v", diags)
	}
	return state
}

func datasetPlan(t *testing.T, ctx context.Context, model log_explorer_dataset.LogExplorerDatasetModel) tfsdk.Plan {
	t.Helper()

	state := datasetState(t, ctx, model)
	return tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}
}

func knownDatasetModel(ctx context.Context) log_explorer_dataset.LogExplorerDatasetModel {
	model := baseDatasetModel(ctx)
	model.ID = types.StringValue(testDatasetID)
	model.DatasetID = types.StringValue(testDatasetID)
	model.DeletionProtection = types.BoolValue(true)
	model.Enabled = types.BoolValue(true)
	model.ObjectType = types.StringValue("zone")
	return model
}

func writeDatasetResponse(t *testing.T, w http.ResponseWriter, response datasetResponse) {
	t.Helper()

	result := map[string]any{
		"dataset_id":          testDatasetID,
		"dataset":             "dns_logs",
		"enabled":             response.Enabled,
		"deletion_protection": response.DeletionProtection,
		"fields":              response.Fields,
		"object_type":         response.ObjectType,
		"object_id":           response.ObjectID,
		"created_at":          "2026-09-01T12:00:00Z",
		"updated_at":          "2026-09-02T12:00:00Z",
	}
	result["filter_attack_traffic"] = true
	if response.Filter != nil {
		result["filter"] = *response.Filter
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]any{
		"success":  true,
		"errors":   []any{},
		"messages": []any{},
		"result":   result,
	}); err != nil {
		t.Errorf("failed to write response: %v", err)
	}
}

func decodeRequestBody(t *testing.T, r *http.Request) map[string]any {
	t.Helper()

	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		t.Errorf("failed to decode %s request: %v", r.Method, err)
		return nil
	}
	return body
}
