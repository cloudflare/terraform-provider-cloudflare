package sdkv2provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		Description: heredoc.Doc(`
			Provides a resource which manages Cloudflare Logpush jobs. For
			Logpush jobs pushing to Amazon S3, Google Cloud Storage, Microsoft
			Azure or Sumo Logic, this resource cannot be automatically created.
			In order to have this automated, you must have:

			- ` + "`cloudflare_logpush_ownership_challenge`" + `: Configured to generate the challenge
			to confirm ownership of the destination.
			- Either manual inspection or another Terraform Provider to get the contents of
			the ` + "`ownership_challenge_filename`" + ` value from the` + "`cloudflare_logpush_ownership_challenge`" + ` resource.
			- ` + "`cloudflare_logpush_job`" + `: Create and manage the Logpush Job itself.
		`),
	}
}

func validateDestination(destinationConf, ownershipChallenge string) error {
	var re = regexp.MustCompile(`^((datadog|splunk|https|r2)://|s3://.+endpoint=)`)

	if ownershipChallenge == "" && !re.MatchString(destinationConf) {
		return fmt.Errorf("ownership_challenge must be set for the provided destination_conf")
	}

	return nil
}

func resourceCloudflareLogpushJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	jobID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("could not extract Logpush job from resource - invalid identifier (%s): %w", d.Id(), err))
	}

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	job, err := client.GetLogpushJob(ctx, identifier, jobID)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Could not find LogpushJob for %s with id: %q", identifier, jobID))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading logpush job %q for %s: %w", jobID, identifier, err))
	}

	if job.ID == 0 {
		d.SetId("")
		return nil
	}

	var filter string

	if job.Filter != nil {
		b, err := json.Marshal(job.Filter)
		if err != nil {
			return diag.FromErr(err)
		}

		filter = string(b)
	}

	d.Set("name", job.Name)
	d.Set("kind", job.Kind)
	d.Set("enabled", job.Enabled)
	d.Set("logpull_options", job.LogpullOptions)
	d.Set("dataset", job.Dataset)
	d.Set("destination_conf", job.DestinationConf)
	d.Set("ownership_challenge", d.Get("ownership_challenge"))
	d.Set("frequency", job.Frequency)
	d.Set("filter", filter)
	d.Set("max_upload_bytes", job.MaxUploadBytes)
	d.Set("max_upload_records", job.MaxUploadRecords)
	d.Set("max_upload_interval_seconds", job.MaxUploadIntervalSeconds)

	return nil
}

func resourceCloudflareLogpushJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to initialise identifiers"))
	}

	destConf := d.Get("destination_conf").(string)
	ownershipChallenge := d.Get("ownership_challenge").(string)

	if err := validateDestination(destConf, ownershipChallenge); err != nil {
		return diag.FromErr(fmt.Errorf("failed to validate destination configuration: %w", err))
	}

	job := cloudflare.CreateLogpushJobParams{
		Enabled:                  d.Get("enabled").(bool),
		Kind:                     d.Get("kind").(string),
		Name:                     d.Get("name").(string),
		Dataset:                  d.Get("dataset").(string),
		LogpullOptions:           d.Get("logpull_options").(string),
		DestinationConf:          destConf,
		OwnershipChallenge:       ownershipChallenge,
		Frequency:                d.Get("frequency").(string),
		MaxUploadBytes:           d.Get("max_upload_bytes").(int),
		MaxUploadRecords:         d.Get("max_upload_records").(int),
		MaxUploadIntervalSeconds: d.Get("max_upload_interval_seconds").(int),
	}

	filter := d.Get("filter")
	if filter != "" {
		var jobFilter cloudflare.LogpushJobFilters
		if err := json.Unmarshal([]byte(filter.(string)), &jobFilter); err != nil {
			return diag.FromErr(fmt.Errorf("failed to unmarshal logpush job filter"))
		}
		err := jobFilter.Where.Validate()
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to validate job filter"))
		}
		job.Filter = &jobFilter
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Logpush job for %s from struct: %+v", identifier, job))

	j, err := client.CreateLogpushJob(ctx, identifier, job)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating logpush job for %s: %w", identifier, err))
	}
	if j.ID == 0 {
		return diag.FromErr(fmt.Errorf("failed to find ID in Create response; resource was empty"))
	}

	d.SetId(strconv.Itoa(j.ID))

	tflog.Info(ctx, fmt.Sprintf("Created Cloudflare Logpush Job for %s: %s", identifier, d.Id()))

	return resourceCloudflareLogpushJobRead(ctx, d, meta)
}

func resourceCloudflareLogpushJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to initialise identifiers"))
	}

	destConf := d.Get("destination_conf").(string)
	ownershipChallenge := d.Get("ownership_challenge").(string)

	if err := validateDestination(destConf, ownershipChallenge); err != nil {
		return diag.FromErr(fmt.Errorf("failed to validate destination configuration: %w", err))
	}

	jobID, _ := strconv.Atoi(d.Id())
	job := cloudflare.UpdateLogpushJobParams{
		ID:                       jobID,
		Enabled:                  d.Get("enabled").(bool),
		Kind:                     d.Get("kind").(string),
		Name:                     d.Get("name").(string),
		Dataset:                  d.Get("dataset").(string),
		LogpullOptions:           d.Get("logpull_options").(string),
		DestinationConf:          destConf,
		OwnershipChallenge:       ownershipChallenge,
		Frequency:                d.Get("frequency").(string),
		MaxUploadBytes:           d.Get("max_upload_bytes").(int),
		MaxUploadRecords:         d.Get("max_upload_records").(int),
		MaxUploadIntervalSeconds: d.Get("max_upload_interval_seconds").(int),
	}

	filter := d.Get("filter")
	if filter != "" {
		var jobFilter cloudflare.LogpushJobFilters
		if err := json.Unmarshal([]byte(filter.(string)), &jobFilter); err != nil {
			return diag.FromErr(fmt.Errorf("failed to unmarshal logpush job filter"))
		}
		err := jobFilter.Where.Validate()
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to validate job filter"))
		}
		job.Filter = &jobFilter
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	tflog.Info(ctx, fmt.Sprintf("Updating Cloudflare Logpush job for %s from struct: %+v", identifier, job))

	err = client.UpdateLogpushJob(ctx, identifier, job)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating logpush job id %q for %s: %w", job.ID, identifier, err))
	}

	return resourceCloudflareLogpushJobRead(ctx, d, meta)
}

func resourceCloudflareLogpushJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing logpush job from resource: %w", err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Logpush job for %s with id: %+v", identifier, d.Id()))

	jobID, _ := strconv.Atoi(d.Id())
	err = client.DeleteLogpushJob(ctx, identifier, jobID)
	if err != nil {
		if strings.Contains(err.Error(), "job not found") {
			tflog.Info(ctx, fmt.Sprintf("Could not find logpush job for %s with id: %q", identifier, d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error deleting logpush job id %v for %s: %w", d.Id(), identifier, err))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareLogpushJobImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.Split(d.Id(), "/")

	if len(idAttr) != 3 || !contains([]string{"zone", "account"}, idAttr[0]) || idAttr[1] == "" || idAttr[2] == "" {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"account/accountID/jobID\" or \"zone/zoneID/jobID\"", d.Id())
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Logpush Job for %s with id %s", idAttr[1], idAttr[2]))

	if idAttr[0] == "account" {
		if err := d.Set(consts.AccountIDSchemaKey, idAttr[1]); err != nil {
			return nil, fmt.Errorf("failed to set account_id: %w", err)
		}
	} else {
		if err := d.Set(consts.ZoneIDSchemaKey, idAttr[1]); err != nil {
			return nil, fmt.Errorf("failed to set zone_id: %w", err)
		}
	}
	d.SetId(idAttr[2])

	resourceCloudflareLogpushJobRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
