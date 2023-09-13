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

func resourceCloudflareAPIShieldOperation() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAPIShieldOperationSchema(),
		CreateContext: resourceCloudflareAPIShieldOperationCreate,
		ReadContext:   resourceCloudflareAPIShieldOperationRead,
		DeleteContext: resourceCloudflareAPIShieldOperationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage an operation in API Shield Endpoint Management.
		`),
	}
}

func resourceCloudflareAPIShieldOperationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	ops, err := client.CreateAPIShieldOperations(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.CreateAPIShieldOperationsParams{
			Operations: []cloudflare.APIShieldBasicOperation{
				{
					Method:   d.Get("method").(string),
					Host:     d.Get("host").(string),
					Endpoint: d.Get("endpoint").(string),
				},
			},
		},
	)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create API Shield Operation"))
	}

	if length := len(ops); length != 1 {
		return diag.FromErr(fmt.Errorf("expected output to have 1 entry but got: %d", length))
	}

	d.SetId(ops[0].ID)
	return resourceCloudflareAPIShieldOperationRead(ctx, d, meta)
}

func resourceCloudflareAPIShieldOperationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	op, err := client.GetAPIShieldOperation(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.GetAPIShieldOperationParams{
			OperationID: d.Id(),
		},
	)

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Operation: %w", err))
	}

	if err := d.Set("method", op.Method); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("host", op.Host); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("endpoint", op.Endpoint); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(op.ID)
	return nil
}

func resourceCloudflareAPIShieldOperationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteAPIShieldOperation(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.DeleteAPIShieldOperationParams{
			OperationID: d.Id(),
		},
	)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch API Shield Operation: %w", err))
	}

	return nil
}
