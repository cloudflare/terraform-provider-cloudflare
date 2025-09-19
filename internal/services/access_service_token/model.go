// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_service_token

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessServiceTokenResultEnvelope struct {
	Result AccessServiceTokenModel `json:"result"`
}

type AccessServiceTokenModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID       types.String      `tfsdk:"zone_id" path:"zone_id,optional"`
	Name         types.String      `tfsdk:"name" json:"name,required"`
	Duration     types.String      `tfsdk:"duration" json:"duration,computed_optional"`
	ClientID     types.String      `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret types.String      `tfsdk:"client_secret" json:"client_secret,computed,no_refresh"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
}

func (m AccessServiceTokenModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessServiceTokenModel) MarshalJSONForUpdate(state AccessServiceTokenModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
