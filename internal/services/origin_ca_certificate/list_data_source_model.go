// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/origin_ca_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificatesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[OriginCACertificatesResultDataSourceModel] `json:"result,computed"`
}

type OriginCACertificatesDataSourceModel struct {
	ZoneID   types.String                                                            `tfsdk:"zone_id" query:"zone_id"`
	MaxItems types.Int64                                                             `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[OriginCACertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *OriginCACertificatesDataSourceModel) toListParams() (params origin_ca_certificates.OriginCACertificateListParams, diags diag.Diagnostics) {
	params = origin_ca_certificates.OriginCACertificateListParams{}

	if !m.ZoneID.IsNull() {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type OriginCACertificatesResultDataSourceModel struct {
	Csr               types.String      `tfsdk:"csr" json:"csr,computed"`
	Hostnames         types.List        `tfsdk:"hostnames" json:"hostnames,computed"`
	RequestType       types.String      `tfsdk:"request_type" json:"request_type,computed"`
	RequestedValidity types.Float64     `tfsdk:"requested_validity" json:"requested_validity,computed"`
	ID                types.String      `tfsdk:"id" json:"id"`
	Certificate       types.String      `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn         timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed"`
}
