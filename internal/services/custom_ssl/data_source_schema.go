// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomSSLDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"custom_certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate from the authority expires.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the certificate was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"policy": schema.StringAttribute{
				Description: "Specify the policy that determines the region where your private key will be held locally. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Any combination of countries, specified by their two letter country code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements) can be chosen, such as 'country: IN', as well as 'region: EU' which refers to the EU region. If there are too few data centers satisfying the policy, it will be rejected.",
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
				Description: "Status of the zone's custom SSL.\nAvailable values: \"active\", \"expired\", \"deleted\", \"pending\", \"initializing\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"expired",
						"deleted",
						"pending",
						"initializing",
					),
				},
			},
			"uploaded_on": schema.StringAttribute{
				Description: "When the certificate was uploaded to Cloudflare.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"hosts": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"geo_restrictions": schema.SingleNestedAttribute{
				Description: "Specify the region where your private key can be held locally for optimal TLS performance. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Options allow distribution to only to U.S. data centers, only to E.U. data centers, or only to highest security data centers. Default distribution is to all Cloudflare datacenters, for optimal performance.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CustomSSLGeoRestrictionsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"label": schema.StringAttribute{
						Description: `Available values: "us", "eu", "highest_security".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"us",
								"eu",
								"highest_security",
							),
						},
					},
				},
			},
			"keyless_server": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[CustomSSLKeylessServerDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Keyless certificate identifier tag.",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "When the Keyless SSL was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
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
						CustomType:  timetypes.RFC3339Type{},
					},
					"name": schema.StringAttribute{
						Description: "The keyless SSL name.",
						Computed:    true,
					},
					"permissions": schema.ListAttribute{
						Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"port": schema.Float64Attribute{
						Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Status of the Keyless SSL.\nAvailable values: \"active\", \"deleted\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("active", "deleted"),
						},
					},
					"tunnel": schema.SingleNestedAttribute{
						Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[CustomSSLKeylessServerTunnelDataSourceModel](ctx),
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
			"bundle_method": schema.StringAttribute{
				Computed:   true,
				CustomType: jsontypes.NormalizedType{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"match": schema.StringAttribute{
						Description: "Whether to match all search requirements or at least one (any).\nAvailable values: \"any\", \"all\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("any", "all"),
						},
					},
					"status": schema.StringAttribute{
						Description: "Status of the zone's custom SSL.\nAvailable values: \"active\", \"expired\", \"deleted\", \"pending\", \"initializing\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"active",
								"expired",
								"deleted",
								"pending",
								"initializing",
							),
						},
					},
				},
			},
		},
	}
}

func (d *CustomSSLDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomSSLDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("custom_certificate_id"), path.MatchRoot("filter")),
	}
}
