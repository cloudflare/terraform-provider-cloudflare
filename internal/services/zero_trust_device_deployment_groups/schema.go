// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_deployment_groups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDeviceDeploymentGroupsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the deployment group.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the deployment group.",
				Required:    true,
			},
			"version_config": schema.ListNestedAttribute{
				Description: "Contains at least one version configuration.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"target_environment": schema.StringAttribute{
							Description: "The target environment for the client version (e.g., windows, macos).",
							Required:    true,
						},
						"version": schema.StringAttribute{
							Description: "The specific client version to deploy.",
							Required:    true,
						},
					},
				},
			},
			"policy_ids": schema.ListAttribute{
				Description: "Contains an optional list of policy IDs assigned to a group.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "The RFC3339Nano timestamp when the deployment group was created.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The RFC3339Nano timestamp when the deployment group was last updated.",
				Computed:    true,
			},
		},
	}
}

func (r *ZeroTrustDeviceDeploymentGroupsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDeviceDeploymentGroupsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
