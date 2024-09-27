// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*RegistrarDomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"domain_name": schema.StringAttribute{
				Description: "Domain name.",
				Optional:    true,
			},
			"success": schema.BoolAttribute{
				Description: "Whether the API call was successful",
				Optional:    true,
			},
			"errors": schema.ListNestedAttribute{
				Optional: true,
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
				Optional: true,
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
			"result_info": schema.SingleNestedAttribute{
				Optional: true,
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
			"result": schema.StringAttribute{
				Optional:   true,
				CustomType: jsontypes.NormalizedType{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *RegistrarDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *RegistrarDomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("domain_name")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("domain_name")),
	}
}
