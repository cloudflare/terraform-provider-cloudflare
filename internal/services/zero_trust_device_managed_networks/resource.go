// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

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
var _ resource.ResourceWithConfigure = (*ZeroTrustDeviceManagedNetworksResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustDeviceManagedNetworksResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustDeviceManagedNetworksResource)(nil)

func NewResource() resource.Resource {
  return &ZeroTrustDeviceManagedNetworksResource{}
}

// ZeroTrustDeviceManagedNetworksResource defines the resource implementation.
type ZeroTrustDeviceManagedNetworksResource struct {
  client *cloudflare.Client
}

func (r *ZeroTrustDeviceManagedNetworksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_zero_trust_device_managed_networks"
}

func (r *ZeroTrustDeviceManagedNetworksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustDeviceManagedNetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *ZeroTrustDeviceManagedNetworksModel

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
  env := ZeroTrustDeviceManagedNetworksResultEnvelope{*data}
  _, err = r.client.ZeroTrust.Devices.Networks.New(
    ctx,
    zero_trust.DeviceNetworkNewParams{
      AccountID: cloudflare.F(data.AccountID.ValueString()),
    },
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
  data.ID = data.NetworkID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceManagedNetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *ZeroTrustDeviceManagedNetworksModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *ZeroTrustDeviceManagedNetworksModel

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
  env := ZeroTrustDeviceManagedNetworksResultEnvelope{*data}
  _, err = r.client.ZeroTrust.Devices.Networks.Update(
    ctx,
    data.NetworkID.ValueString(),
    zero_trust.DeviceNetworkUpdateParams{
      AccountID: cloudflare.F(data.AccountID.ValueString()),
    },
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
  data.ID = data.NetworkID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceManagedNetworksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *ZeroTrustDeviceManagedNetworksModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := ZeroTrustDeviceManagedNetworksResultEnvelope{*data}
  _, err := r.client.ZeroTrust.Devices.Networks.Get(
    ctx,
    data.NetworkID.ValueString(),
    zero_trust.DeviceNetworkGetParams{
      AccountID: cloudflare.F(data.AccountID.ValueString()),
    },
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
  data.ID = data.NetworkID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceManagedNetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *ZeroTrustDeviceManagedNetworksModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  _, err := r.client.ZeroTrust.Devices.Networks.Delete(
    ctx,
    data.NetworkID.ValueString(),
    zero_trust.DeviceNetworkDeleteParams{
      AccountID: cloudflare.F(data.AccountID.ValueString()),
    },
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }
  data.ID = data.NetworkID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceManagedNetworksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *ZeroTrustDeviceManagedNetworksModel = new(ZeroTrustDeviceManagedNetworksModel)

  path_account_id := ""
  path_network_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<account_id>/<network_id>",
    &path_account_id,
    &path_network_id,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.AccountID = types.StringValue(path_account_id)
  data.NetworkID = types.StringValue(path_network_id)

  res := new(http.Response)
  env := ZeroTrustDeviceManagedNetworksResultEnvelope{*data}
  _, err := r.client.ZeroTrust.Devices.Networks.Get(
    ctx,
    path_network_id,
    zero_trust.DeviceNetworkGetParams{
      AccountID: cloudflare.F(path_account_id),
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
  data.ID = data.NetworkID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceManagedNetworksResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
