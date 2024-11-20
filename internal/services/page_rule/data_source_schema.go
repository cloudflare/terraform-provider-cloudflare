// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PageRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"pagerule_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the rule, used to define which Page Rule is processed\nover another. A higher number indicates a higher priority. For example,\nif you have a catch-all Page Rule (rule A: `/images/*`) but want a more\nspecific Page Rule to take precedence (rule B: `/images/special/*`),\nspecify a higher priority for rule B so it overrides rule A.\n",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the Page Rule.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "disabled"),
				},
			},
			"actions": schema.ListNestedAttribute{
				Description: "The set of actions to perform if the targets of this rule match the\nrequest. Actions can redirect to another URL or override settings, but\nnot both.\n",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PageRuleActionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "If enabled, any `http://`` URL is converted to `https://` through a\n301 redirect.\n",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"always_use_https",
									"automatic_https_rewrites",
									"browser_cache_ttl",
									"browser_check",
									"bypass_cache_on_cookie",
									"cache_by_device_type",
									"cache_deception_armor",
									"cache_key",
									"cache_key_fields",
									"cache_level",
									"cache_on_cookie",
									"cache_ttl_by_status",
									"ddos_protection",
									"development_mode",
									"disable_apps",
									"disable_performance",
									"disable_security",
									"disable_zaraz",
									"edge_cache_ttl",
									"email_obfuscation",
									"explicit_cache_control",
									"forwarding_url",
									"host_header_override",
									"hotlink_protection",
									"ip_geolocation",
									"minify",
									"mirage",
									"opportunistic_encryption",
									"origin_error_page_pass_thru",
									"polish",
									"purge_by_page_rule",
									"resolve_override",
									"respect_strong_etag",
									"response_buffering",
									"rocket_loader",
									"security_level",
									"server_side_exclude",
									"sort_query_string_for_cache",
									"ssl",
									"true_client_ip_header",
									"waf",
								),
							},
						},
						"value": schema.Int64Attribute{
							Description: "The number of seconds to cache resources for. The API prohibits\nsetting this to 0 for non-Enterprise domains.\n",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(0, 31536000),
							},
						},
					},
				},
			},
			"targets": schema.ListNestedAttribute{
				Description: "The rule targets to evaluate on each request.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PageRuleTargetsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"constraint": schema.SingleNestedAttribute{
							Description: "String constraint.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[PageRuleTargetsConstraintDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"operator": schema.StringAttribute{
									Description: "The matches operator can use asterisks and pipes as wildcard and 'or' operators.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"matches",
											"contains",
											"equals",
											"not_equal",
											"not_contain",
										),
									},
								},
								"value": schema.StringAttribute{
									Description: "The URL pattern to match against the current request. The pattern may contain up to four asterisks ('*') as placeholders.",
									Computed:    true,
								},
							},
						},
						"target": schema.StringAttribute{
							Description: "A target based on the URL of the request.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("url"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *PageRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PageRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
