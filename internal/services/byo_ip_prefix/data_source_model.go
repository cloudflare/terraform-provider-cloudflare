// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ByoIPPrefixResultDataSourceEnvelope struct {
	Result ByoIPPrefixDataSourceModel `json:"result,computed"`
}

type ByoIPPrefixResultListDataSourceEnvelope struct {
	Result *[]*ByoIPPrefixDataSourceModel `json:"result,computed"`
}

type ByoIPPrefixDataSourceModel struct {
	AccountID            types.String                         `tfsdk:"account_id" path:"account_id"`
	PrefixID             types.String                         `tfsdk:"prefix_id" path:"prefix_id"`
	ID                   types.String                         `tfsdk:"id" json:"id"`
	Advertised           types.Bool                           `tfsdk:"advertised" json:"advertised"`
	AdvertisedModifiedAt types.String                         `tfsdk:"advertised_modified_at" json:"advertised_modified_at"`
	Approved             types.String                         `tfsdk:"approved" json:"approved"`
	ASN                  types.Int64                          `tfsdk:"asn" json:"asn"`
	CIDR                 types.String                         `tfsdk:"cidr" json:"cidr"`
	CreatedAt            types.String                         `tfsdk:"created_at" json:"created_at"`
	Description          types.String                         `tfsdk:"description" json:"description"`
	LOADocumentID        types.String                         `tfsdk:"loa_document_id" json:"loa_document_id"`
	ModifiedAt           types.String                         `tfsdk:"modified_at" json:"modified_at"`
	OnDemandEnabled      types.Bool                           `tfsdk:"on_demand_enabled" json:"on_demand_enabled"`
	OnDemandLocked       types.Bool                           `tfsdk:"on_demand_locked" json:"on_demand_locked"`
	FindOneBy            *ByoIPPrefixFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type ByoIPPrefixFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
