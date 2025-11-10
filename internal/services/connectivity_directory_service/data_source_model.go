// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/connectivity"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConnectivityDirectoryServiceDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ServiceID types.String `tfsdk:"service_id" path:"service_id,required"`
}

func (m *ConnectivityDirectoryServiceDataSourceModel) toReadParams(_ context.Context) (params connectivity.DirectoryServiceGetParams, diags diag.Diagnostics) {
	params = connectivity.DirectoryServiceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
