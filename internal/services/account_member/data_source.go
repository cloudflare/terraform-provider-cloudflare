// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type AccountMemberDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*AccountMemberDataSource)(nil)

func NewAccountMemberDataSource() datasource.DataSource {
	return &AccountMemberDataSource{}
}

func (d *AccountMemberDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_member"
}

func (d *AccountMemberDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *AccountMemberDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccountMemberDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		params, diags := data.toReadParams()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		res := new(http.Response)
		env := AccountMemberResultDataSourceEnvelope{*data}
		_, err := d.client.Accounts.Members.Get(
			ctx,
			data.MemberID.ValueString(),
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
	} else {
		params, diags := data.toListParams()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		items := customfield.NullObjectList[AccountMemberDataSourceModel](ctx)
		env := AccountMemberResultListDataSourceEnvelope{items}

		page, err := d.client.Accounts.Members.List(ctx, params)
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
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
