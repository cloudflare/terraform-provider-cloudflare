// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r ByoIPPrefixResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifier",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"asn": schema.Int64Attribute{
						Description:   "Autonomous System Number (ASN) the prefix will be advertised under.",
						Required:      true,
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"cidr": schema.StringAttribute{
						Description:   "IP Prefix in Classless Inter-Domain Routing format.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"loa_document_id": schema.StringAttribute{
						Description:   "Identifier for the uploaded LOA document.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"description": schema.StringAttribute{
						Description: "Description of the prefix.",
						Optional:    true,
					},
					"advertised": schema.BoolAttribute{
						Description: "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.",
						Computed:    true,
					},
					"advertised_modified_at": schema.StringAttribute{
						Description: "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"approved": schema.StringAttribute{
						Description: "Approval state of the prefix (P = pending, V = active).",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"modified_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"on_demand_enabled": schema.BoolAttribute{
						Description: "Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.",
						Computed:    true,
					},
					"on_demand_locked": schema.BoolAttribute{
						Description: "Whether advertisement status of the prefix is locked, meaning it cannot be changed.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state ByoIPPrefixModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
