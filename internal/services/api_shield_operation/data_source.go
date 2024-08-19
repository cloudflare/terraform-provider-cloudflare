// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type APIShieldOperationDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &APIShieldOperationDataSource{}

func NewAPIShieldOperationDataSource() datasource.DataSource {
	return &APIShieldOperationDataSource{}
}

func (d *APIShieldOperationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_shield_operation"
}

func (d *APIShieldOperationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *APIShieldOperationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *APIShieldOperationDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataFilterHost := []string{}
	for _, item := range *data.Filter.Host {
		dataFilterHost = append(dataFilterHost, item.ValueString())
	}
	dataFilterMethod := []string{}
	for _, item := range *data.Filter.Method {
		dataFilterMethod = append(dataFilterMethod, item.ValueString())
	}
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*APIShieldOperationDataSourceModel{}
	env := APIShieldOperationResultListDataSourceEnvelope{items}

	page, err := d.client.APIGateway.Discovery.Operations.List(ctx, api_gateway.DiscoveryOperationListParams{
		ZoneID:    cloudflare.F(data.Filter.ZoneID.ValueString()),
		Diff:      cloudflare.F(data.Filter.Diff.ValueBool()),
		Direction: cloudflare.F(api_gateway.DiscoveryOperationListParamsDirection(data.Filter.Direction.ValueString())),
		Endpoint:  cloudflare.F(data.Filter.Endpoint.ValueString()),
		Host:      cloudflare.F(dataFilterHost),
		Method:    cloudflare.F(dataFilterMethod),
		Order:     cloudflare.F(api_gateway.DiscoveryOperationListParamsOrder(data.Filter.Order.ValueString())),
		Origin:    cloudflare.F(api_gateway.DiscoveryOperationListParamsOrigin(data.Filter.Origin.ValueString())),
		State:     cloudflare.F(api_gateway.DiscoveryOperationListParamsState(data.Filter.State.ValueString())),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes := []byte(page.JSON.RawJSON())
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
		return
	}

	if count := len(*items); count != 1 {
		resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
		return
	}
	data = (*items)[0]

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
