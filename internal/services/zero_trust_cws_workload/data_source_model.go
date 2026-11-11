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

type ZeroTrustCWSWorkloadResultDataSourceEnvelope struct {
	Result ZeroTrustCWSWorkloadDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSWorkloadDataSourceModel struct {
	ID          types.String                                  `tfsdk:"id" path:"workload_id,computed"`
	WorkloadID  types.String                                  `tfsdk:"workload_id" path:"workload_id,optional"`
	AccountID   types.String                                  `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt   timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String                                  `tfsdk:"description" json:"description,computed"`
	Filter      types.String                                  `tfsdk:"filter" json:"filter,computed"`
	Name        types.String                                  `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339                             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags        customfield.Map[types.String]                 `tfsdk:"tags" json:"tags,computed"`
	FindOneBy   *ZeroTrustCWSWorkloadFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *ZeroTrustCWSWorkloadDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSWorkloadGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSWorkloadGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSWorkloadDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSWorkloadListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSWorkloadListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.FindOneBy.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSWorkloadListParamsDirection(m.FindOneBy.Direction.ValueString()))
	}
	if !m.FindOneBy.Name.IsNull() {
		params.Name = cloudflare.F(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSWorkloadListParamsOrder(m.FindOneBy.Order.ValueString()))
	}
	if !m.FindOneBy.PolicySetID.IsNull() {
		params.PolicySetID = cloudflare.F(m.FindOneBy.PolicySetID.ValueString())
	}

	return
}

type ZeroTrustCWSWorkloadFindOneByDataSourceModel struct {
	Direction   types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name        types.String `tfsdk:"name" query:"name,optional"`
	Order       types.String `tfsdk:"order" query:"order,computed_optional"`
	PolicySetID types.String `tfsdk:"policy_set_id" query:"policy_set_id,optional"`
}
