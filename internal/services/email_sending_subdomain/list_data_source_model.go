// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_sending"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSendingSubdomainsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSendingSubdomainsResultDataSourceModel] `json:"result,computed"`
}

type EmailSendingSubdomainsDataSourceModel struct {
	ZoneID   types.String                                                              `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                               `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[EmailSendingSubdomainsResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailSendingSubdomainsDataSourceModel) toListParams(_ context.Context) (params email_sending.SubdomainListParams, diags diag.Diagnostics) {
	params = email_sending.SubdomainListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type EmailSendingSubdomainsResultDataSourceModel struct {
	ID                       types.String      `tfsdk:"id" json:"tag,computed"`
	Enabled                  types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Name                     types.String      `tfsdk:"name" json:"name,computed"`
	Tag                      types.String      `tfsdk:"tag" json:"tag,computed"`
	Created                  timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	DKIMSelector             types.String      `tfsdk:"dkim_selector" json:"dkim_selector,computed"`
	DropSuppressedRecipients types.Bool        `tfsdk:"drop_suppressed_recipients" json:"drop_suppressed_recipients,computed"`
	Modified                 timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	PreviewEnabled           types.Bool        `tfsdk:"preview_enabled" json:"preview_enabled,computed"`
	ReturnPathDomain         types.String      `tfsdk:"return_path_domain" json:"return_path_domain,computed"`
}
