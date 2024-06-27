// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r ZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "The domain name",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "A full zone implies that DNS is hosted with Cloudflare. A partial zone is\ntypically a partner-hosted zone or a CNAME setup.\n",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("full", "partial", "secondary"),
				},
			},
			"activated_on": schema.StringAttribute{
				Description: "The last time proof of ownership was detected and the zone was made\nactive",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the zone was created",
				Computed:    true,
			},
			"development_mode": schema.Float64Attribute{
				Description: "The interval (in seconds) from when development mode expires\n(positive integer) or last expired (negative integer) for the\ndomain. If development mode has never been enabled, this value is 0.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the zone was last modified",
				Computed:    true,
			},
			"name_servers": schema.ListAttribute{
				Description: "The name servers Cloudflare assigns to a zone",
				Computed:    true,
				ElementType: types.StringType,
			},
			"original_dnshost": schema.StringAttribute{
				Description: "DNS host at the time of switching to Cloudflare",
				Computed:    true,
			},
			"original_name_servers": schema.ListAttribute{
				Description: "Original name servers before moving to Cloudflare",
				Computed:    true,
				ElementType: types.StringType,
			},
			"original_registrar": schema.StringAttribute{
				Description: "Registrar for the domain at the time of switching to Cloudflare",
				Computed:    true,
			},
			"vanity_name_servers": schema.ListAttribute{
				Description: "An array of domains used for custom name servers. This is only available for Business and Enterprise plans.",
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}
