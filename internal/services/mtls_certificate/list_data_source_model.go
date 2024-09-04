// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/mtls_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificatesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MTLSCertificatesResultDataSourceModel] `json:"result,computed"`
}

type MTLSCertificatesDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MTLSCertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *MTLSCertificatesDataSourceModel) toListParams(_ context.Context) (params mtls_certificates.MTLSCertificateListParams, diags diag.Diagnostics) {
	params = mtls_certificates.MTLSCertificateListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MTLSCertificatesResultDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	CA           types.Bool        `tfsdk:"ca" json:"ca,computed"`
	Certificates types.String      `tfsdk:"certificates" json:"certificates,computed"`
	ExpiresOn    timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer       types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Name         types.String      `tfsdk:"name" json:"name,computed"`
	SerialNumber types.String      `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature    types.String      `tfsdk:"signature" json:"signature,computed"`
	UploadedOn   timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}
