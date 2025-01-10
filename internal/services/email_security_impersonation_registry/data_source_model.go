// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityImpersonationRegistryResultDataSourceEnvelope struct {
	Result EmailSecurityImpersonationRegistryDataSourceModel `json:"result,computed"`
}

type EmailSecurityImpersonationRegistryResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityImpersonationRegistryDataSourceModel] `json:"result,computed"`
}

type EmailSecurityImpersonationRegistryDataSourceModel struct {
	AccountID               types.String                                                `tfsdk:"account_id" path:"account_id,optional"`
	DisplayNameID           types.Int64                                                 `tfsdk:"display_name_id" path:"display_name_id,optional"`
	Comments                types.String                                                `tfsdk:"comments" json:"comments,computed"`
	CreatedAt               timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DirectoryID             types.Int64                                                 `tfsdk:"directory_id" json:"directory_id,computed"`
	DirectoryNodeID         types.Int64                                                 `tfsdk:"directory_node_id" json:"directory_node_id,computed"`
	Email                   types.String                                                `tfsdk:"email" json:"email,computed"`
	ExternalDirectoryNodeID types.String                                                `tfsdk:"external_directory_node_id" json:"external_directory_node_id,computed"`
	ID                      types.Int64                                                 `tfsdk:"id" json:"id,computed"`
	IsEmailRegex            types.Bool                                                  `tfsdk:"is_email_regex" json:"is_email_regex,computed"`
	LastModified            timetypes.RFC3339                                           `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Name                    types.String                                                `tfsdk:"name" json:"name,computed"`
	Provenance              types.String                                                `tfsdk:"provenance" json:"provenance,computed"`
	Filter                  *EmailSecurityImpersonationRegistryFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *EmailSecurityImpersonationRegistryDataSourceModel) toReadParams(_ context.Context) (params email_security.SettingImpersonationRegistryGetParams, diags diag.Diagnostics) {
	params = email_security.SettingImpersonationRegistryGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *EmailSecurityImpersonationRegistryDataSourceModel) toListParams(_ context.Context) (params email_security.SettingImpersonationRegistryListParams, diags diag.Diagnostics) {
	params = email_security.SettingImpersonationRegistryListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingImpersonationRegistryListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingImpersonationRegistryListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Provenance.IsNull() {
		params.Provenance = cloudflare.F(email_security.SettingImpersonationRegistryListParamsProvenance(m.Filter.Provenance.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type EmailSecurityImpersonationRegistryFindOneByDataSourceModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction  types.String `tfsdk:"direction" query:"direction,optional"`
	Order      types.String `tfsdk:"order" query:"order,optional"`
	Provenance types.String `tfsdk:"provenance" query:"provenance,optional"`
	Search     types.String `tfsdk:"search" query:"search,optional"`
}
