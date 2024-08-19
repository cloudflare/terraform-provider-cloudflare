// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

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

type ZeroTrustLocalDomainFallbackDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &ZeroTrustLocalDomainFallbackDataSource{}

func NewZeroTrustLocalDomainFallbackDataSource() datasource.DataSource {
	return &ZeroTrustLocalDomainFallbackDataSource{}
}

func (d *ZeroTrustLocalDomainFallbackDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_local_domain_fallback"
}

func (d *ZeroTrustLocalDomainFallbackDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ZeroTrustLocalDomainFallbackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ZeroTrustLocalDomainFallbackDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		res := new(http.Response)
		env := ZeroTrustLocalDomainFallbackResultDataSourceEnvelope{*data}
		_, err := d.client.ZeroTrust.Devices.Policies.FallbackDomains.Get(
			ctx,
			data.PolicyID.ValueString(),
			zero_trust.DevicePolicyFallbackDomainGetParams{
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
		items := &[]*ZeroTrustLocalDomainFallbackDataSourceModel{}
		env := ZeroTrustLocalDomainFallbackResultListDataSourceEnvelope{items}

		page, err := d.client.ZeroTrust.Devices.Policies.FallbackDomains.List(ctx, zero_trust.DevicePolicyFallbackDomainListParams{
			AccountID: cloudflare.F(data.Filter.AccountID.ValueString()),
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
