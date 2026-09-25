// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSendingSubdomainResultEnvelope struct {
	Result EmailSendingSubdomainModel `json:"result"`
}

type EmailSendingSubdomainModel struct {
	ID                       types.String      `tfsdk:"id" json:"-,computed"`
	Tag                      types.String      `tfsdk:"tag" json:"tag,computed"`
	ZoneID                   types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Name                     types.String      `tfsdk:"name" json:"name,required"`
	DropSuppressedRecipients types.Bool        `tfsdk:"drop_suppressed_recipients" json:"drop_suppressed_recipients,computed_optional"`
	PreviewEnabled           types.Bool        `tfsdk:"preview_enabled" json:"preview_enabled,computed_optional"`
	Created                  timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	DKIMSelector             types.String      `tfsdk:"dkim_selector" json:"dkim_selector,computed"`
	Enabled                  types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Modified                 timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	ReturnPathDomain         types.String      `tfsdk:"return_path_domain" json:"return_path_domain,computed"`
}

func (m EmailSendingSubdomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailSendingSubdomainModel) MarshalJSONForUpdate(state EmailSendingSubdomainModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
