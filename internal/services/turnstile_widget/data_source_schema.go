// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*TurnstileWidgetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Widget item identifier tag.",
				Computed:    true,
			},
			"sitekey": schema.StringAttribute{
				Description: "Widget item identifier tag.",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"bot_fight_mode": schema.BoolAttribute{
				Description: "If bot_fight_mode is set to `true`, Cloudflare issues computationally\nexpensive challenges in response to malicious bots (ENT only).",
				Computed:    true,
			},
			"clearance_level": schema.StringAttribute{
				Description: "If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,\nthis setting can determine the clearance level to be set\nAvailable values: \"no_clearance\", \"jschallenge\", \"managed\", \"interactive\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"no_clearance",
						"jschallenge",
						"managed",
						"interactive",
					),
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the widget was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"ephemeral_id": schema.BoolAttribute{
				Description: "Return the Ephemeral ID in /siteverify (ENT only).",
				Computed:    true,
			},
			"mode": schema.StringAttribute{
				Description: "Widget Mode\nAvailable values: \"non-interactive\", \"invisible\", \"managed\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"non-interactive",
						"invisible",
						"managed",
					),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the widget was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Human readable widget name. Not unique. Cloudflare suggests that you\nset this to a meaningful string to make it easier to identify your\nwidget, and where it is used.",
				Computed:    true,
			},
			"offlabel": schema.BoolAttribute{
				Description: "Do not show any Cloudflare branding on the widget (ENT only).",
				Computed:    true,
			},
			"region": schema.StringAttribute{
				Description: "Region where this widget can be used.\nAvailable values: \"world\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("world"),
				},
			},
			"secret": schema.StringAttribute{
				Description: "Secret key for this widget.",
				Computed:    true,
			},
			"domains": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Direction to order widgets.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Field to order widgets by.\nAvailable values: \"id\", \"sitekey\", \"name\", \"created_on\", \"modified_on\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"id",
								"sitekey",
								"name",
								"created_on",
								"modified_on",
							),
						},
					},
				},
			},
		},
	}
}

func (d *TurnstileWidgetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TurnstileWidgetDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("sitekey"), path.MatchRoot("filter")),
	}
}
