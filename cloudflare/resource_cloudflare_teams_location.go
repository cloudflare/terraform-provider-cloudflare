package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareTeamsLocation() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareTeamsLocationSchema(),
		Create: resourceCloudflareTeamsLocationCreate,
		Read:   resourceCloudflareTeamsLocationRead,
		Update: resourceCloudflareTeamsLocationUpdate,
		Delete: resourceCloudflareTeamsLocationDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareTeamsLocationImport,
		},
	}
}

func resourceCloudflareTeamsLocationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	location, err := client.TeamsLocation(context.Background(), accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 400") {
			log.Printf("[INFO] Teams Location %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Teams Location %q: %s", d.Id(), err)
	}

	if err := d.Set("name", location.Name); err != nil {
		return fmt.Errorf("error parsing Location name")
	}
	if err := d.Set("networks", flattenTeamsLocationNetworks(location.Networks)); err != nil {
		return fmt.Errorf("error parsing Location networks")
	}
	if err := d.Set("policy_ids", location.PolicyIDs); err != nil {
		return fmt.Errorf("error parsing Location policy IDs")
	}
	if err := d.Set("ip", location.Ip); err != nil {
		return fmt.Errorf("error parsing Location IP")
	}
	if err := d.Set("doh_subdomain", location.Subdomain); err != nil {
		return fmt.Errorf("error parsing Location DOH subdomain")
	}
	if err := d.Set("anonymized_logs_enabled", location.AnonymizedLogsEnabled); err != nil {
		return fmt.Errorf("error parsing Location anonimized log enablement")
	}
	if err := d.Set("ipv4_destination", location.IPv4Destination); err != nil {
		return fmt.Errorf("error parsing Location IPv4 destination")
	}
	if err := d.Set("client_default", location.ClientDefault); err != nil {
		return fmt.Errorf("error parsing Location client default")
	}

	return nil
}
func resourceCloudflareTeamsLocationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	accountID := d.Get("account_id").(string)
	networks, err := inflateTeamsLocationNetworks(d.Get("networks"))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error creating Teams Location for account %q: %s, %v", accountID, err, networks))
	}

	newTeamLocation := cloudflare.TeamsLocation{
		Name:          d.Get("name").(string),
		Networks:      networks,
		ClientDefault: d.Get("client_default").(bool),
	}

	log.Printf("[DEBUG] Creating Cloudflare Teams Location from struct: %+v", newTeamLocation)

	location, err := client.CreateTeamsLocation(context.Background(), accountID, newTeamLocation)
	if err != nil {
		return fmt.Errorf("error creating Teams Location for account %q: %s, %v", accountID, err, networks)
	}

	d.SetId(location.ID)
	return resourceCloudflareTeamsLocationRead(d, meta)

}
func resourceCloudflareTeamsLocationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	networks, err := inflateTeamsLocationNetworks(d.Get("networks"))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error updating Teams Location for account %q: %s, %v", accountID, err, networks))
	}
	updatedTeamsLocation := cloudflare.TeamsLocation{
		ID:            d.Id(),
		Name:          d.Get("name").(string),
		ClientDefault: d.Get("client_default").(bool),
		Networks:      networks,
	}
	log.Printf("[DEBUG] Updating Cloudflare Teams Location from struct: %+v", updatedTeamsLocation)

	teamsLocation, err := client.UpdateTeamsLocation(context.Background(), accountID, updatedTeamsLocation)
	if err != nil {
		return fmt.Errorf("error updating Teams Location for account %q: %s", accountID, err)
	}
	if teamsLocation.ID == "" {
		return fmt.Errorf("failed to find Teams Location ID in update response; resource was empty")
	}
	return resourceCloudflareTeamsLocationRead(d, meta)
}

func resourceCloudflareTeamsLocationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	id := d.Id()
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Teams Location using ID: %s", id)

	err := client.DeleteTeamsLocation(context.Background(), accountID, id)
	if err != nil {
		return fmt.Errorf("error deleting Teams Location for account %q: %s", accountID, err)
	}

	return resourceCloudflareTeamsLocationRead(d, meta)
}

func resourceCloudflareTeamsLocationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsLocationID\"", d.Id())
	}

	accountID, teamsLocationID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Teams Location: id %s for account %s", teamsLocationID, accountID)

	d.Set("account_id", accountID)
	d.SetId(teamsLocationID)

	err := resourceCloudflareTeamsLocationRead(d, meta)

	return []*schema.ResourceData{d}, err

}

func inflateTeamsLocationNetworks(networks interface{}) ([]cloudflare.TeamsLocationNetwork, error) {
	var networkStructs []cloudflare.TeamsLocationNetwork
	if networks != nil {
		networkSet, ok := networks.(*schema.Set)
		if !ok {
			return nil, fmt.Errorf("error parsing network list")
		}
		for _, i := range networkSet.List() {
			network, ok := i.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("error parsing network")
			}
			networkStructs = append(networkStructs, cloudflare.TeamsLocationNetwork{
				ID:      network["id"].(string),
				Network: network["network"].(string),
			})
		}
	}
	return networkStructs, nil
}

func flattenTeamsLocationNetworks(networks []cloudflare.TeamsLocationNetwork) []interface{} {
	var flattenedNetworks []interface{}
	for _, net := range networks {
		flattenedNetworks = append(flattenedNetworks, map[string]interface{}{
			"id":      net.ID,
			"network": net.Network,
		})
	}
	return flattenedNetworks
}
