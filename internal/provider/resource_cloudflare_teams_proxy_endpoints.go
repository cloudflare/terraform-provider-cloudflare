package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsProxyEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsProxyEndpointSchema(),
		CreateContext: resourceCloudflareTeamsProxyEndpointCreate,
		ReadContext:   resourceCloudflareTeamsProxyEndpointRead,
		UpdateContext: resourceCloudflareTeamsProxyEndpointUpdate,
		DeleteContext: resourceCloudflareTeamsProxyEndpointDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsProxyEndpointImport,
		},
	}
}

func resourceCloudflareTeamsProxyEndpointRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	endpoint, err := client.TeamsProxyEndpoint(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 400") {
			log.Printf("[INFO] Teams Proxy Endpoint %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Teams Proxy Endpoint %q: %w", d.Id(), err))
	}

	if err := d.Set("name", endpoint.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Proxy Endpoint name"))
	}

	if err := d.Set("ips", endpoint.IPs); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Proxy Endpoint IPs"))
	}

	if err := d.Set("subdomain", endpoint.Subdomain); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Proxy Endpoint subdomain"))
	}

	return nil
}

func resourceCloudflareTeamsProxyEndpointCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get("account_id").(string)
	newProxyEndpoint := cloudflare.TeamsProxyEndpoint{
		Name: d.Get("name").(string),
		IPs:  expandInterfaceToStringList(d.Get("ips").(*schema.Set).List()),
	}

	log.Printf("[DEBUG] Creating Cloudflare Teams Proxy Endpoint from struct: %+v", newProxyEndpoint)

	proxyEndpoint, err := client.CreateTeamsProxyEndpoint(ctx, accountID, newProxyEndpoint)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Teams Proxy Endpoint for account %q: %w", accountID, err))
	}

	d.SetId(proxyEndpoint.ID)
	return resourceCloudflareTeamsProxyEndpointRead(ctx, d, meta)
}

func resourceCloudflareTeamsProxyEndpointUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	updatedProxyEndpoint := cloudflare.TeamsProxyEndpoint{
		ID:   d.Id(),
		Name: d.Get("name").(string),
		IPs:  expandInterfaceToStringList(d.Get("ips").(*schema.Set).List()),
	}

	log.Printf("[DEBUG] Updating Cloudflare Teams Proxy Endpoint from struct: %+v", updatedProxyEndpoint)

	teamsProxyEndpoint, err := client.UpdateTeamsProxyEndpoint(ctx, accountID, updatedProxyEndpoint)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Teams Proxy Endpoint for account %q: %w", accountID, err))
	}

	if teamsProxyEndpoint.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Teams Proxy Endpoint ID in update response; resource was empty"))
	}
	return resourceCloudflareTeamsProxyEndpointRead(ctx, d, meta)
}

func resourceCloudflareTeamsProxyEndpointDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	id := d.Id()
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Teams Proxy Endpoint using ID: %s", id)

	err := client.DeleteTeamsProxyEndpoint(ctx, accountID, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Teams Proxy Endpoint for account %q: %w", accountID, err))
	}

	return resourceCloudflareTeamsProxyEndpointRead(ctx, d, meta)
}

func resourceCloudflareTeamsProxyEndpointImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsProxyEndpointID\"", d.Id())
	}

	accountID, teamsProxyEndpointID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Teams Proxy Endpoint: id %s for account %s", teamsProxyEndpointID, accountID)

	d.Set("account_id", accountID)
	d.SetId(teamsProxyEndpointID)

	resourceCloudflareTeamsProxyEndpointRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
