// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewaySettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "account settings.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"activity_log": schema.SingleNestedAttribute{
						Description: "Activity log settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsActivityLogDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable activity logging.",
								Computed:    true,
							},
						},
					},
					"antivirus": schema.SingleNestedAttribute{
						Description: "Anti-virus settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsAntivirusDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled_download_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on downloads.",
								Computed:    true,
							},
							"enabled_upload_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on uploads.",
								Computed:    true,
							},
							"fail_closed": schema.BoolAttribute{
								Description: "Block requests for files that cannot be scanned.",
								Computed:    true,
							},
							"notification_settings": schema.SingleNestedAttribute{
								Description: "Configure a message to display on the user's device when an antivirus search is performed.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Set notification on",
										Computed:    true,
									},
									"msg": schema.StringAttribute{
										Description: "Customize the message shown in the notification.",
										Computed:    true,
									},
									"support_url": schema.StringAttribute{
										Description: "Optional URL to direct users to additional information. If not set, the notification will open a block page.",
										Computed:    true,
									},
								},
							},
						},
					},
					"block_page": schema.SingleNestedAttribute{
						Description: "Block page layout settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsBlockPageDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"background_color": schema.StringAttribute{
								Description: "Block page background color in #rrggbb format.",
								Computed:    true,
							},
							"enabled": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Computed:    true,
							},
							"footer_text": schema.StringAttribute{
								Description: "Block page footer text.",
								Computed:    true,
							},
							"header_text": schema.StringAttribute{
								Description: "Block page header text.",
								Computed:    true,
							},
							"logo_path": schema.StringAttribute{
								Description: "Full URL to the logo file.",
								Computed:    true,
							},
							"mailto_address": schema.StringAttribute{
								Description: "Admin email for users to contact.",
								Computed:    true,
							},
							"mailto_subject": schema.StringAttribute{
								Description: "Subject line for emails created from block page.",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "Block page title.",
								Computed:    true,
							},
							"suppress_footer": schema.BoolAttribute{
								Description: "Suppress detailed info at the bottom of the block page.",
								Computed:    true,
							},
						},
					},
					"body_scanning": schema.SingleNestedAttribute{
						Description: "DLP body scanning settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsBodyScanningDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"inspection_mode": schema.StringAttribute{
								Description: "Set the inspection mode to either `deep` or `shallow`.",
								Computed:    true,
							},
						},
					},
					"browser_isolation": schema.SingleNestedAttribute{
						Description: "Browser isolation settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsBrowserIsolationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"non_identity_enabled": schema.BoolAttribute{
								Description: "Enable non-identity onramp support for Browser Isolation.",
								Computed:    true,
							},
							"url_browser_isolation_enabled": schema.BoolAttribute{
								Description: "Enable Clientless Browser Isolation.",
								Computed:    true,
							},
						},
					},
					"certificate": schema.SingleNestedAttribute{
						Description: "Certificate settings for Gateway TLS interception. If not specified, the Cloudflare Root CA will be used.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsCertificateDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "UUID of certificate to be used for interception. Certificate must be active on the edge. A nil UUID will indicate the Cloudflare Root CA should be used.",
								Computed:    true,
							},
						},
					},
					"custom_certificate": schema.SingleNestedAttribute{
						Description: "Custom certificate settings for BYO-PKI. (deprecated and replaced by `certificate`)",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsCustomCertificateDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable use of custom certificate authority for signing Gateway traffic.",
								Computed:    true,
							},
							"id": schema.StringAttribute{
								Description: "UUID of certificate (ID from MTLS certificate store).",
								Computed:    true,
							},
							"binding_status": schema.StringAttribute{
								Description: "Certificate status (internal).",
								Computed:    true,
							},
							"updated_at": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
						},
					},
					"extended_email_matching": schema.SingleNestedAttribute{
						Description: "Extended e-mail matching settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable matching all variants of user emails (with + or . modifiers) used as criteria in Firewall policies.",
								Computed:    true,
							},
						},
					},
					"fips": schema.SingleNestedAttribute{
						Description: "FIPS settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsFipsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"tls": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Computed:    true,
							},
						},
					},
					"protocol_detection": schema.SingleNestedAttribute{
						Description: "Protocol Detection settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsProtocolDetectionDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable detecting protocol on initial bytes of client traffic.",
								Computed:    true,
							},
						},
					},
					"tls_decrypt": schema.SingleNestedAttribute{
						Description: "TLS interception settings.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsTLSDecryptDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable inspecting encrypted HTTP traffic.",
								Computed:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustGatewaySettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustGatewaySettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
