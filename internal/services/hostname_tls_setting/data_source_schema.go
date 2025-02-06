// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*HostnameTLSSettingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"setting_id": schema.StringAttribute{
				Description: "The TLS Setting name.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ciphers",
						"min_tls_version",
						"http2",
					),
				},
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "This is the time the tls setting was originally created for this hostname.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname for which the tls settings are set.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Deployment status for the given tls setting.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "This is the time the tls setting was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"value": schema.Float64Attribute{
				Description: "The tls setting value.",
				Computed:    true,
			},
		},
	}
}

func (d *HostnameTLSSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *HostnameTLSSettingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
