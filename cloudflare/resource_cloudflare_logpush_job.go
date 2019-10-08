package cloudflare

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareLogpushJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareLogpushJobCreate,
		Read:   resourceCloudflareLogpushJobRead,
		Update: resourceCloudflareLogpushJobUpdate,
		Delete: resourceCloudflareLogpushJobDelete,

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logpull_options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_conf": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ownership_challenge": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getJobFromResource(d *schema.ResourceData) cloudflare.LogpushJob {
	id, err := strconv.Atoi(d.Id())

	if err != nil {
		fmt.Errorf("Could not extract Logpush job from resource - invalid identifier: %+v", id)
	}

	job := cloudflare.LogpushJob{
		ID:                 id,
		Enabled:            d.Get("enabled").(bool),
		Name:               d.Get("name").(string),
		LogpullOptions:     d.Get("logpull_options").(string),
		DestinationConf:    d.Get("destination_conf").(string),
		OwnershipChallenge: d.Get("ownership_challenge").(string),
	}
	return job
}

func resourceCloudflareLogpushJobRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	jobId, err := strconv.Atoi(d.Id())

	job, err := client.LogpushJob(d.Get("zone_id").(string), jobId)

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Printf("[INFO] Could not find LogpushJob with id: %q", jobId)
			return nil
		}
		return fmt.Errorf("error finding logpush job %q: %s", jobId, err)
	}

	if job.ID == 0 {
		d.SetId("")
		return nil
	}

	d.Set("name", job.Name)
	d.Set("enabled", job.Enabled)
	d.Set("logpull_options", job.LogpullOptions)
	d.Set("destination_conf", job.DestinationConf)
	d.Set("ownership_challenge", d.Get("ownership_challenge"))

	return nil
}

func resourceCloudflareLogpushJobCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	job := getJobFromResource(d)

	log.Printf("[DEBUG] Creating Cloudflare Logpush Job from struct: %+v", job)

	j, err := client.CreateLogpushJob(d.Get("zone_id").(string), job)

	if err != nil {
		return fmt.Errorf("error creating logpush job")
	}

	if j.ID == 0 {
		return fmt.Errorf("failed to find ID in Create response; resource was empty")
	}

	d.SetId(strconv.Itoa(j.ID))

	log.Printf("[INFO] Created Cloudflare Logpush Job ID: %s", d.Id())

	return resourceCloudflareLogpushJobRead(d, meta)

}

func resourceCloudflareLogpushJobUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	job := getJobFromResource(d)

	log.Printf("[INFO] Updating Cloudflare Logpush Job from struct: %+v", job)

	updateErr := client.UpdateLogpushJob(d.Get("zone_id").(string), job.ID, job)

	if updateErr != nil {
		return fmt.Errorf("error updating logpush job: %+v", job.ID)
	}

	return resourceCloudflareLogpushJobRead(d, meta)
}

func resourceCloudflareLogpushJobDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	job := getJobFromResource(d)

	log.Printf("[DEBUG] Deleting Cloudflare Logpush job from zone :%+v with id: %+v", d.Get("zone_id"), job.ID)

	deleteErr := client.DeleteLogpushJob(d.Get("zone_id").(string), job.ID)
	if deleteErr != nil {
		return fmt.Errorf("error deleting logpush job: %+v", job.ID)
	}

	return resourceCloudflareLogpushJobRead(d, meta)

}
