// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSFirewallDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"dns_firewall_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"deprecate_any_requests": schema.BoolAttribute{
				Description: "Whether to refuse to answer queries for the ANY type",
				Computed:    true,
			},
			"ecs_fallback": schema.BoolAttribute{
				Description: "Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent",
				Computed:    true,
			},
			"maximum_cache_ttl": schema.Float64Attribute{
				Description: "Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes of caching between DNS Firewall and the upstream servers. Higher TTLs will be decreased to the maximum defined here for caching purposes.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(30, 36000),
				},
			},
			"minimum_cache_ttl": schema.Float64Attribute{
				Description: "Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes of caching between DNS Firewall and the upstream servers. Lower TTLs will be increased to the minimum defined here for caching purposes.",
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
				Description: "Negative DNS cache TTL This setting controls how long DNS Firewall should cache negative responses (e.g., NXDOMAIN) from the upstream servers.",
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
			"dns_firewall_ips": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"upstream_ips": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"attack_mitigation": schema.SingleNestedAttribute{
				Description: "Attack mitigation settings",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[DNSFirewallAttackMitigationDataSourceModel](ctx),
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
	}
}

func (d *DNSFirewallDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DNSFirewallDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
