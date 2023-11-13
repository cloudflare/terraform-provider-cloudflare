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

func resourceCloudflareAPIShieldOperationSchemaValidationSettings() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAPIShieldOperationSchemaValidationSettingsSchema(),
		CreateContext: resourceCloudflareAPIShieldOperationSchemaValidationSettingsCreate,
		ReadContext:   resourceCloudflareAPIShieldOperationSchemaValidationSettingsRead,
		UpdateContext: resourceCloudflareAPIShieldOperationSchemaValidationSettingsUpdate,
		DeleteContext: resourceCloudflareAPIShieldOperationSchemaValidationSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage operation-level settings in API Shield Schema Validation 2.0.
		`),
	}
}

func resourceCloudflareAPIShieldOperationSchemaValidationSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceCloudflareAPIShieldOperationSchemaValidationSettingsUpdate(ctx, d, meta)
}

func resourceCloudflareAPIShieldOperationSchemaValidationSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	operationID := d.Get("operation_id").(string)

	var mitigationAction *string
	if ma, ok := d.GetOk("mitigation_action"); ok {
		mitigationAction = cloudflare.StringPtr(ma.(string))
	}

	_, err := client.UpdateAPIShieldOperationSchemaValidationSettings(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.UpdateAPIShieldOperationSchemaValidationSettings{
			operationID: cloudflare.APIShieldOperationSchemaValidationSettings{
				MitigationAction: mitigationAction,
			},
		},
	)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create API Shield Operation Schema Validation Settings"))
	}

	// Settings are configured at the operation level so using the operationID as the ID
	d.SetId(operationID)
	return resourceCloudflareAPIShieldOperationSchemaValidationSettingsRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldOperationSchemaValidationSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	operationID := d.Get("operation_id").(string)
	settings, err := client.GetAPIShieldOperationSchemaValidationSettings(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.GetAPIShieldOperationSchemaValidationSettingsParams{
			OperationID: operationID,
		},
	)

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Operation Schema Validation Settings: %w", err))
	}

	if err := d.Set("mitigation_action", settings.MitigationAction); err != nil {
		return diag.FromErr(err)
	}

	// Settings are configured at the operation level so using the zoneID as the ID
	d.SetId(operationID)
	return nil
}

func resourceCloudflareAPIShieldOperationSchemaValidationSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	operationID := d.Get("operation_id").(string)

	// There is no DELETE endpoint for schema validation settings,
	// so terraform should reset the state to default settings
	_, err := client.UpdateAPIShieldOperationSchemaValidationSettings(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.UpdateAPIShieldOperationSchemaValidationSettings{
			operationID: cloudflare.APIShieldOperationSchemaValidationSettings{},
		},
	)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete API Shield Operation Schema Validation Settings: %w", err))
	}

	return nil
}
