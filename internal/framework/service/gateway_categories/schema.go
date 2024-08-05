package gateway_categories

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func (d *CloudflareGatewayCategoriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use this data source to retrieve all Gateway categories for an account.",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:    true,
				Description: "The account ID to fetch Gateway Categories from.",
			},
			"categories": schema.ListNestedAttribute{
				Computed:    true,
				Description: "A list of Gateway Categories.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "The identifier for this category. There is only one category per ID.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of the category.",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "A short summary of domains in the category.",
						},
						"class": schema.StringAttribute{
							Computed:    true,
							Description: "Which account types are allowed to create policies based on this category.",
						},
						"beta": schema.BoolAttribute{
							Computed:    true,
							Description: "True if the category is in beta and subject to change.",
						},
						"subcategories": schema.ListNestedAttribute{
							Computed:    true,
							Description: "A list of subcategories.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Computed:    true,
										Description: "The identifier for this subcategory. There is only one subcategory per ID.",
									},
									"name": schema.StringAttribute{
										Computed:    true,
										Description: "The name of the subcategory.",
									},
									"description": schema.StringAttribute{
										Computed:    true,
										Description: "A short summary of domains in the subcategory.",
									},
									"class": schema.StringAttribute{
										Computed:    true,
										Description: "Which account types are allowed to create policies based on this subcategory.",
									},
									"beta": schema.BoolAttribute{
										Computed:    true,
										Description: "True if the subcategory is in beta and subject to change.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
