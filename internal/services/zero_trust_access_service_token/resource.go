// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustAccessServiceTokenResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustAccessServiceTokenResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustAccessServiceTokenResource)(nil)

func NewResource() resource.Resource {
  return &ZeroTrustAccessServiceTokenResource{}
}

// ZeroTrustAccessServiceTokenResource defines the resource implementation.
type ZeroTrustAccessServiceTokenResource struct {
  client *cloudflare.Client
}

func (r *ZeroTrustAccessServiceTokenResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_zero_trust_access_service_token"
}

func (r *ZeroTrustAccessServiceTokenResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustAccessServiceTokenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *ZeroTrustAccessServiceTokenModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  dataBytes, err := data.MarshalJSON()
  if err != nil {
    resp.Diagnostics.AddError("failed to serialize http request", err.Error())
    return
  }
  res := new(http.Response)
  env := ZeroTrustAccessServiceTokenResultEnvelope{*data}
  params := zero_trust.AccessServiceTokenNewParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err = r.client.ZeroTrust.Access.ServiceTokens.New(
    ctx,
    params,
    option.WithRequestBody("application/json", dataBytes),
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

func (r *ZeroTrustAccessServiceTokenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *ZeroTrustAccessServiceTokenModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *ZeroTrustAccessServiceTokenModel

  resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

  if resp.Diagnostics.HasError() {
    return
  }

  dataBytes, err := data.MarshalJSONForUpdate(*state)
  if err != nil {
    resp.Diagnostics.AddError("failed to serialize http request", err.Error())
    return
  }
  res := new(http.Response)
  env := ZeroTrustAccessServiceTokenResultEnvelope{*data}
  params := zero_trust.AccessServiceTokenUpdateParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err = r.client.ZeroTrust.Access.ServiceTokens.Update(
    ctx,
    data.ID.ValueString(),
    params,
    option.WithRequestBody("application/json", dataBytes),
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

func (r *ZeroTrustAccessServiceTokenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *ZeroTrustAccessServiceTokenModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := ZeroTrustAccessServiceTokenResultEnvelope{*data}
  params := zero_trust.AccessServiceTokenGetParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err := r.client.ZeroTrust.Access.ServiceTokens.Get(
    ctx,
    data.ID.ValueString(),
    params,
    option.WithResponseBodyInto(&res),
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if res != nil && res.StatusCode == 404 {
  resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
    resp.State.RemoveResource(ctx)
    return
  }
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustAccessServiceTokenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *ZeroTrustAccessServiceTokenModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params := zero_trust.AccessServiceTokenDeleteParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err := r.client.ZeroTrust.Access.ServiceTokens.Delete(
    ctx,
    data.ID.ValueString(),
    params,
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustAccessServiceTokenResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *ZeroTrustAccessServiceTokenModel = new(ZeroTrustAccessServiceTokenModel)
  params := zero_trust.AccessServiceTokenGetParams{

  }

  path_accounts_or_zones, path_account_id_or_zone_id := "", ""
  path_service_token_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<{accounts|zones}/{account_id|zone_id}>/<service_token_id>",
    &path_accounts_or_zones,
    &path_account_id_or_zone_id,
    &path_service_token_id,
  )
  resp.Diagnostics.Append(diags...)
  switch path_accounts_or_zones {
  case "accounts":
    params.AccountID = cloudflare.F(path_account_id_or_zone_id)
    data.AccountID = types.StringValue(path_account_id_or_zone_id)
  case "zones":
    params.ZoneID = cloudflare.F(path_account_id_or_zone_id)
    data.ZoneID = types.StringValue(path_account_id_or_zone_id)
  default:
    resp.Diagnostics.AddError("invalid discriminator segment - <{accounts|zones}/{account_id|zone_id}>", "expected discriminator to be one of {accounts|zones}")
  }
  if resp.Diagnostics.HasError() {
    return
  }

  data.ID = types.StringValue(path_service_token_id)

  res := new(http.Response)
  env := ZeroTrustAccessServiceTokenResultEnvelope{*data}
  _, err := r.client.ZeroTrust.Access.ServiceTokens.Get(
    ctx,
    path_service_token_id,
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustAccessServiceTokenResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
