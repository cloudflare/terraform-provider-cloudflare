// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r CustomHostnameResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The custom hostname that will point to your hostname via CNAME.",
				Required:    true,
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
							stringvalidator.OneOfCaseInsensitive("ubiquitous", "optimal", "force"),
						},
						Default: stringdefault.StaticString("ubiquitous"),
					},
					"certificate_authority": schema.StringAttribute{
						Description: "The Certificate Authority that will issue the certificate",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("digicert", "google", "lets_encrypt"),
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
							stringvalidator.OneOfCaseInsensitive("http", "txt", "email"),
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
									stringvalidator.OneOfCaseInsensitive("1.0", "1.1", "1.2", "1.3"),
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
			"custom_metadata": schema.SingleNestedAttribute{
				Description: "These are per-hostname (customer) settings.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Description: "Unique metadata for this hostname.",
						Optional:    true,
					},
				},
			},
		},
	}
}
