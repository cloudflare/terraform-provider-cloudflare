// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapResultDataSourceEnvelope struct {
	Result AddressMapDataSourceModel `json:"result,computed"`
}

type AddressMapResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AddressMapDataSourceModel] `json:"result,computed"`
}

type AddressMapDataSourceModel struct {
	AccountID    types.String                             `tfsdk:"account_id" path:"account_id"`
	AddressMapID types.String                             `tfsdk:"address_map_id" path:"address_map_id"`
	IPs          *[]*AddressMapIPsDataSourceModel         `tfsdk:"ips" json:"ips"`
	Memberships  *[]*AddressMapMembershipsDataSourceModel `tfsdk:"memberships" json:"memberships"`
	CanDelete    types.Bool                               `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool                               `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled      types.Bool                               `tfsdk:"enabled" json:"enabled,computed"`
	ModifiedAt   timetypes.RFC3339                        `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	DefaultSNI   types.String                             `tfsdk:"default_sni" json:"default_sni,computed_optional"`
	Description  types.String                             `tfsdk:"description" json:"description,computed_optional"`
	ID           types.String                             `tfsdk:"id" json:"id,computed_optional"`
	Filter       *AddressMapFindOneByDataSourceModel      `tfsdk:"filter"`
}

func (m *AddressMapDataSourceModel) toReadParams() (params addressing.AddressMapGetParams, diags diag.Diagnostics) {
	params = addressing.AddressMapGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AddressMapDataSourceModel) toListParams() (params addressing.AddressMapListParams, diags diag.Diagnostics) {
	params = addressing.AddressMapListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type AddressMapIPsDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IP        types.String      `tfsdk:"ip" json:"ip,computed_optional"`
}

type AddressMapMembershipsDataSourceModel struct {
	CanDelete  types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Identifier types.String      `tfsdk:"identifier" json:"identifier,computed_optional"`
	Kind       types.String      `tfsdk:"kind" json:"kind,computed_optional"`
}

type AddressMapFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
