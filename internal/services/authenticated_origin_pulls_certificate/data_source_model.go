// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificateResultDataSourceEnvelope struct {
	Result AuthenticatedOriginPullsCertificateDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificateResultListDataSourceEnvelope struct {
	Result *[]*AuthenticatedOriginPullsCertificateDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificateDataSourceModel struct {
	ZoneID        types.String                                                 `tfsdk:"zone_id" path:"zone_id"`
	CertificateID types.String                                                 `tfsdk:"certificate_id" path:"certificate_id"`
	ID            types.String                                                 `tfsdk:"id" json:"id"`
	Certificate   types.String                                                 `tfsdk:"certificate" json:"certificate"`
	ExpiresOn     timetypes.RFC3339                                            `tfsdk:"expires_on" json:"expires_on"`
	Issuer        types.String                                                 `tfsdk:"issuer" json:"issuer"`
	Signature     types.String                                                 `tfsdk:"signature" json:"signature"`
	Status        types.String                                                 `tfsdk:"status" json:"status"`
	UploadedOn    timetypes.RFC3339                                            `tfsdk:"uploaded_on" json:"uploaded_on"`
	FindOneBy     *AuthenticatedOriginPullsCertificateFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type AuthenticatedOriginPullsCertificateFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
