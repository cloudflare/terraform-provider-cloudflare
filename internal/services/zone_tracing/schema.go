// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZoneTracingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Specify the zone ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Specify the zone ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether Cloudflare Traces is enabled for the zone.",
				Optional:    true,
			},
			"forward_context": schema.BoolAttribute{
				Description: "Whether trace context is sent externally or across a zone boundary.",
				Optional:    true,
			},
			"persist": schema.BoolAttribute{
				Description: "Whether traces are persisted in Cloudflare.",
				Optional:    true,
			},
			"propagation_policy": schema.StringAttribute{
				Description: "When inbound trace context may be continued. Authenticated propagation is not supported yet.\nAvailable values: \"accept\", \"authenticated\", \"reject\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"accept",
						"authenticated",
						"reject",
					),
				},
			},
			"sampling_ratio": schema.Float64Attribute{
				Description: "The ratio of requests sampled for tracing, from 0 to 1.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"destinations": schema.ListAttribute{
				Description: "Up to 100 OpenTelemetry destination identifiers that receive traces.",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *ZoneTracingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneTracingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
