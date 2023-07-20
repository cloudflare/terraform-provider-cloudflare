package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareAPIShield() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAPIShieldSchema(),
		CreateContext: resourceCloudflareAPIShieldCreate,
		ReadContext:   resourceCloudflareAPIShieldRead,
		UpdateContext: resourceCloudflareAPIShieldUpdate,
		DeleteContext: resourceCloudflareAPIShieldDelete,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage API Shield configurations.
		`),
	}
}

func resourceCloudflareAPIShieldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	as, err := buildAPIShieldConfiguration(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create API Shield Configuration"))
	}

	_, err = client.UpdateAPIShieldConfiguration(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateAPIShieldParams{AuthIdCharacteristics: as.AuthIdCharacteristics})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create API Shield Configuration"))
	}

	return resourceCloudflareAPIShieldRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	as, _, err := client.GetAPIShieldConfiguration(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Configuration: %w", err))
	}

	d.Set("auth_id_characteristics", flattenAPIShieldConfiguration(as.AuthIdCharacteristics))
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(zoneID)

	return nil
}

func resourceCloudflareAPIShieldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey)

	as, err := buildAPIShieldConfiguration(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create API Shield Configuration")))
	}

	_, err = client.UpdateAPIShieldConfiguration(ctx, cloudflare.ZoneIdentifier(zoneID.(string)), cloudflare.UpdateAPIShieldParams{AuthIdCharacteristics: as.AuthIdCharacteristics})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create API Shield Configuration")))
	}

	return resourceCloudflareAPIShieldRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey)

	_, err := client.UpdateAPIShieldConfiguration(ctx, cloudflare.ZoneIdentifier(zoneID.(string)), cloudflare.UpdateAPIShieldParams{AuthIdCharacteristics: []cloudflare.AuthIdCharacteristics{}})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create API Shield Configuration")))
	}

	return resourceCloudflareAPIShieldRead(ctx, d, meta)
}

func buildAPIShieldConfiguration(d *schema.ResourceData) (cloudflare.APIShield, error) {
	as := cloudflare.APIShield{}
	configs, ok := d.Get("auth_id_characteristics").([]interface{})
	if !ok {
		return cloudflare.APIShield{}, errors.New("unable to create interface map type assertion for rule")
	}

	as.AuthIdCharacteristics = []cloudflare.AuthIdCharacteristics{}
	for i := 0; i < len(configs); i++ {
		as.AuthIdCharacteristics = append(as.AuthIdCharacteristics, cloudflare.AuthIdCharacteristics{Name: configs[i].(map[string]interface{})["name"].(string), Type: configs[i].(map[string]interface{})["type"].(string)})
	}

	return as, nil
}

func flattenAPIShieldConfiguration(characteristics []cloudflare.AuthIdCharacteristics) []interface{} {
	var flattened []interface{}
	for _, c := range characteristics {
		flattened = append(flattened, map[string]interface{}{
			"name": c.Name,
			"type": c.Type,
		})
	}
	return flattened
}
