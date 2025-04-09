// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

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

type CustomHostnameDataSource struct {
  client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*CustomHostnameDataSource)(nil)

func NewCustomHostnameDataSource() datasource.DataSource {
  return &CustomHostnameDataSource{}
}

func (d *CustomHostnameDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_custom_hostname"
}

func (d *CustomHostnameDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CustomHostnameDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  var data *CustomHostnameDataSourceModel

  resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  if data.Filter != nil {
    params, diags := data.toListParams(ctx)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
      return
    }

    env := CustomHostnamesResultListDataSourceEnvelope{}
    page, err := d.client.CustomHostnames.List(ctx, params)
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

    if count := len(env.Result.Elements()); count != 1 {
      resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count) + " found")
      return
    }
    ts, diags := env.Result.AsStructSliceT(ctx)
    resp.Diagnostics.Append(diags...)
    data.CustomHostnameID = ts[0].ID
  }

  params, diags := data.toReadParams(ctx)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := CustomHostnameResultDataSourceEnvelope{*data}
  _, err := d.client.CustomHostnames.Get(
    ctx,
    data.CustomHostnameID.ValueString(),
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
