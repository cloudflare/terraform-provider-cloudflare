// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_incoming

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*SecondaryDNSIncomingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"auto_refresh_seconds": schema.Float64Attribute{
				Description: "How often should a secondary zone auto refresh regardless of DNS NOTIFY.\nNot applicable for primary zones.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Zone name.",
				Required:    true,
			},
			"peers": schema.ListAttribute{
				Description: "A list of peer tags.",
				Required:    true,
				ElementType: types.StringType,
			},
			"checked_time": schema.StringAttribute{
				Description: "The time for a specific event.",
				Computed:    true,
			},
			"created_time": schema.StringAttribute{
				Description: "The time for a specific event.",
				Computed:    true,
			},
			"modified_time": schema.StringAttribute{
				Description: "The time for a specific event.",
				Computed:    true,
			},
			"soa_serial": schema.Float64Attribute{
				Description: "The serial number of the SOA for the given zone.",
				Computed:    true,
			},
		},
	}
}

func (r *SecondaryDNSIncomingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SecondaryDNSIncomingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
