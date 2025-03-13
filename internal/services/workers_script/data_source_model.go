// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/workers"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ScriptName types.String `tfsdk:"script_name" path:"script_name,required"`
}

func (m *WorkersScriptDataSourceModel) toReadParams(_ context.Context) (params workers.ScriptGetParams, diags diag.Diagnostics) {
  params = workers.ScriptGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}
