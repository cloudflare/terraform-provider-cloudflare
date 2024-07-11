// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type R2BucketDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &R2BucketDataSource{}

func NewR2BucketDataSource() datasource.DataSource {
	return &R2BucketDataSource{}
}

func (d *R2BucketDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_r2_bucket"
}

func (r *R2BucketDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *R2BucketDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *R2BucketDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy == nil {
		res := new(http.Response)
		env := R2BucketResultDataSourceEnvelope{*data}
		_, err := r.client.R2.Buckets.Get(
			ctx,
			data.BucketName.ValueString(),
			r2.BucketGetParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
			},
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}
		bytes, _ := io.ReadAll(res.Body)
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
			return
		}
		data = &env.Result
	} else {
		items := &[]*R2BucketDataSourceModel{}
		env := R2BucketResultListDataSourceEnvelope{items}

		page, err := r.client.R2.Buckets.List(ctx, r2.BucketListParams{
			AccountID:    cloudflare.F(data.FindOneBy.AccountID.ValueString()),
			Cursor:       cloudflare.F(data.FindOneBy.Cursor.ValueString()),
			Direction:    cloudflare.F(r2.BucketListParamsDirection(data.FindOneBy.Direction.ValueString())),
			NameContains: cloudflare.F(data.FindOneBy.NameContains.ValueString()),
			Order:        cloudflare.F(r2.BucketListParamsOrder(data.FindOneBy.Order.ValueString())),
			PerPage:      cloudflare.F(data.FindOneBy.PerPage.ValueFloat64()),
			StartAfter:   cloudflare.F(data.FindOneBy.StartAfter.ValueString()),
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
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
