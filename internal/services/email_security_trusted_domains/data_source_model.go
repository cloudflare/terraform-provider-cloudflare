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

type EmailSecurityTrustedDomainsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityTrustedDomainsDataSourceModel] `json:"result,computed"`
}

type EmailSecurityTrustedDomainsDataSourceModel struct {
	Comments     types.String                                         `tfsdk:"comments" json:"comments,computed"`
	CreatedAt    timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID           types.Int64                                          `tfsdk:"id" json:"id,computed"`
	IsRecent     types.Bool                                           `tfsdk:"is_recent" json:"is_recent,computed"`
	IsRegex      types.Bool                                           `tfsdk:"is_regex" json:"is_regex,computed"`
	IsSimilarity types.Bool                                           `tfsdk:"is_similarity" json:"is_similarity,computed"`
	LastModified timetypes.RFC3339                                    `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Pattern      types.String                                         `tfsdk:"pattern" json:"pattern,computed"`
	Filter       *EmailSecurityTrustedDomainsFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *EmailSecurityTrustedDomainsDataSourceModel) toListParams(_ context.Context) (params email_security.SettingTrustedDomainListParams, diags diag.Diagnostics) {
	params = email_security.SettingTrustedDomainListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingTrustedDomainListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.IsRecent.IsNull() {
		params.IsRecent = cloudflare.F(m.Filter.IsRecent.ValueBool())
	}
	if !m.Filter.IsSimilarity.IsNull() {
		params.IsSimilarity = cloudflare.F(m.Filter.IsSimilarity.ValueBool())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingTrustedDomainListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type EmailSecurityTrustedDomainsFindOneByDataSourceModel struct {
	AccountID    types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction    types.String `tfsdk:"direction" query:"direction,optional"`
	IsRecent     types.Bool   `tfsdk:"is_recent" query:"is_recent,optional"`
	IsSimilarity types.Bool   `tfsdk:"is_similarity" query:"is_similarity,optional"`
	Order        types.String `tfsdk:"order" query:"order,optional"`
	Search       types.String `tfsdk:"search" query:"search,optional"`
}
