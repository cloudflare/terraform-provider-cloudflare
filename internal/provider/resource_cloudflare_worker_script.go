package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerScript() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerScriptSchema(),
		CreateContext: resourceCloudflareWorkerScriptCreate,
		ReadContext:   resourceCloudflareWorkerScriptRead,
		UpdateContext: resourceCloudflareWorkerScriptUpdate,
		DeleteContext: resourceCloudflareWorkerScriptDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkerScriptImport,
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

	params := cloudflare.WorkerRequestParams{
		ScriptName: scriptName,
	}

	return ScriptData{
		scriptName,
		params,
	}, nil
}

type ScriptBindings map[string]cloudflare.WorkerBinding

func getWorkerScriptBindings(ctx context.Context, scriptName string, client *cloudflare.API) (ScriptBindings, error) {
	resp, err := client.ListWorkerBindings(ctx, &cloudflare.WorkerRequestParams{ScriptName: scriptName})
	if err != nil {
		return nil, fmt.Errorf("cannot list script bindings: %w", err)
	}

	bindings := make(ScriptBindings, len(resp.BindingList))

	for _, b := range resp.BindingList {
		bindings[b.Name] = b.Binding
	}

	return bindings, nil
}

func parseWorkerBindings(d *schema.ResourceData, bindings ScriptBindings) {
	for _, rawData := range d.Get("kv_namespace_binding").(*schema.Set).List() {
		data := rawData.(map[string]interface{})
		bindings[data["name"].(string)] = cloudflare.WorkerKvNamespaceBinding{
			NamespaceID: data["namespace_id"].(string),
		}
	}

	for _, rawData := range d.Get("plain_text_binding").(*schema.Set).List() {
		data := rawData.(map[string]interface{})
		bindings[data["name"].(string)] = cloudflare.WorkerPlainTextBinding{
			Text: data["text"].(string),
		}
	}

	for _, rawData := range d.Get("secret_text_binding").(*schema.Set).List() {
		data := rawData.(map[string]interface{})
		bindings[data["name"].(string)] = cloudflare.WorkerSecretTextBinding{
			Text: data["text"].(string),
		}
	}

	for _, rawData := range d.Get("webassembly_binding").(*schema.Set).List() {
		data := rawData.(map[string]interface{})
		module := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data["module"].(string)))
		bindings[data["name"].(string)] = cloudflare.WorkerWebAssemblyBinding{
			Module: module,
		}
	}

	for _, rawData := range d.Get("service_binding").(*schema.Set).List() {
		data := rawData.(map[string]interface{})
		bindings[data["name"].(string)] = cloudflare.WorkerServiceBinding{
			Service:     data["service"].(string),
			Environment: data["environment"].(*string),
		}
	}
}

func resourceCloudflareWorkerScriptCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return diag.FromErr(err)
	}

	// make sure that the worker does not already exist
	r, _ := client.DownloadWorker(ctx, &scriptData.Params)
	if r.WorkerScript.Script != "" {
		return diag.FromErr(fmt.Errorf("script already exists"))
	}

	scriptBody := d.Get("content").(string)
	if scriptBody == "" {
		return diag.FromErr(fmt.Errorf("script content cannot be empty"))
	}

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare Worker Script from struct: %+v", &scriptData.Params))

	bindings := make(ScriptBindings)

	parseWorkerBindings(d, bindings)

	scriptParams := cloudflare.WorkerScriptParams{
		Script:   scriptBody,
		Bindings: bindings,
	}

	_, err = client.UploadWorkerWithBindings(ctx, &scriptData.Params, &scriptParams)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating worker script"))
	}

	d.SetId(scriptData.ID)

	return nil
}

func resourceCloudflareWorkerScriptRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.DownloadWorker(ctx, &scriptData.Params)
	if err != nil {
		// If the resource is deleted, we should set the ID to "" and not
		// return an error according to the terraform spec
		if strings.Contains(err.Error(), "HTTP status 404") {
			d.SetId("")
			return nil
		}

		return diag.FromErr(errors.Wrap(err,
			fmt.Sprintf("Error reading worker script from API for resource %+v", &scriptData.Params)))
	}

	existingBindings := make(ScriptBindings)

	parseWorkerBindings(d, existingBindings)

	bindings, err := getWorkerScriptBindings(ctx, d.Get("name").(string), client)
	if err != nil {
		return diag.FromErr(err)
	}

	kvNamespaceBindings := &schema.Set{F: schema.HashResource(kvNamespaceBindingResource)}
	plainTextBindings := &schema.Set{F: schema.HashResource(plainTextBindingResource)}
	secretTextBindings := &schema.Set{F: schema.HashResource(secretTextBindingResource)}
	webAssemblyBindings := &schema.Set{F: schema.HashResource(webAssemblyBindingResource)}
	serviceBindings := &schema.Set{F: schema.HashResource(serviceBindingResource)}

	for name, binding := range bindings {
		switch v := binding.(type) {
		case cloudflare.WorkerKvNamespaceBinding:
			kvNamespaceBindings.Add(map[string]interface{}{
				"name":         name,
				"namespace_id": v.NamespaceID,
			})
		case cloudflare.WorkerPlainTextBinding:
			plainTextBindings.Add(map[string]interface{}{
				"name": name,
				"text": v.Text,
			})
		case cloudflare.WorkerSecretTextBinding:
			value := v.Text
			switch v := existingBindings[name].(type) {
			case cloudflare.WorkerSecretTextBinding:
				value = v.Text
			}
			secretTextBindings.Add(map[string]interface{}{
				"name": name,
				"text": value,
			})
		case cloudflare.WorkerWebAssemblyBinding:
			module, err := ioutil.ReadAll(v.Module)
			if err != nil {
				return diag.FromErr(errors.Wrap(err, fmt.Sprintf("cannot read contents of wasm bindings (%s)", name)))
			}
			webAssemblyBindings.Add(map[string]interface{}{
				"name":   name,
				"module": base64.StdEncoding.EncodeToString(module),
			})
		case cloudflare.WorkerServiceBinding:
			serviceBindings.Add(map[string]interface{}{
				"name":        name,
				"service":     v.Service,
				"environment": v.Environment,
			})
		}
	}

	if err := d.Set("content", r.Script); err != nil {
		return diag.FromErr(fmt.Errorf("cannot set content: %w", err))
	}

	if err := d.Set("kv_namespace_binding", kvNamespaceBindings); err != nil {
		return diag.FromErr(fmt.Errorf("cannot set kv namespace bindings (%s): %w", d.Id(), err))
	}

	if err := d.Set("plain_text_binding", plainTextBindings); err != nil {
		return diag.FromErr(fmt.Errorf("cannot set plain text bindings (%s): %w", d.Id(), err))
	}

	if err := d.Set("secret_text_binding", secretTextBindings); err != nil {
		return diag.FromErr(fmt.Errorf("cannot set secret text bindings (%s): %w", d.Id(), err))
	}

	if err := d.Set("webassembly_binding", webAssemblyBindings); err != nil {
		return diag.FromErr(fmt.Errorf("cannot set webassembly bindings (%s): %w", d.Id(), err))
	}

	if err := d.Set("service_binding", serviceBindings); err != nil {
		return diag.FromErr(fmt.Errorf("cannot set service bindings (%s): %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareWorkerScriptUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return diag.FromErr(err)
	}

	scriptBody := d.Get("content").(string)
	if scriptBody == "" {
		return diag.FromErr(fmt.Errorf("script content cannot be empty"))
	}

	tflog.Info(ctx, fmt.Sprintf("Updating Cloudflare Worker Script from struct: %+v", &scriptData.Params))

	bindings := make(ScriptBindings)

	parseWorkerBindings(d, bindings)

	scriptParams := cloudflare.WorkerScriptParams{
		Script:   scriptBody,
		Bindings: bindings,
	}

	_, err = client.UploadWorkerWithBindings(ctx, &scriptData.Params, &scriptParams)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating worker script"))
	}

	return nil
}

func resourceCloudflareWorkerScriptDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Worker Script from struct: %+v", &scriptData.Params))

	_, err = client.DeleteWorker(ctx, &scriptData.Params)
	if err != nil {
		// If the resource is already deleted, we should return without an error
		// according to the terraform spec
		if strings.Contains(err.Error(), "HTTP status 404") {
			return nil
		}

		return diag.FromErr(errors.Wrap(err, "error deleting worker script"))
	}

	return nil
}

func resourceCloudflareWorkerScriptImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	scriptID := d.Id()
	_ = d.Set("name", scriptID)

	_ = resourceCloudflareWorkerScriptRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
