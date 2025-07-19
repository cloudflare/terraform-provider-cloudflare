package logpush_job_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// func TestToAPIOutputOptions(t *testing.T) {
// 	cve202144228 := true
// 	testData := cloudflare.LogpushOutputOptions{
// 		CVE202144228:    &cve202144228,
// 		BatchPrefix:     "a",
// 		BatchSuffix:     "b",
// 		FieldDelimiter:  ",",
// 		FieldNames:      []string{"a", "b", "c"},
// 		OutputType:      "csv",
// 		RecordDelimiter: "a",
// 		RecordPrefix:    "b",
// 		RecordSuffix:    "c",
// 		RecordTemplate:  "d",
// 		SampleRate:      0.5,
// 		TimestampFormat: "unix",
// 	}
// 	resourceDataMap := map[string]interface{}{
// 		"output_options": []interface{}{
// 			map[string]interface{}{
// 				"cve20214428":      *testData.CVE202144228,
// 				"batch_prefix":     testData.BatchPrefix,
// 				"batch_suffix":     testData.BatchSuffix,
// 				"field_delimiter":  testData.FieldDelimiter,
// 				"field_names":      []interface{}{testData.FieldNames[0], testData.FieldNames[1], testData.FieldNames[2]},
// 				"output_type":      testData.OutputType,
// 				"record_delimiter": testData.RecordDelimiter,
// 				"record_prefix":    testData.RecordPrefix,
// 				"record_suffix":    testData.RecordSuffix,
// 				"record_template":  testData.RecordTemplate,
// 				"sample_rate":      testData.SampleRate,
// 				"timestamp_format": testData.TimestampFormat,
// 			},
// 		},
// 	}
// 	resourceData := schema.TestResourceDataRaw(t, resourceCloudflareLogpushJobSchema(), resourceDataMap)
// 	if resourceData == nil {
// 		t.Fatal("failed to create test ResourceData")
// 	}
// 	// outputOptions, ok := resourceData.GetOk("output_options")
// 	// if !ok {
// 	// 	t.Fatal("output_options not found")
// 	// }
// 	// output, err := toAPIOutputOptions(outputOptions)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// compare output to the testData
// 	testJSON, err := json.Marshal(testData)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	outJSON, err := json.Marshal(output)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if string(testJSON) != string(outJSON) {
// 		t.Fatalf("output and testData are not equal: %s != %s", string(outJSON), string(testJSON))
// 	}
// }

func TestAccCloudflareLogpushJob_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_job." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	destinationConf := `https://logpush-receiver.sd.cfplat.com`
	dataset := "gateway_dns"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushJobBasic(rnd, accountID, dataset, destinationConf),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "destination_conf", destinationConf),
					resource.TestCheckResourceAttr(resourceName, "dataset", dataset),
				),
			},
			{
				Config: testCloudflareLogpushJobBasic(rnd, accountID, dataset, destinationConf+"?updated=true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "destination_conf", destinationConf+"?updated=true"),
					resource.TestCheckResourceAttr(resourceName, "dataset", dataset),
				),
			},
		},
	})
}

func testCloudflareLogpushJobBasic(resourceID, accountID, dataset, destinationConf string) string {
	return acctest.LoadTestCase("basic.tf", resourceID, accountID, dataset, destinationConf)
}

func jsonMarshal(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func unquote(s string) string {
	s, _ = strconv.Unquote(`"` + s + `"`)
	return s
}

// logpushJobConfig is a simplified struct for testing.
type logpushJobConfig struct {
	accountID                string
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

func getTestCheckResourceAttrs(resourceName string, logpushJobConfig *logpushJobConfig) []resource.TestCheckFunc {
	testCheckFuncs := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "dataset", fmt.Sprintf("%v", logpushJobConfig.dataset)),
		resource.TestCheckResourceAttr(resourceName, "destination_conf", fmt.Sprintf("%v", logpushJobConfig.destinationConf)),
		resource.TestCheckResourceAttr(resourceName, "enabled", fmt.Sprintf("%v", logpushJobConfig.enabled)),
		resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v", logpushJobConfig.name)),
		resource.TestCheckResourceAttr(resourceName, "filter", fmt.Sprintf("%v", logpushJobConfig.filter)),
		resource.TestCheckResourceAttr(resourceName, "kind", fmt.Sprintf("%v", logpushJobConfig.kind)),
		resource.TestCheckResourceAttr(resourceName, "max_upload_bytes", fmt.Sprintf("%v", logpushJobConfig.maxUploadBytes)),
		resource.TestCheckResourceAttr(resourceName, "max_upload_records", fmt.Sprintf("%v", logpushJobConfig.maxUploadRecords)),
		resource.TestCheckResourceAttr(resourceName, "max_upload_interval_seconds", fmt.Sprintf("%v", logpushJobConfig.maxUploadIntervalSeconds)),
		resource.TestCheckResourceAttr(resourceName, "frequency", fmt.Sprintf("%v", logpushJobConfig.frequency)),
		resource.TestCheckResourceAttr(resourceName, "logpull_options", fmt.Sprintf("%v", logpushJobConfig.logpullOptions)),
		resource.TestCheckResourceAttr(resourceName, "output_options.batch_prefix", unquote(fmt.Sprintf("%v", logpushJobConfig.outputOptions.batchPrefix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.batch_suffix", unquote(fmt.Sprintf("%v", logpushJobConfig.outputOptions.batchSuffix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.cve_2021_44228", fmt.Sprintf("%v", logpushJobConfig.outputOptions.cve2021_44228)),
		resource.TestCheckResourceAttr(resourceName, "output_options.field_delimiter", unquote(fmt.Sprintf("%v", logpushJobConfig.outputOptions.fieldDelimiter))),
		resource.TestCheckResourceAttr(resourceName, "output_options.field_names.#", fmt.Sprintf("%v", len(logpushJobConfig.outputOptions.fieldNames))),
		resource.TestCheckResourceAttr(resourceName, "output_options.output_type", fmt.Sprintf("%v", logpushJobConfig.outputOptions.outputType)),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_delimiter", unquote(fmt.Sprintf("%v", logpushJobConfig.outputOptions.recordDelimiter))),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_prefix", fmt.Sprintf("%v", logpushJobConfig.outputOptions.recordPrefix)),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_suffix", unquote(fmt.Sprintf("%v", logpushJobConfig.outputOptions.recordSuffix))),
		resource.TestCheckResourceAttr(resourceName, "output_options.record_template", unquote(fmt.Sprintf("%v", logpushJobConfig.outputOptions.recordTemplate))),
		resource.TestCheckResourceAttr(resourceName, "output_options.sample_rate", fmt.Sprintf("%v", logpushJobConfig.outputOptions.sampleRate)),
		resource.TestCheckResourceAttr(resourceName, "output_options.timestamp_format", fmt.Sprintf("%v", logpushJobConfig.outputOptions.timestampFormat)),
	}
	for i, fieldName := range logpushJobConfig.outputOptions.fieldNames {
		testCheckFuncs = append(testCheckFuncs, resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("output_options.field_names.%d", i), fieldName))
	}

	return testCheckFuncs
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

	fmt.Printf("logpushJobConfigCreate = %v\n", logpushJobConfigCreate)
	fmt.Printf("testCloudflareLogpushJobFull(rnd, logpushJobConfigCreate) = %s\n", testCloudflareLogpushJobFull(rnd, logpushJobConfigCreate))

	fmt.Printf("logpushJobConfigUpdate = %v\n", logpushJobConfigUpdate)
	fmt.Printf("testCloudflareLogpushJobFull(rnd, logpushJobConfigUpdate) = %s\n", testCloudflareLogpushJobFull(rnd, logpushJobConfigUpdate))

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

func testCloudflareLogpushJobFull(resourceID string, logpushJobConfig *logpushJobConfig) string {
	// Values must be ordered and match exactly to full.tf
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
