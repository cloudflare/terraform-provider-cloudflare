// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersSecretResultListDataSourceEnvelope struct {
	Result *[]*WorkersSecretDataSourceModel `json:"result,computed"`
}

type WorkersSecretDataSourceModel struct {
	Name   types.String                           `tfsdk:"name" json:"name"`
	Type   types.String                           `tfsdk:"type" json:"type"`
	Filter *WorkersSecretFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WorkersSecretDataSourceModel) toListParams() (params workers_for_platforms.DispatchNamespaceScriptSecretListParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceScriptSecretListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type WorkersSecretFindOneByDataSourceModel struct {
	AccountID         types.String `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	ScriptName        types.String `tfsdk:"script_name" path:"script_name"`
}
