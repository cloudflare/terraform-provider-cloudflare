// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainResultDataSourceEnvelope struct {
	Result PagesDomainDataSourceModel `json:"result,computed"`
}

type PagesDomainResultListDataSourceEnvelope struct {
	Result *[]*PagesDomainDataSourceModel `json:"result,computed"`
}

type PagesDomainDataSourceModel struct {
	AccountID            types.String                                                         `tfsdk:"account_id" path:"account_id"`
	DomainName           types.String                                                         `tfsdk:"domain_name" path:"domain_name"`
	ProjectName          types.String                                                         `tfsdk:"project_name" path:"project_name"`
	CertificateAuthority types.String                                                         `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CreatedOn            types.String                                                         `tfsdk:"created_on" json:"created_on,computed"`
	DomainID             types.String                                                         `tfsdk:"domain_id" json:"domain_id,computed"`
	ID                   types.String                                                         `tfsdk:"id" json:"id,computed"`
	Status               types.String                                                         `tfsdk:"status" json:"status,computed"`
	ZoneTag              types.String                                                         `tfsdk:"zone_tag" json:"zone_tag,computed"`
	ValidationData       customfield.NestedObject[PagesDomainValidationDataDataSourceModel]   `tfsdk:"validation_data" json:"validation_data,computed"`
	VerificationData     customfield.NestedObject[PagesDomainVerificationDataDataSourceModel] `tfsdk:"verification_data" json:"verification_data,computed"`
	Name                 types.String                                                         `tfsdk:"name" json:"name"`
	Filter               *PagesDomainFindOneByDataSourceModel                                 `tfsdk:"filter"`
}

type PagesDomainValidationDataDataSourceModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message"`
	Method       types.String `tfsdk:"method" json:"method"`
	Status       types.String `tfsdk:"status" json:"status"`
	TXTName      types.String `tfsdk:"txt_name" json:"txt_name"`
	TXTValue     types.String `tfsdk:"txt_value" json:"txt_value"`
}

type PagesDomainVerificationDataDataSourceModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message"`
	Status       types.String `tfsdk:"status" json:"status"`
}

type PagesDomainFindOneByDataSourceModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	ProjectName types.String `tfsdk:"project_name" path:"project_name"`
}
