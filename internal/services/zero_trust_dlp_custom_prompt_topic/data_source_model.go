// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_prompt_topic

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomPromptTopicResultDataSourceEnvelope struct {
	Result ZeroTrustDLPCustomPromptTopicDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPCustomPromptTopicDataSourceModel struct {
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	EntryID     types.String      `tfsdk:"entry_id" path:"entry_id,required"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	ProfileID   types.String      `tfsdk:"profile_id" json:"profile_id,computed"`
	Topic       types.String      `tfsdk:"topic" json:"topic,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m *ZeroTrustDLPCustomPromptTopicDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPCustomPromptTopicGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPCustomPromptTopicGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
