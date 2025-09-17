// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WorkerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Immutable ID of the Worker.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Worker.",
				Required:    true,
			},
			"logpush": schema.BoolAttribute{
				Description: "Whether logpush is enabled for the Worker.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"tags": schema.SetAttribute{
				Description: "Tags associated with the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"observability": schema.SingleNestedAttribute{
				Description: "Observability settings for the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerObservabilityModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether observability is enabled for the Worker.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"head_sampling_rate": schema.Float64Attribute{
						Description: "The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).",
						Computed:    true,
						Optional:    true,
						Default:     float64default.StaticFloat64(1),
					},
					"logs": schema.SingleNestedAttribute{
						Description: "Log settings for the Worker.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerObservabilityLogsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether logs are enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"head_sampling_rate": schema.Float64Attribute{
								Description: "The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).",
								Computed:    true,
								Optional:    true,
								Default:     float64default.StaticFloat64(1),
							},
							"invocation_logs": schema.BoolAttribute{
								Description: "Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(true),
							},
						},
					},
				},
			},
			"subdomain": schema.SingleNestedAttribute{
				Description: "Subdomain settings for the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerSubdomainModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether the *.workers.dev subdomain is enabled for the Worker.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"previews_enabled": schema.BoolAttribute{
						Description: "Whether [preview URLs](https://developers.cloudflare.com/workers/configuration/previews/) are enabled for the Worker.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
				},
			},
			"tail_consumers": schema.SetNestedAttribute{
				Description: "Other Workers that should consume logs from the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectSetType[WorkerTailConsumersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of the consumer Worker.",
							Required:    true,
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the Worker was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"updated_on": schema.StringAttribute{
				Description: "When the Worker was most recently updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *WorkerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
