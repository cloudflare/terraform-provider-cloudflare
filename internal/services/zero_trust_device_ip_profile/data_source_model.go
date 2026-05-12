// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_ip_profile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceIPProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceIPProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceIPProfileDataSourceModel struct {
	ID          types.String                                      `tfsdk:"id" path:"profile_id,computed"`
	ProfileID   types.String                                      `tfsdk:"profile_id" path:"profile_id,optional"`
	AccountID   types.String                                      `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt   types.String                                      `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String                                      `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                                        `tfsdk:"enabled" json:"enabled,computed"`
	Match       types.String                                      `tfsdk:"match" json:"match,computed"`
	Name        types.String                                      `tfsdk:"name" json:"name,computed"`
	Precedence  types.Int64                                       `tfsdk:"precedence" json:"precedence,computed"`
	SubnetID    types.String                                      `tfsdk:"subnet_id" json:"subnet_id,computed"`
	UpdatedAt   types.String                                      `tfsdk:"updated_at" json:"updated_at,computed"`
	Filter      *ZeroTrustDeviceIPProfileFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustDeviceIPProfileDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DeviceIPProfileGetParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceIPProfileGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDeviceIPProfileDataSourceModel) toListParams(_ context.Context) (params zero_trust.DeviceIPProfileListParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceIPProfileListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.Filter.PerPage.ValueInt64())
	}

	return
}

type ZeroTrustDeviceIPProfileFindOneByDataSourceModel struct {
	PerPage types.Int64 `tfsdk:"per_page" query:"per_page,computed_optional"`
}
