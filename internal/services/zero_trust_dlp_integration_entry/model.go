// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_integration_entry

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPIntegrationEntryResultEnvelope struct {
	Result ZeroTrustDLPIntegrationEntryModel `json:"result"`
}

type ZeroTrustDLPIntegrationEntryModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	EntryID   types.String      `tfsdk:"entry_id" json:"entry_id,required"`
	ProfileID types.String      `tfsdk:"profile_id" json:"profile_id,optional"`
	Enabled   types.Bool        `tfsdk:"enabled" json:"enabled,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustDLPIntegrationEntryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPIntegrationEntryModel) MarshalJSONForUpdate(state ZeroTrustDLPIntegrationEntryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
