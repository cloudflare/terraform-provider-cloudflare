package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkersKVNamespace() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareWorkersKVNamespaceSchema(),
		CreateContext: resourceCloudflareWorkersKVNamespaceCreate,
		ReadContext: resourceCloudflareWorkersKVNamespaceRead,
		UpdateContext: resourceCloudflareWorkersKVNamespaceUpdate,
		DeleteContext: resourceCloudflareWorkersKVNamespaceDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkersKVNamespaceImport,
		},
	}
}

func resourceCloudflareWorkersKVNamespaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	req := &cloudflare.WorkersKVNamespaceRequest{
		Title: d.Get("title").(string),
	}

	log.Printf("[Info] Creating Cloudflare Workers KV Namespace from struct: %+v", req)

	r, err := client.CreateWorkersKVNamespace(context.Background(), req)
	if err != nil {
		return err.Wrap(err, "error creating workers kv namespace")
	}

	if r.Result.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.Result.ID)

	log.Printf("[INFO] Cloudflare Workers KV Namespace ID: %s", d.Id())

	return nil
}

func resourceCloudflareWorkersKVNamespaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	namespaceID := d.Id()

	resp, err := client.ListWorkersKVNamespaces(context.Background())
	if err != nil {
		return err.Wrap(err, "error reading workers kv namespaces")
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

	log.Printf("[INFO] Updating Cloudflare Workers KV Namespace from struct %+v", namespace)

	_, err := client.UpdateWorkersKVNamespace(context.Background(), d.Id(), namespace)
	if err != nil {
		return err.Wrap(err, "error updating workers kv namespace")
	}

	return nil
}

func resourceCloudflareWorkersKVNamespaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting Cloudflare Workers KV Namespace with id: %+v", d.Id())

	_, err := client.DeleteWorkersKVNamespace(context.Background(), d.Id())
	if err != nil {
		return err.Wrap(err, "error deleting workers kv namespace")
	}

	return nil
}

func resourceCloudflareWorkersKVNamespaceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	namespaces, err := client.ListWorkersKVNamespaces(context.Background())
	var title string

	for _, n := range namespaces {
		if n.ID == d.Id() {
			title = n.Title
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error finding workers kv namespace %q: %s", d.Id(), err)
	}

	d.Set("title", title)
	d.SetId(d.Id())

	return []*schema.ResourceData{d}, nil
}
