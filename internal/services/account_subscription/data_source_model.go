// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountSubscriptionResultDataSourceEnvelope struct {
	Result AccountSubscriptionDataSourceModel `json:"result,computed"`
}

type AccountSubscriptionDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m *AccountSubscriptionDataSourceModel) toReadParams(_ context.Context) (params accounts.SubscriptionGetParams, diags diag.Diagnostics) {
	params = accounts.SubscriptionGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
