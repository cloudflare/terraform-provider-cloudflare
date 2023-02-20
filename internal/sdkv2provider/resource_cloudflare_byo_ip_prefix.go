package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareBYOIPPrefix() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareBYOIPPrefixSchema(),
		CreateContext: resourceCloudflareBYOIPPrefixCreate,
		ReadContext:   resourceCloudflareBYOIPPrefixRead,
		UpdateContext: resourceCloudflareBYOIPPrefixUpdate,
		DeleteContext: resourceCloudflareBYOIPPrefixDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareBYOIPPrefixImport,
		},
		Description: heredoc.Doc(`
			Provides the ability to manage Bring-Your-Own-IP prefixes (BYOIP)
			which are used with or without Magic Transit.
		`),
	}
}

func resourceCloudflareBYOIPPrefixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	prefixID := d.Get("prefix_id")
	d.SetId(prefixID.(string))

	if err := resourceCloudflareBYOIPPrefixUpdate(ctx, d, meta); err != nil {
		return err
	}

	return resourceCloudflareBYOIPPrefixRead(ctx, d, meta)
}

func resourceCloudflareBYOIPPrefixImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/prefixID"`, d.Id())
	}

	accountID, prefixID := attributes[0], attributes[1]

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("prefix_id", prefixID)
	d.SetId(prefixID)

	resourceCloudflareBYOIPPrefixRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareBYOIPPrefixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	prefix, err := client.GetPrefix(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading IP prefix information for %q", d.Id())))
	}

	d.Set("description", prefix.Description)

	advertisementStatus, err := client.GetAdvertisementStatus(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading advertisement status of IP prefix for %q", d.Id())))
	}

	d.Set("advertisement", stringFromBool(advertisementStatus.Advertised))

	return nil
}

func resourceCloudflareBYOIPPrefixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	if _, ok := d.GetOk("description"); ok && d.HasChange("description") {
		if _, err := client.UpdatePrefixDescription(ctx, accountID, d.Id(), d.Get("description").(string)); err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("cannot update prefix description for %q", d.Id())))
		}
	}

	if _, ok := d.GetOk("advertisement"); ok && d.HasChange("advertisement") {
		if _, err := client.UpdateAdvertisementStatus(ctx, accountID, d.Id(), boolFromString(d.Get("advertisement").(string))); err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("cannot update prefix advertisement status for %q", d.Id())))
		}
	}

	return nil
}

// Deletion of prefixes is not really supported, so we keep this as a dummy.
func resourceCloudflareBYOIPPrefixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
