// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptDataSourceModel struct {
	ID         types.String                           `tfsdk:"id" path:"script_name,computed"`
	ScriptName types.String                           `tfsdk:"script_name" path:"script_name,optional"`
	AccountID  types.String                           `tfsdk:"account_id" path:"account_id,required"`
	Script     types.String                           `tfsdk:"script" json:"script,computed"`
	Filter     *WorkersScriptFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WorkersScriptDataSourceModel) toReadParams(_ context.Context) (params workers.ScriptGetParams, diags diag.Diagnostics) {
	params = workers.ScriptGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkersScriptDataSourceModel) toListParams(_ context.Context) (params workers.ScriptListParams, diags diag.Diagnostics) {
	params = workers.ScriptListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Tags.IsNull() {
		params.Tags = cloudflare.F(m.Filter.Tags.ValueString())
	}

	return
}

type WorkersScriptFindOneByDataSourceModel struct {
	Tags types.String `tfsdk:"tags" query:"tags,optional"`
}
