package zero_trust_gateway_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// customResourceSchema wraps the generated ResourceSchema and removes the
// Computed flag from user-configurable rule_settings sub-fields.
//
// The generated schema marks these fields as Computed+Optional because they
// appear in both the API request and response schemas. However, they are not
// server-computed — the API returns what the user set, or null when unset.
// The Computed flag causes Terraform to mark them unknown on every update,
// producing noisy plans with (known after apply) on unrelated changes.
//
// Removing Computed makes Terraform treat them as Optional-only: the plan
// uses the config value (or null) directly, with no unknown marking.
func customResourceSchema(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)

	rs := s.Attributes["rule_settings"].(schema.SingleNestedAttribute)
	removeBoolComputed(rs.Attributes, "allow_child_bypass")
	removeBoolComputed(rs.Attributes, "block_page_enabled")
	removeStringComputed(rs.Attributes, "block_reason")
	removeBoolComputed(rs.Attributes, "ignore_cname_category_matches")
	removeBoolComputed(rs.Attributes, "insecure_disable_dnssec_validation")
	removeBoolComputed(rs.Attributes, "ip_categories")
	removeBoolComputed(rs.Attributes, "ip_indicator_feeds")
	removeStringComputed(rs.Attributes, "override_host")
	removeListComputed(rs.Attributes, "override_ips")
	removeBoolComputed(rs.Attributes, "resolve_dns_through_cloudflare")
	s.Attributes["rule_settings"] = rs

	return s
}

func removeBoolComputed(attrs map[string]schema.Attribute, name string) {
	if a, ok := attrs[name].(schema.BoolAttribute); ok {
		a.Computed = false
		attrs[name] = a
	}
}

func removeStringComputed(attrs map[string]schema.Attribute, name string) {
	if a, ok := attrs[name].(schema.StringAttribute); ok {
		a.Computed = false
		attrs[name] = a
	}
}

func removeListComputed(attrs map[string]schema.Attribute, name string) {
	if a, ok := attrs[name].(schema.ListAttribute); ok {
		a.Computed = false
		attrs[name] = a
	}
}
