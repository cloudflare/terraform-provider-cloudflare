// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type APIShieldOperationDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*APIShieldOperationDataSource)(nil)

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

	params, diags := data.toListParams()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	items := customfield.NullObjectList[APIShieldOperationDataSourceModel](ctx)
	env := APIShieldOperationResultListDataSourceEnvelope{items}

	page, err := d.client.APIGateway.Discovery.Operations.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes := []byte(page.JSON.RawJSON())
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
		return
	}

	if count := len(items.Elements()); count != 1 {
		resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
		return
	}
	ts, diags := items.AsStructSliceT(ctx)
	resp.Diagnostics.Append(diags...)
	data = &ts[0]

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
