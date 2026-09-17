package flagship_flag_test

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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/flagship_flag"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const flagshipFlagTestResponse = `{
	"success": true,
	"errors": [],
	"messages": [],
	"result": {
		"key": "example-flag",
		"default_variation": "disabled",
		"enabled": true,
		"variations": {"enabled": true, "disabled": false},
		"rules": [],
		"description": "",
		"type": "boolean",
		"updated_at": "2026-08-13T00:00:00Z",
		"updated_by": "test@example.com"
	}
}`

func newTestFlagshipFlagResource(t *testing.T, baseURL string) *flagship_flag.FlagshipFlagResource {
	t.Helper()
	r := flagship_flag.NewResource().(*flagship_flag.FlagshipFlagResource)
	client := cloudflare.NewClient(option.WithBaseURL(baseURL), option.WithAPIToken("test-token"))
	resp := &resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("configure resource: %v", resp.Diagnostics)
	}
	return r
}

func newTestFlagshipFlagDataSource(t *testing.T, baseURL string) *flagship_flag.FlagshipFlagDataSource {
	t.Helper()
	d := flagship_flag.NewFlagshipFlagDataSource().(*flagship_flag.FlagshipFlagDataSource)
	client := cloudflare.NewClient(option.WithBaseURL(baseURL), option.WithAPIToken("test-token"))
	resp := &datasource.ConfigureResponse{}
	d.Configure(context.Background(), datasource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("configure data source: %v", resp.Diagnostics)
	}
	return d
}

func flagshipFlagTestModel(key string) flagship_flag.FlagshipFlagModel {
	variations := map[string]types.String{
		"enabled":  types.StringValue("true"),
		"disabled": types.StringValue("false"),
	}
	rules := []*flagship_flag.FlagshipFlagRulesModel{}
	return flagship_flag.FlagshipFlagModel{
		AccountID:        types.StringValue("acct"),
		AppID:            types.StringValue("app"),
		DefaultVariation: types.StringValue("disabled"),
		Enabled:          types.BoolValue(true),
		Key:              types.StringValue(key),
		Variations:       &variations,
		Rules:            &rules,
		Description:      types.StringValue(""),
		Type:             types.StringValue("boolean"),
		UpdatedAt:        types.StringNull(),
		UpdatedBy:        types.StringNull(),
	}
}

func flagshipFlagTestState(t *testing.T, ctx context.Context, model flagship_flag.FlagshipFlagModel) tfsdk.State {
	t.Helper()
	state := tfsdk.State{Schema: flagship_flag.ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set state: %v", diags)
	}
	return state
}

func flagshipFlagTestPlan(t *testing.T, ctx context.Context, model flagship_flag.FlagshipFlagModel) tfsdk.Plan {
	t.Helper()
	plan := tfsdk.Plan{Schema: flagship_flag.ResourceSchema(ctx)}
	if diags := plan.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set plan: %v", diags)
	}
	return plan
}

func TestFlagshipFlagCreateThenReadUsesKey(t *testing.T) {
	ctx := context.Background()
	var methods []string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		methods = append(methods, req.Method)
		if req.URL.Path != "/accounts/acct/flagship/apps/app/flags" &&
			req.URL.Path != "/accounts/acct/flagship/apps/app/flags/example-flag" {
			t.Errorf("unexpected path: %s", req.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, flagshipFlagTestResponse)
	}))
	defer ts.Close()

	r := newTestFlagshipFlagResource(t, ts.URL)
	model := flagshipFlagTestModel("example-flag")
	createResp := &resource.CreateResponse{State: tfsdk.State{Schema: flagship_flag.ResourceSchema(ctx)}}
	r.Create(ctx, resource.CreateRequest{Plan: flagshipFlagTestPlan(t, ctx, model)}, createResp)
	if createResp.Diagnostics.HasError() {
		t.Fatalf("create failed: %v", createResp.Diagnostics)
	}

	readResp := &resource.ReadResponse{State: tfsdk.State{Schema: flagship_flag.ResourceSchema(ctx)}}
	r.Read(ctx, resource.ReadRequest{State: createResp.State}, readResp)
	if readResp.Diagnostics.HasError() {
		t.Fatalf("read failed: %v", readResp.Diagnostics)
	}

	if len(methods) != 2 || methods[0] != http.MethodPost || methods[1] != http.MethodGet {
		t.Fatalf("methods = %v, want [POST GET]", methods)
	}
	var got flagship_flag.FlagshipFlagModel
	if diags := readResp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("decode state: %v", diags)
	}
	if got.Key.ValueString() != "example-flag" {
		t.Fatalf("key = %q", got.Key.ValueString())
	}
}

func TestFlagshipFlagUpdateAddressesPriorKey(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", req.Method)
		}
		if req.URL.Path != "/accounts/acct/flagship/apps/app/flags/old-key" {
			t.Errorf("path = %s, want old key", req.URL.Path)
		}
		var body map[string]any
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body["key"] != "example-flag" {
			t.Errorf("body key = %v", body["key"])
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, flagshipFlagTestResponse)
	}))
	defer ts.Close()

	r := newTestFlagshipFlagResource(t, ts.URL)
	resp := &resource.UpdateResponse{State: tfsdk.State{Schema: flagship_flag.ResourceSchema(ctx)}}
	r.Update(ctx, resource.UpdateRequest{
		Plan:  flagshipFlagTestPlan(t, ctx, flagshipFlagTestModel("example-flag")),
		State: flagshipFlagTestState(t, ctx, flagshipFlagTestModel("old-key")),
	}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("update failed: %v", resp.Diagnostics)
	}
}

func TestFlagshipFlagDeleteUsesKey(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", req.Method)
		}
		if req.URL.Path != "/accounts/acct/flagship/apps/app/flags/example-flag" {
			t.Errorf("unexpected path: %s", req.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"success":true,"errors":[],"messages":[],"result":{"key":"example-flag"}}`)
	}))
	defer ts.Close()

	r := newTestFlagshipFlagResource(t, ts.URL)
	state := flagshipFlagTestState(t, ctx, flagshipFlagTestModel("example-flag"))
	resp := &resource.DeleteResponse{State: state}
	r.Delete(ctx, resource.DeleteRequest{State: state}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("delete failed: %v", resp.Diagnostics)
	}
}

func TestFlagshipFlagDataSourceUsesKey(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/accounts/acct/flagship/apps/app/flags/example-flag" {
			t.Errorf("unexpected path: %s", req.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, flagshipFlagTestResponse)
	}))
	defer ts.Close()

	model := flagship_flag.FlagshipFlagDataSourceModel{
		AccountID:  types.StringValue("acct"),
		AppID:      types.StringValue("app"),
		Key:        types.StringValue("example-flag"),
		Variations: customfield.NullMap[types.String](ctx),
		Rules:      customfield.NullObjectList[flagship_flag.FlagshipFlagRulesDataSourceModel](ctx),
	}
	dsSchema := flagship_flag.DataSourceSchema(ctx)
	configState := tfsdk.State{Schema: dsSchema}
	if diags := configState.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set config: %v", diags)
	}
	config := tfsdk.Config{Schema: dsSchema, Raw: configState.Raw}

	d := newTestFlagshipFlagDataSource(t, ts.URL)
	resp := &datasource.ReadResponse{State: tfsdk.State{Schema: dsSchema}}
	d.Read(ctx, datasource.ReadRequest{Config: config}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("read failed: %v", resp.Diagnostics)
	}
	var got flagship_flag.FlagshipFlagDataSourceModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("decode state: %v", diags)
	}
	if got.Key.ValueString() != "example-flag" {
		t.Fatalf("key = %q", got.Key.ValueString())
	}
}
