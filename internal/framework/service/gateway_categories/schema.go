package gateway_categories

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (d *CloudflareGatewayCategoriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use this data source to retrieve all Gateway categories for an account.",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:    true,
				Description: "The account ID to fetch Gateway Categories from.",
			},
			"categories": schema.ListAttribute{
				Computed: true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":          types.Int64Type,
						"name":        types.StringType,
						"description": types.StringType,
						"class":       types.StringType,
						"beta":        types.BoolType,
					},
				},
				Description: "A list of Gateway Categories.",
			},
		},
	}
}
