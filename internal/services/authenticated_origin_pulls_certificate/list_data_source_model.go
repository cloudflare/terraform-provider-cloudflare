// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificatesResultListDataSourceEnvelope struct {
	Result *[]*AuthenticatedOriginPullsCertificatesItemsDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificatesDataSourceModel struct {
	ZoneID   types.String                                                 `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                                                  `tfsdk:"max_items"`
	Items    *[]*AuthenticatedOriginPullsCertificatesItemsDataSourceModel `tfsdk:"items"`
}

type AuthenticatedOriginPullsCertificatesItemsDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Certificate types.String      `tfsdk:"certificate" json:"certificate"`
	ExpiresOn   timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed"`
	Issuer      types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Signature   types.String      `tfsdk:"signature" json:"signature,computed"`
	Status      types.String      `tfsdk:"status" json:"status"`
	UploadedOn  timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on"`
}
