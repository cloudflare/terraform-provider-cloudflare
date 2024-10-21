// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*SnippetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"snippet_name": schema.StringAttribute{
				Description:   "Snippet identifying name",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"files": schema.StringAttribute{
				Description:   "Content files of uploaded snippet",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[SnippetMetadataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"main_module": schema.StringAttribute{
						Description: "Main module name of uploaded snippet",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"created_on": schema.StringAttribute{
				Description: "Creation time of the snippet",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "Modification time of the snippet",
				Computed:    true,
			},
		},
	}
}

func (r *SnippetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SnippetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
