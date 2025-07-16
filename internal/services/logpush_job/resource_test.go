package logpush_job_test

import (
	"os"
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
	jobName := "terraform-test-job"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareLogpushJobBasic(rnd, accountID, destinationConf, dataset, jobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "destination_conf", destinationConf),
					resource.TestCheckResourceAttr(resourceName, "dataset", dataset),
					resource.TestCheckResourceAttr(resourceName, "name", jobName),
				),
			},
			{
				Config: testCloudflareLogpushJobBasic(rnd, accountID, destinationConf, dataset, jobName+"-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "destination_conf", destinationConf),
					resource.TestCheckResourceAttr(resourceName, "dataset", dataset),
					resource.TestCheckResourceAttr(resourceName, "name", jobName+"-updated"),
				),
			},
		},
	})
}

func testCloudflareLogpushJobBasic(resourceID, zoneID, destinationConf, dataset, jobName string) string {
	return acctest.LoadTestCase("basic.tf", resourceID, zoneID, destinationConf, dataset, jobName)
}
