// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &CustomSSLsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &CustomSSLsDataSource{}

func (r CustomSSLsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"match": schema.StringAttribute{
				Description: "Whether to match all search requirements or at least one (any).",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("any", "all"),
				},
			},
			"page": schema.Float64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "Number of zones per page.",
				Computed:    true,
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the zone's custom SSL.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "expired", "deleted", "pending", "initializing"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"bundle_method": schema.StringAttribute{
							Description: "A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("ubiquitous", "optimal", "force"),
							},
						},
						"expires_on": schema.StringAttribute{
							Description: "When the certificate from the authority expires.",
							Computed:    true,
						},
						"hosts": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"issuer": schema.StringAttribute{
							Description: "The certificate authority that issued the certificate.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the certificate was last modified.",
							Computed:    true,
						},
						"priority": schema.Float64Attribute{
							Description: "The order/priority in which the certificate will be used in a request. The higher priority will break ties across overlapping 'legacy_custom' certificates, but 'legacy_custom' certificates will always supercede 'sni_custom' certificates.",
							Computed:    true,
						},
						"signature": schema.StringAttribute{
							Description: "The type of hash used for the certificate.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Status of the zone's custom SSL.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("active", "expired", "deleted", "pending", "initializing"),
							},
						},
						"uploaded_on": schema.StringAttribute{
							Description: "When the certificate was uploaded to Cloudflare.",
							Computed:    true,
						},
						"zone_id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"geo_restrictions": schema.SingleNestedAttribute{
							Description: "Specify the region where your private key can be held locally for optimal TLS performance. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Options allow distribution to only to U.S. data centers, only to E.U. data centers, or only to highest security data centers. Default distribution is to all Cloudflare datacenters, for optimal performance.",
							Computed:    true,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"label": schema.StringAttribute{
									Computed: true,
									Optional: true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("us", "eu", "highest_security"),
									},
								},
							},
						},
						"keyless_server": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Keyless certificate identifier tag.",
									Computed:    true,
								},
								"created_on": schema.StringAttribute{
									Description: "When the Keyless SSL was created.",
									Computed:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether or not the Keyless SSL is on or off.",
									Computed:    true,
								},
								"host": schema.StringAttribute{
									Description: "The keyless SSL name.",
									Computed:    true,
								},
								"modified_on": schema.StringAttribute{
									Description: "When the Keyless SSL was last modified.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The keyless SSL name.",
									Computed:    true,
								},
								"permissions": schema.ListAttribute{
									Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
									Computed:    true,
									ElementType: jsontypes.NewNormalizedNull().Type(ctx),
								},
								"port": schema.Float64Attribute{
									Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
									Computed:    true,
								},
								"status": schema.StringAttribute{
									Description: "Status of the Keyless SSL.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("active", "deleted"),
									},
								},
								"tunnel": schema.SingleNestedAttribute{
									Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
									Computed:    true,
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"private_ip": schema.StringAttribute{
											Description: "Private IP of the Key Server Host",
											Computed:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "Cloudflare Tunnel Virtual Network ID",
											Computed:    true,
										},
									},
								},
							},
						},
						"policy": schema.StringAttribute{
							Description: "Specify the policy that determines the region where your private key will be held locally. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Any combination of countries, specified by their two letter country code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements) can be chosen, such as 'country: IN', as well as 'region: EU' which refers to the EU region. If there are too few data centers satisfying the policy, it will be rejected.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *CustomSSLsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *CustomSSLsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
