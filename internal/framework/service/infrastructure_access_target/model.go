package infrastructure_access_target

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Resource model
type InfrastructureAccessTargetModel struct {
	AccountID  types.String `tfsdk:"account_id"`
	Hostname   types.String `tfsdk:"hostname"`
	ID         types.String `tfsdk:"id"`
	IP         types.Object `tfsdk:"ip"`
	CreatedAt  types.String `tfsdk:"created_at"`
	ModifiedAt types.String `tfsdk:"modified_at"`
}

type InfrastructureAccessTargetIPInfoModel struct {
	IPV4 types.Object `json:"ipv4,omitempty"`
	IPV6 types.Object `json:"ipv6,omitempty"`
}

type InfrastructureAccessTargetIPDetailsModel struct {
	IPAddr           string `json:"ip_addr"`
	VirtualNetworkId string `json:"virtual_network_id"`
}

// Data source model
type InfrastructureAccessTargetsModel struct {
	// Required
	AccountID types.String `tfsdk:"account_id"`
	// Optional
	Hostname         types.String `tfsdk:"hostname"`
	HostnameContains types.String `tfsdk:"hostname_contains"`
	IPV4             types.String `tfsdk:"ip_v4"`
	IPV6             types.String `tfsdk:"ip_v6"`
	VirtualNetworkId types.String `tfsdk:"virtual_network_id"`
	CreatedAfter     types.String `tfsdk:"created_after"`
	ModifiedAfter    types.String `tfsdk:"modified_after"`
	// Readonly
	Targets []InfrastructureAccessTargetModel `tfsdk:"targets"`
}
