// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultEnvelope struct {
	Result OriginCACertificateModel `json:"result,computed"`
}

type OriginCACertificateModel struct {
	CertificateID     types.String    `tfsdk:"certificate_id" path:"certificate_id"`
	Csr               types.String    `tfsdk:"csr" json:"csr"`
	Hostnames         *[]types.String `tfsdk:"hostnames" json:"hostnames"`
	RequestType       types.String    `tfsdk:"request_type" json:"request_type"`
	RequestedValidity types.Float64   `tfsdk:"requested_validity" json:"requested_validity"`
}
