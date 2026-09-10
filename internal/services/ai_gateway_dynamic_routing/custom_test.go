package ai_gateway_dynamic_routing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func newTestResource(t *testing.T, baseURL string) *AIGatewayDynamicRoutingResource {
	t.Helper()

	r := NewResource().(*AIGatewayDynamicRoutingResource)
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

func newTestDataSource(t *testing.T, baseURL string) *AIGatewayDynamicRoutingDataSource {
	t.Helper()

	d := NewAIGatewayDynamicRoutingDataSource().(*AIGatewayDynamicRoutingDataSource)
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
	)
	resp := &datasource.ConfigureResponse{}
	d.Configure(context.Background(), datasource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure failed: %v", resp.Diagnostics)
	}
	return d
}

func newTestState(t *testing.T, ctx context.Context, elements []*AIGatewayDynamicRoutingElementsModel) tfsdk.State {
	t.Helper()

	model := AIGatewayDynamicRoutingModel{
		ID:         types.StringValue("route-1"),
		AccountID:  types.StringValue("account-1"),
		GatewayID:  types.StringValue("gateway-1"),
		Elements:   &elements,
		Name:       types.StringValue("route"),
		CreatedAt:  newNullTime(),
		ModifiedAt: newNullTime(),
		Success:    types.BoolNull(),
		Deployment: customfield.NullObject[AIGatewayDynamicRoutingDeploymentModel](ctx),
		Route:      customfield.NullObject[AIGatewayDynamicRoutingRouteModel](ctx),
		Version:    customfield.NullObject[AIGatewayDynamicRoutingVersionModel](ctx),
	}

	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to create test state: %v", diags)
	}
	return state
}

func newTestDataSourceConfig(t *testing.T, ctx context.Context) tfsdk.Config {
	t.Helper()

	model := AIGatewayDynamicRoutingDataSourceModel{
		AccountID:  types.StringValue("account-1"),
		GatewayID:  types.StringValue("gateway-1"),
		ID:         types.StringValue("route-1"),
		CreatedAt:  newNullTime(),
		ModifiedAt: newNullTime(),
		Name:       types.StringNull(),
		Deployment: customfield.NullObject[AIGatewayDynamicRoutingDeploymentDataSourceModel](ctx),
		Elements:   customfield.NullObjectList[AIGatewayDynamicRoutingElementsDataSourceModel](ctx),
		Version:    customfield.NullObject[AIGatewayDynamicRoutingVersionDataSourceModel](ctx),
	}

	state := tfsdk.State{Schema: DataSourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to create test data source config: %v", diags)
	}
	return tfsdk.Config{Schema: state.Schema, Raw: state.Raw}
}

func newNullTime() timetypes.RFC3339 {
	return timetypes.NewRFC3339Null()
}

func testElements(next string) []*AIGatewayDynamicRoutingElementsModel {
	return []*AIGatewayDynamicRoutingElementsModel{
		{
			ID:   types.StringValue("start"),
			Type: types.StringValue("start"),
			Outputs: &AIGatewayDynamicRoutingElementsOutputsModel{
				Next: &AIGatewayDynamicRoutingElementsOutputsNextModel{
					ElementID: types.StringValue(next),
				},
			},
		},
	}
}

func testResponse(graph, versionData string, documented bool) string {
	elements := ""
	if documented {
		elements = fmt.Sprintf(",\"elements\":%s", graph)
	}

	data := graph
	if documented {
		data = fmt.Sprintf("%q", versionData)
	}

	return fmt.Sprintf(`{"result":{"id":"route-1","name":"route","gateway_id":"gateway-1"%s,"version":{"active":"true","data":%s,"version_id":"version-1"}},"success":true}`, elements, data)
}

func testServer(t *testing.T, response *string) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/accounts/account-1/ai-gateway/gateways/gateway-1/routes/route-1" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(*response))
	}))
}

func readState(t *testing.T, ctx context.Context, resp *resource.ReadResponse) AIGatewayDynamicRoutingModel {
	t.Helper()

	var model AIGatewayDynamicRoutingModel
	if diags := resp.State.Get(ctx, &model); diags.HasError() {
		t.Fatalf("failed to read response state: %v", diags)
	}
	return model
}

func TestNormalizeDynamicRoutingResponse(t *testing.T) {
	graph := `[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"}]`

	tests := map[string]struct {
		input            string
		wantElementCount int
		wantElementID    string
		wantVersionData  string
		wantUnchanged    bool
	}{
		"array version data": {
			input:            testResponse(graph, "", false),
			wantElementCount: 1,
			wantElementID:    "start",
			wantVersionData:  graph,
		},
		"documented response": {
			input:            testResponse(graph, "documented", true),
			wantElementCount: 1,
			wantElementID:    "start",
			wantVersionData:  "documented",
			wantUnchanged:    true,
		},
		"documented elements with array version data": {
			input:            fmt.Sprintf(`{"result":{"elements":[{"id":"canonical","outputs":{},"type":"end"}],"version":{"data":%s}},"success":true}`, graph),
			wantElementCount: 1,
			wantElementID:    "canonical",
			wantVersionData:  graph,
		},
		"null elements with array version data": {
			input:            fmt.Sprintf(`{"result":{"elements":null,"version":{"data":%s}},"success":true}`, graph),
			wantElementCount: 1,
			wantElementID:    "start",
			wantVersionData:  graph,
		},
		"empty elements with array version data": {
			input:            fmt.Sprintf(`{"result":{"elements":[],"version":{"data":%s}},"success":true}`, graph),
			wantElementCount: 1,
			wantElementID:    "start",
			wantVersionData:  graph,
		},
		"whitespace empty elements with array version data": {
			input:            fmt.Sprintf(`{"result":{"elements":[  ],"version":{"data":%s}},"success":true}`, graph),
			wantElementCount: 1,
			wantElementID:    "start",
			wantVersionData:  graph,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := normalizeDynamicRoutingResponse([]byte(tt.input))
			if err != nil {
				t.Fatalf("normalizeDynamicRoutingResponse returned error: %v", err)
			}
			if tt.wantUnchanged && string(got) != tt.input {
				t.Fatalf("documented response was changed: %s", got)
			}

			var envelope struct {
				Result struct {
					Elements []map[string]any `json:"elements"`
					Version  struct {
						Data string `json:"data"`
					} `json:"version"`
				} `json:"result"`
			}
			if err := json.Unmarshal(got, &envelope); err != nil {
				t.Fatalf("normalized response is invalid JSON: %v", err)
			}
			if len(envelope.Result.Elements) != tt.wantElementCount {
				t.Fatalf("unexpected elements: %#v", envelope.Result.Elements)
			}
			if tt.wantElementID != "" && envelope.Result.Elements[0]["id"] != tt.wantElementID {
				t.Errorf("element id = %v, want %q", envelope.Result.Elements[0]["id"], tt.wantElementID)
			}
			if envelope.Result.Version.Data != tt.wantVersionData {
				t.Errorf("version.data = %q, want %q", envelope.Result.Version.Data, tt.wantVersionData)
			}
		})
	}
}

func TestNormalizeDynamicRoutingResponse_PreservesUnsupportedPayloads(t *testing.T) {
	tests := []string{
		`{"result":{"version":{"data":{"not":"an array"}}},"success":true}`,
		`{"result":{"version":{"data":"already a string"}},"success":true}`,
		`{"result":{"version":{"data":[null]}},"success":true}`,
		`{"result":{"version":{"data":[1]}},"success":true}`,
		`{"result":{"version":{"data":[{}]}},"success":true}`,
		`{"result":{"elements":[],"version":{"data":[{}]}},"success":true}`,
		`{"result":`,
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			got, err := normalizeDynamicRoutingResponse([]byte(input))
			if err != nil {
				t.Fatalf("normalizeDynamicRoutingResponse returned error: %v", err)
			}
			if string(got) != input {
				t.Fatalf("unsupported response changed: %s", got)
			}
		})
	}
}

func TestRead_ArrayVersionDataPreservesAndUpdatesElements(t *testing.T) {
	ctx := context.Background()
	response := testResponse(`[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"}]`, "", false)
	ts := testServer(t, &response)
	defer ts.Close()

	r := newTestResource(t, ts.URL)
	state := newTestState(t, ctx, testElements("end"))
	readResp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, readResp)
	if readResp.Diagnostics.HasError() {
		t.Fatalf("Read returned diagnostics: %v", readResp.Diagnostics)
	}

	model := readState(t, ctx, readResp)
	if model.Elements == nil || len(*model.Elements) != 1 || (*model.Elements)[0].ID.ValueString() != "start" {
		t.Fatalf("array-valued version.data did not populate elements: %#v", model.Elements)
	}

	response = testResponse(`[{"id":"changed","outputs":{"next":{"elementId":"other"}},"type":"start"}]`, "", false)
	changedResp := &resource.ReadResponse{State: readResp.State}
	r.Read(ctx, resource.ReadRequest{State: readResp.State}, changedResp)
	if changedResp.Diagnostics.HasError() {
		t.Fatalf("Read with changed graph returned diagnostics: %v", changedResp.Diagnostics)
	}

	changed := readState(t, ctx, changedResp)
	if changed.Elements == nil || len(*changed.Elements) != 1 || (*changed.Elements)[0].ID.ValueString() != "changed" {
		t.Fatalf("remote graph change was not reflected in state: %#v", changed.Elements)
	}
	if (*changed.Elements)[0].Outputs == nil || (*changed.Elements)[0].Outputs.Next == nil || (*changed.Elements)[0].Outputs.Next.ElementID.ValueString() != "other" {
		t.Fatalf("remote graph edge change was not reflected in state: %#v", (*changed.Elements)[0].Outputs)
	}
}

func TestRead_DocumentedResponsePreservesElements(t *testing.T) {
	ctx := context.Background()
	graph := `[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"}]`
	response := testResponse(graph, "documented", true)
	ts := testServer(t, &response)
	defer ts.Close()

	r := newTestResource(t, ts.URL)
	state := newTestState(t, ctx, testElements("end"))
	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read returned diagnostics: %v", resp.Diagnostics)
	}

	model := readState(t, ctx, resp)
	if model.Elements == nil || len(*model.Elements) != 1 || (*model.Elements)[0].ID.ValueString() != "start" {
		t.Fatalf("documented elements response was not preserved: %#v", model.Elements)
	}
}

func TestDataSourceRead_ArrayVersionDataPopulatesElements(t *testing.T) {
	ctx := context.Background()
	response := testResponse(`[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"}]`, "", false)
	ts := testServer(t, &response)
	defer ts.Close()

	d := newTestDataSource(t, ts.URL)
	config := newTestDataSourceConfig(t, ctx)
	resp := &datasource.ReadResponse{State: tfsdk.State{Schema: DataSourceSchema(ctx)}}
	d.Read(ctx, datasource.ReadRequest{Config: config}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read returned diagnostics: %v", resp.Diagnostics)
	}

	var model AIGatewayDynamicRoutingDataSourceModel
	if diags := resp.State.Get(ctx, &model); diags.HasError() {
		t.Fatalf("failed to read data source state: %v", diags)
	}
	elements, diags := model.Elements.AsStructSliceT(ctx)
	if diags.HasError() {
		t.Fatalf("failed to read data source elements: %v", diags)
	}
	if len(elements) != 1 || elements[0].ID.ValueString() != "start" {
		t.Fatalf("array-valued version.data did not populate data source elements: %#v", model.Elements)
	}
	if model.AccountID.ValueString() != "account-1" || model.GatewayID.ValueString() != "gateway-1" || model.ID.ValueString() != "route-1" {
		t.Fatalf("data source identity changed: account=%q gateway=%q id=%q", model.AccountID.ValueString(), model.GatewayID.ValueString(), model.ID.ValueString())
	}
	version, diags := model.Version.Value(ctx)
	if diags.HasError() {
		t.Fatalf("failed to read data source version: %v", diags)
	}
	if version == nil || version.Data.ValueString() != `[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"}]` {
		t.Fatalf("data source version.data was not normalized: %#v", version)
	}
}

func TestImportState_ArrayVersionDataPopulatesElements(t *testing.T) {
	ctx := context.Background()
	graph := `[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"}]`
	response := testResponse(graph, "", false)
	ts := testServer(t, &response)
	defer ts.Close()

	r := newTestResource(t, ts.URL)
	resp := &resource.ImportStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: "account-1/gateway-1/route-1"}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ImportState returned diagnostics: %v", resp.Diagnostics)
	}

	model := readState(t, ctx, &resource.ReadResponse{State: resp.State})
	if model.Elements == nil || len(*model.Elements) != 1 || (*model.Elements)[0].ID.ValueString() != "start" {
		t.Fatalf("array-valued version.data did not populate imported elements: %#v", model.Elements)
	}
	if model.AccountID.ValueString() != "account-1" || model.GatewayID.ValueString() != "gateway-1" || model.ID.ValueString() != "route-1" {
		t.Fatalf("imported identity changed: account=%q gateway=%q id=%q", model.AccountID.ValueString(), model.GatewayID.ValueString(), model.ID.ValueString())
	}
	version, diags := model.Version.Value(ctx)
	if diags.HasError() {
		t.Fatalf("failed to read imported version: %v", diags)
	}
	if version == nil || version.Data.ValueString() != graph {
		t.Fatalf("imported version.data was not normalized: %#v", version)
	}
}
