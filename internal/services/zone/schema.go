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
				Description:   "The domain name.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"paused": schema.BoolAttribute{
				Description: "Indicates whether the zone is only using Cloudflare DNS services. A\ntrue value means the zone will not receive security or performance\nbenefits.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"type": schema.StringAttribute{
				Description: "A full zone implies that DNS is hosted with Cloudflare. A partial zone is\ntypically a partner-hosted zone or a CNAME setup.\nAvailable values: \"full\", \"partial\", \"secondary\", \"internal\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"full",
						"partial",
						"secondary",
						"internal",
					),
				},
				Default: stringdefault.StaticString("full"),
			},
			"vanity_name_servers": schema.ListAttribute{
				Description: "An array of domains used for custom name servers. This is only\navailable for Business and Enterprise plans.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"activated_on": schema.StringAttribute{
				Description: "The last time proof of ownership was detected and the zone was made\nactive.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"cname_suffix": schema.StringAttribute{
				Description: "Allows the customer to use a custom apex.\n*Tenants Only Configuration*.",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the zone was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"development_mode": schema.Float64Attribute{
				Description: "The interval (in seconds) from when development mode expires\n(positive integer) or last expired (negative integer) for the\ndomain. If development mode has never been enabled, this value is 0.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the zone was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"original_dnshost": schema.StringAttribute{
				Description: "DNS host at the time of switching to Cloudflare.",
				Computed:    true,
			},
			"original_registrar": schema.StringAttribute{
				Description: "Registrar for the domain at the time of switching to Cloudflare.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The zone status on Cloudflare.\nAvailable values: \"initializing\", \"pending\", \"active\", \"moved\".",
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
			"verification_key": schema.StringAttribute{
				Description: "Verification key for partial zone setup.",
				Computed:    true,
			},
			"name_servers": schema.ListAttribute{
				Description: "The name servers Cloudflare assigns to a zone.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"original_name_servers": schema.ListAttribute{
				Description: "Original name servers before moving to Cloudflare.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"permissions": schema.ListAttribute{
				Description:        "Legacy permissions based on legacy user membership information.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         customfield.NewListType[types.String](ctx),
				ElementType:        types.StringType,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Metadata about the zone.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneMetaModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cdn_only": schema.BoolAttribute{
						Description: "The zone is only configured for CDN.",
						Computed:    true,
					},
					"custom_certificate_quota": schema.Int64Attribute{
						Description: "Number of Custom Certificates the zone can have.",
						Computed:    true,
					},
					"dns_only": schema.BoolAttribute{
						Description: "The zone is only configured for DNS.",
						Computed:    true,
					},
					"foundation_dns": schema.BoolAttribute{
						Description: "The zone is setup with Foundation DNS.",
						Computed:    true,
					},
					"page_rule_quota": schema.Int64Attribute{
						Description: "Number of Page Rules a zone can have.",
						Computed:    true,
					},
					"phishing_detected": schema.BoolAttribute{
						Description: "The zone has been flagged for phishing.",
						Computed:    true,
					},
					"step": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"owner": schema.SingleNestedAttribute{
				Description: "The owner of the zone.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneOwnerModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the owner.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of owner.",
						Computed:    true,
					},
				},
			},
			"plan": schema.SingleNestedAttribute{
				Description:        "A Zones subscription information.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         customfield.NewNestedObjectType[ZonePlanModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"can_subscribe": schema.BoolAttribute{
						Description: "States if the subscription can be activated.",
						Computed:    true,
					},
					"currency": schema.StringAttribute{
						Description: "The denomination of the customer.",
						Computed:    true,
					},
					"externally_managed": schema.BoolAttribute{
						Description: "If this Zone is managed by another company.",
						Computed:    true,
					},
					"frequency": schema.StringAttribute{
						Description: "How often the customer is billed.",
						Computed:    true,
					},
					"is_subscribed": schema.BoolAttribute{
						Description: "States if the subscription active.",
						Computed:    true,
					},
					"legacy_discount": schema.BoolAttribute{
						Description: "If the legacy discount applies to this Zone.",
						Computed:    true,
					},
					"legacy_id": schema.StringAttribute{
						Description: "The legacy name of the plan.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the owner.",
						Computed:    true,
					},
					"price": schema.Float64Attribute{
						Description: "How much the customer is paying.",
						Computed:    true,
					},
				},
			},
			"tenant": schema.SingleNestedAttribute{
				Description: "The root organizational unit that this zone belongs to (such as a tenant or organization).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneTenantModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the Tenant account.",
						Computed:    true,
					},
				},
			},
			"tenant_unit": schema.SingleNestedAttribute{
				Description: "The immediate parent organizational unit that this zone belongs to (such as under a tenant or sub-organization).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneTenantUnitModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
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
