// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*DNSFirewallResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "DNS Firewall cluster name",
				Required:    true,
			},
			"upstream_ips": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"deprecate_any_requests": schema.BoolAttribute{
				Description: "Whether to refuse to answer queries for the ANY type",
				Optional:    true,
			},
			"ecs_fallback": schema.BoolAttribute{
				Description: "Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent",
				Optional:    true,
			},
			"negative_cache_ttl": schema.Float64Attribute{
				Description: "Negative DNS cache TTL This setting controls how long DNS Firewall should cache negative responses (e.g., NXDOMAIN) from the upstream servers.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(30, 36000),
				},
			},
			"ratelimit": schema.Float64Attribute{
				Description: "Ratelimit in queries per second per datacenter (applies to DNS queries sent to the upstream nameservers configured on the cluster)",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(100, 1000000000),
				},
			},
			"maximum_cache_ttl": schema.Float64Attribute{
				Description: "Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes of caching between DNS Firewall and the upstream servers. Higher TTLs will be decreased to the maximum defined here for caching purposes.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(30, 36000),
				},
				Default: float64default.StaticFloat64(900),
			},
			"minimum_cache_ttl": schema.Float64Attribute{
				Description: "Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes of caching between DNS Firewall and the upstream servers. Lower TTLs will be increased to the minimum defined here for caching purposes.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(30, 36000),
				},
				Default: float64default.StaticFloat64(60),
			},
			"retries": schema.Float64Attribute{
				Description: "Number of retries for fetching DNS responses from upstream nameservers (not counting the initial attempt)",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 2),
				},
				Default: float64default.StaticFloat64(2),
			},
			"attack_mitigation": schema.SingleNestedAttribute{
				Description: "Attack mitigation settings",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[DNSFirewallAttackMitigationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "When enabled, automatically mitigate random-prefix attacks to protect upstream DNS servers",
						Optional:    true,
					},
					"only_when_upstream_unhealthy": schema.BoolAttribute{
						Description: "Only mitigate attacks when upstream servers seem unhealthy",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
					},
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "Last modification of DNS Firewall cluster",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dns_firewall_ips": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (r *DNSFirewallResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *DNSFirewallResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
