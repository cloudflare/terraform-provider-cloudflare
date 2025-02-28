// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*QueueConsumerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "A Resource identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue_id": schema.StringAttribute{
				Description:   "A Resource identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"consumer_id": schema.StringAttribute{
				Description:   "A Resource identifier.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"dead_letter_queue": schema.StringAttribute{
				Optional: true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of a Worker",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: `Available values: "worker".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("worker", "http_pull"),
				},
			},
			"settings": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[QueueConsumerSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"batch_size": schema.Float64Attribute{
						Description: "The maximum number of messages to include in a batch.",
						Optional:    true,
					},
					"max_concurrency": schema.Float64Attribute{
						Description: "Maximum number of concurrent consumers that may consume from this Queue. Set to `null` to automatically opt in to the platform's maximum (recommended).",
						Optional:    true,
					},
					"max_retries": schema.Float64Attribute{
						Description: "The maximum number of retries",
						Optional:    true,
					},
					"max_wait_time_ms": schema.Float64Attribute{
						Description: "The number of milliseconds to wait for a batch to fill up before attempting to deliver it",
						Optional:    true,
					},
					"retry_delay": schema.Float64Attribute{
						Description: "The number of seconds to delay before making the message available for another attempt.",
						Optional:    true,
					},
					"visibility_timeout_ms": schema.Float64Attribute{
						Description: "The number of milliseconds that a message is exclusively leased. After the timeout, the message becomes available for another attempt.",
						Optional:    true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"script": schema.StringAttribute{
				Description: "Name of a Worker",
				Computed:    true,
			},
		},
	}
}

func (r *QueueConsumerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *QueueConsumerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
