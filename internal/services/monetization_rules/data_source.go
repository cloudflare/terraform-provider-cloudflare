// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

<<<<<<<< HEAD:internal/services/monetization_rules/data_source.go
package monetization_rules
========
package mtls_certificate
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source.go

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

<<<<<<<< HEAD:internal/services/monetization_rules/data_source.go
type MonetizationRulesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*MonetizationRulesDataSource)(nil)

func NewMonetizationRulesDataSource() datasource.DataSource {
	return &MonetizationRulesDataSource{}
}

func (d *MonetizationRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monetization_rules"
}

func (d *MonetizationRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
========
type MTLSCertificateDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*MTLSCertificateDataSource)(nil)

func NewMTLSCertificateDataSource() datasource.DataSource {
	return &MTLSCertificateDataSource{}
}

func (d *MTLSCertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mtls_certificate"
}

func (d *MTLSCertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source.go
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

<<<<<<<< HEAD:internal/services/monetization_rules/data_source.go
func (d *MonetizationRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MonetizationRulesDataSourceModel
========
func (d *MTLSCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MTLSCertificateDataSourceModel
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source.go

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toReadParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
<<<<<<<< HEAD:internal/services/monetization_rules/data_source.go
	env := MonetizationRulesResultDataSourceEnvelope{*data}
	_, err := d.client.Monetization.Rules.Get(
		ctx,
========
	env := MTLSCertificateResultDataSourceEnvelope{*data}
	_, err := d.client.MTLSCertificates.Get(
		ctx,
		data.MTLSCertificateID.ValueString(),
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source.go
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
<<<<<<<< HEAD:internal/services/monetization_rules/data_source.go
	data.ID = data.ZoneID
========
	data.ID = data.MTLSCertificateID
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source.go

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
