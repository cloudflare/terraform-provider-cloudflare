// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_secret

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerSecretResultEnvelope struct {
	Result WorkerSecretModel `json:"result,computed"`
}

type WorkerSecretModel struct {
	ID                types.String `tfsdk:"id" json:"-,computed"`
	AccountID         types.String `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	ScriptName        types.String `tfsdk:"script_name" path:"script_name"`
	Name              types.String `tfsdk:"name" json:"name"`
	Text              types.String `tfsdk:"text" json:"text"`
	Type              types.String `tfsdk:"type" json:"type"`
}
