// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultEnvelope struct {
	Result OriginCACertificateModel `json:"result,computed"`
}

type OriginCACertificateModel struct {
	CertificateID     types.String            `tfsdk:"certificate_id" path:"certificate_id"`
	Csr               types.String            `tfsdk:"csr" json:"csr"`
	RequestType       types.String            `tfsdk:"request_type" json:"request_type"`
	Hostnames         *[]jsontypes.Normalized `tfsdk:"hostnames" json:"hostnames"`
	RequestedValidity types.Float64           `tfsdk:"requested_validity" json:"requested_validity"`
	ID                types.String            `tfsdk:"id" json:"id,computed"`
}
