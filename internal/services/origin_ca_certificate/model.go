// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultEnvelope struct {
	Result OriginCACertificateModel `json:"result"`
}

type OriginCACertificateModel struct {
	ID                types.String      `tfsdk:"id" json:"id,computed"`
	Csr               types.String      `tfsdk:"csr" json:"csr"`
	RequestType       types.String      `tfsdk:"request_type" json:"request_type"`
	Hostnames         *[]types.String   `tfsdk:"hostnames" json:"hostnames"`
	RequestedValidity types.Float64     `tfsdk:"requested_validity" json:"requested_validity"`
	Certificate       types.String      `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn         timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed"`
}
