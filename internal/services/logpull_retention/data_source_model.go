// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpullRetentionResultDataSourceEnvelope struct {
	Result LogpullRetentionDataSourceModel `json:"result,computed"`
}

type LogpullRetentionDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Flag   types.Bool   `tfsdk:"flag" json:"flag"`
}
