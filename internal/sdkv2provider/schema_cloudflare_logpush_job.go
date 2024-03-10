package sdkv2provider

import (
	"fmt"
	"regexp"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLogpushJobSchema() map[string]*schema.Schema {
	kindAllowedValues := []string{"edge", "instant-logs", ""}
	datasetAllowedValues := []string{
		"access_requests",
		"casb_findings",
		"firewall_events",
		"http_requests",
		"spectrum_events",
		"nel_reports",
		"audit_logs",
		"gateway_dns",
		"gateway_http",
		"gateway_network",
		"dns_logs",
		"network_analytics_logs",
		"workers_trace_events",
		"device_posture_results",
		"zero_trust_network_sessions",
		"magic_ids_detections",
	}
	frequencyAllowedValues := []string{"high", "low"}
	outputTypeAllowedValues := []string{"ndjson", "csv"}
	timestampFormatAllowedValues := []string{"unixnano", "unix", "rfc3339"}
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:  consts.AccountIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{consts.AccountIDSchemaKey, consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:  consts.ZoneIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{consts.AccountIDSchemaKey, consts.ZoneIDSchemaKey},
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to enable the job.",
		},
		"kind": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(kindAllowedValues, false),
			Description:  fmt.Sprintf("The kind of logpush job to create. %s", renderAvailableDocumentationValuesStringSlice(kindAllowedValues)),
		},
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]*$`), "must contain only alphanumeric characters, hyphens, and periods"),
			Description:  "The name of the logpush job to create.",
		},
		"dataset": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(datasetAllowedValues, false),
			Description: fmt.Sprintf(
				"The kind of the dataset to use with the logpush job. %s",
				renderAvailableDocumentationValuesStringSlice(datasetAllowedValues),
			),
		},
		"logpull_options": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: `Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpush options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).`,
		},
		"destination_conf": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/reference/logpush-api-configuration#destination).",
		},
		"ownership_challenge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: `Ownership challenge token to prove destination ownership, required when destination is Amazon S3, Google Cloud Storage, Microsoft Azure or Sumo Logic. See [Developer documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#usage).`,
		},
		"filter": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Use filters to select the events to include and/or remove from your logs. For more information, refer to [Filters](https://developers.cloudflare.com/logs/reference/logpush-api-configuration/filters/).",
		},
		"frequency": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "high",
			ValidateFunc: validation.StringInSlice(frequencyAllowedValues, false),
			Description:  fmt.Sprintf("A higher frequency will result in logs being pushed on faster with smaller files. `low` frequency will push logs less often with larger files. %s", renderAvailableDocumentationValuesStringSlice(frequencyAllowedValues)),
		},
		"max_upload_bytes": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(5000000, 1000000000),
			Description:  fmt.Sprint("The maximum uncompressed file size of a batch of logs. Value must be between 5MB and 1GB."),
		},
		"max_upload_records": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(1000, 1000000),
			Description:  fmt.Sprint("The maximum number of log lines per batch. Value must be between 1000 and 1,000,000."),
		},
		"max_upload_interval_seconds": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(30, 300),
			Description:  fmt.Sprint("The maximum interval in seconds for log batches. Value must be between 30 and 300."),
		},
		"output_options": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Structured replacement for logpull_options. When including this field, the logpull_option field will be ignored",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"cve20214428": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Mitigation for CVE-2021-44228. If set to true, will cause all occurrences of ${ in the generated files to be replaced with x{",
					},
					"batch_prefix": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "String to be prepended before each batch",
					},
					"batch_suffix": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "String to be appended after each batch",
					},
					"field_delimiter": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     ",",
						Description: "String to join fields. This field be ignored when record_template is set",
					},
					"field_names": {
						Type: schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional:    true,
						Description: "List of field names to be included in the Logpush output",
					},
					"output_type": {
						Type:         schema.TypeString,
						Optional:     true,
						Default:      "ndjson",
						ValidateFunc: validation.StringInSlice(outputTypeAllowedValues, false),
						Description:  fmt.Sprintf("Specifies the output type. %s", renderAvailableDocumentationValuesStringSlice(outputTypeAllowedValues)),
					},
					"record_delimiter": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "String to be inserted in-between the records as separator",
					},
					"record_prefix": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "{",
						Description: "String to be prepended before each record",
					},
					"record_suffix": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "}",
						Description: "String to be appended after each record",
					},
					"record_template": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "String to use as template for each record instead of the default comma-separated list",
					},
					"sample_rate": {
						Type:         schema.TypeFloat,
						Optional:     true,
						Default:      1.0,
						ValidateFunc: validation.FloatBetween(0.0, 1.0),
						Description:  "Specifies the sampling rate",
					},
					"timestamp_format": {
						Type:         schema.TypeString,
						Optional:     true,
						Default:      "unixnano",
						ValidateFunc: validation.StringInSlice(timestampFormatAllowedValues, false),
						Description:  fmt.Sprintf("Specifies the format for timestamps. %s", renderAvailableDocumentationValuesStringSlice(timestampFormatAllowedValues)),
					},
				},
			},
		},
	}
}
