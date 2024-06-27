// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r WorkersForPlatformsNamespaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Computed:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description:   "API Resource UUID tag.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the dispatch namespace",
				Optional:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the script was created.",
				Computed:    true,
			},
			"modified_by": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the script was last modified.",
				Computed:    true,
			},
			"namespace_name": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Computed:    true,
			},
			"script_count": schema.Int64Attribute{
				Description: "The current number of scripts in this Dispatch Namespace",
				Computed:    true,
			},
		},
	}
}
