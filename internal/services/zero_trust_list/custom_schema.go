package zero_trust_list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// customResourceSchema wraps the generated ResourceSchema and adds UseStateForUnknown
// plan modifiers to computed fields that cause permadiffs when unset.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6580
func customResourceSchema(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)

	setStringPlanModifier(s.Attributes, "description", stringplanmodifier.UseStateForUnknown())
	setFloat64PlanModifier(s.Attributes, "list_count", float64planmodifier.UseStateForUnknown())

	return s
}

func setStringPlanModifier(attrs map[string]schema.Attribute, name string, pm planmodifier.String) {
	if a, ok := attrs[name].(schema.StringAttribute); ok {
		a.PlanModifiers = append(a.PlanModifiers, pm)
		attrs[name] = a
	}
}

func setFloat64PlanModifier(attrs map[string]schema.Attribute, name string, pm planmodifier.Float64) {
	if a, ok := attrs[name].(schema.Float64Attribute); ok {
		a.PlanModifiers = append(a.PlanModifiers, pm)
		attrs[name] = a
	}
}
