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
			"default_mode": schema.StringAttribute{
				Description: "The zone-level Precursor enforcement mode applied to requests that do\nnot match a more specific enforcement rule.\nAvailable values: \"off\", \"min-friction\", \"max-security\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"min-friction",
						"max-security",
					),
				},
			},
			"enforcement_rules": schema.ListNestedAttribute{
				Description: "The ordered list of enforcement rules for the zone.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PrecursorEnforcementRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"expression": schema.StringAttribute{
							Description: "The filter expression that determines which requests the rule matches.",
							Computed:    true,
						},
						"mode": schema.StringAttribute{
							Description: "The override mode Precursor applies to requests matching an enforcement\nrule. Unlike `default_mode`, this cannot be `off`.\nAvailable values: \"min-friction\", \"max-security\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("min-friction", "max-security"),
							},
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
