package sdkv2provider

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestToAPIOutputOptions(t *testing.T) {
	cve202144228 := true
	testData := cloudflare.LogpushOutputOptions{
		CVE202144228:    &cve202144228,
		BatchPrefix:     "a",
		BatchSuffix:     "b",
		FieldDelimiter:  ",",
		FieldNames:      []string{"a", "b", "c"},
		OutputType:      "csv",
		RecordDelimiter: "a",
		RecordPrefix:    "b",
		RecordSuffix:    "c",
		RecordTemplate:  "d",
		SampleRate:      0.5,
		TimestampFormat: "unix",
	}
	resourceDataMap := map[string]interface{}{
		"output_options": []interface{}{
			map[string]interface{}{
				"cve20214428":      *testData.CVE202144228,
				"batch_prefix":     testData.BatchPrefix,
				"batch_suffix":     testData.BatchSuffix,
				"field_delimiter":  testData.FieldDelimiter,
				"field_names":      []interface{}{testData.FieldNames[0], testData.FieldNames[1], testData.FieldNames[2]},
				"output_type":      testData.OutputType,
				"record_delimiter": testData.RecordDelimiter,
				"record_prefix":    testData.RecordPrefix,
				"record_suffix":    testData.RecordSuffix,
				"record_template":  testData.RecordTemplate,
				"sample_rate":      testData.SampleRate,
				"timestamp_format": testData.TimestampFormat,
			},
		},
	}
	resourceData := schema.TestResourceDataRaw(t, resourceCloudflareLogpushJobSchema(), resourceDataMap)
	if resourceData == nil {
		t.Fatal("failed to create test ResourceData")
	}
	outputOptions, ok := resourceData.GetOk("output_options")
	if !ok {
		t.Fatal("output_options not found")
	}
	output, err := toAPIOutputOptions(outputOptions)
	if err != nil {
		t.Fatal(err)
	}
	// compare output to the testData
	testJSON, err := json.Marshal(testData)
	if err != nil {
		t.Fatal(err)
	}
	outJSON, err := json.Marshal(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(testJSON) != string(outJSON) {
		t.Fatalf("output and testData are not equal: %s != %s", string(outJSON), string(testJSON))
	}
}
