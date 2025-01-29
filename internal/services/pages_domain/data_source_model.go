// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainResultDataSourceEnvelope struct {
	Result PagesDomainDataSourceModel `json:"result,computed"`
}

type PagesDomainResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PagesDomainDataSourceModel] `json:"result,computed"`
}

type PagesDomainDataSourceModel struct {
	ID                   types.String                                                         `tfsdk:"id" json:"-,computed"`
	DomainName           types.String                                                         `tfsdk:"domain_name" path:"domain_name,optional"`
	AccountID            types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	ProjectName          types.String                                                         `tfsdk:"project_name" path:"project_name,required"`
	CertificateAuthority types.String                                                         `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CreatedOn            types.String                                                         `tfsdk:"created_on" json:"created_on,computed"`
	DomainID             types.String                                                         `tfsdk:"domain_id" json:"domain_id,computed"`
	Name                 types.String                                                         `tfsdk:"name" json:"name,computed"`
	Status               types.String                                                         `tfsdk:"status" json:"status,computed"`
	ZoneTag              types.String                                                         `tfsdk:"zone_tag" json:"zone_tag,computed"`
	ValidationData       customfield.NestedObject[PagesDomainValidationDataDataSourceModel]   `tfsdk:"validation_data" json:"validation_data,computed"`
	VerificationData     customfield.NestedObject[PagesDomainVerificationDataDataSourceModel] `tfsdk:"verification_data" json:"verification_data,computed"`
}

func (m *PagesDomainDataSourceModel) toReadParams(_ context.Context) (params pages.ProjectDomainGetParams, diags diag.Diagnostics) {
	params = pages.ProjectDomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *PagesDomainDataSourceModel) toListParams(_ context.Context) (params pages.ProjectDomainListParams, diags diag.Diagnostics) {
	params = pages.ProjectDomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type PagesDomainValidationDataDataSourceModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message,computed"`
	Method       types.String `tfsdk:"method" json:"method,computed"`
	Status       types.String `tfsdk:"status" json:"status,computed"`
	TXTName      types.String `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue     types.String `tfsdk:"txt_value" json:"txt_value,computed"`
}

type PagesDomainVerificationDataDataSourceModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message,computed"`
	Status       types.String `tfsdk:"status" json:"status,computed"`
}
