// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZoneResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description:   "The domain name",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"vanity_name_servers": schema.ListAttribute{
				Description: "An array of domains used for custom name servers. This is only\navailable for Business and Enterprise plans.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"type": schema.StringAttribute{
				Description: "A full zone implies that DNS is hosted with Cloudflare. A partial zone is\ntypically a partner-hosted zone or a CNAME setup.\n",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"full",
						"partial",
						"secondary",
					),
				},
				Default: stringdefault.StaticString("full"),
			},
			"activated_on": schema.StringAttribute{
				Description: "The last time proof of ownership was detected and the zone was made\nactive",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"created_on": schema.StringAttribute{
				Description: "When the zone was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"development_mode": schema.Float64Attribute{
				Description: "The interval (in seconds) from when development mode expires\n(positive integer) or last expired (negative integer) for the\ndomain. If development mode has never been enabled, this value is 0.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the zone was last modified",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"original_dnshost": schema.StringAttribute{
				Description: "DNS host at the time of switching to Cloudflare",
				Computed:    true,
			},
			"original_registrar": schema.StringAttribute{
				Description: "Registrar for the domain at the time of switching to Cloudflare",
				Computed:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "Indicates whether the zone is only using Cloudflare DNS services. A\ntrue value means the zone will not receive security or performance\nbenefits.\n",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"status": schema.StringAttribute{
				Description: "The zone status on Cloudflare.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"initializing",
						"pending",
						"active",
						"moved",
					),
				},
			},
			"name_servers": schema.ListAttribute{
				Description: "The name servers Cloudflare assigns to a zone",
				Computed:    true,
				ElementType: types.StringType,
			},
			"original_name_servers": schema.ListAttribute{
				Description: "Original name servers before moving to Cloudflare",
				Computed:    true,
				ElementType: types.StringType,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Metadata about the zone",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneMetaModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cdn_only": schema.BoolAttribute{
						Description: "The zone is only configured for CDN",
						Computed:    true,
						Optional:    true,
					},
					"custom_certificate_quota": schema.Int64Attribute{
						Description: "Number of Custom Certificates the zone can have",
						Computed:    true,
						Optional:    true,
					},
					"dns_only": schema.BoolAttribute{
						Description: "The zone is only configured for DNS",
						Computed:    true,
						Optional:    true,
					},
					"foundation_dns": schema.BoolAttribute{
						Description: "The zone is setup with Foundation DNS",
						Computed:    true,
						Optional:    true,
					},
					"page_rule_quota": schema.Int64Attribute{
						Description: "Number of Page Rules a zone can have",
						Computed:    true,
						Optional:    true,
					},
					"phishing_detected": schema.BoolAttribute{
						Description: "The zone has been flagged for phishing",
						Computed:    true,
						Optional:    true,
					},
					"step": schema.Int64Attribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"owner": schema.SingleNestedAttribute{
				Description: "The owner of the zone",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneOwnerModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the owner",
						Computed:    true,
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of owner",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *ZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
