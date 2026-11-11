// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationResultEnvelope struct {
	Result OrganizationModel `json:"result"`
}

type OrganizationModel struct {
	ID         types.String                                    `tfsdk:"id" json:"id,computed"`
	Parent     *OrganizationParentModel                        `tfsdk:"parent" json:"parent,optional"`
	Profile    *OrganizationProfileModel                       `tfsdk:"profile" json:"profile,optional"`
	Name       types.String                                    `tfsdk:"name" json:"name,required"`
	ParentID   types.String                                    `tfsdk:"parent_id" json:"parent_id,optional,no_refresh"`
	CreateTime timetypes.RFC3339                               `tfsdk:"create_time" json:"create_time,computed" format:"date-time"`
	Meta       customfield.NestedObject[OrganizationMetaModel] `tfsdk:"meta" json:"meta,computed"`
}

func (m OrganizationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OrganizationModel) MarshalJSONForUpdate(state OrganizationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type OrganizationParentModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type OrganizationProfileModel struct {
	BusinessAddress  types.String `tfsdk:"business_address" json:"business_address,optional"`
	BusinessEmail    types.String `tfsdk:"business_email" json:"business_email,optional"`
	BusinessName     types.String `tfsdk:"business_name" json:"business_name,optional"`
	BusinessPhone    types.String `tfsdk:"business_phone" json:"business_phone,optional"`
	ExternalMetadata types.String `tfsdk:"external_metadata" json:"external_metadata,optional"`
}

type OrganizationMetaModel struct {
	AreChildAccountsAccessibleByOrgMembers types.Bool                                                 `tfsdk:"are_child_accounts_accessible_by_org_members" json:"are_child_accounts_accessible_by_org_members,computed"`
	HierarchyTags                          customfield.List[types.String]                             `tfsdk:"hierarchy_tags" json:"hierarchy_tags,computed"`
	ManagedBy                              types.String                                               `tfsdk:"managed_by" json:"managed_by,computed"`
	TenantFlags                            customfield.NestedObject[OrganizationMetaTenantFlagsModel] `tfsdk:"tenant_flags" json:"tenant_flags,computed"`
	TenantType                             types.String                                               `tfsdk:"tenant_type" json:"tenant_type,computed"`
}

type OrganizationMetaTenantFlagsModel struct {
	AccountCreation                      types.String `tfsdk:"account_creation" json:"account_creation,computed"`
	AccountCreationAppliesTenantDefaults types.String `tfsdk:"account_creation_applies_tenant_defaults" json:"account_creation_applies_tenant_defaults,computed"`
	AccountDeletion                      types.String `tfsdk:"account_deletion" json:"account_deletion,computed"`
	AccountMigration                     types.String `tfsdk:"account_migration" json:"account_migration,computed"`
	AccountMobility                      types.String `tfsdk:"account_mobility" json:"account_mobility,computed"`
	EnterpriseCapability                 types.String `tfsdk:"enterprise_capability" json:"enterprise_capability,computed"`
	MemberManagement                     types.String `tfsdk:"member_management" json:"member_management,computed"`
	SubOrgCreation                       types.String `tfsdk:"sub_org_creation" json:"sub_org_creation,computed"`
}
