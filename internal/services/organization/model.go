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
	ID         types.String                                      `tfsdk:"id" json:"id,computed"`
	Name       types.String                                      `tfsdk:"name" json:"name,required"`
	Profile    *OrganizationProfileModel                         `tfsdk:"profile" json:"profile,optional"`
	Parent     customfield.NestedObject[OrganizationParentModel] `tfsdk:"parent" json:"parent,computed_optional"`
	CreateTime timetypes.RFC3339                                 `tfsdk:"create_time" json:"create_time,computed" format:"date-time"`
	Meta       customfield.NestedObject[OrganizationMetaModel]   `tfsdk:"meta" json:"meta,computed"`
}

func (m OrganizationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OrganizationModel) MarshalJSONForUpdate(state OrganizationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type OrganizationProfileModel struct {
	BusinessAddress  types.String `tfsdk:"business_address" json:"business_address,required"`
	BusinessEmail    types.String `tfsdk:"business_email" json:"business_email,required"`
	BusinessName     types.String `tfsdk:"business_name" json:"business_name,required"`
	BusinessPhone    types.String `tfsdk:"business_phone" json:"business_phone,required"`
	ExternalMetadata types.String `tfsdk:"external_metadata" json:"external_metadata,required"`
}

type OrganizationParentModel struct {
	ID   types.String `tfsdk:"id" json:"id,required"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type OrganizationMetaModel struct {
	Flags     customfield.NestedObject[OrganizationMetaFlagsModel] `tfsdk:"flags" json:"flags,computed"`
	ManagedBy types.String                                         `tfsdk:"managed_by" json:"managed_by,computed"`
}

type OrganizationMetaFlagsModel struct {
	AccountCreation  types.String `tfsdk:"account_creation" json:"account_creation,computed"`
	AccountDeletion  types.String `tfsdk:"account_deletion" json:"account_deletion,computed"`
	AccountMigration types.String `tfsdk:"account_migration" json:"account_migration,computed"`
	AccountMobility  types.String `tfsdk:"account_mobility" json:"account_mobility,computed"`
	SubOrgCreation   types.String `tfsdk:"sub_org_creation" json:"sub_org_creation,computed"`
}
