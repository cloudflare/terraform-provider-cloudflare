package cloudflare

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareArgoTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareArgoTunnelCreate,
		Read:   resourceCloudflareArgoTunnelRead,
		Update: resourceCloudflareArgoTunnelUpdate,
		Delete: resourceCloudflareArgoTunnelDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareArgoTunnelImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
		},
	}
}

func resourceCloudflareArgoTunnelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accID := d.Get("account_id").(string)
	name := d.Get("name").(string)
	secret := d.Get("secret").(string)

	tunnel, err := client.CreateArgoTunnel(context.Background(), accID, name, secret)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create Argo Tunnel"))
	}

	d.SetId(tunnel.ID)

	return resourceCloudflareArgoTunnelRead(d, meta)
}

func resourceCloudflareArgoTunnelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accID := d.Get("account_id").(string)

	tunnel, err := client.ArgoTunnel(context.Background(), accID, d.Id())
	if err != nil {
		return fmt.Errorf("failed to fetch Argo Tunnel: %w", err)
	}

	d.Set("cname", fmt.Sprintf("%s.argotunnel.com", tunnel.ID))
	return nil
}

func resourceCloudflareArgoTunnelUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudflareArgoTunnelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accID := d.Get("account_id").(string)

	cleanupErr := client.CleanupArgoTunnelConnections(context.Background(), accID, d.Id())
	if cleanupErr != nil {
		return errors.Wrap(cleanupErr, fmt.Sprintf("failed to clean up Argo Tunnel connections"))
	}

	deleteErr := client.DeleteArgoTunnel(context.Background(), accID, d.Id())
	if deleteErr != nil {
		return errors.Wrap(deleteErr, fmt.Sprintf("failed to delete Argo Tunnel"))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareArgoTunnelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	attributes := strings.Split(d.Id(), "/")

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/argoTunnelUUID\"", d.Id())
	}

	accID, tunnelID := attributes[0], attributes[1]

	tunnel, err := client.ArgoTunnel(context.Background(), accID, tunnelID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to fetch Argo Tunnel %s", tunnelID))
	}

	d.Set("name", tunnel.Name)
	d.Set("cname", fmt.Sprintf("%s.argotunnel.com", tunnel.ID))
	d.SetId(tunnel.ID)

	resourceCloudflareArgoTunnelRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
