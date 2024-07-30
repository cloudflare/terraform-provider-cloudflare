// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type AccountMembersDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &AccountMembersDataSource{}

func NewAccountMembersDataSource() datasource.DataSource {
	return &AccountMembersDataSource{}
}

func (d *AccountMembersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_members"
}

func (r *AccountMembersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AccountMembersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccountMembersDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*AccountMembersResultDataSourceModel{}
	env := AccountMembersResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*AccountMembersResultDataSourceModel{}

	page, err := r.client.Accounts.Members.List(ctx, accounts.MemberListParams{
		AccountID: cloudflare.F(data.AccountID.ValueString()),
		Direction: cloudflare.F(accounts.MemberListParamsDirection(data.Direction.ValueString())),
		Order:     cloudflare.F(accounts.MemberListParamsOrder(data.Order.ValueString())),
		Page:      cloudflare.F(data.Page.ValueFloat64()),
		PerPage:   cloudflare.F(data.PerPage.ValueFloat64()),
		Status:    cloudflare.F(accounts.MemberListParamsStatus(data.Status.ValueString())),
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
