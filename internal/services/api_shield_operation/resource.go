// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/api_gateway"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*APIShieldOperationResource)(nil)
var _ resource.ResourceWithModifyPlan = (*APIShieldOperationResource)(nil)
var _ resource.ResourceWithImportState = (*APIShieldOperationResource)(nil)

func NewResource() resource.Resource {
  return &APIShieldOperationResource{}
}

// APIShieldOperationResource defines the resource implementation.
type APIShieldOperationResource struct {
  client *cloudflare.Client
}

func (r *APIShieldOperationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_api_shield_operation"
}

func (r *APIShieldOperationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *APIShieldOperationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *APIShieldOperationModel

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
  env := APIShieldOperationResultEnvelope{*data}
  _, err = r.client.APIGateway.Operations.New(
    ctx,
    api_gateway.OperationNewParams{
      ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
  data.ID = data.OperationID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIShieldOperationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *APIShieldOperationModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *APIShieldOperationModel

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
  env := APIShieldOperationResultEnvelope{*data}
  _, err = r.client.APIGateway.Operations.New(
    ctx,
    api_gateway.OperationNewParams{
      ZoneID: cloudflare.F(data.OperationID.ValueString()),
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
  data.ID = data.OperationID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIShieldOperationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *APIShieldOperationModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := APIShieldOperationResultEnvelope{*data}
  _, err := r.client.APIGateway.Operations.Get(
    ctx,
    data.OperationID.ValueString(),
    api_gateway.OperationGetParams{
      ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
  data.ID = data.OperationID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIShieldOperationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *APIShieldOperationModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  _, err := r.client.APIGateway.Operations.Delete(
    ctx,
    data.OperationID.ValueString(),
    api_gateway.OperationDeleteParams{
      ZoneID: cloudflare.F(data.ZoneID.ValueString()),
    },
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }
  data.ID = data.OperationID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIShieldOperationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *APIShieldOperationModel = new(APIShieldOperationModel)

  path_zone_id := ""
  path_operation_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<zone_id>/<operation_id>",
    &path_zone_id,
    &path_operation_id,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.ZoneID = types.StringValue(path_zone_id)
  data.OperationID = types.StringValue(path_operation_id)

  res := new(http.Response)
  env := APIShieldOperationResultEnvelope{*data}
  _, err := r.client.APIGateway.Operations.Get(
    ctx,
    path_operation_id,
    api_gateway.OperationGetParams{
      ZoneID: cloudflare.F(path_zone_id),
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
  data.ID = data.OperationID

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIShieldOperationResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
