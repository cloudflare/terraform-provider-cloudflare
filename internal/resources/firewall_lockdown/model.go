// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_lockdown

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallLockdownResultEnvelope struct {
	Result FirewallLockdownModel `json:"result,computed"`
}

type FirewallLockdownModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}
