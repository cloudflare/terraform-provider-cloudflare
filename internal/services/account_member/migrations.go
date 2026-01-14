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

var _ resource.ResourceWithUpgradeState = (*AccountMemberResource)(nil)

func (r *AccountMemberResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   priorSchemaV0(ctx),
			StateUpgrader: upgradeAccountMemberStateV0toV1,
		},
	}
}

type V0AccountMemberModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID types.String                   `tfsdk:"account_id" path:"account_id,required"`
	Email     types.String                   `tfsdk:"email" json:"email,required"`
	Status    types.String                   `tfsdk:"status" json:"status,computed_optional"`
	Roles     *[]types.String                `tfsdk:"roles" json:"roles,optional,no_refresh"`
	Policies  []V0AccountMemberPoliciesModel `tfsdk:"policies" json:"policies,computed_optional"`
	User      *V0AccountMemberUserModel      `tfsdk:"user" json:"user,computed"`
}

type V0AccountMemberPoliciesModel struct {
	ID               types.String                                   `tfsdk:"id" json:"id,computed,force_encode,encode_state_for_unknown"`
	Access           types.String                                   `tfsdk:"access" json:"access,required"`
	PermissionGroups []V0AccountMemberPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
	ResourceGroups   []V0AccountMemberPoliciesResourceGroupsModel   `tfsdk:"resource_groups" json:"resource_groups,required"`
}

type V0AccountMemberPoliciesPermissionGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V0AccountMemberPoliciesResourceGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V0AccountMemberUserModel struct {
	Email                          types.String `tfsdk:"email" json:"email,computed"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name,computed"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name,computed"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}

// priorSchemaV0 returns the schema for version 0 - policy IDs and lists instead of sets
func priorSchemaV0(ctx context.Context) *schema.Schema {
	return &schema.Schema{
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
				CustomType:  customfield.NewNestedObjectListType[V0AccountMemberPoliciesModel](ctx),
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

// upgradeAccountMemberStateV0toV1 upgrades the state from version 0 to version 1
// This removes policy IDs and ensures that lists are transformed into sets.
func upgradeAccountMemberStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var oldState V0AccountMemberModel
	diags := req.State.Get(ctx, &oldState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var newPolicies []AccountMemberPoliciesModel
	for _, policy := range oldState.Policies {
		newPermissionGroups := []AccountMemberPoliciesPermissionGroupsModel{}
		for _, permissionGroup := range policy.PermissionGroups {
			newPermissionGroups = append(newPermissionGroups, AccountMemberPoliciesPermissionGroupsModel(permissionGroup))
		}
		permissionsSet, diags := customfield.NewObjectSet(ctx, newPermissionGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		newResourceGroups := []AccountMemberPoliciesResourceGroupsModel{}
		for _, resourceGroup := range policy.ResourceGroups {
			newResourceGroups = append(newResourceGroups, AccountMemberPoliciesResourceGroupsModel(resourceGroup))
		}
		resourceGroupsSet, diags := customfield.NewObjectSet(ctx, newResourceGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		newPolicies = append(newPolicies, AccountMemberPoliciesModel{
			Access:           policy.Access,
			PermissionGroups: permissionsSet,
			ResourceGroups:   resourceGroupsSet,
		})
	}
	policiesSet, diags := customfield.NewObjectSet(ctx, newPolicies)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var newUser AccountMemberUserModel
	if oldState.User != nil {
		newUser = AccountMemberUserModel{
			Email:                          oldState.User.Email,
			ID:                             oldState.User.ID,
			FirstName:                      oldState.User.FirstName,
			LastName:                       oldState.User.LastName,
			TwoFactorAuthenticationEnabled: oldState.User.TwoFactorAuthenticationEnabled,
		}
	} else {
		// older state didn't have a user object, but the only field we really
		// set in it anyway is the email so we can pull it off of the top level
		newUser = AccountMemberUserModel{
			Email: oldState.Email,
		}
	}
	userObject, diags := customfield.NewObject(ctx, &newUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert roles from *[]types.String to customfield.Set[types.String]
	var rolesSlice []types.String
	if oldState.Roles != nil {
		rolesSlice = *oldState.Roles
	}
	rolesSet, diags := customfield.NewSet[types.String](ctx, rolesSlice)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var newState = AccountMemberModel{
		ID:        oldState.ID,
		AccountID: oldState.AccountID,
		Email:     oldState.Email,
		Status:    oldState.Status,
		Roles:     rolesSet,
		Policies:  policiesSet,
		User:      userObject,
	}

	diags = resp.State.Set(ctx, newState)
	resp.Diagnostics.Append(diags...)
}
