// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nel_setting

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*NELSettingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zone Settings Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the zone.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier of the zone.",
				Required:    true,
			},
			"editable": schema.BoolAttribute{
				Description: "Whether the setting is editable. This is false when the zone's plan does not include NEL or the NEL product feature is not enabled.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the setting was last modified. A zero value (0001-01-01T00:00:00Z) indicates the setting has never been explicitly set and is using the default value.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"value": schema.SingleNestedAttribute{
				Description: "The NEL configuration value.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[NELSettingValueDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether Network Error Logging is enabled for the zone. When enabled, browsers report network errors to Cloudflare's NEL endpoint.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *NELSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NELSettingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
