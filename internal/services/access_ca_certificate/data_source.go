// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_ca_certificate

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type AccessCACertificateDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &AccessCACertificateDataSource{}

func NewAccessCACertificateDataSource() datasource.DataSource {
	return &AccessCACertificateDataSource{}
}

func (d *AccessCACertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_ca_certificate"
}

func (r *AccessCACertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AccessCACertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccessCACertificateDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy == nil {
		res := new(http.Response)
		env := AccessCACertificateResultDataSourceEnvelope{*data}
		params := zero_trust.AccessApplicationCAGetParams{}

		if !data.AccountID.IsNull() {
			params.AccountID = cloudflare.F(data.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
		}

		_, err := r.client.ZeroTrust.Access.Applications.CAs.Get(
			ctx,
			data.AppID.ValueString(),
			params,
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
		params := zero_trust.AccessApplicationCAListParams{}
		if !data.AccountID.IsNull() {
			params.AccountID = cloudflare.F(data.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
		}

		items := &[]*AccessCACertificateDataSourceModel{}
		env := AccessCACertificateResultListDataSourceEnvelope{items}

		page, err := r.client.ZeroTrust.Access.Applications.CAs.List(ctx, params)
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
