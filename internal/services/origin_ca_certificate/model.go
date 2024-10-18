// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultEnvelope struct {
	Result OriginCACertificateModel `json:"result"`
}

type OriginCACertificateModel struct {
	ID                types.String    `tfsdk:"id" json:"id,computed"`
	Csr               types.String    `tfsdk:"csr" json:"csr,optional"`
	RequestType       types.String    `tfsdk:"request_type" json:"request_type,optional"`
	Hostnames         *[]types.String `tfsdk:"hostnames" json:"hostnames,optional"`
	RequestedValidity types.Float64   `tfsdk:"requested_validity" json:"requested_validity,computed_optional"`
	Certificate       types.String    `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn         types.String    `tfsdk:"expires_on" json:"expires_on,computed"`
}

func (m OriginCACertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OriginCACertificateModel) MarshalJSONForUpdate(state OriginCACertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
