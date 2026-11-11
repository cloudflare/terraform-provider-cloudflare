package workers_script

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestModifyPlanAssetsOnlyChangeMarksDeployComputedUnknown(t *testing.T) {
	ctx := context.Background()
	s := ResourceSchema(ctx)

	build := func(manifestSHA string) tfsdk.Plan {
		t.Helper()
		p := tfsdk.Plan{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}
		set := func(name string, v any) {
			if diags := p.SetAttribute(ctx, path.Root(name), v); diags.HasError() {
				t.Fatalf("setting %s: %v", name, diags)
			}
		}
		set("account_id", types.StringValue("account"))
		set("script_name", types.StringValue("app"))
		set("id", types.StringValue("app"))
		set("modified_on", timetypes.NewRFC3339ValueMust("2026-11-01T00:00:00Z"))
		set("has_assets", types.BoolValue(true))
		set("startup_time_ms", types.Int64Value(10))
		if diags := p.SetAttribute(ctx, path.Root("assets").AtName("asset_manifest_sha256"), types.StringValue(manifestSHA)); diags.HasError() {
			t.Fatalf("setting asset_manifest_sha256: %v", diags)
		}
		return p
	}

	state := build("old")
	plan := build("new")

	req := resource.ModifyPlanRequest{
		State: tfsdk.State{Schema: s, Raw: state.Raw},
		Plan:  plan,
	}
	resp := &resource.ModifyPlanResponse{Plan: plan}
	(&WorkersScriptResource{}).ModifyPlan(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ModifyPlan: %v", resp.Diagnostics)
	}

	var startupTime types.Int64
	resp.Diagnostics.Append(resp.Plan.GetAttribute(ctx, path.Root("startup_time_ms"), &startupTime)...)
	var modifiedOn timetypes.RFC3339
	resp.Diagnostics.Append(resp.Plan.GetAttribute(ctx, path.Root("modified_on"), &modifiedOn)...)
	if resp.Diagnostics.HasError() {
		t.Fatalf("reading plan: %v", resp.Diagnostics)
	}
	if !startupTime.IsUnknown() {
		t.Errorf("startup_time_ms = %s, want unknown: the API measures it again on every deploy", startupTime)
	}
	if !modifiedOn.IsUnknown() {
		t.Errorf("modified_on = %s, want unknown", modifiedOn)
	}
}

func TestModifyPlanNoChangeKeepsStartupTime(t *testing.T) {
	ctx := context.Background()
	s := ResourceSchema(ctx)

	p := tfsdk.Plan{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}
	if diags := p.SetAttribute(ctx, path.Root("script_name"), types.StringValue("app")); diags.HasError() {
		t.Fatalf("setting script_name: %v", diags)
	}
	if diags := p.SetAttribute(ctx, path.Root("startup_time_ms"), types.Int64Value(10)); diags.HasError() {
		t.Fatalf("setting startup_time_ms: %v", diags)
	}

	req := resource.ModifyPlanRequest{State: tfsdk.State{Schema: s, Raw: p.Raw}, Plan: p}
	resp := &resource.ModifyPlanResponse{Plan: p}
	(&WorkersScriptResource{}).ModifyPlan(ctx, req, resp)

	var startupTime types.Int64
	resp.Diagnostics.Append(resp.Plan.GetAttribute(ctx, path.Root("startup_time_ms"), &startupTime)...)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ModifyPlan: %v", resp.Diagnostics)
	}
	if startupTime.IsUnknown() || startupTime.ValueInt64() != 10 {
		t.Errorf("startup_time_ms = %s, want 10 when nothing changes", startupTime)
	}
}
