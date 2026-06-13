// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_csr

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomCsrResultEnvelope struct {
	Result CustomCsrModel `json:"result"`
}

type CustomCsrModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	AccountID          types.String      `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID             types.String      `tfsdk:"zone_id" path:"zone_id,optional"`
	CommonName         types.String      `tfsdk:"common_name" json:"common_name,required"`
	Country            types.String      `tfsdk:"country" json:"country,required"`
	Locality           types.String      `tfsdk:"locality" json:"locality,required"`
	Organization       types.String      `tfsdk:"organization" json:"organization,required"`
	State              types.String      `tfsdk:"state" json:"state,required"`
	Sans               *[]types.String   `tfsdk:"sans" json:"sans,required"`
	Description        types.String      `tfsdk:"description" json:"description,optional"`
	Name               types.String      `tfsdk:"name" json:"name,optional"`
	OrganizationalUnit types.String      `tfsdk:"organizational_unit" json:"organizational_unit,optional"`
	KeyType            types.String      `tfsdk:"key_type" json:"key_type,computed_optional"`
	AccountTag         types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Csr                types.String      `tfsdk:"csr" json:"csr,computed"`
}

func (m CustomCsrModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomCsrModel) MarshalJSONForUpdate(state CustomCsrModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
