// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*QueueResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"queue_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "A Resource identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue_name": schema.StringAttribute{
				Required: true,
			},
			"settings": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[QueueSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"delivery_delay": schema.Float64Attribute{
						Description: "Number of seconds to delay delivery of all messages to consumers.",
						Optional:    true,
					},
					"message_retention_period": schema.Float64Attribute{
						Description: "Number of seconds after which an unconsumed message will be delayed.",
						Optional:    true,
					},
				},
			},
			"consumers_total_count": schema.Float64Attribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"producers_total_count": schema.Float64Attribute{
				Computed: true,
			},
			"consumers": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[QueueConsumersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"consumer_id": schema.StringAttribute{
							Description: "A Resource identifier.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"queue_id": schema.StringAttribute{
							Description: "A Resource identifier.",
							Computed:    true,
						},
						"script": schema.StringAttribute{
							Description: "Name of a Worker",
							Computed:    true,
						},
						"script_name": schema.StringAttribute{
							Description: "Name of a Worker",
							Computed:    true,
						},
						"settings": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[QueueConsumersSettingsModel](ctx),
							Attributes: map[string]schema.Attribute{
								"batch_size": schema.Float64Attribute{
									Description: "The maximum number of messages to include in a batch.",
									Computed:    true,
								},
								"max_concurrency": schema.Float64Attribute{
									Description: "Maximum number of concurrent consumers that may consume from this Queue. Set to `null` to automatically opt in to the platform's maximum (recommended).",
									Computed:    true,
								},
								"max_retries": schema.Float64Attribute{
									Description: "The maximum number of retries",
									Computed:    true,
								},
								"max_wait_time_ms": schema.Float64Attribute{
									Description: "The number of milliseconds to wait for a batch to fill up before attempting to deliver it",
									Computed:    true,
								},
								"retry_delay": schema.Float64Attribute{
									Description: "The number of seconds to delay before making the message available for another attempt.",
									Computed:    true,
								},
								"visibility_timeout_ms": schema.Float64Attribute{
									Description: "The number of milliseconds that a message is exclusively leased. After the timeout, the message becomes available for another attempt.",
									Computed:    true,
								},
							},
						},
						"type": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("worker", "http_pull"),
							},
						},
					},
				},
			},
			"producers": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[QueueProducersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"script": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("worker", "r2_bucket"),
							},
						},
						"bucket_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *QueueResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *QueueResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
