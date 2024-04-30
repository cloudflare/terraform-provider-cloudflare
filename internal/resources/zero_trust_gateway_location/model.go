// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_location

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayLocationResultEnvelope struct {
	Result ZeroTrustGatewayLocationModel `json:"result,computed"`
}

type ZeroTrustGatewayLocationModel struct {
	ID            types.String                              `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                              `tfsdk:"account_id" path:"account_id"`
	Name          types.String                              `tfsdk:"name" json:"name"`
	ClientDefault types.Bool                                `tfsdk:"client_default" json:"client_default"`
	EcsSupport    types.Bool                                `tfsdk:"ecs_support" json:"ecs_support"`
	Networks      *[]*ZeroTrustGatewayLocationNetworksModel `tfsdk:"networks" json:"networks"`
}

type ZeroTrustGatewayLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}
