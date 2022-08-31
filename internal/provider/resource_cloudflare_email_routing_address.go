package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareEmailRoutingAddress() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareEmailRoutingAddressSchema(),
		ReadContext:   resourceCloudflareEmailRoutingAddressRead,
		CreateContext: resourceCloudflareEmailRoutingAddressCreate,
		DeleteContext: resourceCloudflareEmailRoutingAddressDelete,
		Description: heredoc.Doc(`
			Provides a resource for managing Email Routing Addresses.
		`),
	}
}

func resourceCloudflareEmailRoutingAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	res, err := client.GetEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(accountID), d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error getting email routing destination address %q: %w", d.Id(), err))
	}

	d.SetId(res.Tag)
	d.Set("email", res.Email)
	if res.Verified != nil {
		d.Set("verified", res.Verified.Format(time.RFC3339Nano))
	}
	d.Set("created", res.Created.Format(time.RFC3339Nano))
	d.Set("modified", res.Modified.Format(time.RFC3339Nano))
	return nil
}

func resourceCloudflareEmailRoutingAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	email := d.Get("email").(string)

	createParams := cloudflare.CreateEmailRoutingAddressParameters{
		Email: email,
	}
	res, err := client.CreateEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(accountID), createParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating email routing destination address %q, %w", email, err))
	}

	d.SetId(res.Tag)
	return resourceCloudflareEmailRoutingAddressRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleteing email routing destination address %q, %w", d.Id(), err))
	}

	return nil
}
