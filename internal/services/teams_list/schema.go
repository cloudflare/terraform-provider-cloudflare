// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r TeamsListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "API Resource UUID tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the list.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the list.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of list.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("SERIAL", "URL", "DOMAIN", "EMAIL", "IP"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items in the list.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Description: "The description of the list item, if present",
							Optional:    true,
						},
						"value": schema.StringAttribute{
							Description: "The value of the item in a list.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"list_count": schema.Float64Attribute{
				Description: "The number of items in the list.",
				Computed:    true,
			},
		},
	}
}
