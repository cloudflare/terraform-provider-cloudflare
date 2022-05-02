package cloudflare

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareLogpushJob() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareLogpushJobSchema(),
		CreateContext: resourceCloudflareLogpushJobCreate,
		ReadContext:   resourceCloudflareLogpushJobRead,
		UpdateContext: resourceCloudflareLogpushJobUpdate,
		DeleteContext: resourceCloudflareLogpushJobDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareLogpushJobImport,
		},
	}
}

func getJobFromResource(d *schema.ResourceData) (cloudflare.LogpushJob, *AccessIdentifier, error) {
	id := 0

	identifier, err := initIdentifier(d)
	if err != nil {
		return cloudflare.LogpushJob{}, identifier, err
	}

	if d.Id() != "" {
		var err error
		if id, err = strconv.Atoi(d.Id()); err != nil {
			return cloudflare.LogpushJob{}, identifier, fmt.Errorf("could not extract Logpush job from resource - invalid identifier (%s): %w", d.Id(), err)
		}
	}

	destConf := d.Get("destination_conf").(string)
	ownershipChallenge := d.Get("ownership_challenge").(string)
	var re = regexp.MustCompile(`^((datadog|splunk|https)://|s3://.+endpoint=)`)

	if ownershipChallenge == "" && !re.MatchString(destConf) {
		return cloudflare.LogpushJob{}, identifier, fmt.Errorf("ownership_challenge must be set for the provided destination_conf")
	}

	job := cloudflare.LogpushJob{
		ID:                 id,
		Enabled:            d.Get("enabled").(bool),
		Name:               d.Get("name").(string),
		Dataset:            d.Get("dataset").(string),
		LogpullOptions:     d.Get("logpull_options").(string),
		DestinationConf:    destConf,
		OwnershipChallenge: ownershipChallenge,
	}

	return job, identifier, nil
}

func resourceCloudflareLogpushJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	jobID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("could not extract Logpush job from resource - invalid identifier (%s): %w", d.Id(), err))
	}

	var job cloudflare.LogpushJob
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}
	if identifier.Type == AccountType {
		job, err = client.GetAccountLogpushJob(context.Background(), identifier.Value, jobID)
	} else {
		job, err = client.GetZoneLogpushJob(context.Background(), identifier.Value, jobID)
	}
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Printf("[INFO] Could not find LogpushJob for %s with id: %q", identifier, jobID)
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading logpush job %q for %s: %w", jobID, identifier, err))
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

func resourceCloudflareLogpushJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	job, identifier, err := getJobFromResource(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	log.Printf("[DEBUG] Creating Cloudflare Logpush job for %s from struct: %+v", identifier, job)

	var j *cloudflare.LogpushJob
	if identifier.Type == AccountType {
		j, err = client.CreateAccountLogpushJob(context.Background(), identifier.Value, job)
	} else {
		j, err = client.CreateZoneLogpushJob(context.Background(), identifier.Value, job)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating logpush job for %s: %w", identifier, err))
	}
	if j.ID == 0 {
		return diag.FromErr(fmt.Errorf("failed to find ID in Create response; resource was empty"))
	}

	d.SetId(strconv.Itoa(j.ID))

	log.Printf("[INFO] Created Cloudflare Logpush Job for %s: %s", identifier, d.Id())

	return resourceCloudflareLogpushJobRead(ctx, d, meta)
}

func resourceCloudflareLogpushJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	job, identifier, err := getJobFromResource(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	log.Printf("[INFO] Updating Cloudflare Logpush job for %s from struct: %+v", identifier, job)

	if identifier.Type == AccountType {
		err = client.UpdateAccountLogpushJob(context.Background(), identifier.Value, job.ID, job)
	} else {
		err = client.UpdateZoneLogpushJob(context.Background(), identifier.Value, job.ID, job)
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating logpush job id %q for %s: %w", job.ID, identifier, err))
	}

	return resourceCloudflareLogpushJobRead(ctx, d, meta)
}

func resourceCloudflareLogpushJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	job, identifier, err := getJobFromResource(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	log.Printf("[DEBUG] Deleting Cloudflare Logpush job for %s with id: %+v", identifier, job.ID)

	if identifier.Type == AccountType {
		err = client.DeleteAccountLogpushJob(context.Background(), identifier.Value, job.ID)
	} else {
		err = client.DeleteZoneLogpushJob(context.Background(), identifier.Value, job.ID)
	}
	if err != nil {
		if strings.Contains(err.Error(), "job not found") {
			log.Printf("[INFO] Could not find logpush job for %s with id: %q", identifier, job.ID)
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error deleting logpush job id %v for %s: %w", job.ID, identifier, err))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareLogpushJobImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.Split(d.Id(), "/")

	if len(idAttr) != 3 || (AccessIdentifierType(idAttr[0]) != AccountType && AccessIdentifierType(idAttr[0]) != ZoneType) || idAttr[1] == "" || idAttr[2] == "" {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"account/accountID/jobID\" or \"zone/zoneID/jobID\"", d.Id())
	}

	identifier := AccessIdentifier{
		Type:  AccessIdentifierType(idAttr[0]),
		Value: idAttr[1],
	}
	logpushJobID := idAttr[2]

	log.Printf("[DEBUG] Importing Cloudflare Logpush Job for %s with id %s", identifier, logpushJobID)

	if identifier.Type == AccountType {
		if err := d.Set("account_id", identifier.Value); err != nil {
			return nil, fmt.Errorf("failed to set account_id: %w", err)
		}
	} else {
		if err := d.Set("zone_id", identifier.Value); err != nil {
			return nil, fmt.Errorf("failed to set zone_id: %w", err)
		}
	}
	d.SetId(logpushJobID)

	resourceCloudflareLogpushJobRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
