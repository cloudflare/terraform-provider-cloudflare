package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWorkerScriptCreate,
		Read:   resourceCloudflareWorkerScriptRead,
		Update: resourceCloudflareWorkerScriptUpdate,
		Delete: resourceCloudflareWorkerScriptDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkerScriptImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"content": {
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

	var params cloudflare.WorkerRequestParams

	params = cloudflare.WorkerRequestParams{
		ScriptName: scriptName,
	}

	return ScriptData{
		scriptName,
		params,
	}, nil
}

func resourceCloudflareWorkerScriptCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return err
	}

	// make sure that the worker does not already exist
	r, _ := client.DownloadWorker(&scriptData.Params)
	if r.WorkerScript.Script != "" {
		return fmt.Errorf("script already exists.")
	}

	scriptBody := d.Get("content").(string)
	if scriptBody == "" {
		return fmt.Errorf("script content cannot be empty")
	}

	log.Printf("[INFO] Creating Cloudflare Worker Script from struct: %+v", &scriptData.Params)

	_, err = client.UploadWorker(&scriptData.Params, scriptBody)
	if err != nil {
		return errors.Wrap(err, "error creating worker script")
	}

	d.SetId(scriptData.ID)

	return nil
}

func resourceCloudflareWorkerScriptRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return err
	}

	r, err := client.DownloadWorker(&scriptData.Params)
	if err != nil {
		// If the resource is deleted, we should set the ID to "" and not
		// return an error according to the terraform spec
		if strings.Contains(err.Error(), "HTTP status 404") {
			d.SetId("")
			return nil
		}

		return errors.Wrap(err,
			fmt.Sprintf("Error reading worker script from API for resouce %+v", &scriptData.Params))
	}

	d.Set("content", r.Script)
	return nil
}

func resourceCloudflareWorkerScriptUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return err
	}

	scriptBody := d.Get("content").(string)
	if scriptBody == "" {
		return fmt.Errorf("script content cannot be empty")
	}

	log.Printf("[INFO] Updating Cloudflare Worker Script from struct: %+v", &scriptData.Params)

	_, err = client.UploadWorker(&scriptData.Params, scriptBody)
	if err != nil {
		return errors.Wrap(err, "error updating worker script")
	}

	return nil
}

func resourceCloudflareWorkerScriptDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	scriptData, err := getScriptData(d, client)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting Cloudflare Worker Script from struct: %+v", &scriptData.Params)

	_, err = client.DeleteWorker(&scriptData.Params)
	if err != nil {
		// If the resource is already deleted, we should return without an error
		// according to the terraform spec
		if strings.Contains(err.Error(), "HTTP status 404") {
			return nil
		}

		return errors.Wrap(err, "error deleting worker script")
	}

	return nil
}

func resourceCloudflareWorkerScriptImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	scriptID := d.Id()
	d.Set("name", scriptID)

	resourceCloudflareWorkerScriptRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
