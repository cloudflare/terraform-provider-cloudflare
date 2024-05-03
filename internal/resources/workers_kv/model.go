// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVResultEnvelope struct {
	Result WorkersKVModel `json:"result,computed"`
}

type WorkersKVModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	NamespaceID types.String `tfsdk:"namespace_id" path:"namespace_id"`
	KeyName     types.String `tfsdk:"key_name" path:"key_name"`
	Metadata    types.String `tfsdk:"metadata" json:"metadata"`
	Value       types.String `tfsdk:"value" json:"value"`
}
