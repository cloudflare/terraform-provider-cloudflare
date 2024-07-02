// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ByoIPPrefixResultEnvelope struct {
	Result ByoIPPrefixModel `json:"result,computed"`
}

type ByoIPPrefixResultDataSourceEnvelope struct {
	Result ByoIPPrefixDataSourceModel `json:"result,computed"`
}

type ByoIPPrefixesResultDataSourceEnvelope struct {
	Result ByoIPPrefixesDataSourceModel `json:"result,computed"`
}

type ByoIPPrefixModel struct {
	ID                   types.String `tfsdk:"id" json:"id,computed"`
	AccountID            types.String `tfsdk:"account_id" path:"account_id"`
	ASN                  types.Int64  `tfsdk:"asn" json:"asn"`
	CIDR                 types.String `tfsdk:"cidr" json:"cidr"`
	LOADocumentID        types.String `tfsdk:"loa_document_id" json:"loa_document_id"`
	Description          types.String `tfsdk:"description" json:"description"`
	Advertised           types.Bool   `tfsdk:"advertised" json:"advertised,computed"`
	AdvertisedModifiedAt types.String `tfsdk:"advertised_modified_at" json:"advertised_modified_at,computed"`
	Approved             types.String `tfsdk:"approved" json:"approved,computed"`
	CreatedAt            types.String `tfsdk:"created_at" json:"created_at,computed"`
	ModifiedAt           types.String `tfsdk:"modified_at" json:"modified_at,computed"`
	OnDemandEnabled      types.Bool   `tfsdk:"on_demand_enabled" json:"on_demand_enabled,computed"`
	OnDemandLocked       types.Bool   `tfsdk:"on_demand_locked" json:"on_demand_locked,computed"`
}

type ByoIPPrefixDataSourceModel struct {
}

type ByoIPPrefixesDataSourceModel struct {
}
