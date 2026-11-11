// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AccessRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Firewall Access Rules Read",
				"Account Firewall Access Rules Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique identifier of the IP Access rule.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.\nAvailable values: \"block\", \"challenge\", \"whitelist\", \"js_challenge\", \"managed_challenge\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"block",
						"challenge",
						"whitelist",
						"js_challenge",
						"managed_challenge",
					),
				},
			},
			"configuration": schema.SingleNestedAttribute{
				Description: "The rule configuration.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"target": schema.StringAttribute{
						Description: "For `target: ip`: The configuration target. You must set the target to `ip` when specifying an IP address in the rule.\n\nFor `target: ip6`: The configuration target. You must set the target to `ip6` when specifying an IPv6 address in the rule.\n\nFor `target: ip_range`: The configuration target. You must set the target to `ip_range` when specifying an IP address range in the rule.\n\nFor `target: asn`: The configuration target. You must set the target to `asn` when specifying an Autonomous System Number (ASN) in the rule.\n\nFor `target: country`: The configuration target. You must set the target to `country` when specifying a country code in the rule.\nAvailable values: \"ip\", \"ip6\", \"ip_range\", \"asn\", \"country\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ip",
								"ip6",
								"ip_range",
								"asn",
								"country",
							),
						},
					},
					"value": schema.StringAttribute{
						Description: "For `target: ip`: The IP address to match. This address will be compared to the IP address of incoming requests.\n\nFor `target: ip6`: The IPv6 address to match.\n\nFor `target: ip_range`: The IP address range to match. You can only use prefix lengths `/16` and `/24` for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.\n\nFor `target: asn`: The AS number to match.\n\nFor `target: country`: The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).",
						Optional:    true,
						Validators: []validator.String{
							ipv6Validator(),
							cidrValidator(),
						},
					},
				},
			},
			"notes": schema.StringAttribute{
				Description: "An informative summary of the rule, typically used as a reminder or explanation.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"created_on": schema.StringAttribute{
				Description:   "The timestamp of when the rule was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the rule was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"allowed_modes": schema.ListAttribute{
				Description: "The available actions that a rule can apply to a matched request.",
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"block",
							"challenge",
							"whitelist",
							"js_challenge",
							"managed_challenge",
						),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"scope": schema.SingleNestedAttribute{
				Description: "All zones owned by the user will have the rule applied.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AccessRuleScopeModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Defines an identifier.",
						Computed:    true,
					},
					"email": schema.StringAttribute{
						Description: "The contact email address of the user.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "Defines the scope of the rule.\nAvailable values: \"user\", \"organization\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("user", "organization"),
						},
					},
				},
			},
		},
	}
}

func (r *AccessRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccessRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
