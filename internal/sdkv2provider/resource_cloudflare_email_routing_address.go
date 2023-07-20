package sdkv2provider

import (
	"context"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.DeleteEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting email routing destination address %q, %w", d.Id(), err))
	}

	return nil
}
