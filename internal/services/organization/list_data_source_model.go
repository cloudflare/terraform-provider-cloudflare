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

type OrganizationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[OrganizationsResultDataSourceModel] `json:"result,computed"`
}

type OrganizationsDataSourceModel struct {
	PageSize   types.Int64                                                      `tfsdk:"page_size" query:"page_size,optional"`
	PageToken  types.String                                                     `tfsdk:"page_token" query:"page_token,optional"`
	ID         *[]types.String                                                  `tfsdk:"id" query:"id,optional"`
	Containing *OrganizationsContainingDataSourceModel                          `tfsdk:"containing" query:"containing,optional"`
	Name       *OrganizationsNameDataSourceModel                                `tfsdk:"name" query:"name,optional"`
	Parent     *OrganizationsParentDataSourceModel                              `tfsdk:"parent" query:"parent,optional"`
	MaxItems   types.Int64                                                      `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[OrganizationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *OrganizationsDataSourceModel) toListParams(_ context.Context) (params organizations.OrganizationListParams, diags diag.Diagnostics) {
	mID := []string{}
	if m.ID != nil {
		for _, item := range *m.ID {
			mID = append(mID, item.ValueString())
		}
	}

	params = organizations.OrganizationListParams{
		ID: cloudflare.F(mID),
	}

	if m.Containing != nil {
		paramsContaining := organizations.OrganizationListParamsContaining{}
		if !m.Containing.Account.IsNull() {
			paramsContaining.Account = cloudflare.F(m.Containing.Account.ValueString())
		}
		if !m.Containing.Organization.IsNull() {
			paramsContaining.Organization = cloudflare.F(m.Containing.Organization.ValueString())
		}
		if !m.Containing.User.IsNull() {
			paramsContaining.User = cloudflare.F(m.Containing.User.ValueString())
		}
		params.Containing = cloudflare.F(paramsContaining)
	}
	if m.Name != nil {
		paramsName := organizations.OrganizationListParamsName{}
		if !m.Name.Contains.IsNull() {
			paramsName.Contains = cloudflare.F(m.Name.Contains.ValueString())
		}
		if !m.Name.EndsWith.IsNull() {
			paramsName.EndsWith = cloudflare.F(m.Name.EndsWith.ValueString())
		}
		if !m.Name.StartsWith.IsNull() {
			paramsName.StartsWith = cloudflare.F(m.Name.StartsWith.ValueString())
		}
		params.Name = cloudflare.F(paramsName)
	}
	if !m.PageSize.IsNull() {
		params.PageSize = cloudflare.F(m.PageSize.ValueInt64())
	}
	if !m.PageToken.IsNull() {
		params.PageToken = cloudflare.F(m.PageToken.ValueString())
	}
	if m.Parent != nil {
		paramsParent := organizations.OrganizationListParamsParent{}
		if !m.Parent.ID.IsNull() {
			paramsParent.ID = cloudflare.F(organizations.OrganizationListParamsParentID(m.Parent.ID.ValueString()))
		}
		params.Parent = cloudflare.F(paramsParent)
	}

	return
}

type OrganizationsContainingDataSourceModel struct {
	Account      types.String `tfsdk:"account" json:"account,optional"`
	Organization types.String `tfsdk:"organization" json:"organization,optional"`
	User         types.String `tfsdk:"user" json:"user,optional"`
}

type OrganizationsNameDataSourceModel struct {
	Contains   types.String `tfsdk:"contains" json:"contains,optional"`
	EndsWith   types.String `tfsdk:"ends_with" json:"endsWith,optional"`
	StartsWith types.String `tfsdk:"starts_with" json:"startsWith,optional"`
}

type OrganizationsParentDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type OrganizationsResultDataSourceModel struct {
	ID         types.String                                                  `tfsdk:"id" json:"id,computed"`
	CreateTime timetypes.RFC3339                                             `tfsdk:"create_time" json:"create_time,computed" format:"date-time"`
	Meta       customfield.NestedObject[OrganizationsMetaDataSourceModel]    `tfsdk:"meta" json:"meta,computed"`
	Name       types.String                                                  `tfsdk:"name" json:"name,computed"`
	Parent     customfield.NestedObject[OrganizationsParentDataSourceModel]  `tfsdk:"parent" json:"parent,computed"`
	Profile    customfield.NestedObject[OrganizationsProfileDataSourceModel] `tfsdk:"profile" json:"profile,computed"`
}

type OrganizationsMetaDataSourceModel struct {
	Flags     customfield.NestedObject[OrganizationsMetaFlagsDataSourceModel] `tfsdk:"flags" json:"flags,computed"`
	ManagedBy types.String                                                    `tfsdk:"managed_by" json:"managed_by,computed"`
}

type OrganizationsMetaFlagsDataSourceModel struct {
	AccountCreation  types.String `tfsdk:"account_creation" json:"account_creation,computed"`
	AccountDeletion  types.String `tfsdk:"account_deletion" json:"account_deletion,computed"`
	AccountMigration types.String `tfsdk:"account_migration" json:"account_migration,computed"`
	AccountMobility  types.String `tfsdk:"account_mobility" json:"account_mobility,computed"`
	SubOrgCreation   types.String `tfsdk:"sub_org_creation" json:"sub_org_creation,computed"`
}

type OrganizationsProfileDataSourceModel struct {
	BusinessAddress  types.String `tfsdk:"business_address" json:"business_address,computed"`
	BusinessEmail    types.String `tfsdk:"business_email" json:"business_email,computed"`
	BusinessName     types.String `tfsdk:"business_name" json:"business_name,computed"`
	BusinessPhone    types.String `tfsdk:"business_phone" json:"business_phone,computed"`
	ExternalMetadata types.String `tfsdk:"external_metadata" json:"external_metadata,computed"`
}
