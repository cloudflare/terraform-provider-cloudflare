// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*OrganizationProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"business_address": schema.StringAttribute{
				Required: true,
			},
			"business_email": schema.StringAttribute{
				Required: true,
			},
			"business_name": schema.StringAttribute{
				Required: true,
			},
			"business_phone": schema.StringAttribute{
				Required: true,
			},
			"external_metadata": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *OrganizationProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OrganizationProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
