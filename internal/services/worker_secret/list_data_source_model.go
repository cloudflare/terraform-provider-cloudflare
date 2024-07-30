// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_secret

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerSecretsResultListDataSourceEnvelope struct {
	Result *[]*WorkerSecretsResultDataSourceModel `json:"result,computed"`
}

type WorkerSecretsDataSourceModel struct {
	AccountID         types.String                           `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String                           `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	ScriptName        types.String                           `tfsdk:"script_name" path:"script_name"`
	MaxItems          types.Int64                            `tfsdk:"max_items"`
	Result            *[]*WorkerSecretsResultDataSourceModel `tfsdk:"result"`
}

type WorkerSecretsResultDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}
