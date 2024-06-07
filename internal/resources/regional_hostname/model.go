// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultEnvelope struct {
	Result RegionalHostnameModel `json:"result,computed"`
}

type RegionalHostnameModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	Hostname  types.String `tfsdk:"hostname" json:"hostname"`
	RegionKey types.String `tfsdk:"region_key" json:"region_key"`
}
