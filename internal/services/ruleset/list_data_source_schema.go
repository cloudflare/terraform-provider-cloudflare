// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &RulesetsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &RulesetsDataSource{}

func (r RulesetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the ruleset.",
							Computed:    true,
						},
						"kind": schema.StringAttribute{
							Description: "The kind of the ruleset.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("managed", "custom", "root", "zone"),
							},
						},
						"last_updated": schema.StringAttribute{
							Description: "The timestamp of when the ruleset was last modified.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The human-readable name of the ruleset.",
							Computed:    true,
						},
						"phase": schema.StringAttribute{
							Description: "The phase of the ruleset.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("ddos_l4", "ddos_l7", "http_config_settings", "http_custom_errors", "http_log_custom_fields", "http_ratelimit", "http_request_cache_settings", "http_request_dynamic_redirect", "http_request_firewall_custom", "http_request_firewall_managed", "http_request_late_transform", "http_request_origin", "http_request_redirect", "http_request_sanitize", "http_request_sbfm", "http_request_select_configuration", "http_request_transform", "http_response_compression", "http_response_firewall_managed", "http_response_headers_transform", "magic_transit", "magic_transit_ids_managed", "magic_transit_managed"),
							},
						},
						"version": schema.StringAttribute{
							Description: "The version of the ruleset.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the ruleset.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *RulesetsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *RulesetsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
