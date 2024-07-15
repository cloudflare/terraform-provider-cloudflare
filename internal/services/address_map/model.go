// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapResultEnvelope struct {
	Result AddressMapModel `json:"result,computed"`
}

type AddressMapModel struct {
	ID           types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID    types.String                   `tfsdk:"account_id" path:"account_id"`
	Description  types.String                   `tfsdk:"description" json:"description"`
	Enabled      types.Bool                     `tfsdk:"enabled" json:"enabled"`
	IPs          *[]types.String                `tfsdk:"ips" json:"ips"`
	Memberships  *[]*AddressMapMembershipsModel `tfsdk:"memberships" json:"memberships"`
	DefaultSNI   types.String                   `tfsdk:"default_sni" json:"default_sni"`
	CanDelete    types.Bool                     `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool                     `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed"`
	ModifiedAt   timetypes.RFC3339              `tfsdk:"modified_at" json:"modified_at,computed"`
}

type AddressMapMembershipsModel struct {
	CanDelete  types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at"`
	Identifier types.String      `tfsdk:"identifier" json:"identifier"`
	Kind       types.String      `tfsdk:"kind" json:"kind"`
}
