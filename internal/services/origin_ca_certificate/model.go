// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultEnvelope struct {
	Result OriginCACertificateModel `json:"result"`
}

type OriginCACertificateModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	Csr               types.String                   `tfsdk:"csr" json:"csr,computed_optional"`
	RequestType       types.String                   `tfsdk:"request_type" json:"request_type,computed_optional"`
	RequestedValidity types.Float64                  `tfsdk:"requested_validity" json:"requested_validity,computed_optional"`
	Hostnames         customfield.List[types.String] `tfsdk:"hostnames" json:"hostnames,computed_optional"`
	Certificate       types.String                   `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn         timetypes.RFC3339              `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
}
