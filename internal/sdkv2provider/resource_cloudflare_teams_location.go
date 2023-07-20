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

func resourceCloudflareTeamsLocation() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsLocationSchema(),
		CreateContext: resourceCloudflareTeamsLocationCreate,
		ReadContext:   resourceCloudflareTeamsLocationRead,
		UpdateContext: resourceCloudflareTeamsLocationUpdate,
		DeleteContext: resourceCloudflareTeamsLocationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsLocationImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Teams Location resource. Teams Locations are
			referenced when creating secure web gateway policies.
		`),
	}
}

func resourceCloudflareTeamsLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	location, err := client.TeamsLocation(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Location ID is invalid") {
			tflog.Info(ctx, fmt.Sprintf("Teams Location %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Teams Location %q: %w", d.Id(), err))
	}

	if err := d.Set("name", location.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location name"))
	}
	if err := d.Set("networks", flattenTeamsLocationNetworks(location.Networks)); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location networks"))
	}
	if err := d.Set("policy_ids", location.PolicyIDs); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location policy IDs"))
	}
	if err := d.Set("ip", location.Ip); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location IP"))
	}
	if err := d.Set("doh_subdomain", location.Subdomain); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location DOH subdomain"))
	}
	if err := d.Set("anonymized_logs_enabled", location.AnonymizedLogsEnabled); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location anonimized log enablement"))
	}
	if err := d.Set("ipv4_destination", location.IPv4Destination); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location IPv4 destination"))
	}
	if err := d.Set("client_default", location.ClientDefault); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing Location client default"))
	}

	return nil
}
func resourceCloudflareTeamsLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	networks, err := inflateTeamsLocationNetworks(d.Get("networks"))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Teams Location for account %q: %w, %v", accountID, err, networks))
	}

	newTeamLocation := cloudflare.TeamsLocation{
		Name:          d.Get("name").(string),
		Networks:      networks,
		ClientDefault: d.Get("client_default").(bool),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Teams Location from struct: %+v", newTeamLocation))

	location, err := client.CreateTeamsLocation(ctx, accountID, newTeamLocation)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Teams Location for account %q: %w, %v", accountID, err, networks))
	}

	d.SetId(location.ID)
	return resourceCloudflareTeamsLocationRead(ctx, d, meta)
}
func resourceCloudflareTeamsLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	networks, err := inflateTeamsLocationNetworks(d.Get("networks"))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Teams Location for account %q: %w, %v", accountID, err, networks))
	}
	updatedTeamsLocation := cloudflare.TeamsLocation{
		ID:            d.Id(),
		Name:          d.Get("name").(string),
		ClientDefault: d.Get("client_default").(bool),
		Networks:      networks,
	}
	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Teams Location from struct: %+v", updatedTeamsLocation))

	teamsLocation, err := client.UpdateTeamsLocation(ctx, accountID, updatedTeamsLocation)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Teams Location for account %q: %w", accountID, err))
	}
	if teamsLocation.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Teams Location ID in update response; resource was empty"))
	}
	return resourceCloudflareTeamsLocationRead(ctx, d, meta)
}

func resourceCloudflareTeamsLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	id := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Teams Location using ID: %s", id))

	err := client.DeleteTeamsLocation(ctx, accountID, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Teams Location for account %q: %w", accountID, err))
	}

	return resourceCloudflareTeamsLocationRead(ctx, d, meta)
}

func resourceCloudflareTeamsLocationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsLocationID\"", d.Id())
	}

	accountID, teamsLocationID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Teams Location: id %s for account %s", teamsLocationID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(teamsLocationID)

	resourceCloudflareTeamsLocationRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
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
