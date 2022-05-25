package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
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
	}
}

func resourceCloudflareWorkersKVNamespaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	req := &cloudflare.WorkersKVNamespaceRequest{
		Title: d.Get("title").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("[Info] Creating Cloudflare Workers KV Namespace from struct: %+v", req))

	r, err := client.CreateWorkersKVNamespace(ctx, req)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating workers kv namespace"))
	}

	if r.Result.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.Result.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Workers KV Namespace ID: %s", d.Id()))

	return nil
}

func resourceCloudflareWorkersKVNamespaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	namespaceID := d.Id()

	resp, err := client.ListWorkersKVNamespaces(ctx)
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

	return nil
}

func resourceCloudflareWorkersKVNamespaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	namespace := &cloudflare.WorkersKVNamespaceRequest{
		Title: d.Get("title").(string),
	}

	tflog.Info(ctx, fmt.Sprintf("Updating Cloudflare Workers KV Namespace from struct %+v", namespace))

	_, err := client.UpdateWorkersKVNamespace(ctx, d.Id(), namespace)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating workers kv namespace"))
	}

	return nil
}

func resourceCloudflareWorkersKVNamespaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Workers KV Namespace with id: %+v", d.Id()))

	_, err := client.DeleteWorkersKVNamespace(ctx, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting workers kv namespace"))
	}

	return nil
}

func resourceCloudflareWorkersKVNamespaceImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	namespaces, err := client.ListWorkersKVNamespaces(ctx)
	var title string

	for _, n := range namespaces {
		if n.ID == d.Id() {
			title = n.Title
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error finding workers kv namespace %q: %w", d.Id(), err)
	}

	d.Set("title", title)
	d.SetId(d.Id())

	return []*schema.ResourceData{d}, nil
}
