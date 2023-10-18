package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareAPIShieldSchemas() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAPIShieldSchemaSchema(),
		CreateContext: resourceCloudflareAPIShieldSchemaCreate,
		ReadContext:   resourceCloudflareAPIShieldSchemaRead,
		DeleteContext: resourceCloudflareAPIShieldSchemaDelete,
		UpdateContext: resourceCloudflareAPIShieldSchemaUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage a schema in API Shield Schema Validation 2.0.
		`),
	}
}

func resourceCloudflareAPIShieldSchemaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	sch, err := client.CreateAPIShieldSchema(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.CreateAPIShieldSchemaParams{
			Name:              d.Get("name").(string),
			Kind:              d.Get("kind").(string),
			Source:            strings.NewReader(d.Get("source").(string)),
			ValidationEnabled: cloudflare.BoolPtr(d.Get("validation_enabled").(bool)),
		},
	)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create cloudflare_api_shield_schema"))
	}

	// log warnings that occurred during creation
	for _, w := range sch.Events.Warnings {
		tflog.Warn(ctx, fmt.Sprintf("cloudflare_api_shield_schema: warning encountered when creating schema: %s", w))
	}

	d.SetId(sch.Schema.ID)

	return resourceCloudflareAPIShieldSchemaRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldSchemaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	sch, err := client.GetAPIShieldSchema(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.GetAPIShieldSchemaParams{
			SchemaID: d.Id(),
		},
	)

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Schema: %w", err))
	}

	if err := d.Set("name", sch.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("kind", sch.Kind); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("source", sch.Source); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("validation_enabled", sch.ValidationEnabled); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sch.ID)
	return nil
}

func resourceCloudflareAPIShieldSchemaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	_, err := client.UpdateAPIShieldSchema(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.UpdateAPIShieldSchemaParams{
			SchemaID:          d.Id(),
			ValidationEnabled: cloudflare.BoolPtr(d.Get("validation_enabled").(bool)),
		},
	)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create API Shield Schema"))
	}

	return resourceCloudflareAPIShieldSchemaRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldSchemaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteAPIShieldSchema(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.DeleteAPIShieldSchemaParams{
			SchemaID: d.Id(),
		},
	)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Schema: %w", err))
	}

	return nil
}
