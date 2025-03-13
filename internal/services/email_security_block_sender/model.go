// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_block_sender

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityBlockSenderResultEnvelope struct {
Result EmailSecurityBlockSenderModel `json:"result"`
}

type EmailSecurityBlockSenderModel struct {
ID types.Int64 `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
IsRegex types.Bool `tfsdk:"is_regex" json:"is_regex,required"`
Pattern types.String `tfsdk:"pattern" json:"pattern,required"`
PatternType types.String `tfsdk:"pattern_type" json:"pattern_type,required"`
Comments types.String `tfsdk:"comments" json:"comments,optional"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
LastModified timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
}

func (m EmailSecurityBlockSenderModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m EmailSecurityBlockSenderModel) MarshalJSONForUpdate(state EmailSecurityBlockSenderModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}
