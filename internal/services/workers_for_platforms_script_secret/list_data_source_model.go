// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsScriptSecretsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[WorkersForPlatformsScriptSecretsResultDataSourceModel] `json:"result,computed"`
}

type WorkersForPlatformsScriptSecretsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace,required"`
ScriptName types.String `tfsdk:"script_name" path:"script_name,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[WorkersForPlatformsScriptSecretsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersForPlatformsScriptSecretsDataSourceModel) toListParams(_ context.Context) (params workers_for_platforms.DispatchNamespaceScriptSecretListParams, diags diag.Diagnostics) {
  params = workers_for_platforms.DispatchNamespaceScriptSecretListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type WorkersForPlatformsScriptSecretsResultDataSourceModel struct {
Name types.String `tfsdk:"name" json:"name,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}
