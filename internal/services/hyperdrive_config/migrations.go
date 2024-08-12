// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithUpgradeState = &HyperdriveConfigResource{}

func (r *HyperdriveConfigResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"origin": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"database": schema.StringAttribute{
								Description: "The name of your origin database.",
								Required:    true,
							},
							"host": schema.StringAttribute{
								Description: "The host (hostname or IP) of your origin database.",
								Required:    true,
							},
							"scheme": schema.StringAttribute{
								Description: "Specifies the URL scheme used to connect to your origin database.",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"postgres",
										"postgresql",
										"mysql",
									),
								},
								Default: stringdefault.StaticString("postgres"),
							},
							"user": schema.StringAttribute{
								Description: "The user of your origin database.",
								Required:    true,
							},
							"access_client_id": schema.StringAttribute{
								Description: "The Client ID of the Access token to use when connecting to the origin database",
								Optional:    true,
							},
							"port": schema.Int64Attribute{
								Description: "The port (default: 5432 for Postgres) of your origin database.",
								Optional:    true,
							},
						},
					},
					"caching": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "When set to true, disables the caching of SQL responses. (Default: false)",
								Optional:    true,
							},
							"max_age": schema.Int64Attribute{
								Description: "When present, specifies max duration for which items should persist in the cache. (Default: 60)",
								Optional:    true,
							},
							"stale_while_revalidate": schema.Int64Attribute{
								Description: "When present, indicates the number of seconds cache may serve the response after it becomes stale. (Default: 15)",
								Optional:    true,
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state HyperdriveConfigModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
