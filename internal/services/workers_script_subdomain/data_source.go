// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_subdomain

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
)

type WorkersScriptSubdomainDataSource struct {
  client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*WorkersScriptSubdomainDataSource)(nil)

func NewWorkersScriptSubdomainDataSource() datasource.DataSource {
  return &WorkersScriptSubdomainDataSource{}
}

func (d *WorkersScriptSubdomainDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_workers_script_subdomain"
}

func (d *WorkersScriptSubdomainDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *WorkersScriptSubdomainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  var data *WorkersScriptSubdomainDataSourceModel

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
  _, err := d.client.Workers.Scripts.Subdomain.Get(
    ctx,
    data.ScriptName.ValueString(),
    params,
    option.WithResponseBodyInto(&res),
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }
  bytes, _ := io.ReadAll(res.Body)
  err = apijson.UnmarshalComputed(bytes, &data)
  if err != nil {
    resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
