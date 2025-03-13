// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/magic_transit"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*MagicTransitConnectorResource)(nil)
var _ resource.ResourceWithModifyPlan = (*MagicTransitConnectorResource)(nil)
var _ resource.ResourceWithImportState = (*MagicTransitConnectorResource)(nil)

func NewResource() resource.Resource {
  return &MagicTransitConnectorResource{}
}

// MagicTransitConnectorResource defines the resource implementation.
type MagicTransitConnectorResource struct {
  client *cloudflare.Client
}

func (r *MagicTransitConnectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_magic_transit_connector"
}

func (r *MagicTransitConnectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MagicTransitConnectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *MagicTransitConnectorModel

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
  env := MagicTransitConnectorResultEnvelope{*data}
  _, err = r.client.MagicTransit.Connectors.Update(
    ctx,
    data.ConnectorID.ValueString(),
    magic_transit.ConnectorUpdateParams{
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MagicTransitConnectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *MagicTransitConnectorModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *MagicTransitConnectorModel

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
  env := MagicTransitConnectorResultEnvelope{*data}
  _, err = r.client.MagicTransit.Connectors.Update(
    ctx,
    data.ID.ValueString(),
    magic_transit.ConnectorUpdateParams{
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MagicTransitConnectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *MagicTransitConnectorModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := MagicTransitConnectorResultEnvelope{*data}
  _, err := r.client.MagicTransit.Connectors.Get(
    ctx,
    data.ID.ValueString(),
    magic_transit.ConnectorGetParams{
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MagicTransitConnectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *MagicTransitConnectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *MagicTransitConnectorModel = new(MagicTransitConnectorModel)

  path_account_id := ""
  path_connector_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<account_id>/<connector_id>",
    &path_account_id,
    &path_connector_id,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.AccountID = types.StringValue(path_account_id)
  data.ConnectorID = types.StringValue(path_connector_id)

  res := new(http.Response)
  env := MagicTransitConnectorResultEnvelope{*data}
  _, err := r.client.MagicTransit.Connectors.Get(
    ctx,
    path_connector_id,
    magic_transit.ConnectorGetParams{
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

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MagicTransitConnectorResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
  if req.State.Raw.IsNull() {
      resp.Diagnostics.AddWarning(
          "Resource Destruction Considerations",
          "This resource cannot be destroyed from Terraform. If you create this resource, it will be "+
          "present in the API until manually deleted.",
      )
  }
  if req.Plan.Raw.IsNull() {
      resp.Diagnostics.AddWarning(
          "Resource Destruction Considerations",
          "Applying this resource destruction will remove the resource from the Terraform state "+
              "but will not change it in the API. If you would like to destroy or reset this resource "+
              "in the API, refer to the documentation for how to do it manually.",
      )
  }
}
