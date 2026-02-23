package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareQueueConsumerSchema returns the schema for cloudflare_queue_consumer at schema_version=0.
// This resource was introduced in v5 with no v4 predecessor. The source schema matches the original
// v5 schema before explicit versioning was applied (i.e., before Version: GetSchemaVersion(1, 500)).
//
// Note: PlanModifiers and Validators are intentionally omitted — they are not stored in state
// and are not needed for state parsing.
func SourceCloudflareQueueConsumerSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"queue_id": schema.StringAttribute{
				Required: true,
			},
			"consumer_id": schema.StringAttribute{
				Computed: true,
			},
			"dead_letter_queue": schema.StringAttribute{
				Optional: true,
			},
			"script_name": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},
			"settings": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"batch_size": schema.Float64Attribute{
						Optional: true,
					},
					"max_concurrency": schema.Float64Attribute{
						Optional: true,
					},
					"max_retries": schema.Float64Attribute{
						Optional: true,
					},
					"max_wait_time_ms": schema.Float64Attribute{
						Optional: true,
					},
					"retry_delay": schema.Float64Attribute{
						Optional: true,
					},
					"visibility_timeout_ms": schema.Float64Attribute{
						Optional: true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"script": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
