// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema_validation_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithUpgradeState = &APIShieldSchemaValidationSettingsResource{}

func (r *APIShieldSchemaValidationSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifier",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"zone_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"validation_default_mitigation_action": schema.StringAttribute{
						Description: "The default mitigation action used when there is no mitigation action defined on the operation\n\nMitigation actions are as follows:\n\n  * `log` - log request when request does not conform to schema\n  * `block` - deny access to the site when request does not conform to schema\n\nA special value of of `none` will skip running schema validation entirely for the request when there is no mitigation action defined on the operation\n",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"none",
								"log",
								"block",
							),
						},
					},
					"validation_override_mitigation_action": schema.StringAttribute{
						Description: "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `none` will skip running schema validation entirely for the request\n  - `null` indicates that no override is in place\n\nTo clear any override, use the special value `disable_override` or `null`\n",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("none", "disable_override"),
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state APIShieldSchemaValidationSettingsModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
