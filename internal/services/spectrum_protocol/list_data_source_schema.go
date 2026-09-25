// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_protocol

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*SpectrumProtocolsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zone Settings Read",
				"Zone Settings Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Zone identifier.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[SpectrumProtocolsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Description: "The full name of the application protocol.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The short name of the application protocol.",
							Computed:    true,
						},
						"ports": schema.ListAttribute{
							Description: "The available listening ports for the given protocol.",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueInt64sAre(
									int64validator.Between(1, 65535),
								),
							},
							CustomType:  customfield.NewListType[types.Int64](ctx),
							ElementType: types.Int64Type,
						},
						"transport": schema.StringAttribute{
							Description: "The transport layer protocol used by the application protocol",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *SpectrumProtocolsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *SpectrumProtocolsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
