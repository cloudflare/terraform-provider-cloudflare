package email_routing_rule

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *EmailRoutingRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Email Routing Rule](https://developers.cloudflare.com/email-routing/setup/email-routing-addresses/#email-rule-actions) resource allows you to create and manage email routing rules for a zone.
		`),
		Version: 1,

		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the email routing rule.",
			},
			"tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The tag of the email routing rule.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Routing rule name.",
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "The priority of the email routing rule.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the email routing rule is enabled.",
				Optional:            true,
			},
		},
		Blocks: map[string]schema.Block{
			"matcher": schema.SetNestedBlock{
				MarkdownDescription: "Matching patterns to forward to your actions.",
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
					setvalidator.IsRequired(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: fmt.Sprintf("Type of matcher. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"literal", "all"})),
							Validators: []validator.String{
								stringvalidator.OneOf("literal", "all"),
							},
						},
						"field": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Field to match on. Required for `type` of `literal`.",
						},
						"value": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Value to match on. Required for `type` of `literal`.",
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 90),
							},
						},
					},
				},
			},
			"action": schema.SetNestedBlock{
				MarkdownDescription: "Actions to take when a match is found.",
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
					setvalidator.IsRequired(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: fmt.Sprintf("Type of action. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"forward", "worker", "drop"})),
							Validators: []validator.String{
								stringvalidator.OneOf("forward", "worker", "drop"),
							},
						},
						"value": schema.SetAttribute{
							Optional:            true,
							ElementType:         types.StringType,
							MarkdownDescription: "Value to match on. Required for `type` of `literal`.",
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(stringvalidator.LengthBetween(0, 90)),
							},
						},
					},
				},
			},
		},
	}
}

func (r *EmailRoutingRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					consts.ZoneIDSchemaKey: schema.StringAttribute{
						MarkdownDescription: consts.ZoneIDSchemaDescription,
						Required:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The ID of the email routing rule.",
					},
					"tag": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The tag of the email routing rule.",
					},
					"name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Routing rule name.",
					},
					"priority": schema.Int64Attribute{
						MarkdownDescription: "The priority of the email routing rule.",
						Optional:            true,
						Computed:            true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: "Whether the email routing rule is enabled.",
						Optional:            true,
					},
				},
				Blocks: map[string]schema.Block{
					"matcher": schema.SetNestedBlock{
						MarkdownDescription: "Matching patterns to forward to your actions.",
						Validators: []validator.Set{
							setvalidator.SizeAtLeast(1),
							setvalidator.IsRequired(),
						},
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: fmt.Sprintf("Type of matcher. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"literal", "all"})),
									Validators: []validator.String{
										stringvalidator.OneOf("literal", "all"),
									},
								},
								"field": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Field to match on. Required for `type` of `literal`.",
								},
								"value": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Value to match on. Required for `type` of `literal`.",
									Validators: []validator.String{
										stringvalidator.LengthBetween(0, 90),
									},
								},
							},
						},
					},
					"action": schema.SetNestedBlock{
						MarkdownDescription: "Actions to take when a match is found.",
						Validators: []validator.Set{
							setvalidator.SizeAtLeast(1),
							setvalidator.IsRequired(),
						},
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: fmt.Sprintf("Type of action. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"forward", "worker", "drop"})),
									Validators: []validator.String{
										stringvalidator.OneOf("forward", "worker", "drop"),
									},
								},
								"value": schema.SetAttribute{
									Optional:            true,
									ElementType:         types.StringType,
									MarkdownDescription: "Value to match on. Required for `type` of `literal`.",
									Validators: []validator.Set{
										setvalidator.ValueStringsAre(stringvalidator.LengthBetween(0, 90)),
									},
								},
							},
						},
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData EmailRoutingRuleModel

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				upgradedStateData := EmailRoutingRuleModel{
					ZoneID:   priorStateData.ZoneID,
					ID:       priorStateData.ID,
					Tag:      priorStateData.ID,
					Name:     priorStateData.Name,
					Priority: priorStateData.Priority,
					Enabled:  priorStateData.Enabled,
					Action:   priorStateData.Action,
					Matcher:  priorStateData.Matcher,
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}
