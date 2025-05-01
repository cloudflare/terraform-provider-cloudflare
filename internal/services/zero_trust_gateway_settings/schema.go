// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewaySettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Account settings",
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
								Default:     booldefault.StaticBool(true),
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
								Default:     booldefault.StaticBool(false),
							},
							"enabled_upload_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on uploads.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"fail_closed": schema.BoolAttribute{
								Description: "Block requests for files that cannot be scanned.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
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
										Default:     booldefault.StaticBool(false),
									},
									"include_context": schema.BoolAttribute{
										Description: "If true, context information will be passed as query parameters",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
									"msg": schema.StringAttribute{
										Description: "Customize the message shown in the notification.",
										Optional:    true,
									},
									"support_url": schema.StringAttribute{
										Description: "Optional URL to direct users to additional information. If not set, the notification will open a block page.",
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
								Description: "If mode is customized_block_page: block page background color in #rrggbb format.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"enabled": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"footer_text": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page footer text.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"header_text": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page header text.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"include_context": schema.BoolAttribute{
								Description: "If mode is redirect_uri: when enabled, context will be appended to target_uri as query parameters.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"logo_path": schema.StringAttribute{
								Description: "If mode is customized_block_page: full URL to the logo file.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"mailto_address": schema.StringAttribute{
								Description: "If mode is customized_block_page: admin email for users to contact.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"mailto_subject": schema.StringAttribute{
								Description: "If mode is customized_block_page: subject line for emails created from block page.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"mode": schema.StringAttribute{
								Description: "Controls whether the user is redirected to a Cloudflare-hosted block page or to a customer-provided URI.\nAvailable values: \"customized_block_page\", \"redirect_uri\".",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("customized_block_page", "redirect_uri"),
								},
								Default: stringdefault.StaticString("customized_block_page"),
							},
							"name": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page title.",
								Optional:    true,
							},
							"suppress_footer": schema.BoolAttribute{
								Description: "If mode is customized_block_page: suppress detailed info at the bottom of the block page.",
								Optional:    true,
							},
							"target_uri": schema.StringAttribute{
								Description: "If mode is redirect_uri: URI to which the user should be redirected.",
								Optional:    true,
							},
						},
					},
					"body_scanning": schema.SingleNestedAttribute{
						Description: "DLP body scanning settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"inspection_mode": schema.StringAttribute{
								Description: "Set the inspection mode to either `deep` or `shallow`.",
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
								Default:     booldefault.StaticBool(false),
							},
							"url_browser_isolation_enabled": schema.BoolAttribute{
								Description: "Enable Clientless Browser Isolation.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
					"certificate": schema.SingleNestedAttribute{
						Description: "Certificate settings for Gateway TLS interception. If not specified, the Cloudflare Root CA will be used.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "UUID of certificate to be used for interception. Certificate must be available (previously called 'active') on the edge. A nil UUID will indicate the Cloudflare Root CA should be used.",
								Required:    true,
							},
						},
					},
					"custom_certificate": schema.SingleNestedAttribute{
						Description:        "Custom certificate settings for BYO-PKI. (deprecated and replaced by `certificate`)",
						Computed:           true,
						Optional:           true,
						DeprecationMessage: "This attribute is deprecated.",
						CustomType:         customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsCustomCertificateModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable use of custom certificate authority for signing Gateway traffic.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"id": schema.StringAttribute{
								Description: "UUID of certificate (ID from MTLS certificate store).",
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
								Default:     booldefault.StaticBool(false),
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
								Default:     booldefault.StaticBool(false),
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
								Default:     booldefault.StaticBool(false),
							},
						},
					},
					"sandbox": schema.SingleNestedAttribute{
						Description: "Sandbox settings.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsSandboxModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable sandbox.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"fallback_action": schema.StringAttribute{
								Description: "Action to take when the file cannot be scanned.\nAvailable values: \"allow\", \"block\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("allow", "block"),
								},
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
								Default:     booldefault.StaticBool(true),
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
