// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

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

func (r AccountMemberResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Membership identifier tag.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
					"email": schema.StringAttribute{
						Description: "The contact email address of the user.",
						Required:    true,
					},
					"roles": schema.StringAttribute{
						Description: "Array of roles associated with this member.",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
						},
						Default: stringdefault.StaticString("pending"),
					},
					"policies": schema.ListNestedAttribute{
						Description: "Array of policies associated with this member.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Policy identifier.",
									Computed:    true,
								},
								"access": schema.StringAttribute{
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
												Required:    true,
											},
										},
									},
								},
								"resource_groups": schema.ListNestedAttribute{
									Description: "A list of resource groups that the policy applies to.",
									Required:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "Identifier of the group.",
												Required:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state AccountMemberModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
