// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*WorkersCustomDomainResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Immutable ID of the domain.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description:   "Hostname of the domain. Can be either the zone apex or a subdomain of the zone. Requests to this hostname will be routed to the configured Worker.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"service": schema.StringAttribute{
				Description:   "Name of the Worker associated with the domain. Requests to the configured hostname will be routed to this Worker.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"environment": schema.StringAttribute{
				Description:        "Worker environment associated with the domain.",
				Computed:           true,
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "ID of the zone containing the domain hostname.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"zone_name": schema.StringAttribute{
				Description:   "Name of the zone containing the domain hostname.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"cert_id": schema.StringAttribute{
				Description: "ID of the TLS certificate issued for the domain.",
				Computed:    true,
			},
		},
	}
}

func (r *WorkersCustomDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkersCustomDomainResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
