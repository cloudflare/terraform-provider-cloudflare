// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationResultDataSourceEnvelope struct {
	Result OrganizationDataSourceModel `json:"result,computed"`
}

type OrganizationDataSourceModel struct {
	OrganizationID types.String                                                 `tfsdk:"organization_id" path:"organization_id,required"`
	CreateTime     timetypes.RFC3339                                            `tfsdk:"create_time" json:"create_time,computed" format:"date-time"`
	ID             types.String                                                 `tfsdk:"id" json:"id,computed"`
	Name           types.String                                                 `tfsdk:"name" json:"name,computed"`
	Meta           customfield.NestedObject[OrganizationMetaDataSourceModel]    `tfsdk:"meta" json:"meta,computed"`
	Parent         customfield.NestedObject[OrganizationParentDataSourceModel]  `tfsdk:"parent" json:"parent,computed"`
	Profile        customfield.NestedObject[OrganizationProfileDataSourceModel] `tfsdk:"profile" json:"profile,computed"`
}

type OrganizationMetaDataSourceModel struct {
	Flags     customfield.NestedObject[OrganizationMetaFlagsDataSourceModel] `tfsdk:"flags" json:"flags,computed"`
	ManagedBy types.String                                                   `tfsdk:"managed_by" json:"managed_by,computed"`
}

type OrganizationMetaFlagsDataSourceModel struct {
	AccountCreation  types.String `tfsdk:"account_creation" json:"account_creation,computed"`
	AccountDeletion  types.String `tfsdk:"account_deletion" json:"account_deletion,computed"`
	AccountMigration types.String `tfsdk:"account_migration" json:"account_migration,computed"`
	AccountMobility  types.String `tfsdk:"account_mobility" json:"account_mobility,computed"`
	SubOrgCreation   types.String `tfsdk:"sub_org_creation" json:"sub_org_creation,computed"`
}

type OrganizationParentDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type OrganizationProfileDataSourceModel struct {
	BusinessAddress  types.String `tfsdk:"business_address" json:"business_address,computed"`
	BusinessEmail    types.String `tfsdk:"business_email" json:"business_email,computed"`
	BusinessName     types.String `tfsdk:"business_name" json:"business_name,computed"`
	BusinessPhone    types.String `tfsdk:"business_phone" json:"business_phone,computed"`
	ExternalMetadata types.String `tfsdk:"external_metadata" json:"external_metadata,computed"`
}
