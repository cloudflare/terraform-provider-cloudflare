package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessKeysConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareAccessKeysConfigurationSchema(),
		ReadContext: resourceCloudflareAccessKeysConfigurationRead,
		CreateContext: resourceCloudflareAccessKeysConfigurationCreate,
		UpdateContext: resourceCloudflareAccessKeysConfigurationUpdate,
		DeleteContext: resourceCloudflareKeysConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareKeysConfigurationImport,
		},
	}
}

func resourceCloudflareAccessKeysConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	keysConfig, err := client.AccessKeysConfig(context.Background(), accountID)
	if err != nil {
		var requestError *cloudflare.RequestError
		if errors.As(err, &requestError) {
			if sliceContainsInt(requestError.ErrorCodes(), 12109) {
				log.Printf("[INFO] Access Keys Configuration not enabled for account %s", accountID)
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(fmt.Errorf("error finding Access Keys Configuration %s: %s", accountID, err))
	}

	d.SetId(accountID)
	d.Set("key_rotation_interval_days", keysConfig.KeyRotationIntervalDays)

	return nil
}

func resourceCloudflareAccessKeysConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// keys configuration share the same lifetime as an organization, so creating is a no-op, unless
	// key_rotation_interval_days was explicitly passed, in which case we need to update its value.

	if keyRotationIntervalDays := d.Get("key_rotation_interval_days").(int); keyRotationIntervalDays == 0 {
		return resourceCloudflareAccessKeysConfigurationRead(d, meta)
	}

	return resourceCloudflareAccessKeysConfigurationUpdate(d, meta)
}

func resourceCloudflareAccessKeysConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	keysConfigUpdateReq := cloudflare.AccessKeysConfigUpdateRequest{
		KeyRotationIntervalDays: d.Get("key_rotation_interval_days").(int),
	}

	_, err := client.UpdateAccessKeysConfig(context.Background(), accountID, keysConfigUpdateReq)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Keys Configuration for account %s: %s", accountID, err))
	}

	return resourceCloudflareAccessKeysConfigurationRead(d, meta)
}

func resourceCloudflareKeysConfigurationDelete(ctx context.Context, _ *schema.ResourceData, _ interface{}) error {
	// keys configuration share the same lifetime as an organization, and can not be
	// explicitly deleted by the user. so this is a no-op.
	return nil
}

func resourceCloudflareKeysConfigurationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID := d.Id()

	d.SetId(accountID)
	d.Set("account_id", accountID)

	err := resourceCloudflareAccessKeysConfigurationRead(d, meta)
	return []*schema.ResourceData{d}, err
}
