// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_deployment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCWSDeploymentResource)(nil)

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
			"default_mode": schema.StringAttribute{
				Description: "Default traffic mode: ALLOW or BLOCK\nAvailable values: \"ALLOW\", \"BLOCK\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ALLOW", "BLOCK"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the deployment",
				Required:    true,
			},
			"basic_auth_credential_ids": schema.SetAttribute{
				Description: "IDs of basic-auth credentials to attach",
				Required:    true,
				ElementType: types.StringType,
			},
			"policy_set_ids": schema.SetAttribute{
				Description: "IDs of policy sets to attach",
				Required:    true,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Description: "Optional free-form description.",
				Optional:    true,
			},
			"domain_acl_id": schema.StringAttribute{
				Description: "ID of the domain ACL to attach (omit or null for none)",
				Optional:    true,
			},
			"tags": schema.MapAttribute{
				Description: "Tags to attach. Keys are limited to 256 UTF-8 bytes and values to 1024 UTF-8 bytes.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"on_error_action": schema.StringAttribute{
				Description: "Action for runtime Rego errors. Omit to use BLOCK.\nAvailable values: \"BLOCK\", \"SKIP\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("BLOCK", "SKIP"),
				},
				Default: stringdefault.StaticString("BLOCK"),
			},
			"rego_raise_error": schema.BoolAttribute{
				Description: "Whether OPA builtin failures are raised as runtime evaluation errors. Omit to use the default true value.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
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

func (r *ZeroTrustCWSDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustCWSDeploymentResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
