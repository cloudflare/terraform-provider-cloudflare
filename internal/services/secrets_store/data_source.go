// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type SecretsStoreDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*SecretsStoreDataSource)(nil)

func NewSecretsStoreDataSource() datasource.DataSource {
	return &SecretsStoreDataSource{}
}

func (d *SecretsStoreDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secrets_store"
}

func (d *SecretsStoreDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SecretsStoreDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SecretsStoreDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter != nil {
		params, diags := data.toListParams(ctx)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		env := SecretsStoresResultListDataSourceEnvelope{}
		page, err := d.client.SecretsStore.Stores.List(ctx, params)
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

		if count := len(env.Result.Elements()); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		ts, diags := env.Result.AsStructSliceT(ctx)
		resp.Diagnostics.Append(diags...)
		data.StoreID = ts[0].ID
	}

	params, diags := data.toReadParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := SecretsStoreResultDataSourceEnvelope{*data}
	_, err := d.client.SecretsStore.Stores.Get(
		ctx,
		data.StoreID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.StoreID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
