// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_flag

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*FlagshipFlagResource)(nil)

func (r *FlagshipFlagResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	priorSchema := resourceSchemaV500(ctx)

	return map[int64]resource.StateUpgrader{
		500: {
			PriorSchema: &priorSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior flagshipFlagModelV500
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradeFlagshipFlagV500ToV501(prior))...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.AddWarning(
					"Flagship flag state upgraded",
					"The obsolete `flag_key` attribute was removed, please change any references to it to use `key` instead.",
				)
			},
		},
	}
}

// flagshipFlagModelV500 mirrors state written by provider versions 5.20 through
// 5.23. Those versions incorrectly modeled the flag's API key twice: key was
// the JSON field while flag_key was the URL path parameter.
type flagshipFlagModelV500 struct {
	AccountID        types.String               `tfsdk:"account_id"`
	AppID            types.String               `tfsdk:"app_id"`
	FlagKey          types.String               `tfsdk:"flag_key"`
	DefaultVariation types.String               `tfsdk:"default_variation"`
	Enabled          types.Bool                 `tfsdk:"enabled"`
	Key              types.String               `tfsdk:"key"`
	Variations       *map[string]types.String   `tfsdk:"variations"`
	Rules            *[]*FlagshipFlagRulesModel `tfsdk:"rules"`
	Description      types.String               `tfsdk:"description"`
	Type             types.String               `tfsdk:"type"`
	UpdatedAt        types.String               `tfsdk:"updated_at"`
	UpdatedBy        types.String               `tfsdk:"updated_by"`
}

func upgradeFlagshipFlagV500ToV501(prior flagshipFlagModelV500) FlagshipFlagModel {
	key := prior.Key
	if (key.IsNull() || key.IsUnknown() || key.ValueString() == "") &&
		!prior.FlagKey.IsNull() && !prior.FlagKey.IsUnknown() && prior.FlagKey.ValueString() != "" {
		key = prior.FlagKey
	}

	return FlagshipFlagModel{
		AccountID:        prior.AccountID,
		AppID:            prior.AppID,
		DefaultVariation: prior.DefaultVariation,
		Enabled:          prior.Enabled,
		Key:              key,
		Variations:       prior.Variations,
		Rules:            prior.Rules,
		Description:      prior.Description,
		Type:             prior.Type,
		UpdatedAt:        prior.UpdatedAt,
		UpdatedBy:        prior.UpdatedBy,
	}
}

// resourceSchemaV500 reconstructs the pre-fix schema so Terraform can decode
// old state before dropping the redundant flag_key attribute.
func resourceSchemaV500(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)
	s.Version = 500
	s.Attributes["flag_key"] = schema.StringAttribute{
		Description:   "Flag key (slug).",
		Optional:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	}
	return s
}
