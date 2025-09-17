// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
				Description: "Specify account settings.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"activity_log": schema.SingleNestedAttribute{
						Description: "Specify activity log settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to log activity.",
								Optional:    true,
							},
						},
					},
					"antivirus": schema.SingleNestedAttribute{
						Description: "Specify anti-virus settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled_download_phase": schema.BoolAttribute{
								Description: "Specify whether to enable anti-virus scanning on downloads.",
								Computed:    true,
								Optional:    true,
							},
							"enabled_upload_phase": schema.BoolAttribute{
								Description: "Specify whether to enable anti-virus scanning on uploads.",
								Computed:    true,
								Optional:    true,
							},
							"fail_closed": schema.BoolAttribute{
								Description: "Specify whether to block requests for unscannable files.",
								Computed:    true,
								Optional:    true,
							},
							"notification_settings": schema.SingleNestedAttribute{
								Description: "Configure the message the user's device shows during an antivirus scan.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel](ctx),
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Specify whether to enable notifications.",
										Optional:    true,
									},
									"include_context": schema.BoolAttribute{
										Description: "Specify whether to include context information as query parameters.",
										Optional:    true,
									},
									"msg": schema.StringAttribute{
										Description: "Specify the message to show in the notification.",
										Optional:    true,
									},
									"support_url": schema.StringAttribute{
										Description: "Specify a URL that directs users to more information. If unset, the notification opens a block page.",
										Optional:    true,
									},
								},
							},
						},
					},
					"block_page": schema.SingleNestedAttribute{
						Description: "Specify block page layout settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"background_color": schema.StringAttribute{
								Description: "Specify the block page background color in `#rrggbb` format when the mode is customized_block_page.",
								Optional:    true,
							},
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to enable the custom block page.",
								Optional:    true,
							},
							"footer_text": schema.StringAttribute{
								Description: "Specify the block page footer text when the mode is customized_block_page.",
								Optional:    true,
							},
							"header_text": schema.StringAttribute{
								Description: "Specify the block page header text when the mode is customized_block_page.",
								Optional:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Specify whether to append context to target_uri as query parameters. This applies only when the mode is redirect_uri.",
								Optional:    true,
							},
							"logo_path": schema.StringAttribute{
								Description: "Specify the full URL to the logo file when the mode is customized_block_page.",
								Optional:    true,
							},
							"mailto_address": schema.StringAttribute{
								Description: "Specify the admin email for users to contact when the mode is customized_block_page.",
								Optional:    true,
							},
							"mailto_subject": schema.StringAttribute{
								Description: "Specify the subject line for emails created from the block page when the mode is customized_block_page.",
								Optional:    true,
							},
							"mode": schema.StringAttribute{
								Description: "Specify whether to redirect users to a Cloudflare-hosted block page or a customer-provided URI.\nAvailable values: \"\", \"customized_block_page\", \"redirect_uri\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"",
										"customized_block_page",
										"redirect_uri",
									),
								},
							},
							"name": schema.StringAttribute{
								Description: "Specify the block page title when the mode is customized_block_page.",
								Optional:    true,
							},
							"read_only": schema.BoolAttribute{
								Description: "Indicate that this setting was shared via the Orgs API and read only for the current account.",
								Computed:    true,
								Optional:    true,
							},
							"source_account": schema.StringAttribute{
								Description: "Indicate the account tag of the account that shared this setting.",
								Computed:    true,
								Optional:    true,
							},
							"suppress_footer": schema.BoolAttribute{
								Description: "Specify whether to suppress detailed information at the bottom of the block page when the mode is customized_block_page.",
								Optional:    true,
							},
							"target_uri": schema.StringAttribute{
								Description: "Specify the URI to redirect users to when the mode is redirect_uri.",
								Optional:    true,
							},
							"version": schema.Int64Attribute{
								Description: "Indicate the version number of the setting.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"body_scanning": schema.SingleNestedAttribute{
						Description: "Specify the DLP inspection mode.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"inspection_mode": schema.StringAttribute{
								Description: "Specify the inspection mode as either `deep` or `shallow`.\nAvailable values: \"deep\", \"shallow\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("deep", "shallow"),
								},
							},
						},
					},
					"browser_isolation": schema.SingleNestedAttribute{
						Description: "Specify Clientless Browser Isolation settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"non_identity_enabled": schema.BoolAttribute{
								Description: "Specify whether to enable non-identity onramp support for Browser Isolation.",
								Optional:    true,
							},
							"url_browser_isolation_enabled": schema.BoolAttribute{
								Description: "Specify whether to enable Clientless Browser Isolation.",
								Optional:    true,
							},
						},
					},
					"certificate": schema.SingleNestedAttribute{
						Description: "Specify certificate settings for Gateway TLS interception. If unset, the Cloudflare Root CA handles interception.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "Specify the UUID of the certificate used for interception. Ensure the certificate is available at the edge(previously called 'active'). A nil UUID directs Cloudflare to use the Root CA.",
								Required:    true,
							},
						},
					},
					"custom_certificate": schema.SingleNestedAttribute{
						Description:        "Specify custom certificate settings for BYO-PKI. This field is deprecated; use `certificate` instead.",
						Optional:           true,
						DeprecationMessage: "This attribute is deprecated.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to enable a custom certificate authority for signing Gateway traffic.",
								Required:    true,
							},
							"id": schema.StringAttribute{
								Description: "Specify the UUID of the certificate (ID from MTLS certificate store).",
								Optional:    true,
							},
							"binding_status": schema.StringAttribute{
								Description: "Indicate the internal certificate status.",
								Computed:    true,
								Optional:    true,
							},
							"updated_at": schema.StringAttribute{
								Computed:   true,
								Optional:   true,
								CustomType: timetypes.RFC3339Type{},
							},
						},
					},
					"extended_email_matching": schema.SingleNestedAttribute{
						Description: "Specify user emails settings for the firewall policies.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to match all variants of user emails (with + or . modifiers) used as criteria in Firewall policies.",
								Optional:    true,
							},
							"read_only": schema.BoolAttribute{
								Description: "Indicate that this setting was shared via the Orgs API and read only for the current account.",
								Optional:    true,
								Computed:    true,
							},
							"source_account": schema.StringAttribute{
								Description: "Indicate the account tag of the account that shared this setting.",
								Optional:    true,
								Computed:    true,
							},
							"version": schema.Int64Attribute{
								Description: "Indicate the version number of the setting.",
								Optional:    true,
								Computed:    true,
							},
						},
					},
					"fips": schema.SingleNestedAttribute{
						Description: "Specify FIPS settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"tls": schema.BoolAttribute{
								Description: "Enforce cipher suites and TLS versions compliant with FIPS 140-2.",
								Optional:    true,
							},
						},
					},
					"host_selector": schema.SingleNestedAttribute{
						Description: "Enable host selection in egress policies.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to enable filtering via hosts for egress policies.",
								Optional:    true,
							},
						},
					},
					"inspection": schema.SingleNestedAttribute{
						Description: "Define the proxy inspection mode.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description: "Define the proxy inspection mode.   1. static: Gateway applies static inspection to HTTP on TCP(80). With TLS decryption on, Gateway inspects HTTPS traffic on TCP(443) and UDP(443).   2. dynamic: Gateway applies protocol detection to inspect HTTP and HTTPS traffic on any port. TLS decryption must remain on to inspect HTTPS traffic.\nAvailable values: \"static\", \"dynamic\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("static", "dynamic"),
								},
							},
						},
					},
					"protocol_detection": schema.SingleNestedAttribute{
						Description: "Specify whether to detect protocols from the initial bytes of client traffic.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to detect protocols from the initial bytes of client traffic.",
								Optional:    true,
							},
						},
					},
					"sandbox": schema.SingleNestedAttribute{
						Description: "Specify whether to enable the sandbox.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to enable the sandbox.",
								Optional:    true,
							},
							"fallback_action": schema.StringAttribute{
								Description: "Specify the action to take when the system cannot scan the file.\nAvailable values: \"allow\", \"block\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("allow", "block"),
								},
							},
						},
					},
					"tls_decrypt": schema.SingleNestedAttribute{
						Description: "Specify whether to inspect encrypted HTTP traffic.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Specify whether to inspect encrypted HTTP traffic.",
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
