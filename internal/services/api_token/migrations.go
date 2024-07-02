// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r APITokenResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"token_id": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Description: "Token name.",
						Required:    true,
					},
					"policies": schema.ListNestedAttribute{
						Description: "List of access policies assigned to the token.",
						Required:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Policy identifier.",
									Computed:    true,
								},
								"effect": schema.StringAttribute{
									Description: "Allow or deny operations against the resources.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("allow", "deny"),
									},
								},
								"permission_groups": schema.ListNestedAttribute{
									Description: "A set of permission groups that are specified to the policy.",
									Required:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "Identifier of the group.",
												Computed:    true,
											},
											"meta": schema.StringAttribute{
												Description: "Attributes associated to the permission group.",
												Optional:    true,
											},
											"name": schema.StringAttribute{
												Description: "Name of the group.",
												Computed:    true,
											},
										},
									},
								},
								"resources": schema.StringAttribute{
									Description: "A list of resource names that the policy applies to.",
									Required:    true,
								},
							},
						},
					},
					"condition": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"request_ip": schema.SingleNestedAttribute{
								Description: "Client IP restrictions.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"in": schema.ListAttribute{
										Description: "List of IPv4/IPv6 CIDR addresses.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"not_in": schema.ListAttribute{
										Description: "List of IPv4/IPv6 CIDR addresses.",
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
					},
					"expires_on": schema.StringAttribute{
						Description: "The expiration time on or after which the JWT MUST NOT be accepted for processing.",
						Optional:    true,
					},
					"not_before": schema.StringAttribute{
						Description: "The time before which the token MUST NOT be accepted for processing.",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "Status of the token.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("active", "disabled", "expired"),
						},
					},
					"value": schema.StringAttribute{
						Description: "The token value.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state APITokenModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
