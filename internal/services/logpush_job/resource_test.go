package logpush_job_test

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact(toString(logpushJobConfigCreate.dataset))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
				},
			},
			{
				Config: testCloudflareLogpushJobBasic(rnd, logpushJobConfigUpdate),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigUpdate.destinationConf))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact(toString(logpushJobConfigUpdate.dataset))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigUpdate.destinationConf))),
				},
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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact(toString(logpushJobConfigCreate.dataset))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("output_type"), knownvalue.StringExact(toString(logpushJobConfigCreate.outputOptions.outputType))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("sample_rate"), knownvalue.Float64Exact(logpushJobConfigCreate.outputOptions.sampleRate)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("timestamp_format"), knownvalue.StringExact(toString(logpushJobConfigCreate.outputOptions.timestampFormat))),
				},
			},
			{
				Config: testCloudflareLogpushJobBasicOutputOptions(rnd, logpushJobConfigUpdate),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigUpdate.destinationConf))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact(toString(logpushJobConfigUpdate.dataset))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigUpdate.destinationConf))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("output_type"), knownvalue.StringExact(toString(logpushJobConfigUpdate.outputOptions.outputType))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("sample_rate"), knownvalue.Float64Exact(logpushJobConfigUpdate.outputOptions.sampleRate)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("timestamp_format"), knownvalue.StringExact(toString(logpushJobConfigUpdate.outputOptions.timestampFormat))),
				},
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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
					},
				},
				ConfigStateChecks: getStateChecks(resourceName, logpushJobConfigCreate),
			},
			{
				Config: testCloudflareLogpushJobFull(rnd, logpushJobConfigUpdate),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigUpdate.destinationConf))),
					},
				},
				ConfigStateChecks: getStateChecks(resourceName, logpushJobConfigUpdate),
			},
		},
	})
}

func getStateChecks(resourceName string, logpushJobConfig *logpushJobConfig) []statecheck.StateCheck {
	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact(toString(logpushJobConfig.dataset))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfig.destinationConf))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(logpushJobConfig.enabled)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(toString(logpushJobConfig.name))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("filter"), knownvalue.StringExact(toString(logpushJobConfig.filter))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact(toString(logpushJobConfig.kind))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("max_upload_bytes"), knownvalue.Int64Exact(int64(logpushJobConfig.maxUploadBytes))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("max_upload_records"), knownvalue.Int64Exact(int64(logpushJobConfig.maxUploadRecords))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("max_upload_interval_seconds"), knownvalue.Int64Exact(int64(logpushJobConfig.maxUploadIntervalSeconds))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("frequency"), knownvalue.StringExact(toString(logpushJobConfig.frequency))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("logpull_options"), knownvalue.StringExact(toString(logpushJobConfig.logpullOptions))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("batch_prefix"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.batchPrefix)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("batch_suffix"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.batchSuffix)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("cve_2021_44228"), knownvalue.Bool(logpushJobConfig.outputOptions.cve2021_44228)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("field_delimiter"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.fieldDelimiter)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("output_type"), knownvalue.StringExact(toString(logpushJobConfig.outputOptions.outputType))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("record_delimiter"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.recordDelimiter)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("record_prefix"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.recordPrefix)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("record_suffix"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.recordSuffix)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("record_template"), knownvalue.StringExact(unquote(toString(logpushJobConfig.outputOptions.recordTemplate)))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("sample_rate"), knownvalue.Float64Exact(logpushJobConfig.outputOptions.sampleRate)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("timestamp_format"), knownvalue.StringExact(toString(logpushJobConfig.outputOptions.timestampFormat))),
	}

	if len(logpushJobConfig.outputOptions.fieldNames) == 0 {
		stateChecks = append(stateChecks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("field_names"), knownvalue.Null()))
	} else {
		list := []knownvalue.Check{}
		for _, fieldName := range logpushJobConfig.outputOptions.fieldNames {
			list = append(list, knownvalue.StringExact(fieldName))
		}
		stateChecks = append(stateChecks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("output_options").AtMapKey("field_names"), knownvalue.ListExact(list)))
	}

	return stateChecks
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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dataset"), knownvalue.StringExact(toString(logpushJobConfigCreate.dataset))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact(toString(logpushJobConfigCreate.destinationConf))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact(toString(logpushJobConfigCreate.kind))),
				},
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
