// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_subnet

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceSubnetResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceSubnetDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceSubnetDataSourceModel struct {
	ID               types.String      `tfsdk:"id" path:"subnet_id,computed"`
	SubnetID         types.String      `tfsdk:"subnet_id" path:"subnet_id,required"`
	AccountID        types.String      `tfsdk:"account_id" path:"account_id,required"`
	Comment          types.String      `tfsdk:"comment" json:"comment,computed"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network" json:"is_default_network,computed"`
	Name             types.String      `tfsdk:"name" json:"name,computed"`
	Network          types.String      `tfsdk:"network" json:"network,computed"`
	SubnetType       types.String      `tfsdk:"subnet_type" json:"subnet_type,computed"`
}

func (m *ZeroTrustDeviceSubnetDataSourceModel) toReadParams(_ context.Context) (params zero_trust.NetworkSubnetWARPGetParams, diags diag.Diagnostics) {
	params = zero_trust.NetworkSubnetWARPGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
