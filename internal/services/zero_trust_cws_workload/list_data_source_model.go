// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_workload

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSWorkloadsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSWorkloadsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSWorkloadsDataSourceModel struct {
	AccountID   types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String                                                             `tfsdk:"name" query:"name,optional"`
	PolicySetID types.String                                                             `tfsdk:"policy_set_id" query:"policy_set_id,optional"`
	Direction   types.String                                                             `tfsdk:"direction" query:"direction,computed_optional"`
	Order       types.String                                                             `tfsdk:"order" query:"order,computed_optional"`
	MaxItems    types.Int64                                                              `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[ZeroTrustCWSWorkloadsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSWorkloadsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSWorkloadListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSWorkloadListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSWorkloadListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSWorkloadListParamsOrder(m.Order.ValueString()))
	}
	if !m.PolicySetID.IsNull() {
		params.PolicySetID = cloudflare.F(m.PolicySetID.ValueString())
	}

	return
}

type ZeroTrustCWSWorkloadsResultDataSourceModel struct {
	ID          types.String                  `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Filter      types.String                  `tfsdk:"filter" json:"filter,computed"`
	Name        types.String                  `tfsdk:"name" json:"name,computed"`
	Tags        customfield.Map[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt   timetypes.RFC3339             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String                  `tfsdk:"description" json:"description,computed"`
}
