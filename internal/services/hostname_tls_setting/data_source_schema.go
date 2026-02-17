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
			"id": schema.StringAttribute{
				Description: "The TLS Setting name. The value type depends on the setting:\n- `ciphers`: value is an array of cipher suite strings (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`)\n- `min_tls_version`: value is a TLS version string (`\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"`)\n- `http2`: value is `\"on\"` or `\"off\"`\nAvailable values: \"ciphers\", \"min_tls_version\", \"http2\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ciphers",
						"min_tls_version",
						"http2",
					),
				},
			},
			"setting_id": schema.StringAttribute{
				Description: "The TLS Setting name. The value type depends on the setting:\n- `ciphers`: value is an array of cipher suite strings (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`)\n- `min_tls_version`: value is a TLS version string (`\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"`)\n- `http2`: value is `\"on\"` or `\"off\"`\nAvailable values: \"ciphers\", \"min_tls_version\", \"http2\".",
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
				Description: "Identifier.",
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
			"value": schema.StringAttribute{
				Description: "The TLS setting value. The type depends on the `setting_id` used in the request path:\n- `ciphers`: an array of allowed cipher suite strings in BoringSSL format (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`)\n- `min_tls_version`: a string indicating the minimum TLS version — one of `\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"` (e.g., `\"1.2\"`)\n- `http2`: a string indicating whether HTTP/2 is enabled — `\"on\"` or `\"off\"` (e.g., `\"on\"`)\nAvailable values: \"1.0\", \"1.1\", \"1.2\", \"1.3\", \"on\", \"off\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"1.0",
						"1.1",
						"1.2",
						"1.3",
						"on",
						"off",
					),
				},
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
