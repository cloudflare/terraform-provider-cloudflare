// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/kv"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVDataSourceModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	KeyName     types.String `tfsdk:"key_name" path:"key_name,required"`
	NamespaceID types.String `tfsdk:"namespace_id" path:"namespace_id,required"`
}

func (m *WorkersKVDataSourceModel) toReadParams(_ context.Context) (params kv.NamespaceValueGetParams, diags diag.Diagnostics) {
	params = kv.NamespaceValueGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
