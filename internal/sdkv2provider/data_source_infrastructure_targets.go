package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareZeroTrustInfrastructureTargets() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hostname of a target",
			},
			"hostname_contains": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A partial match to the hostname of a target.",
			},
			"ip_v4": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv4 address of the target.",
			},
			"ip_v6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv6 address of the target",
			},
			"virtual_network_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The private virtual network identifier of the target.",
			},
			"created_after": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The date and time at which the target was created.",
			},
			"modified_after": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The date and time at which the target was modified.",
			},
		},
		Description: "Use this datasource to lookup a tunnel in an account.",
		ReadContext: dataSourceCloudflareInfrastructureTargetRead,
	}
}

func dataSourceCloudflareInfrastructureTargetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Reading Targets")
	client := meta.(*cloudflare.API)
	accID := d.Get(consts.AccountIDSchemaKey).(string)

	hostname := d.Get("hostname").(string)
	hostnameContains := d.Get("hostname_contains").(string)
	ipv4 := d.Get("ip_v4").(string)
	ipv6 := d.Get("ip_v6").(string)
	vnetId := d.Get("virtual_network_id").(string)
	createdAfter := d.Get("created_after").(string)
	modifiedAfter := d.Get("created_after").(string)
	checkSetNil := func(s string) *string {
		if s == "" {
			return nil
		} else {
			return &s
		}
	}

	params := cloudflare.TargetListParams{
		CreatedAfter:     *checkSetNil(createdAfter),
		Hostname:         *checkSetNil(hostname),
		HostnameContains: *checkSetNil(hostnameContains),
		IPV4:             *checkSetNil(ipv4),
		IPV6:             *checkSetNil(ipv6),
		ModifedAfter:     *checkSetNil(modifiedAfter),
		VirtualNetworkId: *checkSetNil(vnetId),
	}

	targets, _, err := client.ListInfrastructureTargets(ctx, cloudflare.AccountIdentifier(accID), params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Infrastructure Targets: %w", err))
	}
	if len(targets) == 0 {
		return diag.FromErr(fmt.Errorf("no Infrastructure Targets matching given query parameters"))
	}
	if err = d.Set("targets", targets); err != nil {
		return diag.FromErr(fmt.Errorf("error setting Infrastructure Targets set: %w", err))
	}
	return nil
}
