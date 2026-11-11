// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PrecursorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"mode": schema.StringAttribute{
				Description: "Controls Precursor detection for the zone. `off` disables detection,\n`monitor` records detections without affecting bot scores, and `arm`\nallows detections to affect bot scores.\nAvailable values: \"off\", \"monitor\", \"arm\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"monitor",
						"arm",
					),
				},
			},
			"show_interstitial": schema.BoolAttribute{
				Description: "Whether Precursor shows an interstitial when it takes action on a\nrequest.",
				Computed:    true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "The ordered list of Precursor rules for the zone.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PrecursorRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"expression": schema.StringAttribute{
							Description: "The filter expression that determines which requests the rule matches.",
							Computed:    true,
						},
						"show_interstitial": schema.BoolAttribute{
							Description: "Whether Precursor shows an interstitial when it takes action on a\nrequest.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "The read-only identifier that Cloudflare assigns to the rule.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the rule.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the rule is active.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *PrecursorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PrecursorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
