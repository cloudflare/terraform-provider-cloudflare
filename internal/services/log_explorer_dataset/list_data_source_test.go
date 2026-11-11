package log_explorer_dataset_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func newLogExplorerDatasetsDataSource(t *testing.T, baseURL string) *log_explorer_dataset.LogExplorerDatasetsDataSource {
	t.Helper()

	d := log_explorer_dataset.NewLogExplorerDatasetsDataSource().(*log_explorer_dataset.LogExplorerDatasetsDataSource)
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
	)
	resp := &datasource.ConfigureResponse{}
	d.Configure(context.Background(), datasource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("configure data source: %v", resp.Diagnostics)
	}
	return d
}

func logExplorerDatasetsConfig(
	t *testing.T,
	ctx context.Context,
	accountID, zoneID types.String,
	includeZones types.Bool,
	maxItems types.Int64,
) tfsdk.Config {
	t.Helper()

	schema := log_explorer_dataset.ListDataSourceSchema(ctx)
	state := tfsdk.State{Schema: schema}
	model := log_explorer_dataset.LogExplorerDatasetsDataSourceModel{
		AccountID:    accountID,
		ZoneID:       zoneID,
		IncludeZones: includeZones,
		MaxItems:     maxItems,
		Result:       customfield.NullObjectList[log_explorer_dataset.LogExplorerDatasetsResultDataSourceModel](ctx),
	}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set data source config: %v", diags)
	}
	return tfsdk.Config{Schema: schema, Raw: state.Raw}
}

func readLogExplorerDatasets(
	t *testing.T,
	ctx context.Context,
	d *log_explorer_dataset.LogExplorerDatasetsDataSource,
	config tfsdk.Config,
) log_explorer_dataset.LogExplorerDatasetsDataSourceModel {
	t.Helper()

	resp := &datasource.ReadResponse{State: tfsdk.State{Schema: config.Schema}}
	d.Read(ctx, datasource.ReadRequest{Config: config}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("read data source: %v", resp.Diagnostics)
	}
	var result log_explorer_dataset.LogExplorerDatasetsDataSourceModel
	if diags := resp.State.Get(ctx, &result); diags.HasError() {
		t.Fatalf("get data source state: %v", diags)
	}
	return result
}

func datasetSummaryJSON(index int) string {
	return fmt.Sprintf(`{
		"created_at":"2026-01-02T03:04:05Z",
		"dataset":"dns_logs",
		"dataset_id":"dataset-%d",
		"deletion_protection":true,
		"enabled":false,
		"object_id":"object-%d",
		"object_type":"zone",
		"updated_at":"2026-02-03T04:05:06Z"
	}`, index, index)
}

func datasetListResponse(count int) string {
	items := make([]string, count)
	for index := range count {
		items[index] = datasetSummaryJSON(index)
	}
	return fmt.Sprintf(`{"success":true,"errors":[],"messages":[],"result":[%s]}`, strings.Join(items, ","))
}

func TestLogExplorerDatasetsDataSourceRejectsNegativeMaxItems(t *testing.T) {
	ctx := context.Background()
	schema := log_explorer_dataset.ListDataSourceSchema(ctx)
	attribute, ok := schema.Attributes["max_items"].(datasourceschema.Int64Attribute)
	if !ok {
		t.Fatalf("max_items attribute type = %T, want schema.Int64Attribute", schema.Attributes["max_items"])
	}
	config := logExplorerDatasetsConfig(
		t,
		ctx,
		types.StringValue("account-id"),
		types.StringNull(),
		types.BoolNull(),
		types.Int64Value(-1),
	)
	for _, attributeValidator := range attribute.Validators {
		resp := &validator.Int64Response{}
		attributeValidator.ValidateInt64(ctx, validator.Int64Request{
			Path:           path.Root("max_items"),
			PathExpression: path.MatchRoot("max_items"),
			Config:         config,
			ConfigValue:    types.Int64Value(-1),
		}, resp)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	t.Fatal("expected max_items=-1 to be rejected")
}

func TestLogExplorerDatasetsDataSourceMaxItems(t *testing.T) {
	testCases := map[string]struct {
		maxItems types.Int64
		want     int
	}{
		"omitted defaults to 1000": {maxItems: types.Int64Null(), want: 1000},
		"zero defaults to 1000":    {maxItems: types.Int64Value(0), want: 1000},
		"positive limit":           {maxItems: types.Int64Value(1), want: 1},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			var requests atomic.Int64
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				requests.Add(1)
				if req.URL.RawQuery != "" {
					t.Errorf("query = %q, want max_items to remain provider-only", req.URL.RawQuery)
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, datasetListResponse(1001))
			}))
			defer server.Close()

			d := newLogExplorerDatasetsDataSource(t, server.URL)
			config := logExplorerDatasetsConfig(
				t,
				ctx,
				types.StringValue("account-id"),
				types.StringNull(),
				types.BoolNull(),
				testCase.maxItems,
			)
			state := readLogExplorerDatasets(t, ctx, d, config)
			if got := len(state.Result.Elements()); got != testCase.want {
				t.Fatalf("result count = %d, want %d", got, testCase.want)
			}
			if requests.Load() != 1 {
				t.Fatalf("request count = %d, want 1 for unpaginated endpoint", requests.Load())
			}
		})
	}
}

func TestLogExplorerDatasetsDataSourceUsesDatasetIDForIdentity(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, datasetListResponse(1))
	}))
	defer server.Close()

	d := newLogExplorerDatasetsDataSource(t, server.URL)
	config := logExplorerDatasetsConfig(
		t,
		ctx,
		types.StringValue("account-id"),
		types.StringNull(),
		types.BoolNull(),
		types.Int64Null(),
	)
	state := readLogExplorerDatasets(t, ctx, d, config)
	results, diags := state.Result.AsStructSliceT(ctx)
	if diags.HasError() {
		t.Fatalf("decode result: %v", diags)
	}
	if len(results) != 1 {
		t.Fatalf("result count = %d, want 1", len(results))
	}
	result := results[0]
	if result.ID.ValueString() != "dataset-0" {
		t.Errorf("id = %q, want dataset-0", result.ID.ValueString())
	}
	if result.DatasetID.ValueString() != "dataset-0" {
		t.Errorf("dataset_id = %q, want dataset-0", result.DatasetID.ValueString())
	}
}
