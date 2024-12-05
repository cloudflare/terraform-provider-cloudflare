package content_scanning

import "github.com/hashicorp/terraform-plugin-framework/types"

type ContentScanningModel struct {
	ZoneID  types.String `tfsdk:"zone_id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}
