// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_dns_destination_ips

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayDNSDestinationIPsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayDNSDestinationIPsListResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayDNSDestinationIPsListDataSourceModel struct {
	AccountID types.String                                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustGatewayDNSDestinationIPsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayDNSDestinationIPsListDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayDNSDestinationIPListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayDNSDestinationIPListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayDNSDestinationIPsListResultDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	BackupIP  types.String `tfsdk:"backup_ip" json:"backup_ip,computed"`
	PairType  types.String `tfsdk:"pair_type" json:"pair_type,computed"`
	PrimaryIP types.String `tfsdk:"primary_ip" json:"primary_ip,computed"`
}
