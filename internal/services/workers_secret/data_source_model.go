// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
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

type WorkersSecretFindOneByDataSourceModel struct {
	AccountID         types.String `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	ScriptName        types.String `tfsdk:"script_name" path:"script_name"`
}
