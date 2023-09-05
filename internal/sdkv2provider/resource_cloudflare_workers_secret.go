package sdkv2provider

import (
	"context"
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareWorkerSecretCreate,
		ReadContext:   resourceCloudflareWorkerSecretRead,
		UpdateContext: resourceCloudflareWorkerSecretCreate,
		DeleteContext: resourceCloudflareWorkerSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: heredoc.Doc("Provides a Cloudflare worker secret resource."),
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
			},
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

func resourceCloudflareWorkerSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Always return nil, as secrets cannot be read back from the Cloudflare Worker API as it currently stands.
	return nil
}

func resourceCloudflareWorkerSecretCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	scriptName := d.Get("script_name").(string)
	name := d.Get("name").(string)
	secretText := d.Get("secret_text").(string)

	params := cloudflare.SetWorkersSecretParams{
		Secret: &cloudflare.WorkersPutSecretRequest{
			Name: name,
			Text: secretText,
			Type: cloudflare.WorkerSecretTextBindingType,
		},
		ScriptName: scriptName,
	}

	_, err := client.SetWorkersSecret(context.Background(), cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating worker secret"))
	}

	d.SetId(stringChecksum(fmt.Sprintf("%s/%s", scriptName, name)))

	log.Printf("[INFO] Cloudflare Workers Secret ID: %s", d.Id())

	return nil
}

func resourceCloudflareWorkerSecretDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	scriptName := d.Get("script_name").(string)
	name := d.Get("name").(string)

	params := cloudflare.DeleteWorkersSecretParams{
		SecretName: name,
		ScriptName: scriptName,
	}

	log.Printf("[INFO] Deleting Cloudflare Workers secret with id: %+v", d.Id())

	_, err := client.DeleteWorkersSecret(context.Background(), cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting worker secret"))
	}

	return nil
}
