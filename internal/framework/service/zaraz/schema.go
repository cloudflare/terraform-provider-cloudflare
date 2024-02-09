package zaraz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type ZarazConfig struct {
	DebugKey string               `json:"debugKey"`
	Tools    map[string]ZarazTool `json:"tools"`
}

type ZarazTool struct {
	Enabled   *bool  `json:"enabled"`
	Component string `json:"component,omitempty"`
}

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
		},
		Blocks: map[string]schema.Block{
			"config": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"debug_key": schema.StringAttribute{
							Required: true,
						},
						"tools": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "",
										// ... potentially other fields ...
									},
									"name": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "",
										// ... potentially other fields ...
									},
									"default_fields": schema.MapAttribute{
										ElementType: types.StringType,
										Required:    true,
									},
									"default_purpose": schema.StringAttribute{
										Optional: true,
									},
									"library": schema.StringAttribute{
										Optional: true,
									},
									"component": schema.StringAttribute{
										Required: true,
										// ... potentially other fields ...
									},
									"permissions": schema.ListAttribute{
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
										Required: true,
									},
									"type": schema.StringAttribute{
										Required: true,
									},
									"worker": schema.MapAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
								},
							},
							Required: true,
							// ... potentially other fields ...
						},
						"triggers": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: triggerSchema().NestedObject.Attributes,
							},
							Required: true,
						},
					},
				},
			},
		},
	}
}

func triggerSchema() schema.MapNestedAttribute {
	return schema.MapNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Optional: true,
				},
				"description": schema.StringAttribute{
					Optional: true,
				},
				"load_rules": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: triggerRuleSchema().NestedObject.Attributes,
					},
					Required: true,
				},
				"exclude_rules": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: triggerRuleSchema().NestedObject.Attributes,
					},
					Required: true,
				},
			},
		},
	}
}

func triggerRuleSchema() schema.MapNestedAttribute {
	return schema.MapNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Required: true,
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
				"settings": schema.StringAttribute{
					Optional: true,
				},
			},
		},
	}
}

func clientRuleSettings() schema.MapNestedAttribute {
	return schema.MapNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					Required: false,
					Validators: []validator.String{
						stringvalidator.OneOf(
							"xpath",
							"css",
						),
					},
					Description: "ClickListener setting",
				},
				"selector": schema.StringAttribute{
					Required:    false,
					Description: "Click Listener/Element Visibility/Form Submission/ setting",
				},
				"wait_for_tags": schema.StringAttribute{
					Required:    false,
					Description: "ClickListener setting",
				},
				"interval": schema.Int64Attribute{
					Required:    false,
					Description: "Timer setting",
				},
				"limit": schema.Int64Attribute{
					Required:    false,
					Description: "Timer setting",
				},
				"validate": schema.BoolAttribute{
					Required:    false,
					Description: "Form Submission setting",
				},
				"positions": schema.StringAttribute{
					Required:    false,
					Description: "Scroll Depth setting",
				},
				"variable": schema.StringAttribute{
					Required:    false,
					Description: "Variable Match setting",
				},
				"match": schema.StringAttribute{
					Required:    false,
					Description: "Variable Match setting",
				},
			},
		},
	}
}

// func configSchema() schema.MapAttribute {
// 	return schema.MapAttribute{
// 		ElementType: types.StringType,
// 		Required:    true,
// 		// ... potentially other fields ...
// 	}
// }

// func configToolSchema() schema.MapNestedAttribute {
// 	return schema.MapNestedAttribute{
// 		NestedObject: schema.NestedAttributeObject{
// 			Attributes: map[string]schema.Attribute{
// 				"enabled": schema.BoolAttribute{
// 					Required: true,
// 				},
// 				"name": schema.StringAttribute{
// 					Required: true,
// 				},
// 			},
// 		},
// 		Optional: true,
// 	}
// }
