// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*QueueConsumerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"consumer_id": schema.StringAttribute{
				Description:   "Identifier.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"dead_letter_queue": schema.StringAttribute{
				Computed: true,
			},
			"environment": schema.StringAttribute{
				Computed: true,
			},
			"queue_name": schema.StringAttribute{
				Computed: true,
			},
			"script_name": schema.StringAttribute{
				Computed: true,
			},
			"settings": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[QueueConsumerSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"batch_size": schema.Float64Attribute{
						Computed: true,
					},
					"max_retries": schema.Float64Attribute{
						Description: "The maximum number of retries",
						Computed:    true,
					},
					"max_wait_time_ms": schema.Float64Attribute{
						Computed: true,
					},
				},
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
