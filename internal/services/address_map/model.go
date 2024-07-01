// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapResultEnvelope struct {
	Result AddressMapModel `json:"result,computed"`
}

type AddressMapResultDataSourceEnvelope struct {
	Result AddressMapDataSourceModel `json:"result,computed"`
}

type AddressMapsResultDataSourceEnvelope struct {
	Result AddressMapsDataSourceModel `json:"result,computed"`
}

type AddressMapModel struct {
	ID           types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID    types.String                   `tfsdk:"account_id" path:"account_id"`
	Description  types.String                   `tfsdk:"description" json:"description"`
	Enabled      types.Bool                     `tfsdk:"enabled" json:"enabled"`
	IPs          *[]types.String                `tfsdk:"ips" json:"ips"`
	Memberships  *[]*AddressMapMembershipsModel `tfsdk:"memberships" json:"memberships"`
	CanDelete    types.Bool                     `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool                     `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    types.String                   `tfsdk:"created_at" json:"created_at,computed"`
	DefaultSNI   types.String                   `tfsdk:"default_sni" json:"default_sni,computed"`
	ModifiedAt   types.String                   `tfsdk:"modified_at" json:"modified_at,computed"`
}

type AddressMapMembershipsModel struct {
	CanDelete  types.Bool   `tfsdk:"can_delete" json:"can_delete,computed"`
	CreatedAt  types.String `tfsdk:"created_at" json:"created_at"`
	Identifier types.String `tfsdk:"identifier" json:"identifier"`
	Kind       types.String `tfsdk:"kind" json:"kind"`
}

type AddressMapDataSourceModel struct {
}

type AddressMapsDataSourceModel struct {
}
