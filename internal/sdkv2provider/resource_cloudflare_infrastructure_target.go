package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZeroTrustInfrastructureTarget() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareInfrastructureTargetSchema(),
		CreateContext: resourceCloudflareInfrastructureTargetCreate,
		ReadContext:   resourceCloudflareInfrastructureTargetRead,
		UpdateContext: resourceCloudflareInfrastructureTargetUpdate,
		DeleteContext: resourceCloudflareInfrastructureTargetDelete,
		Description: heredoc.Doc(`
			Provides a Cloudflare Infrastructure Target resource.
		`),
	}
}

func resourceCloudflareInfrastructureTargetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	ipInfo, err := parseIPInfo(d)
	if err != nil {
		return diag.FromErr(err)
	}
	createTargetParams := cloudflare.CreateInfrastructureTargetParams{
		InfrastructureTargetParams: cloudflare.InfrastructureTargetParams{
			Hostname: d.Get("hostname").(string),
			IP:       ipInfo,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Infrastructure Target from struct %+v", createTargetParams))
	target, err := client.CreateInfrastructureTarget(ctx, identifier, createTargetParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Infrastructure Target for account %q: %w", d.Get(consts.AccountIDSchemaKey).(string), err))
	}

	d.SetId(target.ID)
	d.Set("created_at", target.CreatedAt)
	d.Set("modified_at", target.ModifiedAt)

	return resourceCloudflareInfrastructureTargetRead(ctx, d, meta)
}

func resourceCloudflareInfrastructureTargetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Retrieving Cloudflare Infrastructure Target with ID %s", d.Id()))
	target, err := client.GetInfrastructureTarget(ctx, identifier, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Infrastructure Target with ID %s does not exist", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Infrastructure Target with ID %s: %w", d.Id(), err))
	}

	d.Set("hostname", target.Hostname)
	d.Set("ip", target.IP)
	d.Set("created_at", target.CreatedAt)
	d.Set("modified_at", target.ModifiedAt)
	return nil
}

func resourceCloudflareInfrastructureTargetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	ipInfo, err := parseIPInfo(d)
	if err != nil {
		return diag.FromErr(err)
	}
	updatedTargetParams := cloudflare.UpdateInfrastructureTargetParams{
		ID: d.Id(),
		ModifyParams: cloudflare.InfrastructureTargetParams{
			Hostname: d.Get("hostname").(string),
			IP:       ipInfo,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Infrastructure Target from struct: %+v", updatedTargetParams))
	updatedTarget, err := client.UpdateInfrastructureTarget(ctx, identifier, updatedTargetParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Infrastructure Target with ID %s for account %q: %w", d.Id(), accountID, err))
	}

	d.Set("modified_at", updatedTarget.ModifiedAt)
	return resourceCloudflareInfrastructureTargetRead(ctx, d, meta)
}

func resourceCloudflareInfrastructureTargetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Infrastructure Target with ID: %s", d.Id()))
	err = client.DeleteInfrastructureTarget(ctx, identifier, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Infrastructure Target with ID %s for account %q: %w", d.Id(), d.Get(consts.AccountIDSchemaKey).(string), err))
	}

	d.SetId("")
	return nil
}

func parseIPInfo(d *schema.ResourceData) (cloudflare.IPInfo, error) {
	ip := d.Get("ip").(map[string]interface{})
	ip_v4, ip_v4_exists := ip["ipv4"]
	ip_v6, ip_v6_exists := ip["ipv6"]

	if !ip_v4_exists && !ip_v6_exists {
		return cloudflare.IPInfo{}, fmt.Errorf("error creating target resource: one of ipv4 or ipv6 must be configured")
	}

	if ip_v4_exists && ip_v6_exists {
		return cloudflare.IPInfo{
			IPV4: &cloudflare.IPDetails{
				IpAddr:           ip_v4.(map[string]interface{})["ip_addr"].(string),
				VirtualNetworkId: ip_v4.(map[string]interface{})["virtual_network_id"].(string),
			},
			IPV6: &cloudflare.IPDetails{
				IpAddr:           ip_v6.(map[string]interface{})["ip_addr"].(string),
				VirtualNetworkId: ip_v6.(map[string]interface{})["virtual_network_id"].(string),
			},
		}, nil
	} else if ip_v4_exists {
		return cloudflare.IPInfo{
			IPV4: &cloudflare.IPDetails{
				IpAddr:           ip_v4.(map[string]interface{})["ip_addr"].(string),
				VirtualNetworkId: ip_v4.(map[string]interface{})["virtual_network_id"].(string),
			},
		}, nil
	} else {
		return cloudflare.IPInfo{
			IPV6: &cloudflare.IPDetails{
				IpAddr:           ip_v6.(map[string]interface{})["ip_addr"].(string),
				VirtualNetworkId: ip_v6.(map[string]interface{})["virtual_network_id"].(string),
			},
		}, nil
	}
}
