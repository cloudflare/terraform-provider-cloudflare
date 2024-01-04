package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerSecret() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerSecretSchema(),
		CreateContext: resourceCloudflareWorkerSecretCreate,
		ReadContext:   resourceCloudflareWorkerSecretRead,
		UpdateContext: resourceCloudflareWorkerSecretCreate,
		DeleteContext: resourceCloudflareWorkerSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkerSecretImport,
		},
		Description: heredoc.Doc("Provides a Cloudflare Worker secret resource."),
	}
}

func resourceCloudflareWorkerSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	secrets, err := client.ListWorkersSecrets(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWorkersSecretsParams{
		ScriptName: d.Get("script_name").(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error listing worker secrets")))
	}

	for _, secret := range secrets.Result {
		if secret.Name == d.Get("name") {
			return nil
		}
	}

	return diag.Errorf(fmt.Sprintf("worker secret %q not found for script %q", d.Get("name"), d.Get("script_name")))
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

	_, err := client.SetWorkersSecret(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating worker secret"))
	}

	d.SetId(stringChecksum(fmt.Sprintf("%s/%s/%s", accountID, scriptName, name)))

	tflog.Info(ctx, fmt.Sprintf("created Cloudflare Workers secret with ID: %s", d.Id()))

	return resourceCloudflareWorkerSecretRead(ctx, d, meta)
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

	tflog.Info(ctx, fmt.Sprintf("deleting Cloudflare Workers secret with id: %s", d.Id()))

	_, err := client.DeleteWorkersSecret(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting worker secret"))
	}

	return nil
}

func resourceCloudflareWorkerSecretImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 3 {
		return []*schema.ResourceData{nil}, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/scriptName/secretName"`, d.Id())
	}

	accountID, scriptName, secretName := attributes[0], attributes[1], attributes[2]

	d.SetId(stringChecksum(fmt.Sprintf("%s/%s/%s", accountID, scriptName, secretName)))
	d.Set("name", secretName)
	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("script_name", scriptName)

	resourceCloudflareWorkerSecretRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
