// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_subdomain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/workers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptSubdomainDataSourceModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id,required"`
	ScriptName      types.String `tfsdk:"script_name" path:"script_name,required"`
	Enabled         types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	PreviewsEnabled types.Bool   `tfsdk:"previews_enabled" json:"previews_enabled,computed"`
}

func (m *WorkersScriptSubdomainDataSourceModel) toReadParams(_ context.Context) (params workers.ScriptSubdomainGetParams, diags diag.Diagnostics) {
	params = workers.ScriptSubdomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
