// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressMapResultDataSourceEnvelope struct {
	Result AddressMapDataSourceModel `json:"result,computed"`
}

type AddressMapDataSourceModel struct {
	ID           types.String                                                       `tfsdk:"id" path:"address_map_id,computed"`
	AddressMapID types.String                                                       `tfsdk:"address_map_id" path:"address_map_id,optional"`
	AccountID    types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	CanDelete    types.Bool                                                         `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool                                                         `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DefaultSNI   types.String                                                       `tfsdk:"default_sni" json:"default_sni,computed"`
	Description  types.String                                                       `tfsdk:"description" json:"description,computed"`
	Enabled      types.Bool                                                         `tfsdk:"enabled" json:"enabled,computed"`
	ModifiedAt   timetypes.RFC3339                                                  `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	IPs          customfield.NestedObjectList[AddressMapIPsDataSourceModel]         `tfsdk:"ips" json:"ips,computed"`
	Memberships  customfield.NestedObjectList[AddressMapMembershipsDataSourceModel] `tfsdk:"memberships" json:"memberships,computed"`
}

func (m *AddressMapDataSourceModel) toReadParams(_ context.Context) (params addressing.AddressMapGetParams, diags diag.Diagnostics) {
	params = addressing.AddressMapGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AddressMapIPsDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IP        types.String      `tfsdk:"ip" json:"ip,computed"`
}

type AddressMapMembershipsDataSourceModel struct {
	CanDelete  types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Identifier types.String      `tfsdk:"identifier" json:"identifier,computed"`
	Kind       types.String      `tfsdk:"kind" json:"kind,computed"`
}
