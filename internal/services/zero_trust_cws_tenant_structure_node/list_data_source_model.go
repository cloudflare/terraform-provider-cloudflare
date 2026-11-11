// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure_node

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSTenantStructureNodesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSTenantStructureNodesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSTenantStructureNodesDataSourceModel struct {
	AccountID   types.String                                                                        `tfsdk:"account_id" path:"account_id,required"`
	StructureID types.String                                                                        `tfsdk:"structure_id" path:"structure_id,required"`
	Name        types.String                                                                        `tfsdk:"name" query:"name,optional"`
	NativeID    types.String                                                                        `tfsdk:"native_id" query:"native_id,optional"`
	NodeKind    types.String                                                                        `tfsdk:"node_kind" query:"node_kind,optional"`
	Parent      types.String                                                                        `tfsdk:"parent" query:"parent,optional"`
	Direction   types.String                                                                        `tfsdk:"direction" query:"direction,computed_optional"`
	Order       types.String                                                                        `tfsdk:"order" query:"order,computed_optional"`
	MaxItems    types.Int64                                                                         `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[ZeroTrustCWSTenantStructureNodesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSTenantStructureNodesDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSTenantStructureNodeListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSTenantStructureNodeListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSTenantStructureNodeListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.NativeID.IsNull() {
		params.NativeID = cloudflare.F(m.NativeID.ValueString())
	}
	if !m.NodeKind.IsNull() {
		params.NodeKind = cloudflare.F(m.NodeKind.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSTenantStructureNodeListParamsOrder(m.Order.ValueString()))
	}
	if !m.Parent.IsNull() {
		params.Parent = cloudflare.F(m.Parent.ValueString())
	}

	return
}

type ZeroTrustCWSTenantStructureNodesResultDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	NativeID  types.String      `tfsdk:"native_id" json:"native_id,computed"`
	NodeKind  types.String      `tfsdk:"node_kind" json:"node_kind,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	Parent    types.String      `tfsdk:"parent" json:"parent,computed"`
}
