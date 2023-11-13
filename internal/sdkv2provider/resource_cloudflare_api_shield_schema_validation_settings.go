package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAPIShieldSchemaValidationSettings() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAPIShieldSchemaValidationSettingsSchema(),
		CreateContext: resourceCloudflareAPIShieldSchemaValidationSettingsCreate,
		ReadContext:   resourceCloudflareAPIShieldSchemaValidationSettingsRead,
		UpdateContext: resourceCloudflareAPIShieldSchemaValidationSettingsUpdate,
		DeleteContext: resourceCloudflareAPIShieldSchemaValidationSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage settings in API Shield Schema Validation 2.0.
		`),
	}
}

func resourceCloudflareAPIShieldSchemaValidationSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceCloudflareAPIShieldSchemaValidationSettingsUpdate(ctx, d, meta)
}

func resourceCloudflareAPIShieldSchemaValidationSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	dm := d.Get("validation_default_mitigation_action").(string)

	var overrideAction *string
	if oa, ok := d.GetOk("validation_override_mitigation_action"); ok {
		overrideAction = cloudflare.StringPtr(oa.(string))
	}

	_, err := client.UpdateAPIShieldSchemaValidationSettings(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.UpdateAPIShieldSchemaValidationSettingsParams{
			DefaultMitigationAction:  &dm,
			OverrideMitigationAction: overrideAction,
		},
	)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create API Shield Schema Validation Settings"))
	}

	// Settings are configured at the zone level so using the zoneID as the ID
	d.SetId(zoneID)
	return resourceCloudflareAPIShieldSchemaValidationSettingsRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldSchemaValidationSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	settings, err := client.GetAPIShieldSchemaValidationSettings(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
	)

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Schema Validation Settings: %w", err))
	}

	if err := d.Set("validation_default_mitigation_action", settings.DefaultMitigationAction); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("validation_override_mitigation_action", settings.OverrideMitigationAction); err != nil {
		return diag.FromErr(err)
	}

	// Settings are configured at the zone level so using the zoneID as the ID
	d.SetId(zoneID)
	return nil
}

func resourceCloudflareAPIShieldSchemaValidationSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	defaultSettings := cloudflareAPIShieldSchemaValidationSettingsDefault()

	// There is no DELETE endpoint for schema validation settings,
	// so terraform should reset the state to default settings
	_, err := client.UpdateAPIShieldSchemaValidationSettings(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.UpdateAPIShieldSchemaValidationSettingsParams{
			DefaultMitigationAction:  &defaultSettings.DefaultMitigationAction,
			OverrideMitigationAction: defaultSettings.OverrideMitigationAction,
		},
	)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete API Shield Schema Validation Settings: %w", err))
	}

	return nil
}

func cloudflareAPIShieldSchemaValidationSettingsDefault() *cloudflare.APIShieldSchemaValidationSettings {
	disableOverride := "disable_override"
	return &cloudflare.APIShieldSchemaValidationSettings{
		DefaultMitigationAction:  "none",
		OverrideMitigationAction: &disableOverride,
	}
}
