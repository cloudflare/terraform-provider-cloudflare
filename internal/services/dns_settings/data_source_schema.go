// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"zone_defaults": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[DNSSettingsZoneDefaultsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"flatten_all_cnames": schema.BoolAttribute{
						Description: "Whether to flatten all CNAME records in the zone. Note that, due to DNS limitations, a CNAME record at the zone apex will always be flattened.",
						Computed:    true,
					},
					"foundation_dns": schema.BoolAttribute{
						Description: "Whether to enable Foundation DNS Advanced Nameservers on the zone.",
						Computed:    true,
					},
					"internal_dns": schema.SingleNestedAttribute{
						Description: "Settings for this internal zone.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[DNSSettingsZoneDefaultsInternalDNSDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"reference_zone_id": schema.StringAttribute{
								Description: "The ID of the zone to fallback to.",
								Computed:    true,
							},
						},
					},
					"multi_provider": schema.BoolAttribute{
						Description: "Whether to enable multi-provider DNS, which causes Cloudflare to activate the zone even when non-Cloudflare NS records exist, and to respect NS records at the zone apex during outbound zone transfers.",
						Computed:    true,
					},
					"nameservers": schema.SingleNestedAttribute{
						Description: "Settings determining the nameservers through which the zone should be available.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[DNSSettingsZoneDefaultsNameserversDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: "Nameserver type\nAvailable values: \"cloudflare.standard\", \"cloudflare.standard.random\", \"custom.account\", \"custom.tenant\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"cloudflare.standard",
										"cloudflare.standard.random",
										"custom.account",
										"custom.tenant",
									),
								},
							},
						},
					},
					"ns_ttl": schema.Float64Attribute{
						Description: "The time to live (TTL) of the zone's nameserver (NS) records.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(30, 86400),
						},
					},
					"secondary_overrides": schema.BoolAttribute{
						Description: "Allows a Secondary DNS zone to use (proxied) override records and CNAME flattening at the zone apex.",
						Computed:    true,
					},
					"soa": schema.SingleNestedAttribute{
						Description: "Components of the zone's SOA record.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[DNSSettingsZoneDefaultsSOADataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"expire": schema.Float64Attribute{
								Description: "Time in seconds of being unable to query the primary server after which secondary servers should stop serving the zone.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.Between(86400, 2419200),
								},
							},
							"min_ttl": schema.Float64Attribute{
								Description: "The time to live (TTL) for negative caching of records within the zone.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.Between(60, 86400),
								},
							},
							"mname": schema.StringAttribute{
								Description: "The primary nameserver, which may be used for outbound zone transfers.",
								Computed:    true,
							},
							"refresh": schema.Float64Attribute{
								Description: "Time in seconds after which secondary servers should re-check the SOA record to see if the zone has been updated.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.Between(600, 86400),
								},
							},
							"retry": schema.Float64Attribute{
								Description: "Time in seconds after which secondary servers should retry queries after the primary server was unresponsive.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.Between(600, 86400),
								},
							},
							"rname": schema.StringAttribute{
								Description: "The email address of the zone administrator, with the first label representing the local part of the email address.",
								Computed:    true,
							},
							"ttl": schema.Float64Attribute{
								Description: "The time to live (TTL) of the SOA record itself.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.Between(300, 86400),
								},
							},
						},
					},
					"zone_mode": schema.StringAttribute{
						Description: "Whether the zone mode is a regular or CDN/DNS only zone.\nAvailable values: \"standard\", \"cdn_only\", \"dns_only\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"standard",
								"cdn_only",
								"dns_only",
							),
						},
					},
				},
			},
		},
	}
}

func (d *DNSSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DNSSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
