// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r QueueResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"queue_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"queue_name": schema.StringAttribute{
				Computed: true,
			},
			"consumers": schema.StringAttribute{
				Computed: true,
			},
			"consumers_total_count": schema.StringAttribute{
				Computed: true,
			},
			"producers": schema.StringAttribute{
				Computed: true,
			},
			"producers_total_count": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
