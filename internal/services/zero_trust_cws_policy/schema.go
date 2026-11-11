// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCWSPolicyResource)(nil)

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
			"entrypoint": schema.StringAttribute{
				Description: "Non-function rule declared in rego_source. Evaluation uses data.<package_path>.<entrypoint>.",
				Required:    true,
			},
			"filter": schema.StringAttribute{
				Description: "Filter expression that determines whether this policy applies. Maximum 4096 UTF-8 bytes.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the policy",
				Required:    true,
			},
			"rego_source": schema.StringAttribute{
				Description: `Inline Rego source with a package beginning with "policy." Maximum 16384 UTF-8 bytes.`,
				Required:    true,
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
			"created_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"package_path": schema.StringAttribute{
				Description: "Dotted path derived from rego_source. Begins with policy. and is unique per account.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustCWSPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustCWSPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
