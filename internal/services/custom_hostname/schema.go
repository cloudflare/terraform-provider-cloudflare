// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*CustomHostnameResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description:   "The custom hostname that will point to your hostname via CNAME.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ssl": schema.SingleNestedAttribute{
				Description: "SSL properties used when creating the custom hostname.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"bundle_method": schema.StringAttribute{
						Description: "A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ubiquitous",
								"optimal",
								"force",
							),
						},
						Default: stringdefault.StaticString("ubiquitous"),
					},
					"certificate_authority": schema.StringAttribute{
						Description: "The Certificate Authority that will issue the certificate",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"digicert",
								"google",
								"lets_encrypt",
								"ssl_com",
							),
						},
					},
					"custom_certificate": schema.StringAttribute{
						Description: "If a custom uploaded certificate is used.",
						Optional:    true,
					},
					"custom_key": schema.StringAttribute{
						Description: "The key for a custom uploaded certificate.",
						Optional:    true,
					},
					"method": schema.StringAttribute{
						Description: "Domain control validation (DCV) method used for this hostname.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"http",
								"txt",
								"email",
							),
						},
					},
					"settings": schema.SingleNestedAttribute{
						Description: "SSL specific settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ciphers": schema.ListAttribute{
								Description: "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"early_hints": schema.StringAttribute{
								Description: "Whether or not Early Hints is enabled.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("on", "off"),
								},
							},
							"http2": schema.StringAttribute{
								Description: "Whether or not HTTP2 is enabled.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("on", "off"),
								},
							},
							"min_tls_version": schema.StringAttribute{
								Description: "The minimum TLS version supported.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"1.0",
										"1.1",
										"1.2",
										"1.3",
									),
								},
							},
							"tls_1_3": schema.StringAttribute{
								Description: "Whether or not TLS 1.3 is enabled.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("on", "off"),
								},
							},
						},
					},
					"type": schema.StringAttribute{
						Description: "Level of validation to be used for this hostname. Domain validation (dv) must be used.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("dv"),
						},
					},
					"wildcard": schema.BoolAttribute{
						Description: "Indicates whether the certificate covers a wildcard.",
						Optional:    true,
					},
				},
			},
			"custom_origin_server": schema.StringAttribute{
				Description: "a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME record.",
				Optional:    true,
			},
			"custom_origin_sni": schema.StringAttribute{
				Description: "A hostname that will be sent to your custom origin server as SNI for TLS handshake. This can be a valid subdomain of the zone or custom origin server name or the string ':request_host_header:' which will cause the host header in the request to be used as SNI. Not configurable with default/fallback origin server.",
				Optional:    true,
			},
			"custom_metadata": schema.MapAttribute{
				Description: "Unique key/value metadata for this hostname. These are per-hostname (customer) settings.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "This is the time the hostname was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Status of the hostname's activation.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"pending",
						"active_redeploying",
						"moved",
						"pending_deletion",
						"deleted",
						"pending_blocked",
						"pending_migration",
						"pending_provisioned",
						"test_pending",
						"test_active",
						"test_active_apex",
						"test_blocked",
						"test_failed",
						"provisioned",
						"blocked",
					),
				},
			},
			"verification_errors": schema.ListAttribute{
				Description: "These are errors that were encountered while trying to activate a hostname.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"ownership_verification": schema.SingleNestedAttribute{
				Description: "This is a record which can be placed to activate a hostname.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CustomHostnameOwnershipVerificationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "DNS Name for record.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "DNS Record type.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("txt"),
						},
					},
					"value": schema.StringAttribute{
						Description: "Content for the record.",
						Computed:    true,
					},
				},
			},
			"ownership_verification_http": schema.SingleNestedAttribute{
				Description: "This presents the token to be served by the given http url to activate a hostname.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CustomHostnameOwnershipVerificationHTTPModel](ctx),
				Attributes: map[string]schema.Attribute{
					"http_body": schema.StringAttribute{
						Description: "Token to be served.",
						Computed:    true,
					},
					"http_url": schema.StringAttribute{
						Description: "The HTTP URL that will be checked during custom hostname verification and where the customer should host the token.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *CustomHostnameResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CustomHostnameResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
