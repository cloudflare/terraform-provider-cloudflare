package infrastructure_access_target_deprecated

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type InfrastructureAccessTargetDeprecatedModel struct {
	AccountID  types.String `tfsdk:"account_id"`
	Hostname   types.String `tfsdk:"hostname"`
	ID         types.String `tfsdk:"id"`
	IP         types.Object `tfsdk:"ip"`
	CreatedAt  types.String `tfsdk:"created_at"`
	ModifiedAt types.String `tfsdk:"modified_at"`
}

type InfrastructureAccessTargetIPInfoDeprecatedModel struct {
	IPV4 types.Object `tfsdk:"ipv4"`
	IPV6 types.Object `tfsdk:"ipv6"`
}

type InfrastructureAccessTargetIPDetailsDeprecatedModel struct {
	IPAddr           types.String `tfsdk:"ip_addr"`
	VirtualNetworkId types.String `tfsdk:"virtual_network_id"`
}

type InfrastructureAccessTargetsDeprecatedModel struct {
	AccountID        types.String                                `tfsdk:"account_id"`
	Hostname         types.String                                `tfsdk:"hostname"`
	HostnameContains types.String                                `tfsdk:"hostname_contains"`
	IPV4             types.String                                `tfsdk:"ipv4"`
	IPV6             types.String                                `tfsdk:"ipv6"`
	VirtualNetworkId types.String                                `tfsdk:"virtual_network_id"`
	CreatedAfter     types.String                                `tfsdk:"created_after"`
	ModifiedAfter    types.String                                `tfsdk:"modified_after"`
	Targets          []InfrastructureAccessTargetDeprecatedModel `tfsdk:"targets"`
}
