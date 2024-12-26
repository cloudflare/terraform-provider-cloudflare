package content_scanning_expression

import "github.com/hashicorp/terraform-plugin-framework/types"

type ContentScanningExpressionModel struct {
	ZoneID  types.String `tfsdk:"zone_id"`
	ID      types.String `tfsdk:"id"`
	Payload types.String `tfsdk:"payload"`
}
