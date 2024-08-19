// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersSecretResultEnvelope struct {
	Result WorkersSecretModel `json:"result,computed"`
}

type WorkersSecretModel struct {
	ID                types.String `tfsdk:"id" json:"-,computed"`
	ScriptName        types.String `tfsdk:"script_name" path:"script_name"`
	AccountID         types.String `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	Name              types.String `tfsdk:"name" json:"name"`
	Text              types.String `tfsdk:"text" json:"text"`
	Type              types.String `tfsdk:"type" json:"type"`
}
