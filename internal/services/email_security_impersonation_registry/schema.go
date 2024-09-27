// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*EmailSecurityImpersonationRegistryResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.StringAttribute{
							Required: true,
						},
						"is_email_regex": schema.BoolAttribute{
							Required: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"email": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"is_email_regex": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"directory_id": schema.Int64Attribute{
				Computed: true,
			},
			"directory_node_id": schema.StringAttribute{
				Computed: true,
			},
			"last_modified": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"provenance": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *EmailSecurityImpersonationRegistryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailSecurityImpersonationRegistryResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
