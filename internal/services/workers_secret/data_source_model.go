// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersSecretResultDataSourceEnvelope struct {
	Result WorkersSecretDataSourceModel `json:"result,computed"`
}

type WorkersSecretResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersSecretDataSourceModel] `json:"result,computed"`
}

type WorkersSecretDataSourceModel struct {
	ID                types.String `tfsdk:"id" json:"-,computed"`
	SecretName        types.String `tfsdk:"secret_name" path:"secret_name,optional"`
	AccountID         types.String `tfsdk:"account_id" path:"account_id,required"`
	DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace,required"`
	ScriptName        types.String `tfsdk:"script_name" path:"script_name,required"`
	Name              types.String `tfsdk:"name" json:"name,computed"`
	Type              types.String `tfsdk:"type" json:"type,computed"`
}

func (m *WorkersSecretDataSourceModel) toReadParams(_ context.Context) (params workers_for_platforms.DispatchNamespaceScriptSecretGetParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceScriptSecretGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkersSecretDataSourceModel) toListParams(_ context.Context) (params workers_for_platforms.DispatchNamespaceScriptSecretListParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceScriptSecretListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
