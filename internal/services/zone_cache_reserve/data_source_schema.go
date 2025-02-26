// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneCacheReserveDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"editable": schema.BoolAttribute{
				Description: "Whether the setting is editable",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "ID of the zone setting.\navailable values: \"cache_reserve\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("cache_reserve"),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "Last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"value": schema.StringAttribute{
				Description: "The value of the feature\navailable values: \"on\", \"off\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("on", "off"),
				},
			},
		},
	}
}

func (d *ZoneCacheReserveDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneCacheReserveDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
