// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dns_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZoneDNSSettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"flatten_all_cnames": schema.BoolAttribute{
				Description:   "Whether to flatten all CNAME records in the zone. Note that, due to DNS limitations, a CNAME record at the zone apex will always be flattened.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"foundation_dns": schema.BoolAttribute{
				Description:   "Whether to enable Foundation DNS Advanced Nameservers on the zone.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"multi_provider": schema.BoolAttribute{
				Description:   "Whether to enable multi-provider DNS, which causes Cloudflare to activate the zone even when non-Cloudflare NS records exist, and to respect NS records at the zone apex during outbound zone transfers.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"ns_ttl": schema.Float64Attribute{
				Description: "The time to live (TTL) of the zone's nameserver (NS) records.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(30, 86400),
				},
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"secondary_overrides": schema.BoolAttribute{
				Description:   "Allows a Secondary DNS zone to use (proxied) override records and CNAME flattening at the zone apex.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"zone_mode": schema.StringAttribute{
				Description: "Whether the zone mode is a regular or CDN/DNS only zone.\nAvailable values: \"standard\", \"cdn_only\", \"dns_only\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"standard",
						"cdn_only",
						"dns_only",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"internal_dns": schema.SingleNestedAttribute{
				Description: "Settings for this internal zone.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneDNSSettingsInternalDNSModel](ctx),
				Attributes: map[string]schema.Attribute{
					"reference_zone_id": schema.StringAttribute{
						Description: "The ID of the zone to fallback to.",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"nameservers": schema.SingleNestedAttribute{
				Description: "Settings determining the nameservers through which the zone should be available.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneDNSSettingsNameserversModel](ctx),
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: "Nameserver type\nAvailable values: \"cloudflare.standard\", \"custom.account\", \"custom.tenant\", \"custom.zone\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"cloudflare.standard",
								"custom.account",
								"custom.tenant",
								"custom.zone",
							),
						},
					},
					"ns_set": schema.Int64Attribute{
						Description: "Configured nameserver set to be used for this zone",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 5),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"soa": schema.SingleNestedAttribute{
				Description: "Components of the zone's SOA record.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneDNSSettingsSOAModel](ctx),
				Attributes: map[string]schema.Attribute{
					"expire": schema.Float64Attribute{
						Description: "Time in seconds of being unable to query the primary server after which secondary servers should stop serving the zone.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(86400, 2419200),
						},
					},
					"min_ttl": schema.Float64Attribute{
						Description: "The time to live (TTL) for negative caching of records within the zone.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(60, 86400),
						},
					},
					"mname": schema.StringAttribute{
						Description: "The primary nameserver, which may be used for outbound zone transfers.",
						Required:    true,
					},
					"refresh": schema.Float64Attribute{
						Description: "Time in seconds after which secondary servers should re-check the SOA record to see if the zone has been updated.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(600, 86400),
						},
					},
					"retry": schema.Float64Attribute{
						Description: "Time in seconds after which secondary servers should retry queries after the primary server was unresponsive.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(600, 86400),
						},
					},
					"rname": schema.StringAttribute{
						Description: "The email address of the zone administrator, with the first label representing the local part of the email address.",
						Required:    true,
					},
					"ttl": schema.Float64Attribute{
						Description: "The time to live (TTL) of the SOA record itself.",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(300, 86400),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *ZoneDNSSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneDNSSettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
