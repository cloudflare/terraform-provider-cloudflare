// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ContentScanningResultEnvelope struct {
	Result ContentScanningModel `json:"result"`
}

type ContentScanningModel struct {
	ZoneID   types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Value    types.String `tfsdk:"value" json:"value,required"`
	Modified types.String `tfsdk:"modified" json:"modified,computed"`
}

func (m ContentScanningModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ContentScanningModel) MarshalJSONForUpdate(state ContentScanningModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
