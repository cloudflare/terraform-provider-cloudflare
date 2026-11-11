// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_rego_metadata

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSRegoMetadataResultEnvelope struct {
	Result ZeroTrustCWSRegoMetadataModel `json:"result"`
}

type ZeroTrustCWSRegoMetadataModel struct {
	ID         types.String         `tfsdk:"id" json:"id,computed"`
	AccountID  types.String         `tfsdk:"account_id" path:"account_id,required"`
	ImportPath types.String         `tfsdk:"import_path" json:"import_path,required"`
	Name       types.String         `tfsdk:"name" json:"name,required"`
	Data       jsontypes.Normalized `tfsdk:"data" json:"data,required"`
	CreatedAt  timetypes.RFC3339    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt  timetypes.RFC3339    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSRegoMetadataModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSRegoMetadataModel) MarshalJSONForUpdate(state ZeroTrustCWSRegoMetadataModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
