package zero_trust_gateway_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// customResourceSchema wraps the generated ResourceSchema and adds UseStateForUnknown
// plan modifiers to computed fields that cause permadiffs when unset.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6580
func customResourceSchema(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)

	// Top-level server-managed computed fields.
	setStringPlanModifier(s.Attributes, "created_at", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "updated_at", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "deleted_at", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "source_account", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "warning_status", stringplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(s.Attributes, "read_only", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(s.Attributes, "sharable", boolplanmodifier.UseStateForUnknown())
	setInt64PlanModifier(s.Attributes, "version", int64planmodifier.UseStateForUnknown())

	// rule_settings computed+optional sub-fields.
	rs := s.Attributes["rule_settings"].(schema.SingleNestedAttribute)
	setBoolPlanModifier(rs.Attributes, "allow_child_bypass", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(rs.Attributes, "block_page_enabled", boolplanmodifier.UseStateForUnknown())
	setStringPlanModifier(rs.Attributes, "block_reason", stringplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(rs.Attributes, "ignore_cname_category_matches", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(rs.Attributes, "insecure_disable_dnssec_validation", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(rs.Attributes, "ip_categories", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(rs.Attributes, "ip_indicator_feeds", boolplanmodifier.UseStateForUnknown())
	setStringPlanModifier(rs.Attributes, "override_host", stringplanmodifier.UseStateForUnknown())
	setListPlanModifier(rs.Attributes, "override_ips", listplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(rs.Attributes, "resolve_dns_through_cloudflare", boolplanmodifier.UseStateForUnknown())
	s.Attributes["rule_settings"] = rs

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

func setInt64PlanModifier(attrs map[string]schema.Attribute, name string, pm planmodifier.Int64) {
	if a, ok := attrs[name].(schema.Int64Attribute); ok {
		a.PlanModifiers = append(a.PlanModifiers, pm)
		attrs[name] = a
	}
}

func setListPlanModifier(attrs map[string]schema.Attribute, name string, pm planmodifier.List) {
	if a, ok := attrs[name].(schema.ListAttribute); ok {
		a.PlanModifiers = append(a.PlanModifiers, pm)
		attrs[name] = a
	}
}
