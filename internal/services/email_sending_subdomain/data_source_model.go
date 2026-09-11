// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_sending"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSendingSubdomainResultDataSourceEnvelope struct {
	Result EmailSendingSubdomainDataSourceModel `json:"result,computed"`
}

type EmailSendingSubdomainDataSourceModel struct {
	ID                       types.String      `tfsdk:"id" path:"subdomain_id,computed"`
	SubdomainID              types.String      `tfsdk:"subdomain_id" path:"subdomain_id,required"`
	ZoneID                   types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Created                  timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	DKIMSelector             types.String      `tfsdk:"dkim_selector" json:"dkim_selector,computed"`
	DropSuppressedRecipients types.Bool        `tfsdk:"drop_suppressed_recipients" json:"drop_suppressed_recipients,computed"`
	Enabled                  types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Modified                 timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name                     types.String      `tfsdk:"name" json:"name,computed"`
	PreviewEnabled           types.Bool        `tfsdk:"preview_enabled" json:"preview_enabled,computed"`
	ReturnPathDomain         types.String      `tfsdk:"return_path_domain" json:"return_path_domain,computed"`
	Tag                      types.String      `tfsdk:"tag" json:"tag,computed"`
}

func (m *EmailSendingSubdomainDataSourceModel) toReadParams(_ context.Context) (params email_sending.SubdomainGetParams, diags diag.Diagnostics) {
	params = email_sending.SubdomainGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
