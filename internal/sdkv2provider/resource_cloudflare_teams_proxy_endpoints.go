package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		Description: heredoc.Doc(`
			Provides a Cloudflare Teams Proxy Endpoint resource. Teams Proxy
			Endpoints are used for pointing proxy clients at Cloudflare Secure
			Gateway.
		`),
	}
}

func resourceCloudflareTeamsProxyEndpointRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	endpoint, err := client.TeamsProxyEndpoint(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Proxy Endpoint ID is invalid") {
			tflog.Info(ctx, fmt.Sprintf("Teams Proxy Endpoint %s no longer exists", d.Id()))
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

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	newProxyEndpoint := cloudflare.TeamsProxyEndpoint{
		Name: d.Get("name").(string),
		IPs:  expandInterfaceToStringList(d.Get("ips").(*schema.Set).List()),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Teams Proxy Endpoint from struct: %+v", newProxyEndpoint))

	proxyEndpoint, err := client.CreateTeamsProxyEndpoint(ctx, accountID, newProxyEndpoint)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Teams Proxy Endpoint for account %q: %w", accountID, err))
	}

	d.SetId(proxyEndpoint.ID)
	return resourceCloudflareTeamsProxyEndpointRead(ctx, d, meta)
}

func resourceCloudflareTeamsProxyEndpointUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	updatedProxyEndpoint := cloudflare.TeamsProxyEndpoint{
		ID:   d.Id(),
		Name: d.Get("name").(string),
		IPs:  expandInterfaceToStringList(d.Get("ips").(*schema.Set).List()),
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Teams Proxy Endpoint from struct: %+v", updatedProxyEndpoint))

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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Teams Proxy Endpoint using ID: %s", id))

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

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Teams Proxy Endpoint: id %s for account %s", teamsProxyEndpointID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(teamsProxyEndpointID)

	resourceCloudflareTeamsProxyEndpointRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
