// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_auto_origin_tls_kex

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneAutoOriginTLSKexResultEnvelope struct {
	Result ZoneAutoOriginTLSKexModel `json:"result"`
}

type ZoneAutoOriginTLSKexModel struct {
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,required"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m ZoneAutoOriginTLSKexModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneAutoOriginTLSKexModel) MarshalJSONForUpdate(state ZoneAutoOriginTLSKexModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
