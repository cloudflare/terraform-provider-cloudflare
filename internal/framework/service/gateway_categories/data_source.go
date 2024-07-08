package gateway_categories

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &CloudflareGatewayCategoriesDataSource{}

func NewDataSource() datasource.DataSource {
	return &CloudflareGatewayCategoriesDataSource{}
}

type CloudflareGatewayCategoriesDataSource struct {
	client *muxclient.Client
}

func (d *CloudflareGatewayCategoriesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_categories"
}

func (d *CloudflareGatewayCategoriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *CloudflareGatewayCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GatewayCategoriesModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	params := zero_trust.GatewayCategoryListParams{
		AccountID: cloudflare.F(data.AccountID.ValueString()),
	}

	// Create a new Gateway Category Service
	service := zero_trust.NewGatewayCategoryService(d.client.V2.Options...)

	// Retrieve categories - v2 SDK way
	iter := service.ListAutoPaging(ctx, params)
	var categories []GatewayCategoryModel

	for iter.Next() {
		category := iter.Current()

		// Map subcategories
		var subcategories []GatewaySubCategoryModel
		for _, subcategory := range category.Subcategories {
			subcategories = append(subcategories, GatewaySubCategoryModel{
				ID:          types.Int64Value(subcategory.ID),
				Name:        types.StringValue(subcategory.Name),
				Description: types.StringValue(subcategory.Description),
				Class:       types.StringValue(string(subcategory.Class)),
				Beta:        types.BoolValue(subcategory.Beta),
			})
		}

		categories = append(categories, GatewayCategoryModel{
			ID:            types.Int64Value(category.ID),
			Name:          types.StringValue(category.Name),
			Description:   types.StringValue(category.Description),
			Class:         types.StringValue(string(category.Class)),
			Beta:          types.BoolValue(category.Beta),
			Subcategories: subcategories,
		})
	}
	if err := iter.Err(); err != nil {
		resp.Diagnostics.AddError("Failed to fetch Gateway Categories", err.Error())
		return
	}

	data.Categories = categories
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
