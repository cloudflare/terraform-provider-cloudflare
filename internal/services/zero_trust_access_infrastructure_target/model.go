// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessInfrastructureTargetResultEnvelope struct {
	Result ZeroTrustAccessInfrastructureTargetModel `json:"result"`
}

type ZeroTrustAccessInfrastructureTargetModel struct {
	ID         types.String                                `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                                `tfsdk:"account_id" path:"account_id,required"`
	Hostname   types.String                                `tfsdk:"hostname" json:"hostname,required"`
	IP         *ZeroTrustAccessInfrastructureTargetIPModel `tfsdk:"ip" json:"ip,required"`
	CreatedAt  timetypes.RFC3339                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339                           `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessInfrastructureTargetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessInfrastructureTargetModel) MarshalJSONForUpdate(state ZeroTrustAccessInfrastructureTargetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessInfrastructureTargetIPModel struct {
	IPV4 *ZeroTrustAccessInfrastructureTargetIPIPV4Model `tfsdk:"ipv4" json:"ipv4,optional"`
	IPV6 *ZeroTrustAccessInfrastructureTargetIPIPV6Model `tfsdk:"ipv6" json:"ipv6,optional"`
}

type ZeroTrustAccessInfrastructureTargetIPIPV4Model struct {
	IPAddr           types.String `tfsdk:"ip_addr" json:"ip_addr,optional"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,optional"`
}

type ZeroTrustAccessInfrastructureTargetIPIPV6Model struct {
	IPAddr           types.String `tfsdk:"ip_addr" json:"ip_addr,optional"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,optional"`
}
