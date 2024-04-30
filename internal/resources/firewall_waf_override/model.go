// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_waf_override

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallWAFOverrideResultEnvelope struct {
	Result FirewallWAFOverrideModel `json:"result,computed"`
}

type FirewallWAFOverrideModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}
