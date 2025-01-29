// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PagesDomainsResultDataSourceModel] `json:"result,computed"`
}

type PagesDomainsDataSourceModel struct {
	AccountID   types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	ProjectName types.String                                                    `tfsdk:"project_name" path:"project_name,required"`
	MaxItems    types.Int64                                                     `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[PagesDomainsResultDataSourceModel] `tfsdk:"result"`
}

func (m *PagesDomainsDataSourceModel) toListParams(_ context.Context) (params pages.ProjectDomainListParams, diags diag.Diagnostics) {
	params = pages.ProjectDomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type PagesDomainsResultDataSourceModel struct {
	ID                   types.String                                                          `tfsdk:"id" json:"id,computed"`
	CertificateAuthority types.String                                                          `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CreatedOn            types.String                                                          `tfsdk:"created_on" json:"created_on,computed"`
	DomainID             types.String                                                          `tfsdk:"domain_id" json:"domain_id,computed"`
	Name                 types.String                                                          `tfsdk:"name" json:"name,computed"`
	Status               types.String                                                          `tfsdk:"status" json:"status,computed"`
	ValidationData       customfield.NestedObject[PagesDomainsValidationDataDataSourceModel]   `tfsdk:"validation_data" json:"validation_data,computed"`
	VerificationData     customfield.NestedObject[PagesDomainsVerificationDataDataSourceModel] `tfsdk:"verification_data" json:"verification_data,computed"`
	ZoneTag              types.String                                                          `tfsdk:"zone_tag" json:"zone_tag,computed"`
}

type PagesDomainsValidationDataDataSourceModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message,computed"`
	Method       types.String `tfsdk:"method" json:"method,computed"`
	Status       types.String `tfsdk:"status" json:"status,computed"`
	TXTName      types.String `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue     types.String `tfsdk:"txt_value" json:"txt_value,computed"`
}

type PagesDomainsVerificationDataDataSourceModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message,computed"`
	Status       types.String `tfsdk:"status" json:"status,computed"`
}
