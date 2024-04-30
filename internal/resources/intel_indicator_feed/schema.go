// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_indicator_feed

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func (r IntelIndicatorFeedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The unique identifier for the indicator feed",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the example test",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the indicator feed",
				Optional:    true,
			},
		},
	}
}
