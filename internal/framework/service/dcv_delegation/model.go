package dcv_delegation

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudflareDCVDelegationModel struct {
	ID       types.String `tfsdk:"id"`
	ZoneID   types.String `tfsdk:"zone_id"`
	Hostname types.String `tfsdk:"hostname"`
}
