package logpush_job_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	cfold "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_logpush_job", &resource.Sweeper{
		Name: "cloudflare_logpush_job",
		F:    testSweepCloudflareLogpushJob,
	})
}

func testSweepCloudflareLogpushJob(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	jobs, err := client.ListLogpushJobs(ctx, cfold.AccountIdentifier(accountID), cfold.ListLogpushJobsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Logpush Jobs: %s", err))
		return err
	}

	if len(jobs) == 0 {
		tflog.Debug(ctx, "[DEBUG] No Cloudflare Logpush Jobs to sweep")
		return nil
	}

	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Found %d Cloudflare Logpush Jobs to sweep", len(jobs)))

	// Track deletion results
	deleted := 0
	failed := 0

	for _, job := range jobs {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Logpush Job ID: %d, Name: %s", job.ID, job.Name))

		err := client.DeleteLogpushJob(ctx, cfold.AccountIdentifier(accountID), job.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Logpush Job %d (%s): %v", job.ID, job.Name, err))
			failed++
			// Continue with other jobs
		} else {
			tflog.Info(ctx, fmt.Sprintf("Successfully deleted Logpush Job %d (%s)", job.ID, job.Name))
			deleted++
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Logpush Job sweep completed: %d deleted, %d failed", deleted, failed))
	return nil
}

func jsonMarshal(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func toString(v any) string {
	return fmt.Sprintf("%v", v)
}

func unquote(s string) string {
	s, _ = strconv.Unquote(`"` + s + `"`)
	return s
}

// logpushJobConfig is a simplified struct for testing.
type logpushJobConfig struct {
	accountID                string
	zoneID                   string
	dataset                  string
	destinationConf          string
	enabled                  bool
	name                     string
	filter                   string
	kind                     string
	maxUploadBytes           int
	maxUploadRecords         int
	maxUploadIntervalSeconds int
	frequency                string
	logpullOptions           string
	outputOptions            *logpushJobConfigOutputOptions
}
type logpushJobConfigOutputOptions struct {
	batchPrefix     string
	batchSuffix     string
	cve2021_44228   bool
	fieldDelimiter  string
	fieldNames      []string
	outputType      string
	recordDelimiter string
	recordPrefix    string
	recordSuffix    string
	recordTemplate  string
	sampleRate      float64
	timestampFormat string
}

// This tests with basic fields to create / update a Logpush job.
func TestAccCloudflareLogpushJob_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// Logpush job config to create, with minimal fields.
	logpushJobConfigCreate := &logpushJobConfig{
		accountID:       accountID,
		dataset:         "gateway_dns", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com`,
	}

	// Logpush job config to update, with different values (where possible).
	logpushJobConfigUpdate := &logpushJobConfig{
		accountID:       accountID,
		dataset:         "gateway_dns", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com?updated=true`,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushJobBasic(rnd, logpushJobConfigCreate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dataset", toString(logpushJobConfigCreate.dataset)),
					resource.TestCheckResourceAttr(resourceName, "destination_conf", toString(logpushJobConfigCreate.destinationConf)),
				),
			},
			{
				Config: testCloudflareLogpushJobBasic(rnd, logpushJobConfigUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dataset", toString(logpushJobConfigUpdate.dataset)),
					resource.TestCheckResourceAttr(resourceName, "destination_conf", toString(logpushJobConfigUpdate.destinationConf)),
				),
			},
		},
	})
}

func testCloudflareLogpushJobBasic(resourceID string, logpushJobConfig *logpushJobConfig) string {
	// Values must be ordered to match the .tf file exactly.
	params := []any{
		resourceID,
		logpushJobConfig.accountID,
		logpushJobConfig.dataset,
		logpushJobConfig.destinationConf,
	}
	return acctest.LoadTestCase("basic.tf", params...)
}

// This tests with basic output_options fields to create / update a Logpush job.
func TestAccCloudflareLogpushJob_BasicOutputOptions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// Logpush job config to create, with minimal fields.
	logpushJobConfigCreate := &logpushJobConfig{
		accountID:       accountID,
		dataset:         "gateway_dns", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com`,
		outputOptions: &logpushJobConfigOutputOptions{
			outputType:      "ndjson",
			sampleRate:      1.0,
			timestampFormat: "unixnano",
		},
	}

	// Logpush job config to update, with different values (where possible).
	logpushJobConfigUpdate := &logpushJobConfig{
		accountID:       accountID,
		dataset:         "gateway_dns", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com?updated=true`,
		outputOptions: &logpushJobConfigOutputOptions{
			outputType:      "csv",
			sampleRate:      0.01,
			timestampFormat: "rfc3339",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushJobBasicOutputOptions(rnd, logpushJobConfigCreate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dataset", toString(logpushJobConfigCreate.dataset)),
					resource.TestCheckResourceAttr(resourceName, "destination_conf", toString(logpushJobConfigCreate.destinationConf)),
					resource.TestCheckResourceAttr(resourceName, "output_options.output_type", toString(logpushJobConfigCreate.outputOptions.outputType)),
					resource.TestCheckResourceAttr(resourceName, "output_options.sample_rate", toString(logpushJobConfigCreate.outputOptions.sampleRate)),
					resource.TestCheckResourceAttr(resourceName, "output_options.timestamp_format", toString(logpushJobConfigCreate.outputOptions.timestampFormat)),
				),
			},
			{
				Config: testCloudflareLogpushJobBasicOutputOptions(rnd, logpushJobConfigUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dataset", toString(logpushJobConfigUpdate.dataset)),
					resource.TestCheckResourceAttr(resourceName, "destination_conf", toString(logpushJobConfigUpdate.destinationConf)),
					resource.TestCheckResourceAttr(resourceName, "output_options.output_type", toString(logpushJobConfigUpdate.outputOptions.outputType)),
					resource.TestCheckResourceAttr(resourceName, "output_options.sample_rate", toString(logpushJobConfigUpdate.outputOptions.sampleRate)),
					resource.TestCheckResourceAttr(resourceName, "output_options.timestamp_format", toString(logpushJobConfigUpdate.outputOptions.timestampFormat)),
				),
			},
		},
	})
}

func testCloudflareLogpushJobBasicOutputOptions(resourceID string, logpushJobConfig *logpushJobConfig) string {
	// Values must be ordered to match the .tf file exactly.
	params := []any{
		resourceID,
		logpushJobConfig.accountID,
		logpushJobConfig.dataset,
		logpushJobConfig.destinationConf,
		logpushJobConfig.outputOptions.outputType,
		logpushJobConfig.outputOptions.sampleRate,
		logpushJobConfig.outputOptions.timestampFormat,
	}
	return acctest.LoadTestCase("output_options.tf", params...)
}

// This tests with all fields to create / update a Logpush job, except
// ownership_challenge (tested in logpush_ownership_challenge).
// Some field values or their combination is not realistic, but it is just for
// testing the schema and Logpush API behavior here.
func TestAccCloudflareLogpushJob_Full(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// Logpush job config to create, with minimal fields.
	logpushJobConfigCreate := &logpushJobConfig{
		accountID:       accountID,
		dataset:         "gateway_dns", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com`,
		frequency:       "high", // deprecated, but testing
		outputOptions: &logpushJobConfigOutputOptions{
			outputType:      "ndjson",
			sampleRate:      1.0,
			timestampFormat: "unixnano",
		},
	}

	// Logpush job config to update, with different values (where possible).
	logpushJobConfigUpdate := &logpushJobConfig{
		accountID:                accountID,
		dataset:                  "gateway_dns", // cannot be changed
		destinationConf:          `https://logpush-receiver.sd.cfplat.com?updated=true`,
		enabled:                  true,
		name:                     "terraform-test-job-updated",
		filter:                   `{"where":{"and":[{"key":"ColoCode","operator":"!eq","value":"IAD"}]}}`,
		kind:                     "", // cannot be changed
		maxUploadBytes:           5000000,
		maxUploadRecords:         1000,
		maxUploadIntervalSeconds: 30,
		frequency:                "low",                                                   // deprecated, but testing
		logpullOptions:           "fields=AccountID,Datetime,ColoCode&timestamps=rfc3339", // deprecated, but testing
		outputOptions: &logpushJobConfigOutputOptions{
			batchPrefix:     `FirstColumn\tAccountID\tDatetime\tColoCode\tLastColumn\n`,
			batchSuffix:     `\n`,
			cve2021_44228:   true,
			fieldDelimiter:  `\t`,
			fieldNames:      []string{"AccountID", "Datetime", "ColoCode"},
			outputType:      "csv",
			recordDelimiter: `\n`,
			recordPrefix:    `FirstColumn\t`,
			recordSuffix:    `\tLastColumn`,
			recordTemplate:  `FirstColumn\t{{.ClientIP}}\t{{.Datetime}}\t{{.ColoCode}}\tLastColumn`,
			sampleRate:      0.01,
			timestampFormat: "rfc3339",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushJobFull(rnd, logpushJobConfigCreate),
				Check:  resource.ComposeTestCheckFunc(getTestCheckResourceAttrs(resourceName, logpushJobConfigCreate)...),
			},
			{
				Config: testCloudflareLogpushJobFull(rnd, logpushJobConfigUpdate),
				Check:  resource.ComposeTestCheckFunc(getTestCheckResourceAttrs(resourceName, logpushJobConfigUpdate)...),
			},
		},
	})
}

func getTestCheckResourceAttrs(resourceName string, logpushJobConfig *logpushJobConfig) []resource.TestCheckFunc {
	testCheckFuncs := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "dataset", toString(logpushJobConfig.dataset)),
		resource.TestCheckResourceAttr(resourceName, "destination_conf", toString(logpushJobConfig.destinationConf)),
		resource.TestCheckResourceAttr(resourceName, "enabled", toString(logpushJobConfig.enabled)),
		resource.TestCheckResourceAttr(resourceName, "name", toString(logpushJobConfig.name)),
		resource.TestCheckResourceAttr(resourceName, "filter", toString(logpushJobConfig.filter)),
		resource.TestCheckResourceAttr(resourceName, "kind", toString(logpushJobConfig.kind)),
		resource.TestCheckResourceAttr(resourceName, "max_upload_bytes", toString(logpushJobConfig.maxUploadBytes)),
		resource.TestCheckResourceAttr(resourceName, "max_upload_records", toString(logpushJobConfig.maxUploadRecords)),
		resource.TestCheckResourceAttr(resourceName, "max_upload_interval_seconds", toString(logpushJobConfig.maxUploadIntervalSeconds)),
		resource.TestCheckResourceAttr(resourceName, "frequency", toString(logpushJobConfig.frequency)),
		resource.TestCheckResourceAttr(resourceName, "logpull_options", toString(logpushJobConfig.logpullOptions)),
		resource.TestCheckResourceAttr(resourceName, "output_options.batch_prefix", unquote(toString(logpushJobConfig.outputOptions.batchPrefix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.batch_suffix", unquote(toString(logpushJobConfig.outputOptions.batchSuffix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.cve_2021_44228", toString(logpushJobConfig.outputOptions.cve2021_44228)),
		resource.TestCheckResourceAttr(resourceName, "output_options.field_delimiter", unquote(toString(logpushJobConfig.outputOptions.fieldDelimiter))),
		resource.TestCheckResourceAttr(resourceName, "output_options.field_names.#", toString(len(logpushJobConfig.outputOptions.fieldNames))),
		resource.TestCheckResourceAttr(resourceName, "output_options.output_type", toString(logpushJobConfig.outputOptions.outputType)),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_delimiter", unquote(toString(logpushJobConfig.outputOptions.recordDelimiter))),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_prefix", unquote(toString(logpushJobConfig.outputOptions.recordPrefix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_suffix", unquote(toString(logpushJobConfig.outputOptions.recordSuffix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_template", unquote(toString(logpushJobConfig.outputOptions.recordTemplate))),
		resource.TestCheckResourceAttr(resourceName, "output_options.sample_rate", toString(logpushJobConfig.outputOptions.sampleRate)),
		resource.TestCheckResourceAttr(resourceName, "output_options.timestamp_format", toString(logpushJobConfig.outputOptions.timestampFormat)),
	}
	for i, fieldName := range logpushJobConfig.outputOptions.fieldNames {
		testCheckFuncs = append(testCheckFuncs, resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("output_options.field_names.%d", i), fieldName))
	}

	return testCheckFuncs
}

func testCloudflareLogpushJobFull(resourceID string, logpushJobConfig *logpushJobConfig) string {
	// Values must be ordered to match the .tf file exactly.
	params := []any{
		resourceID,
		logpushJobConfig.accountID,
		logpushJobConfig.dataset,
		logpushJobConfig.destinationConf,
		logpushJobConfig.enabled,
		logpushJobConfig.name,
		logpushJobConfig.filter,
		logpushJobConfig.kind,
		logpushJobConfig.maxUploadBytes,
		logpushJobConfig.maxUploadRecords,
		logpushJobConfig.maxUploadIntervalSeconds,
		logpushJobConfig.frequency,
		logpushJobConfig.logpullOptions,
		logpushJobConfig.outputOptions.batchPrefix,
		logpushJobConfig.outputOptions.batchSuffix,
		logpushJobConfig.outputOptions.cve2021_44228,
		logpushJobConfig.outputOptions.fieldDelimiter,
		jsonMarshal(logpushJobConfig.outputOptions.fieldNames),
		logpushJobConfig.outputOptions.outputType,
		logpushJobConfig.outputOptions.recordDelimiter,
		logpushJobConfig.outputOptions.recordPrefix,
		logpushJobConfig.outputOptions.recordSuffix,
		logpushJobConfig.outputOptions.recordTemplate,
		logpushJobConfig.outputOptions.sampleRate,
		logpushJobConfig.outputOptions.timestampFormat,
	}
	return acctest.LoadTestCase("full.tf", params...)
}

// This tests with immutable fields to create / update a Logpush job.
// dataset cannot be tested, as it has PlanModifiers in schema.go to replace.
// This needs to bes tested on a zone, to test both kinds with http_requests.
func TestAccCloudflareLogpushJob_ImmutableFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// Logpush job config to create, with minimal fields.
	logpushJobConfigCreate := &logpushJobConfig{
		zoneID:          zoneID,
		dataset:         "http_requests", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com`,
		kind:            "edge", // cannot be changed
	}

	// Logpush job config to update, with different values (where possible).
	logpushJobConfigUpdate := &logpushJobConfig{
		zoneID:          zoneID,
		dataset:         "http_requests", // cannot be changed
		destinationConf: `https://logpush-receiver.sd.cfplat.com?updated=true`,
		kind:            "", // cannot be changed
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushJobImmutableFields(rnd, logpushJobConfigCreate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dataset", toString(logpushJobConfigCreate.dataset)),
					resource.TestCheckResourceAttr(resourceName, "destination_conf", toString(logpushJobConfigCreate.destinationConf)),
					resource.TestCheckResourceAttr(resourceName, "kind", toString(logpushJobConfigCreate.kind)),
				),
			},
			{
				Config:      testCloudflareLogpushJobImmutableFields(rnd, logpushJobConfigUpdate),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("400 Bad Request")),
			},
		},
	})
}

func testCloudflareLogpushJobImmutableFields(resourceID string, logpushJobConfig *logpushJobConfig) string {
	// Values must be ordered to match the .tf file exactly.
	params := []any{
		resourceID,
		logpushJobConfig.zoneID,
		logpushJobConfig.dataset,
		logpushJobConfig.destinationConf,
		logpushJobConfig.kind,
	}
	return acctest.LoadTestCase("immutable_fields.tf", params...)
}
