// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure_node

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSTenantStructureNodeResultDataSourceEnvelope struct {
	Result ZeroTrustCWSTenantStructureNodeDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSTenantStructureNodeDataSourceModel struct {
	ID          types.String                                             `tfsdk:"id" path:"node_id,computed"`
	NodeID      types.String                                             `tfsdk:"node_id" path:"node_id,optional"`
	AccountID   types.String                                             `tfsdk:"account_id" path:"account_id,required"`
	StructureID types.String                                             `tfsdk:"structure_id" path:"structure_id,required"`
	CreatedAt   timetypes.RFC3339                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String                                             `tfsdk:"name" json:"name,computed"`
	NativeID    types.String                                             `tfsdk:"native_id" json:"native_id,computed"`
	NodeKind    types.String                                             `tfsdk:"node_kind" json:"node_kind,computed"`
	Parent      types.String                                             `tfsdk:"parent" json:"parent,computed"`
	UpdatedAt   timetypes.RFC3339                                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Filter      *ZeroTrustCWSTenantStructureNodeFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustCWSTenantStructureNodeDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSTenantStructureNodeGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSTenantStructureNodeGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSTenantStructureNodeDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSTenantStructureNodeListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSTenantStructureNodeListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSTenantStructureNodeListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.NativeID.IsNull() {
		params.NativeID = cloudflare.F(m.Filter.NativeID.ValueString())
	}
	if !m.Filter.NodeKind.IsNull() {
		params.NodeKind = cloudflare.F(m.Filter.NodeKind.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSTenantStructureNodeListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Parent.IsNull() {
		params.Parent = cloudflare.F(m.Filter.Parent.ValueString())
	}

	return
}

type ZeroTrustCWSTenantStructureNodeFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
	NativeID  types.String `tfsdk:"native_id" query:"native_id,optional"`
	NodeKind  types.String `tfsdk:"node_kind" query:"node_kind,optional"`
	Order     types.String `tfsdk:"order" query:"order,computed_optional"`
	Parent    types.String `tfsdk:"parent" query:"parent,optional"`
}
