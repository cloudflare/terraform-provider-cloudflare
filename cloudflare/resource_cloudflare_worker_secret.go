package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWorkerSecretCreate,
		Read:   resourceCloudflareWorkerSecretRead,
		Delete: resourceCloudflareWorkerSecretDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkersSecretImport,
		},

		Schema: map[string]*schema.Schema{
			"script_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The name of the Worker script to associate the secret with.",
			},
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The name of the Worker secret.",
			},
			"secret_text": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "The text of the Worker secret, this cannot be read back after creation and is stored encrypted .",
			},
		},
	}
}

func resourceCloudflareWorkerSecretRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	scriptName := d.Get("script_name").(string)
	name := d.Get("name").(string)

	secrets, err := client.ListWorkersSecrets(context.Background(), scriptName)
	if err != nil {
		return errors.Wrap(err, "error reading worker secrets")
	}

	for _, secret := range secrets.Result {
		if secret.Name == name {
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceCloudflareWorkerSecretCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	scriptName := d.Get("script_name").(string)
	name := d.Get("name").(string)
	secretText := d.Get("secret_text").(string)

	request := cloudflare.WorkersPutSecretRequest{
		Name: name,
		Text: secretText,
		Type: cloudflare.WorkerSecretTextBindingType,
	}

	_, err := client.SetWorkersSecret(context.Background(), scriptName, &request)
	if err != nil {
		return errors.Wrap(err, "error creating worker secret")
	}

	d.SetId(fmt.Sprintf("%s/%s", scriptName, name))

	log.Printf("[INFO] Cloudflare Workers Secret ID: %s", d.Id())

	return nil
}

func resourceCloudflareWorkerSecretDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	scriptName, name := parseSecretId(d.Id())

	log.Printf("[INFO] Deleting Cloudflare Workers secret with id: %+v", d.Id())

	_, err := client.DeleteWorkersSecret(context.Background(), scriptName, name)
	if err != nil {
		return errors.Wrap(err, "error deleting worker secret")
	}

	return nil
}

func resourceCloudflareWorkersSecretImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	scriptName, name := parseSecretId(d.Id())

	d.Set("script_name", scriptName)
	d.Set("name", name)

	resourceCloudflareWorkerSecretRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func parseSecretId(id string) (string, string) {
	parts := strings.SplitN(id, "/", 2)
	return parts[0], parts[1]
}
