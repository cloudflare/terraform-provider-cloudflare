// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language

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

type StreamCaptionLanguageDataSource struct {
  client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*StreamCaptionLanguageDataSource)(nil)

func NewStreamCaptionLanguageDataSource() datasource.DataSource {
  return &StreamCaptionLanguageDataSource{}
}

func (d *StreamCaptionLanguageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_stream_caption_language"
}

func (d *StreamCaptionLanguageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *StreamCaptionLanguageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  var data *StreamCaptionLanguageDataSourceModel

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
  env := StreamCaptionLanguageResultDataSourceEnvelope{*data}
  _, err := d.client.Stream.Captions.Language.Get(
    ctx,
    data.Identifier.ValueString(),
    data.Language.ValueString(),
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
