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

type EmailSecurityImpersonationRegistriesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityImpersonationRegistriesResultDataSourceModel] `json:"result,computed"`
}

type EmailSecurityImpersonationRegistriesDataSourceModel struct {
	AccountID  types.String                                                                            `tfsdk:"account_id" path:"account_id,required"`
	Direction  types.String                                                                            `tfsdk:"direction" query:"direction,optional"`
	Order      types.String                                                                            `tfsdk:"order" query:"order,optional"`
	Provenance types.String                                                                            `tfsdk:"provenance" query:"provenance,optional"`
	Search     types.String                                                                            `tfsdk:"search" query:"search,optional"`
	MaxItems   types.Int64                                                                             `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[EmailSecurityImpersonationRegistriesResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailSecurityImpersonationRegistriesDataSourceModel) toListParams(_ context.Context) (params email_security.SettingImpersonationRegistryListParams, diags diag.Diagnostics) {
	params = email_security.SettingImpersonationRegistryListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingImpersonationRegistryListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingImpersonationRegistryListParamsOrder(m.Order.ValueString()))
	}
	if !m.Provenance.IsNull() {
		params.Provenance = cloudflare.F(email_security.SettingImpersonationRegistryListParamsProvenance(m.Provenance.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type EmailSecurityImpersonationRegistriesResultDataSourceModel struct {
	ID                      types.Int64       `tfsdk:"id" json:"id,computed"`
	CreatedAt               timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsEmailRegex            types.Bool        `tfsdk:"is_email_regex" json:"is_email_regex,computed"`
	LastModified            timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Name                    types.String      `tfsdk:"name" json:"name,computed"`
	Comments                types.String      `tfsdk:"comments" json:"comments,computed"`
	DirectoryID             types.Int64       `tfsdk:"directory_id" json:"directory_id,computed"`
	DirectoryNodeID         types.Int64       `tfsdk:"directory_node_id" json:"directory_node_id,computed"`
	Email                   types.String      `tfsdk:"email" json:"email,computed"`
	ExternalDirectoryNodeID types.String      `tfsdk:"external_directory_node_id" json:"external_directory_node_id,computed"`
	Provenance              types.String      `tfsdk:"provenance" json:"provenance,computed"`
}
