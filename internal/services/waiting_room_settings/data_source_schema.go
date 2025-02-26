// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WaitingRoomSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"search_engine_crawler_bypass": schema.BoolAttribute{
				Description: "Whether to allow verified search engine crawlers to bypass all waiting rooms on this zone.\nVerified search engine crawlers will not be tracked or counted by the waiting room system,\nand will not appear in waiting room analytics.",
				Computed:    true,
			},
		},
	}
}

func (d *WaitingRoomSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WaitingRoomSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
