// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*SnippetsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"snippet_name": schema.StringAttribute{
				Description:   "The identifying name of the snippet.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"files": schema.ListAttribute{
				Description: "The list of files belonging to the snippet.",
				Required:    true,
				ElementType: types.StringType,
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "Metadata about the snippet.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"main_module": schema.StringAttribute{
						Description: "Name of the file that contains the main module of the snippet.",
						Required:    true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the snippet was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the snippet was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *SnippetsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SnippetsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
