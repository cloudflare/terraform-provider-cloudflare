package zero_trust_gateway_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// customResourceSchema wraps the generated ResourceSchema and adds UseStateForUnknown
// plan modifiers to computed fields that cause permadiffs when unset.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6580
func customResourceSchema(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)

	setStringPlanModifier(s.Attributes, "binding_status", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "certificate", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "fingerprint", stringplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(s.Attributes, "in_use", boolplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "issuer_org", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "issuer_raw", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "type", stringplanmodifier.UseStateForUnknown())

	return s
}

func setStringPlanModifier(attrs map[string]schema.Attribute, name string, pm planmodifier.String) {
	if a, ok := attrs[name].(schema.StringAttribute); ok {
		a.PlanModifiers = append(a.PlanModifiers, pm)
		attrs[name] = a
	}
}

func setBoolPlanModifier(attrs map[string]schema.Attribute, name string, pm planmodifier.Bool) {
	if a, ok := attrs[name].(schema.BoolAttribute); ok {
		a.PlanModifiers = append(a.PlanModifiers, pm)
		attrs[name] = a
	}
}
