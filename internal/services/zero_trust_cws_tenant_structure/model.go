// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSTenantStructureResultEnvelope struct {
	Result ZeroTrustCWSTenantStructureModel `json:"result"`
}

type ZeroTrustCWSTenantStructureModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	Application types.String      `tfsdk:"application" json:"application,required"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	ImportPath  types.String      `tfsdk:"import_path" json:"import_path,optional"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSTenantStructureModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSTenantStructureModel) MarshalJSONForUpdate(state ZeroTrustCWSTenantStructureModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
