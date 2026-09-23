package zone_dns_settings

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithMoveState = (*ZoneDNSSettingsResource)(nil)

// MoveState migrates a cname_flattening cloudflare_zone_setting to
// cloudflare_zone_dns_settings when Terraform 1.8+ processes a moved block.
func (r *ZoneDNSSettingsResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := zone_setting.ResourceSchema(ctx)

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   moveCNAMEFlatteningState,
		},
	}
}

func moveCNAMEFlatteningState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_zone_setting" {
		return
	}
	if migrations.DiagnoseMoveStateNilSourceState(req, resp) {
		return
	}

	var source zone_setting.ZoneSettingModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &source)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !strings.EqualFold(source.SettingID.ValueString(), zone_setting.CNAMEFlatteningSettingID) {
		resp.Diagnostics.AddError(
			"Only cname_flattening can be moved to DNS settings",
			fmt.Sprintf("The %q zone setting cannot be moved to cloudflare_zone_dns_settings.", source.SettingID.ValueString()),
		)
		return
	}

	flattenAllCNAMEs, diags := flattenAllCNAMEsFromLegacyValue(ctx, source.Value)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	target := ZoneDNSSettingsModel{
		ZoneID:           source.ZoneID,
		FlattenAllCNAMEs: flattenAllCNAMEs,
	}
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, target)...)
}

func flattenAllCNAMEsFromLegacyValue(ctx context.Context, value customfield.NormalizedDynamicValue) (types.Bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	if value.IsNull() || value.IsUnknown() {
		diags.AddError(
			"Unable to migrate CNAME flattening state",
			"The legacy cname_flattening value must be known. Set flatten_all_cnames explicitly and remove the old resource from state.",
		)
		return types.BoolNull(), diags
	}

	dynamicValue, dynamicDiags := value.ToDynamicValue(ctx)
	diags.Append(dynamicDiags...)
	if diags.HasError() {
		return types.BoolNull(), diags
	}

	stringValue, ok := dynamicValue.UnderlyingValue().(types.String)
	if !ok || stringValue.IsNull() || stringValue.IsUnknown() {
		diags.AddError(
			"Unable to migrate CNAME flattening state",
			"The legacy cname_flattening value must be a known string. Set flatten_all_cnames explicitly and remove the old resource from state.",
		)
		return types.BoolNull(), diags
	}

	switch strings.ToLower(stringValue.ValueString()) {
	case "on", "flatten_all":
		return types.BoolValue(true), diags
	case "off", "apex", "flatten_at_root":
		return types.BoolValue(false), diags
	default:
		diags.AddError(
			"Unable to migrate CNAME flattening state",
			fmt.Sprintf("The legacy cname_flattening value %q is not supported. Set flatten_all_cnames explicitly and remove the old resource from state.", stringValue.ValueString()),
		)
		return types.BoolNull(), diags
	}
}
