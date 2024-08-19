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

type APIShieldOperationsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &APIShieldOperationsDataSource{}

func NewAPIShieldOperationsDataSource() datasource.DataSource {
	return &APIShieldOperationsDataSource{}
}

func (d *APIShieldOperationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_shield_operations"
}

func (d *APIShieldOperationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *APIShieldOperationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *APIShieldOperationsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataHost := []string{}
	for _, item := range *data.Host {
		dataHost = append(dataHost, item.ValueString())
	}
	dataMethod := []string{}
	for _, item := range *data.Method {
		dataMethod = append(dataMethod, item.ValueString())
	}
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*APIShieldOperationsResultDataSourceModel{}
	env := APIShieldOperationsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*APIShieldOperationsResultDataSourceModel{}

	page, err := d.client.APIGateway.Discovery.Operations.List(ctx, api_gateway.DiscoveryOperationListParams{
		ZoneID:    cloudflare.F(data.ZoneID.ValueString()),
		Diff:      cloudflare.F(data.Diff.ValueBool()),
		Direction: cloudflare.F(api_gateway.DiscoveryOperationListParamsDirection(data.Direction.ValueString())),
		Endpoint:  cloudflare.F(data.Endpoint.ValueString()),
		Host:      cloudflare.F(dataHost),
		Method:    cloudflare.F(dataMethod),
		Order:     cloudflare.F(api_gateway.DiscoveryOperationListParamsOrder(data.Order.ValueString())),
		Origin:    cloudflare.F(api_gateway.DiscoveryOperationListParamsOrigin(data.Origin.ValueString())),
		State:     cloudflare.F(api_gateway.DiscoveryOperationListParamsState(data.State.ValueString())),
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
