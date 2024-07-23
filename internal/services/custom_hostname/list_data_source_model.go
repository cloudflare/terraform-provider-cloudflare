// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnamesResultListDataSourceEnvelope struct {
	Result *[]*CustomHostnamesItemsDataSourceModel `json:"result,computed"`
}

type CustomHostnamesDataSourceModel struct {
	ZoneID    types.String                            `tfsdk:"zone_id" path:"zone_id"`
	ID        types.String                            `tfsdk:"id" query:"id"`
	Direction types.String                            `tfsdk:"direction" query:"direction"`
	Hostname  types.String                            `tfsdk:"hostname" query:"hostname"`
	Order     types.String                            `tfsdk:"order" query:"order"`
	Page      types.Float64                           `tfsdk:"page" query:"page"`
	PerPage   types.Float64                           `tfsdk:"per_page" query:"per_page"`
	SSL       types.Float64                           `tfsdk:"ssl" query:"ssl"`
	MaxItems  types.Int64                             `tfsdk:"max_items"`
	Items     *[]*CustomHostnamesItemsDataSourceModel `tfsdk:"items"`
}

type CustomHostnamesItemsDataSourceModel struct {
	ID                        types.String                                                  `tfsdk:"id" json:"id,computed"`
	Hostname                  types.String                                                  `tfsdk:"hostname" json:"hostname,computed"`
	CreatedAt                 types.String                                                  `tfsdk:"created_at" json:"created_at"`
	CustomMetadata            *CustomHostnamesItemsCustomMetadataDataSourceModel            `tfsdk:"custom_metadata" json:"custom_metadata"`
	CustomOriginServer        types.String                                                  `tfsdk:"custom_origin_server" json:"custom_origin_server"`
	CustomOriginSNI           types.String                                                  `tfsdk:"custom_origin_sni" json:"custom_origin_sni"`
	OwnershipVerification     *CustomHostnamesItemsOwnershipVerificationDataSourceModel     `tfsdk:"ownership_verification" json:"ownership_verification"`
	OwnershipVerificationHTTP *CustomHostnamesItemsOwnershipVerificationHTTPDataSourceModel `tfsdk:"ownership_verification_http" json:"ownership_verification_http"`
	Status                    types.String                                                  `tfsdk:"status" json:"status"`
	VerificationErrors        *[]types.String                                               `tfsdk:"verification_errors" json:"verification_errors"`
}

type CustomHostnamesItemsCustomMetadataDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key"`
}

type CustomHostnamesItemsOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type CustomHostnamesItemsOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url"`
}
