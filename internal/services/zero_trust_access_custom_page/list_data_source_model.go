// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPagesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessCustomPagesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessCustomPagesDataSourceModel struct {
	AccountID types.String                                                                  `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                   `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessCustomPagesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessCustomPagesDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessCustomPageListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCustomPageListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustAccessCustomPagesResultDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
	UID  types.String `tfsdk:"uid" json:"uid,computed"`
}
