// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type APIShieldSchemasDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &APIShieldSchemasDataSource{}

func NewAPIShieldSchemasDataSource() datasource.DataSource {
	return &APIShieldSchemasDataSource{}
}

func (d *APIShieldSchemasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_shield_schemas"
}

func (r *APIShieldSchemasDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *APIShieldSchemasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *APIShieldSchemasDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*APIShieldSchemasResultDataSourceModel{}
	env := APIShieldSchemasResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*APIShieldSchemasResultDataSourceModel{}

	page, err := r.client.APIGateway.UserSchemas.List(ctx, api_gateway.UserSchemaListParams{
		ZoneID:            cloudflare.F(data.ZoneID.ValueString()),
		OmitSource:        cloudflare.F(data.OmitSource.ValueBool()),
		Page:              cloudflare.F(data.Page.ValueInt64()),
		PerPage:           cloudflare.F(data.PerPage.ValueInt64()),
		ValidationEnabled: cloudflare.F(data.ValidationEnabled.ValueBool()),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Result) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, *items...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:maxItems]
	data.Result = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
