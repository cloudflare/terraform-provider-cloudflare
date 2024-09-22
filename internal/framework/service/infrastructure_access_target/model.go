package infrastructure_access_target

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Resource model
type InfrastructureAccessTargetModel struct {
	AccountID  types.String      `tfsdk:"account_id"`
	Hostname   types.String      `tfsdk:"hostname"`
	ID         types.String      `tfsdk:"id"`
	IP         cloudflare.IPInfo `tfsdk:"ip"`
	CreatedAt  types.String      `tfsdk:"hostname"`
	ModifiedAt types.String      `tfsdk:"hostname"`
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
