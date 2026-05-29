// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_prompt_topic

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomPromptTopicsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPCustomPromptTopicsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPCustomPromptTopicsDataSourceModel struct {
	AccountID types.String                                                                      `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDLPCustomPromptTopicsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPCustomPromptTopicsDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPCustomPromptTopicListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPCustomPromptTopicListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPCustomPromptTopicsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled     types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Topic       types.String      `tfsdk:"topic" json:"topic,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	ProfileID   types.String      `tfsdk:"profile_id" json:"profile_id,computed"`
}
