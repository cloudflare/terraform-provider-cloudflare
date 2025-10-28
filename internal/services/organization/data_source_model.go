// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/organizations"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationResultDataSourceEnvelope struct {
	Result OrganizationDataSourceModel `json:"result,computed"`
}

type OrganizationDataSourceModel struct {
	ID             types.String                                                 `tfsdk:"id" path:"organization_id,computed"`
	OrganizationID types.String                                                 `tfsdk:"organization_id" path:"organization_id,optional"`
	CreateTime     timetypes.RFC3339                                            `tfsdk:"create_time" json:"create_time,computed" format:"date-time"`
	Name           types.String                                                 `tfsdk:"name" json:"name,computed"`
	Meta           customfield.NestedObject[OrganizationMetaDataSourceModel]    `tfsdk:"meta" json:"meta,computed"`
	Parent         customfield.NestedObject[OrganizationParentDataSourceModel]  `tfsdk:"parent" json:"parent,computed"`
	Profile        customfield.NestedObject[OrganizationProfileDataSourceModel] `tfsdk:"profile" json:"profile,computed"`
	Filter         *OrganizationFindOneByDataSourceModel                        `tfsdk:"filter"`
}

func (m *OrganizationDataSourceModel) toListParams(_ context.Context) (params organizations.OrganizationListParams, diags diag.Diagnostics) {
	mFilterID := []string{}
	if m.Filter.ID != nil {
		for _, item := range *m.Filter.ID {
			mFilterID = append(mFilterID, item.ValueString())
		}
	}

	params = organizations.OrganizationListParams{
		ID: cloudflare.F(mFilterID),
	}

	if m.Filter.Containing != nil {
		paramsContaining := organizations.OrganizationListParamsContaining{}
		if !m.Filter.Containing.Account.IsNull() {
			paramsContaining.Account = cloudflare.F(m.Filter.Containing.Account.ValueString())
		}
		if !m.Filter.Containing.Organization.IsNull() {
			paramsContaining.Organization = cloudflare.F(m.Filter.Containing.Organization.ValueString())
		}
		if !m.Filter.Containing.User.IsNull() {
			paramsContaining.User = cloudflare.F(m.Filter.Containing.User.ValueString())
		}
		params.Containing = cloudflare.F(paramsContaining)
	}
	if m.Filter.Name != nil {
		paramsName := organizations.OrganizationListParamsName{}
		if !m.Filter.Name.Contains.IsNull() {
			paramsName.Contains = cloudflare.F(m.Filter.Name.Contains.ValueString())
		}
		if !m.Filter.Name.EndsWith.IsNull() {
			paramsName.EndsWith = cloudflare.F(m.Filter.Name.EndsWith.ValueString())
		}
		if !m.Filter.Name.StartsWith.IsNull() {
			paramsName.StartsWith = cloudflare.F(m.Filter.Name.StartsWith.ValueString())
		}
		params.Name = cloudflare.F(paramsName)
	}
	if !m.Filter.PageSize.IsNull() {
		params.PageSize = cloudflare.F(m.Filter.PageSize.ValueInt64())
	}
	if !m.Filter.PageToken.IsNull() {
		params.PageToken = cloudflare.F(m.Filter.PageToken.ValueString())
	}
	if m.Filter.Parent != nil {
		paramsParent := organizations.OrganizationListParamsParent{}
		if !m.Filter.Parent.ID.IsNull() {
			paramsParent.ID = cloudflare.F(organizations.OrganizationListParamsParentID(m.Filter.Parent.ID.ValueString()))
		}
		params.Parent = cloudflare.F(paramsParent)
	}

	return
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

type OrganizationFindOneByDataSourceModel struct {
	ID         *[]types.String                         `tfsdk:"id" query:"id,computed"`
	Containing *OrganizationsContainingDataSourceModel `tfsdk:"containing" query:"containing,optional"`
	Name       *OrganizationsNameDataSourceModel       `tfsdk:"name" query:"name,optional"`
	PageSize   types.Int64                             `tfsdk:"page_size" query:"page_size,optional"`
	PageToken  types.String                            `tfsdk:"page_token" query:"page_token,optional"`
	Parent     *OrganizationsParentDataSourceModel     `tfsdk:"parent" query:"parent,optional"`
}
