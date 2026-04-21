package zero_trust_gateway_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// customResourceSchema wraps the generated ResourceSchema and adds UseStateForUnknown
// plan modifiers to computed fields that cause permadiffs when unset.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6580
func customResourceSchema(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)

	// Top-level timestamps.
	setStringPlanModifier(s.Attributes, "created_at", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(s.Attributes, "updated_at", stringplanmodifier.UseStateForUnknown())

	// Navigate into settings -> antivirus.
	settings := s.Attributes["settings"].(schema.SingleNestedAttribute)

	antivirus := settings.Attributes["antivirus"].(schema.SingleNestedAttribute)
	setBoolPlanModifier(antivirus.Attributes, "enabled_download_phase", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(antivirus.Attributes, "enabled_upload_phase", boolplanmodifier.UseStateForUnknown())
	setBoolPlanModifier(antivirus.Attributes, "fail_closed", boolplanmodifier.UseStateForUnknown())
	settings.Attributes["antivirus"] = antivirus

	// settings -> block_page.
	blockPage := settings.Attributes["block_page"].(schema.SingleNestedAttribute)
	setBoolPlanModifier(blockPage.Attributes, "read_only", boolplanmodifier.UseStateForUnknown())
	setStringPlanModifier(blockPage.Attributes, "source_account", stringplanmodifier.UseStateForUnknown())
	setInt64PlanModifier(blockPage.Attributes, "version", int64planmodifier.UseStateForUnknown())
	settings.Attributes["block_page"] = blockPage

	// settings -> custom_certificate.
	customCert := settings.Attributes["custom_certificate"].(schema.SingleNestedAttribute)
	setStringPlanModifier(customCert.Attributes, "binding_status", stringplanmodifier.UseStateForUnknown())
	setStringPlanModifier(customCert.Attributes, "updated_at", stringplanmodifier.UseStateForUnknown())
	settings.Attributes["custom_certificate"] = customCert

	// settings -> extended_email_matching.
	extEmail := settings.Attributes["extended_email_matching"].(schema.SingleNestedAttribute)
	setBoolPlanModifier(extEmail.Attributes, "read_only", boolplanmodifier.UseStateForUnknown())
	setStringPlanModifier(extEmail.Attributes, "source_account", stringplanmodifier.UseStateForUnknown())
	setInt64PlanModifier(extEmail.Attributes, "version", int64planmodifier.UseStateForUnknown())
	settings.Attributes["extended_email_matching"] = extEmail

	s.Attributes["settings"] = settings

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
