package sdkv2provider

import (
	"context"
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"strings"
)

func resourceCloudflareWorkerSecret() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerSecretSchema(),
		CreateContext: resourceCloudflareWorkerSecretCreate,
		ReadContext:   resourceCloudflareWorkerSecretRead,
		UpdateContext: resourceCloudflareWorkerSecretCreate,
		DeleteContext: resourceCloudflareWorkerSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: heredoc.Doc("Provides a Cloudflare Worker secret resource."),
	}
}

func resourceCloudflareWorkerSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID, scriptName, secretName, err := resourceCloudflareWorkerSecretParseId(d.Id())
	secrets, err := client.ListWorkersSecrets(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWorkersSecretsParams{
		ScriptName: scriptName,
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("Error listing worker secrets")))
	}

	for _, secret := range secrets.Result {
		tflog.Info(ctx, fmt.Sprintf("Found secret %s", secret.Name))
		if secret.Name == secretName {
			d.Set("name", secretName)
			d.Set(consts.AccountIDSchemaKey, accountID)
			d.Set("script_name", scriptName)
			return nil
		}
	}

	return diag.Errorf(fmt.Sprintf("worker secret %s not found for script %s", secretName, scriptName))
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

	d.SetId(fmt.Sprintf("%s/%s/%s", accountID, scriptName, name))

	tflog.Info(ctx, fmt.Sprintf("Created Cloudflare Workers secret with ID: %s", d.Id()))

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

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Workers secret with id: %s", d.Id()))

	_, err := client.DeleteWorkersSecret(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting worker secret"))
	}

	return nil
}

func resourceCloudflareWorkerSecretParseId(id string) (string, string, string, error) {
	attributes := strings.SplitN(id, "/", 3)

	if len(attributes) != 3 {
		return "", "", "", fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/scriptName/secretName"`, id)
	}

	return attributes[0], attributes[1], attributes[2], nil
}
