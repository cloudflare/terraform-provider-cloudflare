package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkersKVNamespace() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkersKVNamespaceSchema(),
		CreateContext: resourceCloudflareWorkersKVNamespaceCreate,
		ReadContext:   resourceCloudflareWorkersKVNamespaceRead,
		UpdateContext: resourceCloudflareWorkersKVNamespaceUpdate,
		DeleteContext: resourceCloudflareWorkersKVNamespaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkersKVNamespaceImport,
		},
		Description: "Provides the ability to manage Cloudflare Workers KV Namespace features.",
	}
}

func resourceCloudflareWorkersKVNamespaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	req := cloudflare.CreateWorkersKVNamespaceParams{
		Title: d.Get("title").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("[Info] Creating Cloudflare Workers KV Namespace from struct: %+v", req))

	r, err := client.CreateWorkersKVNamespace(ctx, cloudflare.AccountIdentifier(accountID), req)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating workers kv namespace"))
	}

	if r.Result.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.Result.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Workers KV Namespace ID: %s", d.Id()))

	return resourceCloudflareWorkersKVNamespaceRead(ctx, d, meta)
}

func resourceCloudflareWorkersKVNamespaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	namespaceID := d.Id()

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	resp, _, err := client.ListWorkersKVNamespaces(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWorkersKVNamespacesParams{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error reading workers kv namespaces"))
	}

	var namespace cloudflare.WorkersKVNamespace
	for _, r := range resp {
		if r.ID == namespaceID {
			namespace = r
			break
		}
	}

	if namespace.ID == "" {
		d.SetId("")
		return nil
	}

	d.Set(consts.AccountIDSchemaKey, accountID)

	return nil
}

func resourceCloudflareWorkersKVNamespaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.UpdateWorkersKVNamespace(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateWorkersKVNamespaceParams{
		NamespaceID: d.Id(),
		Title:       d.Get("title").(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating workers kv namespace"))
	}

	return resourceCloudflareWorkersKVNamespaceRead(ctx, d, meta)
}

func resourceCloudflareWorkersKVNamespaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Workers KV Namespace with id: %+v", d.Id()))

	_, err := client.DeleteWorkersKVNamespace(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting workers kv namespace"))
	}

	d.SetId("")
	return nil
}

func resourceCloudflareWorkersKVNamespaceImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/namespaceID\"", d.Id())
	}

	accountID, namespaceID := attributes[0], attributes[1]
	d.Set(consts.AccountIDSchemaKey, accountID)

	client := meta.(*cloudflare.API)

	namespaces, _, err := client.ListWorkersKVNamespaces(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWorkersKVNamespacesParams{})
	var title string

	for _, n := range namespaces {
		if n.ID == namespaceID {
			title = n.Title
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error finding workers kv namespace %q: %w", namespaceID, err)
	}

	d.Set("title", title)
	d.SetId(namespaceID)

	resourceCloudflareWorkersKVNamespaceRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
