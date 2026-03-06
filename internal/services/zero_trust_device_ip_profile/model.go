// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_ip_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceIPProfileResultEnvelope struct {
	Result ZeroTrustDeviceIPProfileModel `json:"result"`
}

type ZeroTrustDeviceIPProfileModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	Match       types.String `tfsdk:"match" json:"match,required"`
	Name        types.String `tfsdk:"name" json:"name,required"`
	Precedence  types.Int64  `tfsdk:"precedence" json:"precedence,required"`
	SubnetID    types.String `tfsdk:"subnet_id" json:"subnet_id,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	CreatedAt   types.String `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt   types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}

func (m ZeroTrustDeviceIPProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceIPProfileModel) MarshalJSONForUpdate(state ZeroTrustDeviceIPProfileModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
