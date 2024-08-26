// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/custom_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomSSLsResultDataSourceModel] `json:"result,computed"`
}

type CustomSSLsDataSourceModel struct {
	ZoneID   types.String                                                  `tfsdk:"zone_id" path:"zone_id"`
	Status   types.String                                                  `tfsdk:"status" query:"status"`
	Match    types.String                                                  `tfsdk:"match" query:"match,computed_optional"`
	MaxItems types.Int64                                                   `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[CustomSSLsResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomSSLsDataSourceModel) toListParams() (params custom_certificates.CustomCertificateListParams, diags diag.Diagnostics) {
	params = custom_certificates.CustomCertificateListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Match.IsNull() {
		params.Match = cloudflare.F(custom_certificates.CustomCertificateListParamsMatch(m.Match.ValueString()))
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(custom_certificates.CustomCertificateListParamsStatus(m.Status.ValueString()))
	}

	return
}

type CustomSSLsResultDataSourceModel struct {
	ID              types.String                              `tfsdk:"id" json:"id,computed"`
	BundleMethod    types.String                              `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       timetypes.RFC3339                         `tfsdk:"expires_on" json:"expires_on,computed"`
	Hosts           types.List                                `tfsdk:"hosts" json:"hosts,computed"`
	Issuer          types.String                              `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                         `tfsdk:"modified_on" json:"modified_on,computed"`
	Priority        types.Float64                             `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                              `tfsdk:"signature" json:"signature,computed"`
	Status          types.String                              `tfsdk:"status" json:"status,computed"`
	UploadedOn      timetypes.RFC3339                         `tfsdk:"uploaded_on" json:"uploaded_on,computed"`
	ZoneID          types.String                              `tfsdk:"zone_id" json:"zone_id,computed"`
	GeoRestrictions *CustomSSLsGeoRestrictionsDataSourceModel `tfsdk:"geo_restrictions" json:"geo_restrictions,computed_optional"`
	KeylessServer   *CustomSSLsKeylessServerDataSourceModel   `tfsdk:"keyless_server" json:"keyless_server,computed_optional"`
	Policy          types.String                              `tfsdk:"policy" json:"policy,computed_optional"`
}

type CustomSSLsGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label,computed_optional"`
}

type CustomSSLsKeylessServerDataSourceModel struct {
	ID          types.String                                  `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                             `tfsdk:"created_on" json:"created_on,computed"`
	Enabled     types.Bool                                    `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                  `tfsdk:"host" json:"host,computed"`
	ModifiedOn  timetypes.RFC3339                             `tfsdk:"modified_on" json:"modified_on,computed"`
	Name        types.String                                  `tfsdk:"name" json:"name,computed"`
	Permissions types.List                                    `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                 `tfsdk:"port" json:"port,computed"`
	Status      types.String                                  `tfsdk:"status" json:"status,computed"`
	Tunnel      *CustomSSLsKeylessServerTunnelDataSourceModel `tfsdk:"tunnel" json:"tunnel,computed_optional"`
}

type CustomSSLsKeylessServerTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}
