package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
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
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// zone is used for single-script, name is used for multi-script
				ConflictsWith: []string{"name"},
			},

			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// zone is used for single-script, name is used for multi-script
				ConflictsWith: []string{"zone"},
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
	zoneName := d.Get("zone").(string)
	scriptName := d.Get("name").(string)
	if zoneName == "" && scriptName == "" {
		return ScriptData{}, fmt.Errorf("either `zone` or `name` field must be set")
	}

	var params cloudflare.WorkerRequestParams
	var id string

	if scriptName != "" {
		params = cloudflare.WorkerRequestParams{
			ScriptName: scriptName,
		}
		id = "name:" + scriptName
	} else {
		zoneID, err := client.ZoneIDByName(zoneName)
		if err != nil {
			return ScriptData{}, fmt.Errorf("error finding zone %q: %s", zoneName, err)
		}
		d.Set("zone_id", zoneID)
		params = cloudflare.WorkerRequestParams{
			ZoneID: zoneID,
		}
		id = "zone:" + zoneName
	}

	return ScriptData{
		id,
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
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), ":", 2)
	var scriptType string
	var scriptId string
	if len(idAttr) == 2 {
		scriptType = idAttr[0]
		scriptId = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"scriptType:scriptId\"", d.Id())
	}

	if scriptType == "name" {
		d.Set("name", scriptId)
	} else if scriptType == "zone" {
		zoneName := scriptId
		zoneId, err := client.ZoneIDByName(zoneName)
		if err != nil {
			return nil, fmt.Errorf("error finding zone %q: %s", zoneName, err)
		}
		d.Set("zone", zoneName)
		d.Set("zone_id", zoneId)
	} else {
		return nil, fmt.Errorf("invalid scriptType (\"%s\") specified, should be either \"name\" or \"zone\"", scriptType)
	}

	return []*schema.ResourceData{d}, nil
}
