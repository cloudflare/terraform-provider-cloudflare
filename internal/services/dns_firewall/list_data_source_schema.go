// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSFirewallsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[DNSFirewallsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier.",
							Computed:    true,
						},
						"deprecate_any_requests": schema.BoolAttribute{
							Description: "Whether to refuse to answer queries for the ANY type",
							Computed:    true,
						},
						"dns_firewall_ips": schema.SetAttribute{
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"ecs_fallback": schema.BoolAttribute{
							Description: "Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent",
							Computed:    true,
						},
						"maximum_cache_ttl": schema.Float64Attribute{
							Description: "By default, Cloudflare attempts to cache responses for as long as\nindicated by the TTL received from upstream nameservers. This setting\nsets an upper bound on this duration. For caching purposes, higher TTLs\nwill be decreased to the maximum value defined by this setting.\n\nThis setting does not affect the TTL value in the DNS response\nCloudflare returns to clients. Cloudflare will always forward the TTL\nvalue received from upstream nameservers.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(30, 36000),
							},
						},
						"minimum_cache_ttl": schema.Float64Attribute{
							Description: "By default, Cloudflare attempts to cache responses for as long as\nindicated by the TTL received from upstream nameservers. This setting\nsets a lower bound on this duration. For caching purposes, lower TTLs\nwill be increased to the minimum value defined by this setting.\n\nThis setting does not affect the TTL value in the DNS response\nCloudflare returns to clients. Cloudflare will always forward the TTL\nvalue received from upstream nameservers.\n\nNote that, even with this setting, there is no guarantee that a\nresponse will be cached for at least the specified duration. Cached\nresponses may be removed earlier for capacity or other operational\nreasons.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(30, 36000),
							},
						},
						"modified_on": schema.StringAttribute{
							Description: "Last modification of DNS Firewall cluster",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "DNS Firewall cluster name",
							Computed:    true,
						},
						"negative_cache_ttl": schema.Float64Attribute{
							Description: "This setting controls how long DNS Firewall should cache negative\nresponses (e.g., NXDOMAIN) from the upstream servers.\n\nThis setting does not affect the TTL value in the DNS response\nCloudflare returns to clients. Cloudflare will always forward the TTL\nvalue received from upstream nameservers.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(30, 36000),
							},
						},
						"ratelimit": schema.Float64Attribute{
							Description: "Ratelimit in queries per second per datacenter (applies to DNS queries sent to the upstream nameservers configured on the cluster)",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(100, 1000000000),
							},
						},
						"retries": schema.Float64Attribute{
							Description: "Number of retries for fetching DNS responses from upstream nameservers (not counting the initial attempt)",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 2),
							},
						},
						"upstream_ips": schema.SetAttribute{
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"attack_mitigation": schema.SingleNestedAttribute{
							Description: "Attack mitigation settings",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[DNSFirewallsAttackMitigationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "When enabled, automatically mitigate random-prefix attacks to protect upstream DNS servers",
									Computed:    true,
								},
								"only_when_upstream_unhealthy": schema.BoolAttribute{
									Description: "Only mitigate attacks when upstream servers seem unhealthy",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DNSFirewallsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DNSFirewallsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
