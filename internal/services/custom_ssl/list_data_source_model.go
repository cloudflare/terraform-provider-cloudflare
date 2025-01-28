// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/custom_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomSSLsResultDataSourceModel] `json:"result,computed"`
}

type CustomSSLsDataSourceModel struct {
	ZoneID   types.String                                                  `tfsdk:"zone_id" path:"zone_id,required"`
	Status   types.String                                                  `tfsdk:"status" query:"status,optional"`
	Match    types.String                                                  `tfsdk:"match" query:"match,computed_optional"`
	MaxItems types.Int64                                                   `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[CustomSSLsResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomSSLsDataSourceModel) toListParams(_ context.Context) (params custom_certificates.CustomCertificateListParams, diags diag.Diagnostics) {
	params = custom_certificates.CustomCertificateListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Match.IsNull() {
		params.Match = cloudflare.F(CustomCertificateListParamsMatch(m.Match.ValueString()))
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(CustomCertificateListParamsStatus(m.Status.ValueString()))
	}

	return
}

type CustomSSLsResultDataSourceModel struct {
	ID              types.String                                                       `tfsdk:"id" json:"id,computed"`
	BundleMethod    types.String                                                       `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       timetypes.RFC3339                                                  `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Hosts           customfield.List[types.String]                                     `tfsdk:"hosts" json:"hosts,computed"`
	Issuer          types.String                                                       `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                                                  `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority        types.Float64                                                      `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                                       `tfsdk:"signature" json:"signature,computed"`
	Status          types.String                                                       `tfsdk:"status" json:"status,computed"`
	UploadedOn      timetypes.RFC3339                                                  `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	ZoneID          types.String                                                       `tfsdk:"zone_id" json:"zone_id,computed"`
	GeoRestrictions customfield.NestedObject[CustomSSLsGeoRestrictionsDataSourceModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed"`
	KeylessServer   customfield.NestedObject[CustomSSLsKeylessServerDataSourceModel]   `tfsdk:"keyless_server" json:"keyless_server,computed"`
	Policy          types.String                                                       `tfsdk:"policy" json:"policy,computed"`
}

type CustomSSLsGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
}

type CustomSSLsKeylessServerDataSourceModel struct {
	ID          types.String                                                           `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Enabled     types.Bool                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                                           `tfsdk:"host" json:"host,computed"`
	ModifiedOn  timetypes.RFC3339                                                      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                           `tfsdk:"name" json:"name,computed"`
	Permissions customfield.List[types.String]                                         `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                                          `tfsdk:"port" json:"port,computed"`
	Status      types.String                                                           `tfsdk:"status" json:"status,computed"`
	Tunnel      customfield.NestedObject[CustomSSLsKeylessServerTunnelDataSourceModel] `tfsdk:"tunnel" json:"tunnel,computed"`
}

type CustomSSLsKeylessServerTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}
