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

type AddressMapsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AddressMapsResultDataSourceModel] `json:"result,computed"`
}

type AddressMapsDataSourceModel struct {
	AccountID types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AddressMapsResultDataSourceModel] `tfsdk:"result"`
}

func (m *AddressMapsDataSourceModel) toListParams(_ context.Context) (params addressing.AddressMapListParams, diags diag.Diagnostics) {
	params = addressing.AddressMapListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AddressMapsResultDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	CanDelete    types.Bool        `tfsdk:"can_delete" json:"can_delete,computed"`
	CanModifyIPs types.Bool        `tfsdk:"can_modify_ips" json:"can_modify_ips,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DefaultSNI   types.String      `tfsdk:"default_sni" json:"default_sni,computed"`
	Description  types.String      `tfsdk:"description" json:"description,computed"`
	Enabled      types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ModifiedAt   timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}
