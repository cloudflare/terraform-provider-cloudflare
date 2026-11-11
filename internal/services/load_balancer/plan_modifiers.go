package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// safeDefault retains only values that are API defaults, not arbitrary values
// that may have been removed from configuration. Null is a meaningful default
// for some API fields and must not be excluded merely because it is null.
type safeDefault struct {
	kind       string
	dependency string
	also       string
	rule       bool
	ruleObject bool
}

var (
	_ planmodifier.Bool    = safeDefault{}
	_ planmodifier.Float64 = safeDefault{}
	_ planmodifier.List    = safeDefault{}
	_ planmodifier.Map     = safeDefault{}
	_ planmodifier.Object  = safeDefault{}
	_ planmodifier.String  = safeDefault{}
)

func (m safeDefault) Description(context.Context) string {
	return "Retains an unchanged API default when the attribute is omitted."
}

func (m safeDefault) MarkdownDescription(ctx context.Context) string { return m.Description(ctx) }

func (m safeDefault) keep(
	ctx context.Context, p path.Path, config tfsdk.Config, plan tfsdk.Plan, state tfsdk.State,
	configValue, planned, prior attr.Value,
) bool {
	if state.Raw.IsNull() || plan.Raw.IsNull() || !configValue.IsNull() ||
		!planned.IsUnknown() || prior.IsUnknown() {
		return false
	}
	for _, dependency := range [...]string{m.dependency, m.also} {
		if dependency == "" {
			continue
		}
		var before, after attr.Value
		dep := path.Root(dependency)
		if state.GetAttribute(ctx, dep, &before).HasError() ||
			plan.GetAttribute(ctx, dep, &after).HasError() ||
			before.IsNull() || before.IsUnknown() || after.IsNull() || after.IsUnknown() ||
			!before.Equal(after) {
			return false
		}
	}
	if m.rule && !sameRule(ctx, p, config, plan, state, m.ruleObject) {
		return false
	}
	return m.isDefault(ctx, plan, prior)
}

func (m safeDefault) isDefault(ctx context.Context, plan tfsdk.Plan, v attr.Value) bool {
	switch m.kind {
	case "null":
		return v.IsNull()
	case "description":
		return v.IsNull() || v.Equal(types.StringValue(""))
	case "map":
		if v.IsNull() {
			return true
		}
		value, ok := v.(types.Map)
		return ok && len(value.Elements()) == 0
	case "rules":
		if v.IsNull() {
			return true
		}
		value, ok := v.(types.List)
		return ok && len(value.Elements()) == 0
	case "pool_sets":
		if v.IsNull() {
			return true
		}
		value, ok := v.(types.List)
		return ok && len(value.Elements()) == 0
	case "networks":
		value, ok := v.(types.List)
		if !ok || v.IsNull() {
			return false
		}
		elements := value.Elements()
		return len(elements) == 1 && elements[0].Equal(types.StringValue("cloudflare"))
	case "ttl":
		// ttl applies only to gray-clouded (unproxied) load balancers. For a
		// proxied LB ttl is server-managed: the GET omits it and the PUT echoes
		// the server default, so there is no meaningful config-driven value to
		// diff. Keep whatever prior stable value state holds (null or the
		// preserved server default) rather than only null. keep()'s "proxied"
		// dependency guard already restricts this to proxied-unchanged updates,
		// so a proxied<->unproxied transition still re-resolves ttl normally.
		var proxied types.Bool
		if plan.GetAttribute(ctx, path.Root("proxied"), &proxied).HasError() ||
			proxied.IsNull() || proxied.IsUnknown() {
			return false
		}
		return proxied.ValueBool()
	case "affinity_ttl":
		var affinity types.String
		if plan.GetAttribute(ctx, path.Root("session_affinity"), &affinity).HasError() ||
			affinity.IsNull() || affinity.IsUnknown() {
			return false
		}
		switch affinity.ValueString() {
		case "none":
			return v.IsNull()
		case "cookie", "ip_cookie":
			return v.Equal(types.Float64Value(82800))
		case "header":
			return v.Equal(types.Float64Value(1800))
		}
		return false
	case "drain":
		return v.Equal(types.Float64Value(0))
	case "headers":
		return v.IsNull()
	case "require_all":
		return v.IsNull() || v.Equal(types.BoolValue(false))
	case "auto":
		return v.Equal(types.StringValue("Auto"))
	case "none":
		return v.Equal(types.StringValue("none"))
	}
	object, ok := v.(types.Object)
	if !ok || v.IsNull() {
		return false
	}
	a := object.Attributes()
	switch m.kind {
	case "adaptive":
		return a["failover_across_pools"].Equal(types.BoolValue(false))
	case "location":
		return a["mode"].Equal(types.StringValue("pop")) &&
			a["prefer_ecs"].Equal(types.StringValue("proximity"))
	case "random":
		return a["default_weight"].Equal(types.Float64Value(1)) && a["pool_weights"].IsNull()
	case "affinity_attributes":
		return a["drain_duration"].Equal(types.Float64Value(0)) && a["headers"].IsNull() &&
			(a["require_all_headers"].IsNull() || a["require_all_headers"].Equal(types.BoolValue(false))) &&
			a["samesite"].Equal(types.StringValue("Auto")) &&
			a["secure"].Equal(types.StringValue("Auto")) &&
			a["zero_downtime_failover"].Equal(types.StringValue("none"))
	}
	return false
}

func (m safeDefault) PlanModifyBool(
	ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse,
) {
	if m.keep(
		ctx, req.Path, req.Config, req.Plan, req.State,
		req.ConfigValue, req.PlanValue, req.StateValue,
	) {
		resp.PlanValue = req.StateValue
	}
}
func (m safeDefault) PlanModifyFloat64(
	ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response,
) {
	if m.keep(
		ctx, req.Path, req.Config, req.Plan, req.State,
		req.ConfigValue, req.PlanValue, req.StateValue,
	) {
		resp.PlanValue = req.StateValue
	}
}
func (m safeDefault) PlanModifyList(
	ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse,
) {
	if m.keep(
		ctx, req.Path, req.Config, req.Plan, req.State,
		req.ConfigValue, req.PlanValue, req.StateValue,
	) {
		resp.PlanValue = req.StateValue
	}
}
func (m safeDefault) PlanModifyMap(
	ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse,
) {
	if m.keep(
		ctx, req.Path, req.Config, req.Plan, req.State,
		req.ConfigValue, req.PlanValue, req.StateValue,
	) {
		resp.PlanValue = req.StateValue
	}
}
func (m safeDefault) PlanModifyObject(
	ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse,
) {
	if m.keep(
		ctx, req.Path, req.Config, req.Plan, req.State,
		req.ConfigValue, req.PlanValue, req.StateValue,
	) {
		resp.PlanValue = req.StateValue
	}
}
func (m safeDefault) PlanModifyString(
	ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse,
) {
	if m.keep(
		ctx, req.Path, req.Config, req.Plan, req.State,
		req.ConfigValue, req.PlanValue, req.StateValue,
	) {
		resp.PlanValue = req.StateValue
	}
}

// List element state is positional. Never use another rule's state after a
// reorder, or a rule's old defaults after it has otherwise changed.
func sameRule(
	ctx context.Context, p path.Path, config tfsdk.Config, plan tfsdk.Plan,
	state tfsdk.State, ruleObject bool,
) bool {
	rule := p.ParentPath()
	if !ruleObject {
		rule = rule.ParentPath()
	}
	var name types.String
	if plan.GetAttribute(ctx, rule.AtName("name"), &name).HasError() ||
		name.IsNull() || name.IsUnknown() || name.ValueString() == "" {
		return false
	}
	for _, field := range []string{
		"name", "condition", "disabled", "fixed_response", "priority", "terminates",
	} {
		var before, after attr.Value
		fieldPath := rule.AtName(field)
		if state.GetAttribute(ctx, fieldPath, &before).HasError() ||
			plan.GetAttribute(ctx, fieldPath, &after).HasError() ||
			before.IsUnknown() || after.IsUnknown() || !before.Equal(after) {
			return false
		}
	}
	for _, field := range []string{
		"adaptive_routing", "country_pools", "default_pools", "fallback_pool", "location_strategy",
		"pop_pools", "random_steering", "region_pools", "session_affinity",
		"session_affinity_attributes", "session_affinity_ttl", "steering_policy", "ttl",
	} {
		var before, after, configured attr.Value
		fieldPath := rule.AtName("overrides").AtName(field)
		if fieldPath.Equal(p) {
			continue
		}
		if state.GetAttribute(ctx, fieldPath, &before).HasError() ||
			plan.GetAttribute(ctx, fieldPath, &after).HasError() ||
			config.GetAttribute(ctx, fieldPath, &configured).HasError() ||
			before.IsUnknown() || configured.IsUnknown() {
			return false
		}
		if after.IsUnknown() {
			if !before.IsNull() {
				return false
			}
		} else if !before.Equal(after) {
			return false
		}
	}
	return true
}
