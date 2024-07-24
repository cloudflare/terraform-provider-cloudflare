// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapResultDataSourceEnvelope struct {
	Result AddressMapDataSourceModel `json:"result,computed"`
}

type AddressMapResultListDataSourceEnvelope struct {
	Result *[]*AddressMapDataSourceModel `json:"result,computed"`
}

type AddressMapDataSourceModel struct {
	AccountID    types.String                             `tfsdk:"account_id" path:"account_id"`
	AddressMapID types.String                             `tfsdk:"address_map_id" path:"address_map_id"`
	ID           types.String                             `tfsdk:"id" json:"id"`
	CanDelete    types.Bool                               `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool                               `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed"`
	DefaultSNI   types.String                             `tfsdk:"default_sni" json:"default_sni"`
	Description  types.String                             `tfsdk:"description" json:"description"`
	Enabled      types.Bool                               `tfsdk:"enabled" json:"enabled,computed"`
	IPs          *[]*AddressMapIPsDataSourceModel         `tfsdk:"ips" json:"ips"`
	Memberships  *[]*AddressMapMembershipsDataSourceModel `tfsdk:"memberships" json:"memberships"`
	ModifiedAt   timetypes.RFC3339                        `tfsdk:"modified_at" json:"modified_at,computed"`
	FindOneBy    *AddressMapFindOneByDataSourceModel      `tfsdk:"find_one_by"`
}

type AddressMapIPsDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	IP        types.String      `tfsdk:"ip" json:"ip"`
}

type AddressMapMembershipsDataSourceModel struct {
	CanDelete  types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Identifier types.String      `tfsdk:"identifier" json:"identifier"`
	Kind       types.String      `tfsdk:"kind" json:"kind"`
}

type AddressMapFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
