// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainResultEnvelope struct {
	Result PagesDomainModel `json:"result"`
}

type PagesDomainModel struct {
	ID                   types.String                                               `tfsdk:"id" json:"id,computed"`
	AccountID            types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	ProjectName          types.String                                               `tfsdk:"project_name" path:"project_name,required"`
	Name                 types.String                                               `tfsdk:"name" json:"name,computed_optional"`
	CertificateAuthority types.String                                               `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CreatedOn            types.String                                               `tfsdk:"created_on" json:"created_on,computed"`
	DomainID             types.String                                               `tfsdk:"domain_id" json:"domain_id,computed"`
	Status               types.String                                               `tfsdk:"status" json:"status,computed"`
	ZoneTag              types.String                                               `tfsdk:"zone_tag" json:"zone_tag,computed"`
	ValidationData       customfield.NestedObject[PagesDomainValidationDataModel]   `tfsdk:"validation_data" json:"validation_data,computed"`
	VerificationData     customfield.NestedObject[PagesDomainVerificationDataModel] `tfsdk:"verification_data" json:"verification_data,computed"`
}

type PagesDomainValidationDataModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message,computed"`
	Method       types.String `tfsdk:"method" json:"method,computed"`
	Status       types.String `tfsdk:"status" json:"status,computed"`
	TXTName      types.String `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue     types.String `tfsdk:"txt_value" json:"txt_value,computed"`
}

type PagesDomainVerificationDataModel struct {
	ErrorMessage types.String `tfsdk:"error_message" json:"error_message,computed"`
	Status       types.String `tfsdk:"status" json:"status,computed"`
}
