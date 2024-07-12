// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameResultDataSourceEnvelope struct {
	Result CustomHostnameDataSourceModel `json:"result,computed"`
}

type CustomHostnameResultListDataSourceEnvelope struct {
	Result *[]*CustomHostnameDataSourceModel `json:"result,computed"`
}

type CustomHostnameDataSourceModel struct {
	ZoneID                    types.String                                            `tfsdk:"zone_id" path:"zone_id"`
	CustomHostnameID          types.String                                            `tfsdk:"custom_hostname_id" path:"custom_hostname_id"`
	ID                        types.String                                            `tfsdk:"id" json:"id,computed"`
	Hostname                  types.String                                            `tfsdk:"hostname" json:"hostname,computed"`
	CreatedAt                 types.String                                            `tfsdk:"created_at" json:"created_at"`
	CustomMetadata            *CustomHostnameCustomMetadataDataSourceModel            `tfsdk:"custom_metadata" json:"custom_metadata"`
	CustomOriginServer        types.String                                            `tfsdk:"custom_origin_server" json:"custom_origin_server"`
	CustomOriginSNI           types.String                                            `tfsdk:"custom_origin_sni" json:"custom_origin_sni"`
	OwnershipVerification     *CustomHostnameOwnershipVerificationDataSourceModel     `tfsdk:"ownership_verification" json:"ownership_verification"`
	OwnershipVerificationHTTP *CustomHostnameOwnershipVerificationHTTPDataSourceModel `tfsdk:"ownership_verification_http" json:"ownership_verification_http"`
	Status                    types.String                                            `tfsdk:"status" json:"status"`
	VerificationErrors        *[]types.String                                         `tfsdk:"verification_errors" json:"verification_errors"`
	FindOneBy                 *CustomHostnameFindOneByDataSourceModel                 `tfsdk:"find_one_by"`
}

type CustomHostnameCustomMetadataDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key"`
}

type CustomHostnameOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type CustomHostnameOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url"`
}

type CustomHostnameFindOneByDataSourceModel struct {
	ZoneID    types.String  `tfsdk:"zone_id" path:"zone_id"`
	ID        types.String  `tfsdk:"id" query:"id"`
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Hostname  types.String  `tfsdk:"hostname" query:"hostname"`
	Order     types.String  `tfsdk:"order" query:"order"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
	SSL       types.Float64 `tfsdk:"ssl" query:"ssl"`
}
