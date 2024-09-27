// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ResourceGroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier of the group.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier tag.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"scope": schema.SingleNestedAttribute{
				Description: "A scope is a combination of scope objects which provides additional context.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
						Required:    true,
					},
					"objects": schema.ListNestedAttribute{
						Description: "A list of scope objects for additional context. The number of Scope objects should not be zero.",
						Required:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description: "This is a combination of pre-defined resource name and identifier (like Zone ID etc.)",
									Required:    true,
								},
							},
						},
					},
				},
			},
			"meta": schema.StringAttribute{
				Description: "Attributes associated to the resource group.",
				Optional:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the resource group.",
				Computed:    true,
			},
		},
	}
}

func (r *ResourceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ResourceGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
