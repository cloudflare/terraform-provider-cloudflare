package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsProxyEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareTeamsProxyEndpointSchema(),
		Create: resourceCloudflareTeamsProxyEndpointCreate,
		Read:   resourceCloudflareTeamsProxyEndpointRead,
		Update: resourceCloudflareTeamsProxyEndpointUpdate,
		Delete: resourceCloudflareTeamsProxyEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareTeamsProxyEndpointImport,
		},
	}
}

func resourceCloudflareTeamsProxyEndpointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	endpoint, err := client.TeamsProxyEndpoint(context.Background(), accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 400") {
			log.Printf("[INFO] Teams Proxy Endpoint %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Teams Proxy Endpoint %q: %s", d.Id(), err)
	}

	if err := d.Set("name", endpoint.Name); err != nil {
		return fmt.Errorf("error parsing Proxy Endpoint name")
	}

	if err := d.Set("ips", endpoint.IPs); err != nil {
		return fmt.Errorf("error parsing Proxy Endpoint IPs")
	}

	if err := d.Set("subdomain", endpoint.Subdomain); err != nil {
		return fmt.Errorf("error parsing Proxy Endpoint subdomain")
	}

	return nil
}

func resourceCloudflareTeamsProxyEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	accountID := d.Get("account_id").(string)
	newProxyEndpoint := cloudflare.TeamsProxyEndpoint{
		Name: d.Get("name").(string),
		IPs:  expandInterfaceToStringList(d.Get("ips").(*schema.Set).List()),
	}

	log.Printf("[DEBUG] Creating Cloudflare Teams Proxy Endpoint from struct: %+v", newProxyEndpoint)

	proxyEndpoint, err := client.CreateTeamsProxyEndpoint(context.Background(), accountID, newProxyEndpoint)
	if err != nil {
		return fmt.Errorf("error creating Teams Proxy Endpoint for account %q: %s", accountID, err)
	}

	d.SetId(proxyEndpoint.ID)
	return resourceCloudflareTeamsProxyEndpointRead(d, meta)

}

func resourceCloudflareTeamsProxyEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	updatedProxyEndpoint := cloudflare.TeamsProxyEndpoint{
		ID:   d.Id(),
		Name: d.Get("name").(string),
		IPs:  expandInterfaceToStringList(d.Get("ips").(*schema.Set).List()),
	}

	log.Printf("[DEBUG] Updating Cloudflare Teams Proxy Endpoint from struct: %+v", updatedProxyEndpoint)

	teamsProxyEndpoint, err := client.UpdateTeamsProxyEndpoint(context.Background(), accountID, updatedProxyEndpoint)

	if err != nil {
		return fmt.Errorf("error updating Teams Proxy Endpoint for account %q: %s", accountID, err)
	}

	if teamsProxyEndpoint.ID == "" {
		return fmt.Errorf("failed to find Teams Proxy Endpoint ID in update response; resource was empty")
	}
	return resourceCloudflareTeamsProxyEndpointRead(d, meta)
}

func resourceCloudflareTeamsProxyEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	id := d.Id()
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Teams Proxy Endpoint using ID: %s", id)

	err := client.DeleteTeamsProxyEndpoint(context.Background(), accountID, id)
	if err != nil {
		return fmt.Errorf("error deleting Teams Proxy Endpoint for account %q: %s", accountID, err)
	}

	return resourceCloudflareTeamsProxyEndpointRead(d, meta)
}

func resourceCloudflareTeamsProxyEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsProxyEndpointID\"", d.Id())
	}

	accountID, teamsProxyEndpointID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Teams Proxy Endpoint: id %s for account %s", teamsProxyEndpointID, accountID)

	d.Set("account_id", accountID)
	d.SetId(teamsProxyEndpointID)

	err := resourceCloudflareTeamsProxyEndpointRead(d, meta)

	return []*schema.ResourceData{d}, err

}
