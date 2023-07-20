package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDeviceManagedNetworks() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDeviceManagedNetworksSchema(),
		CreateContext: resourceCloudflareDeviceManagedNetworksCreate,
		ReadContext:   resourceCloudflareDeviceManagedNetworksRead,
		UpdateContext: resourceCloudflareDeviceManagedNetworksUpdate,
		DeleteContext: resourceCloudflareDeviceManagedNetworksDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDeviceManagedNetworksImport,
		},
		Description: "Provides a Cloudflare Device Managed Network resource. Device managed networks allow for building location-aware device settings policies.",
	}
}

func resourceCloudflareDeviceManagedNetworksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))
	tflog.Debug(ctx, fmt.Sprintf("Reading Cloudflare Device Managed Network for Id: %+v", d.Id()))

	managedNetwork, err := client.GetDeviceManagedNetwork(ctx, identifier, d.Id())

	var notFoundError *cloudflare.NotFoundError
	if errors.As(err, &notFoundError) {
		tflog.Info(ctx, fmt.Sprintf("Device Managed Network %s no longer exists", d.Id()))
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading Device Managed Network: %w", err))
	}

	d.Set("name", managedNetwork.Name)
	d.Set("type", managedNetwork.Type)
	d.Set("config", convertDeviceManagedNetworkConfigToSchema(managedNetwork.Config))

	return nil
}

func resourceCloudflareDeviceManagedNetworksCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))

	params := cloudflare.CreateDeviceManagedNetworkParams{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
		Config: &cloudflare.Config{
			TlsSockAddr: d.Get("config.0.tls_sockaddr").(string),
			Sha256:      d.Get("config.0.sha256").(string),
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Device Managed Network with params: %+v", params))

	managedNetwork, err := client.CreateDeviceManagedNetwork(ctx, identifier, params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Managed Network with provided config: %w", err))
	}

	d.SetId(managedNetwork.NetworkID)

	return resourceCloudflareDeviceManagedNetworksRead(ctx, d, meta)
}

func resourceCloudflareDeviceManagedNetworksUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))

	updatedDeviceManagedNetworkParams := cloudflare.UpdateDeviceManagedNetworkParams{
		NetworkID: d.Id(),
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
		Config: &cloudflare.Config{
			TlsSockAddr: d.Get("config.0.tls_sockaddr").(string),
			Sha256:      d.Get("config.0.sha256").(string),
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Device Managed Network with params: %+v", updatedDeviceManagedNetworkParams))

	managedNetwork, err := client.UpdateDeviceManagedNetwork(ctx, identifier, updatedDeviceManagedNetworkParams)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Device Managed Network for ID %q: %w", d.Id(), err))
	}
	if managedNetwork.NetworkID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Network ID in update response; resource was empty"))
	}

	return resourceCloudflareDeviceManagedNetworksRead(ctx, d, meta)
}

func resourceCloudflareDeviceManagedNetworksDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))
	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Device Managed Network using ID: %s", d.Id()))

	if _, err := client.DeleteManagedNetworks(ctx, identifier, d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting DLP Profile for ID %q: %w", d.Id(), err))
	}

	resourceCloudflareDeviceManagedNetworksRead(ctx, d, meta)
	return nil
}

func resourceCloudflareDeviceManagedNetworksImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID, managedNetworkID, err := parseDeviceManagedNetworksIDImport(d.Id())
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Device Managed Network: id %s for account %s", managedNetworkID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(managedNetworkID)

	resourceCloudflareDeviceManagedNetworksRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func parseDeviceManagedNetworksIDImport(id string) (string, string, error) {
	attributes := strings.SplitN(id, "/", 2)

	if len(attributes) != 2 {
		return "", "", fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/networkID\"", id)
	}

	return attributes[0], attributes[1], nil
}

func convertDeviceManagedNetworkConfigToSchema(input *cloudflare.Config) []interface{} {
	m := map[string]interface{}{
		"sha256":       &input.Sha256,
		"tls_sockaddr": &input.TlsSockAddr,
	}
	return []interface{}{m}
}
