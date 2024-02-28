package zaraz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *ZarazConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Zaraz Config](https://developers.cloudflare.com/zaraz/) resource allows you to manage your Cloudflare Zaraz config.
	`),

		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.ZoneIDSchemaKey))),
					),
				},
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.AccountIDSchemaKey))),
					),
				},
			},
			"config": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"debug_key": schema.StringAttribute{
						Required: true,
					},
					"zaraz_version": schema.Int64Attribute{
						Required: true,
					},
					"triggers": schema.MapNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: triggerSchema().NestedObject.Attributes,
						},
						Required: true,
					},
					"settings": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"auto_inject_script": schema.BoolAttribute{
								Optional: true,
							},
							"inject_iframes": schema.BoolAttribute{
								Optional: true,
							},
							"ecommerce": schema.BoolAttribute{
								Optional: true,
							},
							"hide_query_params": schema.BoolAttribute{
								Optional: true,
							},
							"hide_ip_address": schema.BoolAttribute{
								Optional: true,
							},
							"hide_user_agent": schema.BoolAttribute{
								Optional: true,
							},
							"hide_external_referer": schema.BoolAttribute{
								Optional: true,
							},
							"cookie_domain": schema.StringAttribute{
								Optional: true,
							},
							"init_path": schema.StringAttribute{
								Optional: true,
							},
							"script_path": schema.StringAttribute{
								Optional: true,
							},
							"track_path": schema.StringAttribute{
								Optional: true,
							},
							"events_api_path": schema.StringAttribute{
								Optional: true,
							},
							"mc_root_path": schema.StringAttribute{
								Optional: true,
							},
							"context_enricher": schema.StringAttribute{
								Optional: true,
							},
						},
						Optional: true,
					},
					"history_change": schema.BoolAttribute{
						Optional: true,
					},
					"tools": schema.MapNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mode": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"light": schema.BoolAttribute{
											Optional: true,
										},
										"cloud": schema.BoolAttribute{
											Optional: true,
										},
										"sample": schema.BoolAttribute{
											Optional: true,
										},
										"segment": schema.MapAttribute{
											ElementType: types.Float64Type,
											Optional:    true,
										},
										"trigger": schema.StringAttribute{
											Optional: true,
										},
										"ignore_spa": schema.BoolAttribute{
											Optional: true,
										},
									},
									Optional: true,
								},
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "",
									// ... potentially other fields ...
								},
								"type": schema.StringAttribute{
									Required: true,
								},
								"enabled": schema.BoolAttribute{
									Optional:            true,
									MarkdownDescription: "",
									// ... potentially other fields ...
								},

								"default_fields": schema.MapAttribute{
									ElementType: types.StringType,
									Required:    true,
								},
								"settings": schema.MapAttribute{
									ElementType: types.StringType,
									Required:    true,
									// TODO QQ how do we set the type to any ???
								},
								"actions": schema.MapNestedAttribute{
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"blocking_triggers": schema.ListAttribute{
												ElementType: types.StringType,
												Required:    true,
											},
											"firing_triggers": schema.ListAttribute{
												ElementType: types.StringType,
												Required:    true,
											},
											"data": schema.MapAttribute{
												ElementType: types.StringType,
												Required:    true,
											},
											"action_type": schema.StringAttribute{
												Required: true,
											},
										},
									},
									Optional: true,
								},
								"default_purpose": schema.StringAttribute{
									Optional: true,
								},
								"blocking_triggers": schema.ListAttribute{
									ElementType: types.StringType,
									Required:    true,
								},
								"library": schema.StringAttribute{
									Optional: true,
								},
								"component": schema.StringAttribute{
									Optional: true,
								},
								"permissions": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"worker": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"escaped_worker_name": schema.StringAttribute{
											Required: true,
										},
										"worker_tag": schema.StringAttribute{
											Required: true,
										},
										"mutable_id": schema.StringAttribute{
											Optional: true,
										},
									},
									Optional: true,
								},
							},
						},
						Required: true,
						// ... potentially other fields ...
					},
				},
				Required: true,
			},
		},
	}
}

func triggerSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Optional: true,
				},
				"description": schema.StringAttribute{
					Optional: true,
				},
				"system": schema.StringAttribute{
					Optional: true,
				},
				"load_rules": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: triggerRuleSchema().NestedObject.Attributes,
					},
					Required: true,
				},
			},
		},
	}
}

func triggerRuleSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Optional: true,
				},
				"match": schema.StringAttribute{
					Optional: true,
				},
				"value": schema.StringAttribute{
					Optional: true,
				},
				"op": schema.StringAttribute{
					Optional: true,
				},
				"action": schema.StringAttribute{
					Optional: true,
				},
				"settings": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional: true,
						},
						"selector": schema.StringAttribute{
							Optional: true,
						},
						"wait_for_tags": schema.Int64Attribute{
							Optional: true,
						},
						"interval": schema.Int64Attribute{
							Optional: true,
						},
						"limit": schema.Int64Attribute{
							Optional: true,
						},
						"validate": schema.BoolAttribute{
							Optional: true,
						},
						"variable": schema.StringAttribute{
							Optional: true,
						},
						"match": schema.StringAttribute{
							Optional: true,
						},
						"positions": schema.StringAttribute{
							Optional: true,
						},
						"op": schema.StringAttribute{
							Optional: true,
						},
						"value": schema.StringAttribute{
							Optional: true,
						},
					},
					Optional: true,
				},
			},
		},
	}
}
