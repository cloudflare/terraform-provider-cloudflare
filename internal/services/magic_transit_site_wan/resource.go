// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

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
var _ resource.ResourceWithConfigure = (*MagicTransitSiteWANResource)(nil)
var _ resource.ResourceWithModifyPlan = (*MagicTransitSiteWANResource)(nil)
var _ resource.ResourceWithImportState = (*MagicTransitSiteWANResource)(nil)

func NewResource() resource.Resource {
  return &MagicTransitSiteWANResource{}
}

// MagicTransitSiteWANResource defines the resource implementation.
type MagicTransitSiteWANResource struct {
  client *cloudflare.Client
}

func (r *MagicTransitSiteWANResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_magic_transit_site_wan"
}

func (r *MagicTransitSiteWANResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MagicTransitSiteWANResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *MagicTransitSiteWANModel

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
  env := MagicTransitSiteWANResultEnvelope{*data}
  _, err = r.client.MagicTransit.Sites.WANs.New(
    ctx,
    data.SiteID.ValueString(),
    magic_transit.SiteWANNewParams{
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

func (r *MagicTransitSiteWANResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *MagicTransitSiteWANModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *MagicTransitSiteWANModel

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
  env := MagicTransitSiteWANResultEnvelope{*data}
  _, err = r.client.MagicTransit.Sites.WANs.Update(
    ctx,
    data.SiteID.ValueString(),
    data.ID.ValueString(),
    magic_transit.SiteWANUpdateParams{
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

func (r *MagicTransitSiteWANResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *MagicTransitSiteWANModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := MagicTransitSiteWANResultEnvelope{*data}
  _, err := r.client.MagicTransit.Sites.WANs.Get(
    ctx,
    data.SiteID.ValueString(),
    data.ID.ValueString(),
    magic_transit.SiteWANGetParams{
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

func (r *MagicTransitSiteWANResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *MagicTransitSiteWANModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  _, err := r.client.MagicTransit.Sites.WANs.Delete(
    ctx,
    data.SiteID.ValueString(),
    data.ID.ValueString(),
    magic_transit.SiteWANDeleteParams{
      AccountID: cloudflare.F(data.AccountID.ValueString()),
    },
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MagicTransitSiteWANResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *MagicTransitSiteWANModel = new(MagicTransitSiteWANModel)

  path_account_id := ""
  path_site_id := ""
  path_wan_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<account_id>/<site_id>/<wan_id>",
    &path_account_id,
    &path_site_id,
    &path_wan_id,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.AccountID = types.StringValue(path_account_id)
  data.SiteID = types.StringValue(path_site_id)
  data.ID = types.StringValue(path_wan_id)

  res := new(http.Response)
  env := MagicTransitSiteWANResultEnvelope{*data}
  _, err := r.client.MagicTransit.Sites.WANs.Get(
    ctx,
    path_site_id,
    path_wan_id,
    magic_transit.SiteWANGetParams{
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

func (r *MagicTransitSiteWANResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
