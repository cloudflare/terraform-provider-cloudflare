// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*RegistrarDomainsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
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
				CustomType:  customfield.NewNestedObjectListType[RegistrarDomainsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"errors": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[RegistrarDomainsErrorsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"code": schema.Int64Attribute{
										Computed: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1000),
										},
									},
									"message": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"messages": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[RegistrarDomainsMessagesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"code": schema.Int64Attribute{
										Computed: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1000),
										},
									},
									"message": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"result": schema.StringAttribute{
							Computed:   true,
							CustomType: jsontypes.NormalizedType{},
						},
						"success": schema.BoolAttribute{
							Description: "Whether the API call was successful",
							Computed:    true,
						},
						"result_info": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[RegistrarDomainsResultInfoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"count": schema.Float64Attribute{
									Description: "Total number of results for the requested service",
									Computed:    true,
								},
								"page": schema.Float64Attribute{
									Description: "Current page within paginated list of results",
									Computed:    true,
								},
								"per_page": schema.Float64Attribute{
									Description: "Number of results per page of results",
									Computed:    true,
								},
								"total_count": schema.Float64Attribute{
									Description: "Total results available without any search parameters",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *RegistrarDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *RegistrarDomainsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
