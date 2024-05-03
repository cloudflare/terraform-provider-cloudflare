// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r WorkerDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifer of the Worker Domain.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"environment": schema.StringAttribute{
				Description: "Worker environment associated with the zone and hostname.",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "Hostname of the Worker Domain.",
				Required:    true,
			},
			"service": schema.StringAttribute{
				Description: "Worker service associated with the zone and hostname.",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier of the zone.",
				Required:    true,
			},
		},
	}
}
