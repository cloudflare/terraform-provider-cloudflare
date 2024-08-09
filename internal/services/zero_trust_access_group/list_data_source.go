// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ZeroTrustAccessGroupsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &ZeroTrustAccessGroupsDataSource{}

func NewZeroTrustAccessGroupsDataSource() datasource.DataSource {
	return &ZeroTrustAccessGroupsDataSource{}
}

func (d *ZeroTrustAccessGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_access_groups"
}

func (d *ZeroTrustAccessGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ZeroTrustAccessGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ZeroTrustAccessGroupsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := zero_trust.AccessGroupListParams{}
	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	items := &[]*ZeroTrustAccessGroupsResultDataSourceModel{}
	env := ZeroTrustAccessGroupsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*ZeroTrustAccessGroupsResultDataSourceModel{}

	page, err := d.client.ZeroTrust.Access.Groups.List(ctx, params)
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
