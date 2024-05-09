// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r RulesetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique ID of the ruleset.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"kind": schema.StringAttribute{
				Description: "The kind of the ruleset.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("managed", "custom", "root", "zone"),
				},
			},
			"name": schema.StringAttribute{
				Description: "The human-readable name of the ruleset.",
				Required:    true,
			},
			"phase": schema.StringAttribute{
				Description: "The phase of the ruleset.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ddos_l4", "ddos_l7", "http_config_settings", "http_custom_errors", "http_log_custom_fields", "http_ratelimit", "http_request_cache_settings", "http_request_dynamic_redirect", "http_request_firewall_custom", "http_request_firewall_managed", "http_request_late_transform", "http_request_origin", "http_request_redirect", "http_request_sanitize", "http_request_sbfm", "http_request_select_configuration", "http_request_transform", "http_response_compression", "http_response_firewall_managed", "http_response_headers_transform", "magic_transit", "magic_transit_ids_managed", "magic_transit_managed"),
				},
			},
			"rules": schema.ListNestedAttribute{
				Description: "The list of rules in the ruleset.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"last_updated": schema.StringAttribute{
							Description: "The timestamp of when the rule was last modified.",
							Computed:    true,
						},
						"version": schema.StringAttribute{
							Description: "The version of the rule.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "The unique ID of the rule.",
							Optional:    true,
						},
						"action": schema.StringAttribute{
							Description: "The action to perform when the rule matches.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("block", "challenge", "compress_response", "execute", "js_challenge", "log", "managed_challenge", "redirect", "rewrite", "route", "score", "serve_error", "set_config", "skip", "set_cache_settings", "log_custom_field"),
							},
						},
						"action_parameters": schema.SingleNestedAttribute{
							Description: "The parameters configuring the rule's action.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"response": schema.SingleNestedAttribute{
									Description: "The response to show when the block is applied.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description: "The content to return.",
											Required:    true,
										},
										"content_type": schema.StringAttribute{
											Description: "The type of the content to return.",
											Required:    true,
										},
										"status_code": schema.Int64Attribute{
											Description: "The status code to return.",
											Required:    true,
										},
									},
								},
							},
						},
						"categories": schema.StringAttribute{
							Description: "The categories of the rule.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the rule.",
							Computed:    true,
							Default:     stringdefault.StaticString(""),
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the rule should be executed.",
							Computed:    true,
							Default:     booldefault.StaticBool(true),
						},
						"expression": schema.StringAttribute{
							Description: "The expression defining which traffic will match the rule.",
							Optional:    true,
						},
						"logging": schema.SingleNestedAttribute{
							Description: "An object configuring the rule's logging behavior.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether to generate a log when the rule matches.",
									Required:    true,
								},
							},
						},
						"ref": schema.StringAttribute{
							Description: "The reference of the rule (the rule ID by default).",
							Optional:    true,
						},
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "An informative description of the ruleset.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
		},
	}
}
