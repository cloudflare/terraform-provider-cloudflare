// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/cloudflare-go/v4/workers"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkersDeploymentResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkersDeploymentResource)(nil)
var _ resource.ResourceWithImportState = (*WorkersDeploymentResource)(nil)

func NewResource() resource.Resource {
  return &WorkersDeploymentResource{}
}

// WorkersDeploymentResource defines the resource implementation.
type WorkersDeploymentResource struct {
  client *cloudflare.Client
}

func (r *WorkersDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_workers_deployment"
}

func (r *WorkersDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WorkersDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *WorkersDeploymentModel

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
  env := WorkersDeploymentResultEnvelope{*data}
  _, err = r.client.Workers.Scripts.Deployments.New(
    ctx,
    data.ScriptName.ValueString(),
    workers.ScriptDeploymentNewParams{
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

func (r *WorkersDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  // Update is not supported for this resource
}

func (r *WorkersDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *WorkersDeploymentModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  env := WorkersDeploymentResultEnvelope{*data}
  _, err := r.client.Workers.Scripts.Deployments.Get(
    ctx,
    data.ID.ValueString(),
    workers.ScriptDeploymentGetParams{
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

func (r *WorkersDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *WorkersDeploymentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *WorkersDeploymentModel = new(WorkersDeploymentModel)

  path_account_id := ""
  path_script_name := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<account_id>/<script_name>",
    &path_account_id,
    &path_script_name,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.AccountID = types.StringValue(path_account_id)
  data.ScriptName = types.StringValue(path_script_name)

  res := new(http.Response)
  env := WorkersDeploymentResultEnvelope{*data}
  _, err := r.client.Workers.Scripts.Deployments.Get(
    ctx,
    path_script_name,
    workers.ScriptDeploymentGetParams{
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

func (r *WorkersDeploymentResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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
