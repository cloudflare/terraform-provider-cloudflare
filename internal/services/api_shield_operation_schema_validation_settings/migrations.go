// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_schema_validation_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r APIShieldOperationSchemaValidationSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"operation_id": schema.StringAttribute{
						Required: true,
					},
					"mitigation_action": schema.StringAttribute{
						Description: "When set, this applies a mitigation action to this operation\n\n  - `log` log request when request does not conform to schema for this operation\n  - `block` deny access to the site when request does not conform to schema for this operation\n  - `none` will skip mitigation for this operation\n  - `null` indicates that no operation level mitigation is in place, see Zone Level Schema Validation Settings for mitigation action that will be applied\n",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("log", "block", "none"),
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state APIShieldOperationSchemaValidationSettingsModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
