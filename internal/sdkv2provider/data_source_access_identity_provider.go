package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Schema:      dataSourceCloudflareAccessIdentityProviderSchema(),
		ReadContext: dataSourceCloudflareAccessIdentityProviderRead,
		Description: "Use this data source to lookup a single [Access Identity Provider](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration) by name.",
	}
}

func dataSourceCloudflareAccessIdentityProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	name := d.Get("name").(string)
	if err != nil {
		return diag.FromErr(err)
	}

	var providers []cloudflare.AccessIdentityProvider
	if identifier.Type == AccountType {
		providers, err = client.AccessIdentityProviders(ctx, identifier.Value)
	} else {
		providers, err = client.ZoneLevelAccessIdentityProviders(ctx, identifier.Value)
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Access Identity Providers: %w", err))
	}

	if len(providers) == 0 {
		return diag.FromErr(fmt.Errorf("no Access Identity Providers found"))
	}

	var accessIdentityProvider cloudflare.AccessIdentityProvider
	for _, provider := range providers {
		if provider.Name == name {
			accessIdentityProvider = provider
			break
		}
	}

	if accessIdentityProvider.ID == "" {
		return diag.FromErr(fmt.Errorf("no Access Identity Provider matching name %q", name))
	}

	d.SetId(accessIdentityProvider.ID)
	d.Set("name", accessIdentityProvider.Name)
	d.Set("type", accessIdentityProvider.Type)

	return nil
}
