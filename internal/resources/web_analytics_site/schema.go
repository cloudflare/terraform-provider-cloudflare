// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r WebAnalyticsSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"site_tag": schema.StringAttribute{
				Description:   "The Web Analytics site identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"auto_install": schema.BoolAttribute{
				Description: "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
				Optional:    true,
			},
			"host": schema.StringAttribute{
				Description: "The hostname to use for gray-clouded sites.",
				Optional:    true,
			},
			"zone_tag": schema.StringAttribute{
				Description: "The zone identifier.",
				Optional:    true,
			},
		},
	}
}
