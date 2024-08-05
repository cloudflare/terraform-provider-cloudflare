// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/custom_hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type CustomHostnamesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &CustomHostnamesDataSource{}

func NewCustomHostnamesDataSource() datasource.DataSource {
	return &CustomHostnamesDataSource{}
}

func (d *CustomHostnamesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_hostnames"
}

func (d *CustomHostnamesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CustomHostnamesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CustomHostnamesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*CustomHostnamesResultDataSourceModel{}
	env := CustomHostnamesResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*CustomHostnamesResultDataSourceModel{}

	page, err := d.client.CustomHostnames.List(ctx, custom_hostnames.CustomHostnameListParams{
		ZoneID:    cloudflare.F(data.ZoneID.ValueString()),
		ID:        cloudflare.F(data.ID.ValueString()),
		Direction: cloudflare.F(custom_hostnames.CustomHostnameListParamsDirection(data.Direction.ValueString())),
		Hostname:  cloudflare.F(data.Hostname.ValueString()),
		Order:     cloudflare.F(custom_hostnames.CustomHostnameListParamsOrder(data.Order.ValueString())),
		Page:      cloudflare.F(data.Page.ValueFloat64()),
		PerPage:   cloudflare.F(data.PerPage.ValueFloat64()),
		SSL:       cloudflare.F(custom_hostnames.CustomHostnameListParamsSSL(data.SSL.ValueFloat64())),
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
