// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldPolicyResultDataSourceEnvelope struct {
	Result PageShieldPolicyDataSourceModel `json:"result,computed"`
}

type PageShieldPolicyResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PageShieldPolicyDataSourceModel] `json:"result,computed"`
}

type PageShieldPolicyDataSourceModel struct {
	PolicyID    types.String                              `tfsdk:"policy_id" path:"policy_id,optional"`
	ZoneID      types.String                              `tfsdk:"zone_id" path:"zone_id,optional"`
	Action      types.String                              `tfsdk:"action" json:"action,computed"`
	Description types.String                              `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                                `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String                              `tfsdk:"expression" json:"expression,computed"`
	ID          types.String                              `tfsdk:"id" json:"id,computed"`
	Value       types.String                              `tfsdk:"value" json:"value,computed"`
	Filter      *PageShieldPolicyFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *PageShieldPolicyDataSourceModel) toReadParams(_ context.Context) (params page_shield.PolicyGetParams, diags diag.Diagnostics) {
	params = page_shield.PolicyGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *PageShieldPolicyDataSourceModel) toListParams(_ context.Context) (params page_shield.PolicyListParams, diags diag.Diagnostics) {
	params = page_shield.PolicyListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type PageShieldPolicyFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}
