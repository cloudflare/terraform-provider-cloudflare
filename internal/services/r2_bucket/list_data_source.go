// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type R2BucketsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &R2BucketsDataSource{}

func NewR2BucketsDataSource() datasource.DataSource {
	return &R2BucketsDataSource{}
}

func (d *R2BucketsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_r2_buckets"
}

func (r *R2BucketsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *R2BucketsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *R2BucketsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*R2BucketsResultDataSourceModel{}
	env := R2BucketsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*R2BucketsResultDataSourceModel{}

	page, err := r.client.R2.Buckets.List(ctx, r2.BucketListParams{
		AccountID:    cloudflare.F(data.AccountID.ValueString()),
		Cursor:       cloudflare.F(data.Cursor.ValueString()),
		Direction:    cloudflare.F(r2.BucketListParamsDirection(data.Direction.ValueString())),
		NameContains: cloudflare.F(data.NameContains.ValueString()),
		Order:        cloudflare.F(r2.BucketListParamsOrder(data.Order.ValueString())),
		PerPage:      cloudflare.F(data.PerPage.ValueFloat64()),
		StartAfter:   cloudflare.F(data.StartAfter.ValueString()),
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
