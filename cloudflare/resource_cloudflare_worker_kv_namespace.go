package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerKVNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWorkerKVNamespaceCreate,
		Read:   resourceCloudflareWorkerKVNamespaceRead,
		Update: resourceCloudflareWorkerKVNamespaceUpdate,
		Delete: resourceCloudflareWorkerKVNamespaceDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkerKVNamespaceImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func kvNamespaceRequestFromResource(d *schema.ResourceData) *cloudflare.WorkersKVNamespaceRequest {
	return &cloudflare.WorkersKVNamespaceRequest{
		Title: d.Get("title").(string),
	}
}

func kvNamespaceFromResource(d *schema.ResourceData) cloudflare.WorkersKVNamespace {
	namespace := cloudflare.WorkersKVNamespace{
		ID:    d.Id(),
		Title: d.Get("title").(string),
	}
	return namespace
}

func resourceCloudflareWorkerKVNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	req := kvNamespaceRequestFromResource(d)

	log.Printf("[INFO] Creating Cloudflare Workers KV Namespace from struct: %+v", req)

	client, err := clientWithOrg(meta)
	var resp cloudflare.WorkersKVNamespaceResponse
	if client != nil {
		resp, err = client.CreateWorkersKVNamespace(context.Background(), req)
	}

	if err != nil {
		return errors.Wrap(err, "error creating workers kv namespace")
	}

	id := resp.Result.ID

	if id == "" {
		return fmt.Errorf("failed to find id in Create response; resource was empty")
	}

	d.SetId(id)

	log.Printf("[INFO] Cloudflare Workers KV Namespace ID: %s", d.Id())

	return nil
}

func resourceCloudflareWorkerKVNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	namespace, err := resourceCloudflareWorkerKVNamespaceReadById(d.Id(), meta)
	if err != nil {
		return errors.Wrap(err, "error Reading workers kv namespaces")
	}

	if namespace.ID == "" {
		d.SetId("")
	} else {
		d.Set("title", namespace.Title)
	}

	return nil
}

func resourceCloudflareWorkerKVNamespaceReadById(namespaceId string, meta interface{}) (*cloudflare.WorkersKVNamespace, error) {
	client, err := clientWithOrg(meta)

	var resp cloudflare.ListWorkersKVNamespacesResponse
	if client != nil {
		resp, err = client.ListWorkersKVNamespaces(context.Background())
	}

	if err != nil {
		return nil, err
	}

	for _, ns := range resp.Result {
		if ns.ID == namespaceId {
			return &ns, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func resourceCloudflareWorkerKVNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	namespace := kvNamespaceFromResource(d)
	req := kvNamespaceRequestFromResource(d)

	log.Printf("[INFO] Updating Cloudflare Workers KV Namespace from struct: %+v", namespace)

	client, err := clientWithOrg(meta)
	if client != nil {
		_, err = client.UpdateWorkersKVNamespace(context.Background(), namespace.ID, req)
	}

	if err != nil {
		return errors.Wrap(err, "error updating workers kv namespace")
	}

	return nil
}

func resourceCloudflareWorkerKVNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	namespace := kvNamespaceFromResource(d)

	log.Printf("[INFO] Deleting Cloudflare Workers KV Namespace with id: %+v", namespace.ID)

	client, err := clientWithOrg(meta)
	if client != nil {
		_, err = client.DeleteWorkersKVNamespace(context.Background(), namespace.ID)
	}

	if err != nil {
		return errors.Wrap(err, "error deleting workers kv namespace")
	}

	return nil
}

func resourceCloudflareWorkerKVNamespaceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	namespace, err := resourceCloudflareWorkerKVNamespaceReadById(d.Id(), meta)
	if err != nil {
		return nil, errors.Wrap(err, "error Importing workers kv namespace")
	}

	d.Set("title", namespace.Title)

	return []*schema.ResourceData{d}, nil
}
