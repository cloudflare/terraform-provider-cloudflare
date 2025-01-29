// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameFallbackOriginResultEnvelope struct {
	Result CustomHostnameFallbackOriginModel `json:"result"`
}

type CustomHostnameFallbackOriginModel struct {
	ID        types.String                   `tfsdk:"id" json:"-,computed"`
	ZoneID    types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Origin    types.String                   `tfsdk:"origin" json:"origin,required"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Status    types.String                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Errors    customfield.List[types.String] `tfsdk:"errors" json:"errors,computed"`
}

func (m CustomHostnameFallbackOriginModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomHostnameFallbackOriginModel) MarshalJSONForUpdate(state CustomHostnameFallbackOriginModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
