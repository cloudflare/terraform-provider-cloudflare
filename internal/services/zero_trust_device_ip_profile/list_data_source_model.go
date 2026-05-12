// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_ip_profile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceIPProfilesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDeviceIPProfilesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDeviceIPProfilesDataSourceModel struct {
	AccountID types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	PerPage   types.Int64                                                                  `tfsdk:"per_page" query:"per_page,computed_optional"`
	MaxItems  types.Int64                                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDeviceIPProfilesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDeviceIPProfilesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DeviceIPProfileListParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceIPProfileListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.PerPage.ValueInt64())
	}

	return
}

type ZeroTrustDeviceIPProfilesResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Match       types.String `tfsdk:"match" json:"match,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Precedence  types.Int64  `tfsdk:"precedence" json:"precedence,computed"`
	SubnetID    types.String `tfsdk:"subnet_id" json:"subnet_id,computed"`
	UpdatedAt   types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
