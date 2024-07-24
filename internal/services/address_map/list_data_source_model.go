// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapsResultListDataSourceEnvelope struct {
	Result *[]*AddressMapsItemsDataSourceModel `json:"result,computed"`
}

type AddressMapsDataSourceModel struct {
	AccountID types.String                        `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                         `tfsdk:"max_items"`
	Items     *[]*AddressMapsItemsDataSourceModel `tfsdk:"items"`
}

type AddressMapsItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id"`
	CanDelete    types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool        `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	DefaultSNI   types.String      `tfsdk:"default_sni" json:"default_sni"`
	Description  types.String      `tfsdk:"description" json:"description"`
	Enabled      types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ModifiedAt   timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed"`
}
