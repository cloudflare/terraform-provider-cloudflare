// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/custom_pages"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*CustomPagesResource)(nil)
var _ resource.ResourceWithModifyPlan = (*CustomPagesResource)(nil)

func NewResource() resource.Resource {
  return &CustomPagesResource{}
}

// CustomPagesResource defines the resource implementation.
type CustomPagesResource struct {
  client *cloudflare.Client
}

func (r *CustomPagesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_custom_pages"
}

func (r *CustomPagesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomPagesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *CustomPagesModel

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
  env := CustomPagesResultEnvelope{*data}
  params := custom_pages.CustomPageUpdateParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err = r.client.CustomPages.Update(
    ctx,
    data.Identifier.ValueString(),
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

func (r *CustomPagesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *CustomPagesModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *CustomPagesModel

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
  env := CustomPagesResultEnvelope{*data}
  params := custom_pages.CustomPageUpdateParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err = r.client.CustomPages.Update(
    ctx,
    data.Identifier.ValueString(),
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

func (r *CustomPagesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *CustomPagesModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := CustomPagesResultEnvelope{*data}
  params := custom_pages.CustomPageGetParams{

  }

  if !data.AccountID.IsNull() {
    params.AccountID = cloudflare.F(data.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
  }

  _, err := r.client.CustomPages.Get(
    ctx,
    data.Identifier.ValueString(),
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

func (r *CustomPagesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *CustomPagesResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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
