// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewaySettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "account settings.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"activity_log": schema.SingleNestedAttribute{
						Description: "Activity log settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsActivityLogModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable activity logging.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"antivirus": schema.SingleNestedAttribute{
						Description: "Anti-virus settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsAntivirusModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled_download_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on downloads.",
								Computed:    true,
								Optional:    true,
							},
							"enabled_upload_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on uploads.",
								Computed:    true,
								Optional:    true,
							},
							"fail_closed": schema.BoolAttribute{
								Description: "Block requests for files that cannot be scanned.",
								Computed:    true,
								Optional:    true,
							},
							"notification_settings": schema.SingleNestedAttribute{
								Description: "Configure a message to display on the user's device when an antivirus search is performed.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel](ctx),
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Set notification on",
										Computed:    true,
										Optional:    true,
									},
									"msg": schema.StringAttribute{
										Description: "Customize the message shown in the notification.",
										Computed:    true,
										Optional:    true,
									},
									"support_url": schema.StringAttribute{
										Description: "Optional URL to direct users to additional information. If not set, the notification will open a block page.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
						},
					},
					"block_page": schema.SingleNestedAttribute{
						Description: "Block page layout settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsBlockPageModel](ctx),
						Attributes: map[string]schema.Attribute{
							"background_color": schema.StringAttribute{
								Description: "Block page background color in #rrggbb format.",
								Computed:    true,
								Optional:    true,
							},
							"enabled": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Computed:    true,
								Optional:    true,
							},
							"footer_text": schema.StringAttribute{
								Description: "Block page footer text.",
								Computed:    true,
								Optional:    true,
							},
							"header_text": schema.StringAttribute{
								Description: "Block page header text.",
								Computed:    true,
								Optional:    true,
							},
							"logo_path": schema.StringAttribute{
								Description: "Full URL to the logo file.",
								Computed:    true,
								Optional:    true,
							},
							"mailto_address": schema.StringAttribute{
								Description: "Admin email for users to contact.",
								Computed:    true,
								Optional:    true,
							},
							"mailto_subject": schema.StringAttribute{
								Description: "Subject line for emails created from block page.",
								Computed:    true,
								Optional:    true,
							},
							"name": schema.StringAttribute{
								Description: "Block page title.",
								Computed:    true,
								Optional:    true,
							},
							"suppress_footer": schema.BoolAttribute{
								Description: "Suppress detailed info at the bottom of the block page.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"body_scanning": schema.SingleNestedAttribute{
						Description: "DLP body scanning settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsBodyScanningModel](ctx),
						Attributes: map[string]schema.Attribute{
							"inspection_mode": schema.StringAttribute{
								Description: "Set the inspection mode to either `deep` or `shallow`.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"browser_isolation": schema.SingleNestedAttribute{
						Description: "Browser isolation settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsBrowserIsolationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"non_identity_enabled": schema.BoolAttribute{
								Description: "Enable non-identity onramp support for Browser Isolation.",
								Computed:    true,
								Optional:    true,
							},
							"url_browser_isolation_enabled": schema.BoolAttribute{
								Description: "Enable Clientless Browser Isolation.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"certificate": schema.SingleNestedAttribute{
						Description: "Certificate settings for Gateway TLS interception. If not specified, the Cloudflare Root CA will be used.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsCertificateModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "UUID of certificate to be used for interception. Certificate must be active on the edge. A nil UUID will indicate the Cloudflare Root CA should be used.",
								Required:    true,
							},
						},
					},
					"custom_certificate": schema.SingleNestedAttribute{
						Description: "Custom certificate settings for BYO-PKI. (deprecated and replaced by `certificate`)",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsCustomCertificateModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable use of custom certificate authority for signing Gateway traffic.",
								Required:    true,
							},
							"id": schema.StringAttribute{
								Description: "UUID of certificate (ID from MTLS certificate store).",
								Computed:    true,
								Optional:    true,
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
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable matching all variants of user emails (with + or . modifiers) used as criteria in Firewall policies.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"fips": schema.SingleNestedAttribute{
						Description: "FIPS settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsFipsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"tls": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"protocol_detection": schema.SingleNestedAttribute{
						Description: "Protocol Detection settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsProtocolDetectionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable detecting protocol on initial bytes of client traffic.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"tls_decrypt": schema.SingleNestedAttribute{
						Description: "TLS interception settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsTLSDecryptModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable inspecting encrypted HTTP traffic.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustGatewaySettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewaySettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
