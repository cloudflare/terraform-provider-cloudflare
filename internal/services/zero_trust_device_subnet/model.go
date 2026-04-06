// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_subnet

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceSubnetResultEnvelope struct {
	Result ZeroTrustDeviceSubnetModel `json:"result"`
}

type ZeroTrustDeviceSubnetModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	AccountID        types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name             types.String      `tfsdk:"name" json:"name,required"`
	Network          types.String      `tfsdk:"network" json:"network,required"`
	Comment          types.String      `tfsdk:"comment" json:"comment,computed_optional"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network" json:"is_default_network,computed_optional"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	SubnetType       types.String      `tfsdk:"subnet_type" json:"subnet_type,computed"`
}

func (m ZeroTrustDeviceSubnetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceSubnetModel) MarshalJSONForUpdate(state ZeroTrustDeviceSubnetModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
