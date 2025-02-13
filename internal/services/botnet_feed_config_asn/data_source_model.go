// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package botnet_feed_config_asn

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/botnet_feed"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotnetFeedConfigASNResultDataSourceEnvelope struct {
	Result BotnetFeedConfigASNDataSourceModel `json:"result,computed"`
}

type BotnetFeedConfigASNDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ASN       types.Int64  `tfsdk:"asn" json:"asn,computed"`
}

func (m *BotnetFeedConfigASNDataSourceModel) toReadParams(_ context.Context) (params botnet_feed.ConfigASNGetParams, diags diag.Diagnostics) {
	params = botnet_feed.ConfigASNGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
