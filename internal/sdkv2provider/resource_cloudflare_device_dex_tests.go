package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDeviceDexTest() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDeviceDexTestSchema(),
		CreateContext: resourceCloudflareDeviceDexTestCreate,
		ReadContext:   resourceCloudflareDeviceDexTestRead,
		UpdateContext: resourceCloudflareDeviceDexTestUpdate,
		DeleteContext: resourceCloudflareDeviceDexTestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDeviceDexTestImport,
		},
		Description: "Provides a Cloudflare Device Dex Test resource. Device Dex Tests allow for building location-aware device settings policies.",
	}
}

func resourceCloudflareDeviceDexTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))
	tflog.Debug(ctx, fmt.Sprintf("Reading Cloudflare Device Dex Test for Id: %+v", d.Id()))

	dexTest, err := client.GetDeviceDexTest(ctx, identifier, d.Id())

	var notFoundError *cloudflare.NotFoundError
	if errors.As(err, &notFoundError) {
		tflog.Info(ctx, fmt.Sprintf("Device Dex Test %s no longer exists", d.Id()))
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading Device Dex Test: %w", err))
	}

	d.Set("name", dexTest.Name)
	d.Set("description", dexTest.Description)
	d.Set("enabled", dexTest.Enabled)
	d.Set("interval", dexTest.Interval)
	d.Set("updated", dexTest.Updated.Format(time.RFC3339Nano))
	d.Set("created", dexTest.Created.Format(time.RFC3339Nano))
	d.Set("data", convertDeviceDexTestDataToSchema(dexTest.Data))

	return nil
}

func resourceCloudflareDeviceDexTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))

	kind := d.Get("data.0.kind").(string)

	params := cloudflare.CreateDeviceDexTestParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Interval:    d.Get("interval").(string),
		Enabled:     d.Get("enabled").(bool),
		Data: &cloudflare.DeviceDexTestData{
			"kind": kind,
			"host": d.Get("data.0.host").(string),
		},
	}

	if kind == "http" {
		(*params.Data)["method"] = d.Get("data.0.method").(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Device Dex Test with params: %+v", params))

	dexTest, err := client.CreateDeviceDexTest(ctx, identifier, params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Dex Test with provided config: %w", err))
	}

	d.SetId(dexTest.TestID)

	return resourceCloudflareDeviceDexTestRead(ctx, d, meta)
}

func resourceCloudflareDeviceDexTestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))

	updatedDeviceDexTestParams := cloudflare.UpdateDeviceDexTestParams{
		TestID:      d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Interval:    d.Get("interval").(string),
		Enabled:     d.Get("enabled").(bool),
		Data: &cloudflare.DeviceDexTestData{
			"kind":   d.Get("data.0.kind").(string),
			"host":   d.Get("data.0.host").(string),
			"method": d.Get("data.0.method").(string),
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Device Dex Test with params: %+v", updatedDeviceDexTestParams))

	dexTest, err := client.UpdateDeviceDexTest(ctx, identifier, updatedDeviceDexTestParams)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Device Dex Test for ID %q: %w", d.Id(), err))
	}
	if dexTest.TestID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Dex Test ID in update response; resource was empty"))
	}

	return resourceCloudflareDeviceDexTestRead(ctx, d, meta)
}

func resourceCloudflareDeviceDexTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))
	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Device Dex Test using ID: %s", d.Id()))

	if _, err := client.DeleteDexTest(ctx, identifier, d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting DLP Profile for ID %q: %w", d.Id(), err))
	}

	resourceCloudflareDeviceDexTestRead(ctx, d, meta)
	return nil
}

func resourceCloudflareDeviceDexTestImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID, dexTestID, err := parseDeviceDexTestsIDImport(d.Id())
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Device Dex Test: id %s for account %s", dexTestID, accountID))

	d.Set("account_id", accountID)
	d.SetId(dexTestID)

	resourceCloudflareDeviceDexTestRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func parseDeviceDexTestsIDImport(id string) (string, string, error) {
	attributes := strings.SplitN(id, "/", 2)

	if len(attributes) != 2 {
		return "", "", fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/testID\"", id)
	}

	return attributes[0], attributes[1], nil
}

func convertDeviceDexTestDataToSchema(input *cloudflare.DeviceDexTestData) []interface{} {
	kind, _ := (*input)["kind"]
	host, _ := (*input)["host"]
	method, _ := (*input)["method"]

	m := map[string]interface{}{
		"kind": kind,
		"host": host,
	}

	// The `method` field doesn't exist under all conditions. Only add it if we are converting an `http` kind of test.
	if kind == "http" && method != "" {
		m["method"] = method
	}
	return []interface{}{m}
}
