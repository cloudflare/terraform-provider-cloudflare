// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type TunnelsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &TunnelsDataSource{}

func NewTunnelsDataSource() datasource.DataSource {
	return &TunnelsDataSource{}
}

func (d *TunnelsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnels"
}

func (r *TunnelsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TunnelsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TunnelsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataExistedAt, err := time.Parse(time.RFC3339, data.ExistedAt.ValueString())
	resp.Diagnostics.AddError("failed to parse time", err.Error())
	dataWasActiveAt, err := time.Parse(time.RFC3339, data.WasActiveAt.ValueString())
	resp.Diagnostics.AddError("failed to parse time", err.Error())
	dataWasInactiveAt, err := time.Parse(time.RFC3339, data.WasInactiveAt.ValueString())
	resp.Diagnostics.AddError("failed to parse time", err.Error())
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*TunnelsItemsDataSourceModel{}
	env := TunnelsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*TunnelsItemsDataSourceModel{}

	page, err := r.client.ZeroTrust.Tunnels.List(ctx, zero_trust.TunnelListParams{
		AccountID:     cloudflare.F(data.AccountID.ValueString()),
		ExcludePrefix: cloudflare.F(data.ExcludePrefix.ValueString()),
		ExistedAt:     cloudflare.F(dataExistedAt),
		IncludePrefix: cloudflare.F(data.IncludePrefix.ValueString()),
		IsDeleted:     cloudflare.F(data.IsDeleted.ValueBool()),
		Name:          cloudflare.F(data.Name.ValueString()),
		Page:          cloudflare.F(data.Page.ValueFloat64()),
		PerPage:       cloudflare.F(data.PerPage.ValueFloat64()),
		Status:        cloudflare.F(zero_trust.TunnelListParamsStatus(data.Status.ValueString())),
		TunTypes:      cloudflare.F(data.TunTypes.ValueString()),
		UUID:          cloudflare.F(data.UUID.ValueString()),
		WasActiveAt:   cloudflare.F(dataWasActiveAt),
		WasInactiveAt: cloudflare.F(dataWasInactiveAt),
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
	data.Items = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
