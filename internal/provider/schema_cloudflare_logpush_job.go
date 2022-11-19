package provider

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLogpushJobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:  "The account identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"zone_id": {
			Description:  "The zone identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to enable the job.",
		},
		"kind": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"edge", "instant-logs", ""}, false),
			Description:  fmt.Sprintf("The kind of logpush job to create. %s", renderAvailableDocumentationValuesStringSlice([]string{"edge", "instant-logs", `""`})),
		},
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]*$`), "must contain only alphanumeric characters, hyphens, and periods"),
			Description:  "The name of the logpush job to create.",
		},
		"dataset": {
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"access_requests",
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
			}, false),
			Description: fmt.Sprintf(
				"Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/reference/logpush-api-configuration#destination). %s",
				renderAvailableDocumentationValuesStringSlice([]string{
					"access_requests",
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
				}),
			),
		},
		"logpull_options": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: `Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpull options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).`,
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
			ValidateFunc: validation.StringInSlice([]string{"high", "low"}, false),
			Description:  fmt.Sprintf("A higher frequency will result in logs being pushed on faster with smaller files. `low` frequency will push logs less often with larger files. %s", renderAvailableDocumentationValuesStringSlice([]string{"high", "low"})),
		},
	}
}
