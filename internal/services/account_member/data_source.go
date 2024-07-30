// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type AccountMemberDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &AccountMemberDataSource{}

func NewAccountMemberDataSource() datasource.DataSource {
	return &AccountMemberDataSource{}
}

func (d *AccountMemberDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_member"
}

func (r *AccountMemberDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AccountMemberDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccountMemberDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		res := new(http.Response)
		env := AccountMemberResultDataSourceEnvelope{*data}
		_, err := r.client.Accounts.Members.Get(
			ctx,
			data.MemberID.ValueString(),
			accounts.MemberGetParams{
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
		items := &[]*AccountMemberDataSourceModel{}
		env := AccountMemberResultListDataSourceEnvelope{items}

		page, err := r.client.Accounts.Members.List(ctx, accounts.MemberListParams{
			AccountID: cloudflare.F(data.Filter.AccountID.ValueString()),
			Direction: cloudflare.F(accounts.MemberListParamsDirection(data.Filter.Direction.ValueString())),
			Order:     cloudflare.F(accounts.MemberListParamsOrder(data.Filter.Order.ValueString())),
			Page:      cloudflare.F(data.Filter.Page.ValueFloat64()),
			PerPage:   cloudflare.F(data.Filter.PerPage.ValueFloat64()),
			Status:    cloudflare.F(accounts.MemberListParamsStatus(data.Filter.Status.ValueString())),
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
