// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

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

type AccessMutualTLSCertificateDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &AccessMutualTLSCertificateDataSource{}

func NewAccessMutualTLSCertificateDataSource() datasource.DataSource {
	return &AccessMutualTLSCertificateDataSource{}
}

func (d *AccessMutualTLSCertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_mutual_tls_certificate"
}

func (r *AccessMutualTLSCertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AccessMutualTLSCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccessMutualTLSCertificateDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy == nil {
		res := new(http.Response)
		env := AccessMutualTLSCertificateResultDataSourceEnvelope{*data}
		params := zero_trust.AccessCertificateGetParams{}

		if !data.AccountID.IsNull() {
			params.AccountID = cloudflare.F(data.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
		}

		_, err := r.client.ZeroTrust.Access.Certificates.Get(
			ctx,
			data.CertificateID.ValueString(),
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
		params := zero_trust.AccessCertificateListParams{}
		if !data.AccountID.IsNull() {
			params.AccountID = cloudflare.F(data.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
		}

		items := &[]*AccessMutualTLSCertificateDataSourceModel{}
		env := AccessMutualTLSCertificateResultListDataSourceEnvelope{items}

		page, err := r.client.ZeroTrust.Access.Certificates.List(ctx, params)
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
