package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerKV() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerKVSchema(),
		CreateContext: resourceCloudflareWorkersKVUpdate,
		ReadContext:   resourceCloudflareWorkersKVRead,
		UpdateContext: resourceCloudflareWorkersKVUpdate,
		DeleteContext: resourceCloudflareWorkersKVDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkersKVImport,
		},
		Description: "Provides a resource to manage a Cloudflare Workers KV Pair.",
	}
}

func resourceCloudflareWorkersKVRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	namespaceID, key, err := parseId(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	value, err := client.GetWorkersKV(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.GetWorkersKVParams{
		NamespaceID: namespaceID,
		Key:         key,
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error reading workers kv"))
	}

	if value == nil {
		d.SetId("")
		return nil
	}

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("value", string(value))
	return nil
}

func resourceCloudflareWorkersKVUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	namespaceID := d.Get("namespace_id").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.WriteWorkersKVEntry(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.WriteWorkersKVEntryParams{
		NamespaceID: namespaceID,
		Key:         key,
		Value:       []byte(value),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating workers kv"))
	}

	d.SetId(fmt.Sprintf("%s/%s", namespaceID, key))

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Workers KV Namespace ID: %s", d.Id()))

	return resourceCloudflareWorkersKVRead(ctx, d, meta)
}

func resourceCloudflareWorkersKVDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	namespaceID, key, err := parseId(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Workers KV with id: %+v", d.Id()))

	_, err = client.DeleteWorkersKVEntry(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.DeleteWorkersKVEntryParams{
		NamespaceID: namespaceID,
		Key:         key,
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting workers kv"))
	}

	return nil
}

func resourceCloudflareWorkersKVImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", -1)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/namespaceID/keyName\"", d.Id())
	}

	d.Set(consts.AccountIDSchemaKey, parts[0])
	d.Set("namespace_id", parts[1])
	d.Set("key", parts[2])
	d.SetId(fmt.Sprintf("%s/%s", parts[1], parts[2]))
	resourceCloudflareWorkersKVRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func parseId(id string) (string, string, error) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("error parsing workers kv id: %s", id)
	}
	return parts[0], parts[1], nil
}
