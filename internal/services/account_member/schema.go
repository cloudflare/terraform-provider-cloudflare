// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AccountMemberResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Membership identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier tag.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"email": schema.StringAttribute{
				Description:   "The contact email address of the user.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Description: `Available values: "accepted", "pending".`,
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("pending"),
			},
			"roles": schema.ListAttribute{
				Description: "Array of roles associated with this member.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"policies": schema.ListNestedAttribute{
				Description: "Array of policies associated with this member.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[AccountMemberPoliciesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Policy identifier.",
							Computed:    true,
						},
						"access": schema.StringAttribute{
							Description: "Allow or deny operations against the resources.\nAvailable values: \"allow\", \"deny\".",
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
			"user": schema.SingleNestedAttribute{
				Description: "Details of the user associated to the membership.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AccountMemberUserModel](ctx),
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description: "The contact email address of the user.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"first_name": schema.StringAttribute{
						Description: "User's first name",
						Computed:    true,
					},
					"last_name": schema.StringAttribute{
						Description: "User's last name",
						Computed:    true,
					},
					"two_factor_authentication_enabled": schema.BoolAttribute{
						Description: "Indicates whether two-factor authentication is enabled for the user account. Does not apply to API authentication.",
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
				},
			},
		},
	}
}

func (r *AccountMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccountMemberResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
