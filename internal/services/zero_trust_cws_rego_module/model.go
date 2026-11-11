// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_rego_module

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSRegoModuleResultEnvelope struct {
	Result ZeroTrustCWSRegoModuleModel `json:"result"`
}

type ZeroTrustCWSRegoModuleModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	RegoSource  types.String      `tfsdk:"rego_source" json:"rego_source,required"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	PackagePath types.String      `tfsdk:"package_path" json:"package_path,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSRegoModuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSRegoModuleModel) MarshalJSONForUpdate(state ZeroTrustCWSRegoModuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
