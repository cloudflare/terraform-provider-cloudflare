package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

// TODO: add expiring key support when cloudflare-go updates.
func resourceCloudflareWorkerKVPair() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWorkerKVPairUpdate,
		Read:   resourceCloudflareWorkerKVPairRead,
		Update: resourceCloudflareWorkerKVPairUpdate,
		Delete: resourceCloudflareWorkerKVPairDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkerKVPairImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCloudflareWorkerKVPairGet(namespace string, key string, meta interface{}) ([]byte, error) {
	client, err := clientWithOrg(meta)
	if err != nil {
		return nil, err
	}

	return client.ReadWorkersKV(context.Background(), namespace, key)
}

func resourceCloudflareWorkerKVPairRead(d *schema.ResourceData, meta interface{}) error {
	namespace := d.Get("namespace").(string)
	key := d.Get("key").(string)

	val, err := resourceCloudflareWorkerKVPairGet(namespace, key, meta)
	if err != nil {
		return errors.Wrap(err, "error Reading workers kv pair")
	}

	// Since the resource schema does not support byte arrays,
	// encode the value as a string.
	d.Set("value", string(val))

	return nil
}

func resourceCloudflareWorkerKVPairUpdate(d *schema.ResourceData, meta interface{}) error {
	namespace := d.Get("namespace").(string)
	key := d.Get("key").(string)
	value := []byte(d.Get("value").(string))

	log.Printf("[INFO] Updating Cloudflare Workers KV Pair of %+v/%+v = %+v", namespace, key, value)

	client, err := clientWithOrg(meta)
	if client != nil {
		_, err = client.WriteWorkersKV(context.Background(), namespace, key, value)
	}

	if err != nil {
		return err
	}

	d.SetId(namespace + "/" + key)

	return nil
}

func resourceCloudflareWorkerKVPairDelete(d *schema.ResourceData, meta interface{}) error {
	namespace := d.Get("namespace").(string)
	key := d.Get("key").(string)

	log.Printf("[INFO] Deleting Cloudflare Workers KV Pair with key: %+v", key)

	client, err := clientWithOrg(meta)
	if client != nil {
		_, err = client.DeleteWorkersKV(context.Background(), namespace, key)
	}

	if err != nil {
		return errors.Wrap(err, "error Deleting workers kv pair")
	}

	return nil
}

func resourceCloudflareWorkerKVPairImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	attr := strings.SplitN(d.Id(), "/", 2)
	var namespace string
	var key string
	if len(attr) == 2 {
		namespace = attr[0]
		key = attr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"namespaceId/key\"", d.Id())
	}

	d.Set("namespace", namespace)
	d.Set("key", key)

	return []*schema.ResourceData{d}, nil
}
