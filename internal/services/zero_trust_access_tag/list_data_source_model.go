// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessTagsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessTagsDataSourceModel struct {
	AccountID types.String                                                           `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessTagsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessTagsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessTagListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessTagListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustAccessTagsResultDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}
