// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure_node

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSTenantStructureNodeResultEnvelope struct {
	Result ZeroTrustCWSTenantStructureNodeModel `json:"result"`
}

type ZeroTrustCWSTenantStructureNodeModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	StructureID types.String      `tfsdk:"structure_id" path:"structure_id,required"`
	NativeID    types.String      `tfsdk:"native_id" json:"native_id,required"`
	NodeKind    types.String      `tfsdk:"node_kind" json:"node_kind,required"`
	Name        types.String      `tfsdk:"name" json:"name,optional"`
	Parent      types.String      `tfsdk:"parent" json:"parent,optional"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSTenantStructureNodeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSTenantStructureNodeModel) MarshalJSONForUpdate(state ZeroTrustCWSTenantStructureNodeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
