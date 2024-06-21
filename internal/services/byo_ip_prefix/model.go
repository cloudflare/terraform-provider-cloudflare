// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ByoIPPrefixResultEnvelope struct {
	Result ByoIPPrefixModel `json:"result,computed"`
}

type ByoIPPrefixModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	AccountID     types.String `tfsdk:"account_id" path:"account_id"`
	ASN           types.Int64  `tfsdk:"asn" json:"asn"`
	CIDR          types.String `tfsdk:"cidr" json:"cidr"`
	LOADocumentID types.String `tfsdk:"loa_document_id" json:"loa_document_id"`
}
