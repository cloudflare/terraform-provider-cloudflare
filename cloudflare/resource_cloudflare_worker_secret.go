package cloudflare

import (
	"context"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWorkerSecretUpdate,
		Read:   resourceCloudflareWorkerSecretRead,
		Update: resourceCloudflareWorkerSecretUpdate,
		Delete: resourceCloudflareWorkerSecretDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkersSecretImport,
		},

		Schema: map[string]*schema.Schema{
			"script_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type ScriptData struct {
	// The script id will be the `name` for named script
	// or the `zone_name` for zone-scoped scripts
	ID     string
	Params cloudflare.WorkerRequestParams
}

func getScriptData(d *schema.ResourceData, client *cloudflare.API) (ScriptData, error) {
	scriptName := d.Get("name").(string)

func resourceCloudflareWorkerSecretRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	namespaceID, key := parseId(d.Id())

	value, err := client.ReadWorkersKV(context.Background(), namespaceID, key)
	if err != nil {
		return errors.Wrap(err, "error reading workers kv")
	}

	if value == nil {
		d.SetId("")
		return nil
	}

	d.Set("value", string(value))
	return nil
}

func resourceCloudflareWorkerSecretUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	namespaceID := d.Get("namespace_id").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)

	_, err := client.WriteWorkersKV(context.Background(), namespaceID, key, []byte(value))
	if err != nil {
		return errors.Wrap(err, "error creating workers kv")
	}

	d.SetId(fmt.Sprintf("%s/%s", namespaceID, key))

	log.Printf("[INFO] Cloudflare Workers KV Namespace ID: %s", d.Id())

	return resourceCloudflareWorkersKVRead(d, meta)
}

func resourceCloudflareWorkerSecretDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	namespaceID, key := parseId(d.Id())

	log.Printf("[INFO] Deleting Cloudflare Workers KV with id: %+v", d.Id())

	_, err := client.DeleteWorkersKV(context.Background(), namespaceID, key)
	if err != nil {
		return errors.Wrap(err, "error deleting workers kv")
	}

	return nil
}

func resourceCloudflareWorkersSecretImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	namespaceID, key := parseId(d.Id())

	d.Set("namespace_id", namespaceID)
	d.Set("key", key)

	resourceCloudflareWorkersKVRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
