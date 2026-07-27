// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityTrustedDomainsResultEnvelope struct {
	Result EmailSecurityTrustedDomainsModel `json:"result"`
}

type EmailSecurityTrustedDomainsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id,required"`
	IsRecent     types.Bool        `tfsdk:"is_recent" json:"is_recent,required"`
	IsRegex      types.Bool        `tfsdk:"is_regex" json:"is_regex,required"`
	IsSimilarity types.Bool        `tfsdk:"is_similarity" json:"is_similarity,required"`
	Pattern      types.String      `tfsdk:"pattern" json:"pattern,required"`
	Comments     types.String      `tfsdk:"comments" json:"comments,optional"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	ModifiedAt   timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

func (m EmailSecurityTrustedDomainsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailSecurityTrustedDomainsModel) MarshalJSONForUpdate(state EmailSecurityTrustedDomainsModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
