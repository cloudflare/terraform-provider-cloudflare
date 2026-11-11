package load_balancer

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func defaultPlanContext(t *testing.T) (tfsdk.State, tfsdk.Plan) {
	t.Helper()
	ctx := context.Background()
	s := ResourceSchema(ctx)
	state := tfsdk.State{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}
	diags := state.SetAttribute(ctx, path.Root("session_affinity"), types.StringValue("cookie"))
	if diags.HasError() {
		t.Fatalf("set prior affinity: %v", diags)
	}
	if diags := state.SetAttribute(ctx, path.Root("proxied"), types.BoolValue(true)); diags.HasError() {
		t.Fatalf("set prior proxy mode: %v", diags)
	}
	return state, tfsdk.Plan{Schema: s, Raw: state.Raw}
}

func TestSafeDefaultDescriptionConfigBoundary(t *testing.T) {
	state, plan := defaultPlanContext(t)
	for _, tc := range []struct {
		name                          string
		config, prior, proposed, want types.String
	}{
		{
			name: "omitted default", config: types.StringNull(), prior: types.StringValue(""),
			proposed: types.StringUnknown(), want: types.StringValue(""),
		},
		{
			name: "omitted null", config: types.StringNull(), prior: types.StringNull(),
			proposed: types.StringUnknown(), want: types.StringNull(),
		},
		{
			name: "removed nondefault", config: types.StringNull(), prior: types.StringValue("custom"),
			proposed: types.StringUnknown(), want: types.StringUnknown(),
		},
		{
			name: "unknown config", config: types.StringUnknown(), prior: types.StringValue(""),
			proposed: types.StringUnknown(), want: types.StringUnknown(),
		},
		{
			name: "explicit config", config: types.StringValue("new"), prior: types.StringValue(""),
			proposed: types.StringValue("new"), want: types.StringValue("new"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := planmodifier.StringRequest{
				State: state, Plan: plan, ConfigValue: tc.config,
				StateValue: tc.prior, PlanValue: tc.proposed,
			}
			resp := planmodifier.StringResponse{PlanValue: req.PlanValue}
			safeDefault{kind: "description"}.PlanModifyString(context.Background(), req, &resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Fatalf("planned %s, want %s", resp.PlanValue, tc.want)
			}
		})
	}
}

func TestSafeDefaultRemovedPoolMappings(t *testing.T) {
	state, plan := defaultPlanContext(t)
	elementType := types.ListType{ElemType: types.StringType}
	empty := types.MapValueMust(elementType, map[string]attr.Value{})
	configured := types.MapValueMust(elementType, map[string]attr.Value{
		"SJC": types.ListValueMust(types.StringType, []attr.Value{types.StringValue("old-pool")}),
	})
	for _, tc := range []struct {
		name        string
		prior, want types.Map
	}{
		{name: "empty mapping", prior: empty, want: empty},
		{name: "removed mapping", prior: configured, want: types.MapUnknown(elementType)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := planmodifier.MapRequest{
				State: state, Plan: plan, ConfigValue: types.MapNull(elementType),
				StateValue: tc.prior, PlanValue: types.MapUnknown(elementType),
			}
			resp := planmodifier.MapResponse{PlanValue: req.PlanValue}
			safeDefault{kind: "map"}.PlanModifyMap(context.Background(), req, &resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Fatalf("planned %s, want %s", resp.PlanValue, tc.want)
			}
		})
	}
}

func TestSafeDefaultAffinityDependencyChanges(t *testing.T) {
	ctx := context.Background()
	for _, tc := range []struct {
		name        string
		affinity    types.String
		proxied     types.Bool
		prior, want types.Float64
	}{
		{
			name: "cookie unchanged", affinity: types.StringValue("cookie"),
			proxied: types.BoolValue(true), prior: types.Float64Value(82800),
			want: types.Float64Value(82800),
		},
		{
			name: "header transition", affinity: types.StringValue("header"),
			proxied: types.BoolValue(true), prior: types.Float64Value(82800),
			want: types.Float64Unknown(),
		},
		{
			name: "unknown affinity", affinity: types.StringUnknown(),
			proxied: types.BoolValue(true), prior: types.Float64Value(82800),
			want: types.Float64Unknown(),
		},
		{
			name: "proxy transition", affinity: types.StringValue("cookie"),
			proxied: types.BoolValue(false), prior: types.Float64Value(82800),
			want: types.Float64Unknown(),
		},
		{
			name: "removed custom ttl", affinity: types.StringValue("cookie"),
			proxied: types.BoolValue(true), prior: types.Float64Value(3600),
			want: types.Float64Unknown(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state, plan := defaultPlanContext(t)
			if diags := plan.SetAttribute(ctx, path.Root("session_affinity"), tc.affinity); diags.HasError() {
				t.Fatalf("set planned affinity: %v", diags)
			}
			if diags := plan.SetAttribute(ctx, path.Root("proxied"), tc.proxied); diags.HasError() {
				t.Fatalf("set planned proxy mode: %v", diags)
			}
			req := planmodifier.Float64Request{
				State: state, Plan: plan, ConfigValue: types.Float64Null(),
				StateValue: tc.prior, PlanValue: types.Float64Unknown(),
			}
			resp := planmodifier.Float64Response{PlanValue: req.PlanValue}
			safeDefault{
				kind: "affinity_ttl", dependency: "session_affinity", also: "proxied",
			}.PlanModifyFloat64(ctx, req, &resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Fatalf("planned %s, want %s", resp.PlanValue, tc.want)
			}
		})
	}
}

func TestSafeDefaultTTLProxiedKeepsPriorValue(t *testing.T) {
	ctx := context.Background()
	for _, tc := range []struct {
		name        string
		proxied     types.Bool
		prior, want types.Float64
	}{
		{
			name: "proxied keeps server default", proxied: types.BoolValue(true),
			prior: types.Float64Value(30), want: types.Float64Value(30),
		},
		{
			name: "proxied keeps prior null", proxied: types.BoolValue(true),
			prior: types.Float64Null(), want: types.Float64Null(),
		},
		{
			name: "unproxied does not retain", proxied: types.BoolValue(false),
			prior: types.Float64Value(30), want: types.Float64Unknown(),
		},
		{
			name: "unknown proxied does not retain", proxied: types.BoolUnknown(),
			prior: types.Float64Value(30), want: types.Float64Unknown(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state, plan := defaultPlanContext(t)
			if diags := state.SetAttribute(ctx, path.Root("proxied"), tc.proxied); diags.HasError() {
				t.Fatalf("set prior proxy mode: %v", diags)
			}
			if diags := plan.SetAttribute(ctx, path.Root("proxied"), tc.proxied); diags.HasError() {
				t.Fatalf("set planned proxy mode: %v", diags)
			}
			req := planmodifier.Float64Request{
				State: state, Plan: plan, ConfigValue: types.Float64Null(),
				StateValue: tc.prior, PlanValue: types.Float64Unknown(),
			}
			resp := planmodifier.Float64Response{PlanValue: req.PlanValue}
			safeDefault{kind: "ttl", dependency: "proxied"}.PlanModifyFloat64(ctx, req, &resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Fatalf("planned %s, want %s", resp.PlanValue, tc.want)
			}
		})
	}
}

func TestSafeDefaultHeadersNeverRetainRemovedNames(t *testing.T) {
	state, plan := defaultPlanContext(t)
	for _, tc := range []struct {
		name        string
		prior, want types.List
	}{
		{
			name: "null headers", prior: types.ListNull(types.StringType),
			want: types.ListNull(types.StringType),
		},
		{
			name: "removed headers", want: types.ListUnknown(types.StringType),
			prior: types.ListValueMust(types.StringType, []attr.Value{types.StringValue("x-session")}),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := planmodifier.ListRequest{
				State: state, Plan: plan, ConfigValue: types.ListNull(types.StringType),
				StateValue: tc.prior, PlanValue: types.ListUnknown(types.StringType),
			}
			resp := planmodifier.ListResponse{PlanValue: req.PlanValue}
			safeDefault{kind: "headers"}.PlanModifyList(context.Background(), req, &resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Fatalf("planned %s, want %s", resp.PlanValue, tc.want)
			}
		})
	}
}
