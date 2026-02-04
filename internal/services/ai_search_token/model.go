// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_token

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchTokenResultEnvelope struct {
	Result AISearchTokenModel `json:"result"`
}

type AISearchTokenModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	AccountID  types.String      `tfsdk:"account_id" path:"account_id,required"`
	CfAPIID    types.String      `tfsdk:"cf_api_id" json:"cf_api_id,required"`
	CfAPIKey   types.String      `tfsdk:"cf_api_key" json:"cf_api_key,required"`
	Name       types.String      `tfsdk:"name" json:"name,required"`
	AccountTag types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy  types.String      `tfsdk:"created_by" json:"created_by,computed"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Legacy     types.Bool        `tfsdk:"legacy" json:"legacy,computed"`
	ModifiedAt timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy types.String      `tfsdk:"modified_by" json:"modified_by,computed"`
	SyncedAt   timetypes.RFC3339 `tfsdk:"synced_at" json:"synced_at,computed" format:"date-time"`
}

func (m AISearchTokenModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AISearchTokenModel) MarshalJSONForUpdate(state AISearchTokenModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
