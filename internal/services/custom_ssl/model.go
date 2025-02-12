// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultEnvelope struct {
	Result CustomSSLModel `json:"result"`
}

type CustomSSLModel struct {
	ID              types.String                                            `tfsdk:"id" json:"id,computed"`
	ZoneID          types.String                                            `tfsdk:"zone_id" path:"zone_id,required"`
	Type            types.String                                            `tfsdk:"type" json:"type,computed_optional"`
	Certificate     types.String                                            `tfsdk:"certificate" json:"certificate,required"`
	PrivateKey      types.String                                            `tfsdk:"private_key" json:"private_key,required"`
	Policy          types.String                                            `tfsdk:"policy" json:"policy,optional"`
	BundleMethod    types.String                                            `tfsdk:"bundle_method" json:"bundle_method,computed_optional"`
	GeoRestrictions customfield.NestedObject[CustomSSLGeoRestrictionsModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed_optional"`
	ExpiresOn       timetypes.RFC3339                                       `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer          types.String                                            `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                                       `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority        types.Float64                                           `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                            `tfsdk:"signature" json:"signature,computed"`
	Status          types.String                                            `tfsdk:"status" json:"status,computed"`
	UploadedOn      timetypes.RFC3339                                       `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	Hosts           customfield.List[types.String]                          `tfsdk:"hosts" json:"hosts,computed"`
	KeylessServer   jsontypes.Normalized                                    `tfsdk:"keyless_server" json:"keyless_server,computed"`
}

func (m CustomSSLModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomSSLModel) MarshalJSONForUpdate(state CustomSSLModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type CustomSSLGeoRestrictionsModel struct {
	Label types.String `tfsdk:"label" json:"label,optional"`
}
