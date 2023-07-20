package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessKeysConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessKeysConfigurationSchema(),
		ReadContext:   resourceCloudflareAccessKeysConfigurationRead,
		CreateContext: resourceCloudflareAccessKeysConfigurationCreate,
		UpdateContext: resourceCloudflareAccessKeysConfigurationUpdate,
		DeleteContext: resourceCloudflareKeysConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareKeysConfigurationImport,
		},
		Description: heredoc.Doc(`
			Access Keys Configuration defines the rotation policy for the keys
			that access will use to sign data.
		`),
	}
}

func resourceCloudflareAccessKeysConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	keysConfig, err := client.AccessKeysConfig(ctx, accountID)
	if err != nil {
		var requestError *cloudflare.RequestError
		if errors.As(err, &requestError) {
			if sliceContainsInt(requestError.ErrorCodes(), 12109) {
				tflog.Info(ctx, fmt.Sprintf("Access Keys Configuration not enabled for account %s", accountID))
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(fmt.Errorf("error finding Access Keys Configuration %s: %w", accountID, err))
	}

	d.SetId(accountID)
	d.Set("key_rotation_interval_days", keysConfig.KeyRotationIntervalDays)

	return nil
}

func resourceCloudflareAccessKeysConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// keys configuration share the same lifetime as an organization, so creating is a no-op, unless
	// key_rotation_interval_days was explicitly passed, in which case we need to update its value.

	if keyRotationIntervalDays := d.Get("key_rotation_interval_days").(int); keyRotationIntervalDays == 0 {
		return resourceCloudflareAccessKeysConfigurationRead(ctx, d, meta)
	}

	return resourceCloudflareAccessKeysConfigurationUpdate(ctx, d, meta)
}

func resourceCloudflareAccessKeysConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	keysConfigUpdateReq := cloudflare.AccessKeysConfigUpdateRequest{
		KeyRotationIntervalDays: d.Get("key_rotation_interval_days").(int),
	}

	_, err := client.UpdateAccessKeysConfig(ctx, accountID, keysConfigUpdateReq)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Keys Configuration for account %s: %w", accountID, err))
	}

	return resourceCloudflareAccessKeysConfigurationRead(ctx, d, meta)
}

func resourceCloudflareKeysConfigurationDelete(ctx context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	// keys configuration share the same lifetime as an organization, and can not be
	// explicitly deleted by the user. so this is a no-op.
	return nil
}

func resourceCloudflareKeysConfigurationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID := d.Id()

	d.SetId(accountID)
	d.Set(consts.AccountIDSchemaKey, accountID)

	if err := resourceCloudflareAccessKeysConfigurationRead(ctx, d, meta); err != nil {
		return nil, errors.New(err[0].Summary)
	}

	return []*schema.ResourceData{d}, nil
}
