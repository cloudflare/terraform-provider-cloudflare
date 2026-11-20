// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_auth_dmarc_report

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailAuthDMARCReportResultEnvelope struct {
	Result EmailAuthDMARCReportModel `json:"result"`
}

type EmailAuthDMARCReportModel struct {
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed_optional"`
	SkipWizard types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard,computed_optional"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	RuaPrefix  types.String      `tfsdk:"rua_prefix" json:"rua_prefix,computed"`
	Status     types.String      `tfsdk:"status" json:"status,computed"`
}

func (m EmailAuthDMARCReportModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailAuthDMARCReportModel) MarshalJSONForUpdate(state EmailAuthDMARCReportModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
