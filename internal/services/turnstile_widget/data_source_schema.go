// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &TurnstileWidgetDataSource{}
var _ datasource.DataSourceWithValidateConfig = &TurnstileWidgetDataSource{}

func (r TurnstileWidgetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"sitekey": schema.StringAttribute{
				Description: "Widget item identifier tag.",
				Computed:    true,
				Optional:    true,
			},
			"secret": schema.StringAttribute{
				Description: "Secret key for this widget.",
				Optional:    true,
			},
			"bot_fight_mode": schema.BoolAttribute{
				Description: "If bot_fight_mode is set to `true`, Cloudflare issues computationally\nexpensive challenges in response to malicious bots (ENT only).\n",
				Computed:    true,
			},
			"clearance_level": schema.StringAttribute{
				Description: "If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,\nthis setting can determine the clearance level to be set\n",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("no_clearance", "jschallenge", "managed", "interactive"),
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the widget was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"mode": schema.StringAttribute{
				Description: "Widget Mode",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("non-interactive", "invisible", "managed"),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the widget was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Human readable widget name. Not unique. Cloudflare suggests that you\nset this to a meaningful string to make it easier to identify your\nwidget, and where it is used.\n",
				Computed:    true,
			},
			"offlabel": schema.BoolAttribute{
				Description: "Do not show any Cloudflare branding on the widget (ENT only).\n",
				Computed:    true,
			},
			"region": schema.StringAttribute{
				Description: "Region where this widget can be used.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("world"),
				},
			},
			"domains": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Direction to order widgets.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Field to order widgets by.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "sitekey", "name", "created_on", "modified_on"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.AtLeast(1),
						},
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of items per page.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(5, 1000),
						},
					},
				},
			},
		},
	}
}

func (r *TurnstileWidgetDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *TurnstileWidgetDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
