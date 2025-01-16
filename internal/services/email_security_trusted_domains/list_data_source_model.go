// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityTrustedDomainsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityTrustedDomainsListResultDataSourceModel] `json:"result,computed"`
}

type EmailSecurityTrustedDomainsListDataSourceModel struct {
	AccountID    types.String                                                                       `tfsdk:"account_id" path:"account_id,required"`
	Direction    types.String                                                                       `tfsdk:"direction" query:"direction,optional"`
	IsRecent     types.Bool                                                                         `tfsdk:"is_recent" query:"is_recent,optional"`
	IsSimilarity types.Bool                                                                         `tfsdk:"is_similarity" query:"is_similarity,optional"`
	Order        types.String                                                                       `tfsdk:"order" query:"order,optional"`
	Search       types.String                                                                       `tfsdk:"search" query:"search,optional"`
	MaxItems     types.Int64                                                                        `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[EmailSecurityTrustedDomainsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailSecurityTrustedDomainsListDataSourceModel) toListParams(_ context.Context) (params email_security.SettingTrustedDomainListParams, diags diag.Diagnostics) {
	params = email_security.SettingTrustedDomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingTrustedDomainListParamsDirection(m.Direction.ValueString()))
	}
	if !m.IsRecent.IsNull() {
		params.IsRecent = cloudflare.F(m.IsRecent.ValueBool())
	}
	if !m.IsSimilarity.IsNull() {
		params.IsSimilarity = cloudflare.F(m.IsSimilarity.ValueBool())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingTrustedDomainListParamsOrder(m.Order.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type EmailSecurityTrustedDomainsListResultDataSourceModel struct {
	ID           types.Int64       `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsRecent     types.Bool        `tfsdk:"is_recent" json:"is_recent,computed"`
	IsRegex      types.Bool        `tfsdk:"is_regex" json:"is_regex,computed"`
	IsSimilarity types.Bool        `tfsdk:"is_similarity" json:"is_similarity,computed"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Pattern      types.String      `tfsdk:"pattern" json:"pattern,computed"`
	Comments     types.String      `tfsdk:"comments" json:"comments,computed"`
}
