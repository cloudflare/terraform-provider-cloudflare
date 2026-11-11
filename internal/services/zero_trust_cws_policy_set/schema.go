// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy_set

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCWSPolicySetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Server-assigned UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Cloudflare account ID (32-character hexadecimal string)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the policy set",
				Required:    true,
			},
			"policies": schema.SetNestedAttribute{
				Description: "Policies to include with per-membership enforcement mode",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"policy_id": schema.StringAttribute{
							Description: "ID of the member policy",
							Required:    true,
						},
						"enforce": schema.BoolAttribute{
							Description: "Whether this policy should enforce when it matches. Defaults to true; false keeps the policy attached for non-enforcing evaluation.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(true),
						},
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "Optional free-form description.",
				Optional:    true,
			},
			"tags": schema.MapAttribute{
				Description: "Tags to attach. Keys are limited to 256 UTF-8 bytes and values to 1024 UTF-8 bytes.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"exclude_workload_ids": schema.SetAttribute{
				Description:   "Workload IDs that exclude matching requests. Exclusions take precedence.",
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewSetType[types.String](ctx),
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseNonNullStateForUnknown()},
			},
			"include_workload_ids": schema.SetAttribute{
				Description:   "Workload IDs that include matching requests. Empty or omitted includes all requests before exclusions.",
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewSetType[types.String](ctx),
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseNonNullStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustCWSPolicySetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustCWSPolicySetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
