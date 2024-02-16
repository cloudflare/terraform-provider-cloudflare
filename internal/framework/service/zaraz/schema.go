package zaraz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

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
						"triggers": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: triggerSchema().NestedObject.Attributes,
							},
							Required: true,
						},
						"tools": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mode": schema.ObjectAttribute{
										AttributeTypes: map[string]attr.Type{
											"light":  types.BoolType,
											"cloud":  types.BoolType,
											"sample": types.BoolType,
											"segment": types.MapType{
												ElemType: types.NumberType,
											},
											"trigger": types.ListType{
												ElemType: types.StringType,
											},
											"ignore_spa": types.BoolType,
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
									"categories": schema.ListAttribute{
										ElementType: types.StringType,
										Optional:    true,
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
									"neo_events": schema.ListNestedAttribute{
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
									"worker": schema.MapAttribute{
										ElementType: types.StringType,
										Optional:    true,
										Validators: []validator.Map{mapvalidator.KeysAre(
											stringvalidator.OneOf(
												"escaped_worker_name",
												"worker_tag",
												"mutable_id",
											),
										),
										},
									},
								},
							},
							Required: true,
							// ... potentially other fields ...
						},
					},
				},
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

func triggerRuleSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
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
				"settings": schema.MapAttribute{
					ElementType: types.StringType,
					Optional:    true,
					Validators: []validator.Map{mapvalidator.KeysAre(
						stringvalidator.OneOf(
							"type",
							"selector",
							"wait_for_tags",
							"interval",
							"limit",
							"validate",
							"positions",
							"variable",
							"match",
						),
					),
					},
				},
			},
		},
	}
}
