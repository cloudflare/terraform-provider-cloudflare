// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustResourceLibraryApplicationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Returns the application ID.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"category_id": schema.Int64Attribute{
				Description: "Returns the category ID.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"human_id": schema.StringAttribute{
				Description:   "Returns the human readable ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Returns the application name.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostnames": schema.SetAttribute{
				Description: "Hostnames matched by the application.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ip_subnets": schema.SetAttribute{
				Description: "IP subnets matched by the application.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"port_protocols": schema.SetAttribute{
				Description: "Port and protocol pairs matched by the application.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"support_domains": schema.SetAttribute{
				Description: "Support domains matched by the application.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"application_confidence_score": schema.Float64Attribute{
				Description: "Confidence score for the application. Returns -1 when no score is available.",
				Computed:    true,
			},
			"application_source": schema.StringAttribute{
				Description: "Returns the application source.",
				Computed:    true,
			},
			"application_type": schema.StringAttribute{
				Description: "Returns the application type.",
				Computed:    true,
			},
			"application_type_description": schema.StringAttribute{
				Description: "Returns the application type description.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description:   "Returns the application creation time.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"gen_ai_score": schema.Float64Attribute{
				Description: "GenAI score for the application. Returns -1 when no score is available.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Returns the application update time.",
				Computed:    true,
			},
			"version": schema.StringAttribute{
				Description: "Returns the application version.",
				Computed:    true,
			},
			"supported": schema.SetAttribute{
				Description: "Cloudflare products that support this application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"application_score_composition": schema.StringAttribute{
				Description: "Returns the score composition breakdown for the application.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (r *ZeroTrustResourceLibraryApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustResourceLibraryApplicationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
