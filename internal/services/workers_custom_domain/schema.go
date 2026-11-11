// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*WorkersCustomDomainResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers Scripts Read",
				"Workers Scripts Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Immutable ID of the domain.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
			"enabled": schema.BoolAttribute{
				Description:   "Whether the configured hostname is routable for the domain. If disabled, then requests to the configured hostname will not be routed to the configured Worker. If disabled, then preview routing must be enabled, and only subdomains of the configured hostname will be routable.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(true),
			},
			"previews_enabled": schema.BoolAttribute{
				Description:   "Whether previews are routable for the domain. If enabled, previews will be routable as subdomains of the configured hostname. For example, if the configured hostname is `app.example.com`, then requests to `my-feature.app.example.com` will be routed to the preview with slug `my-feature`, if such a preview exists. If preview routing is disabled, then hostname routing must be enabled.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"zone_name": schema.StringAttribute{
				Description:   "Name of the zone containing the domain hostname.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplaceIfConfigured()},
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
