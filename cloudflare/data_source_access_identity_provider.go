package cloudflare

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Schema: dataSourceCloudflareAccessIdentityProviderSchema(),
		Read:   dataSourceCloudflareAccessIdentityProviderRead,
	}
}

func dataSourceCloudflareAccessIdentityProviderRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	name := d.Get("name").(string)
	if err != nil {
		return err
	}

	var providers []cloudflare.AccessIdentityProvider
	if identifier.Type == AccountType {
		providers, err = client.AccessIdentityProviders(context.Background(), identifier.Value)
	} else {
		providers, err = client.ZoneLevelAccessIdentityProviders(context.Background(), identifier.Value)
	}

	if err != nil {
		return fmt.Errorf("error listing Access Identity Providers: %s", err)
	}

	if len(providers) == 0 {
		return fmt.Errorf("no Access Identity Providers found")
	}

	var accessIdentityProvider cloudflare.AccessIdentityProvider
	for _, provider := range providers {
		if provider.Name == name {
			accessIdentityProvider = provider
			break
		}
	}

	if accessIdentityProvider.ID == "" {
		return fmt.Errorf("no Access Identity Provider matching name %q", name)
	}

	d.SetId(accessIdentityProvider.ID)
	d.Set("id", accessIdentityProvider.ID)
	d.Set("name", accessIdentityProvider.Name)
	d.Set("type", accessIdentityProvider.Type)

	return nil
}
