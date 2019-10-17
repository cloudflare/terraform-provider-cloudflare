package cloudflare

import (
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareAccessServiceToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccessServiceTokenCreate,
		Read:   resourceCloudflareAccessServiceTokenRead,
		Update: resourceCloudflareAccessServiceTokenUpdate,
		Delete: resourceCloudflareAccessServiceTokenDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccessServiceTokenImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceCloudflareAccessServiceTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	// The Cloudflare API doesn't support fetching a single service token
	// so instead we loop over all the service tokens and only continue
	// when we have a match.
	serviceTokens, _, err := client.AccessServiceTokens(accountID)
	if err != nil {
		return fmt.Errorf("error fetching access service tokens: %s", err)
	}
	for _, token := range serviceTokens {
		if token.ID == d.Id() {
			d.Set("name", token.Name)
			d.Set("client_id", token.ClientID)
		}
	}

	return nil
}

func resourceCloudflareAccessServiceTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tokenName := d.Get("name").(string)

	serviceToken, err := client.CreateAccessServiceToken(accountID, tokenName)
	if err != nil {
		return fmt.Errorf("error creating access service token: %s", err)
	}

	d.SetId(serviceToken.ID)
	d.Set("name", serviceToken.Name)
	d.Set("client_id", serviceToken.ClientID)
	d.Set("client_secret", serviceToken.ClientSecret)

	resourceCloudflareAccessServiceTokenRead(d, meta)

	return nil
}

func resourceCloudflareAccessServiceTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tokenName := d.Get("name").(string)

	serviceToken, err := client.UpdateAccessServiceToken(accountID, d.Id(), tokenName)
	if err != nil {
		return fmt.Errorf("error updating access service token: %s", err)
	}

	d.Set("name", serviceToken.Name)

	return resourceCloudflareAccessServiceTokenRead(d, meta)
}

func resourceCloudflareAccessServiceTokenDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteAccessServiceToken(accountID, d.Id())
	if err != nil {
		return fmt.Errorf("error deleting access service token: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceCloudflareAccessServiceTokenImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/serviceTokenID\"", d.Id())
	}

	d.Set("account_id", attributes[0])
	d.SetId(attributes[1])

	resourceCloudflareAccessServiceTokenRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
