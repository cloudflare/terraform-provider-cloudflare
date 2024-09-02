// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue_name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"environment": schema.StringAttribute{
							Computed: true,
						},
						"queue_name": schema.StringAttribute{
							Computed: true,
						},
						"service": schema.StringAttribute{
							Computed: true,
						},
						"settings": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[QueueConsumersSettingsModel](ctx),
							Attributes: map[string]schema.Attribute{
								"batch_size": schema.Float64Attribute{
									Description: "The maximum number of messages to include in a batch.",
									Computed:    true,
									Optional:    true,
								},
								"max_retries": schema.Float64Attribute{
									Description: "The maximum number of retries",
									Computed:    true,
									Optional:    true,
								},
								"max_wait_time_ms": schema.Float64Attribute{
									Computed: true,
									Optional: true,
								},
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
						"environment": schema.StringAttribute{
							Computed: true,
						},
						"service": schema.StringAttribute{
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
