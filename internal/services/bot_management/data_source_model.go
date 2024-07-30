// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultDataSourceEnvelope struct {
	Result BotManagementDataSourceModel `json:"result,computed"`
}

type BotManagementDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
