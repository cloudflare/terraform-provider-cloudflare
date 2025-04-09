// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_dns_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*AccountDNSSettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_defaults": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AccountDNSSettingsZoneDefaultsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"flatten_all_cnames": schema.BoolAttribute{
						Description: "Whether to flatten all CNAME records in the zone. Note that, due to DNS limitations, a CNAME record at the zone apex will always be flattened.",
						Optional:    true,
					},
					"foundation_dns": schema.BoolAttribute{
						Description: "Whether to enable Foundation DNS Advanced Nameservers on the zone.",
						Optional:    true,
					},
					"internal_dns": schema.SingleNestedAttribute{
						Description: "Settings for this internal zone.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[AccountDNSSettingsZoneDefaultsInternalDNSModel](ctx),
						Attributes: map[string]schema.Attribute{
							"reference_zone_id": schema.StringAttribute{
								Description: "The ID of the zone to fallback to.",
								Optional:    true,
							},
						},
					},
					"multi_provider": schema.BoolAttribute{
						Description: "Whether to enable multi-provider DNS, which causes Cloudflare to activate the zone even when non-Cloudflare NS records exist, and to respect NS records at the zone apex during outbound zone transfers.",
						Optional:    true,
					},
					"nameservers": schema.SingleNestedAttribute{
						Description: "Settings determining the nameservers through which the zone should be available.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[AccountDNSSettingsZoneDefaultsNameserversModel](ctx),
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: "Nameserver type\nAvailable values: \"cloudflare.standard\", \"cloudflare.standard.random\", \"custom.account\", \"custom.tenant\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"cloudflare.standard",
										"cloudflare.standard.random",
										"custom.account",
										"custom.tenant",
									),
								},
							},
						},
					},
					"ns_ttl": schema.Float64Attribute{
						Description: "The time to live (TTL) of the zone's nameserver (NS) records.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(30, 86400),
						},
					},
					"secondary_overrides": schema.BoolAttribute{
						Description: "Allows a Secondary DNS zone to use (proxied) override records and CNAME flattening at the zone apex.",
						Optional:    true,
					},
					"soa": schema.SingleNestedAttribute{
						Description: "Components of the zone's SOA record.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[AccountDNSSettingsZoneDefaultsSOAModel](ctx),
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
					},
				},
			},
		},
	}
}

func (r *AccountDNSSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccountDNSSettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
