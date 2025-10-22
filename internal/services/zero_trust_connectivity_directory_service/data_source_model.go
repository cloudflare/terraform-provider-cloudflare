// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustConnectivityDirectoryServiceDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ServiceID types.String `tfsdk:"service_id" path:"service_id,required"`
}

func (m *ZeroTrustConnectivityDirectoryServiceDataSourceModel) toReadParams(_ context.Context) (params zero_trust.ConnectivityDirectoryServiceGetParams, diags diag.Diagnostics) {
	params = zero_trust.ConnectivityDirectoryServiceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
