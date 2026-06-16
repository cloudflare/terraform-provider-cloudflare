// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_compliance_modes

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/origin_tls_compliance_modes"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginTLSComplianceModesResultDataSourceEnvelope struct {
	Result OriginTLSComplianceModesDataSourceModel `json:"result,computed"`
}

type OriginTLSComplianceModesDataSourceModel struct {
	ID         types.String                   `tfsdk:"id" path:"zone_id,computed"`
	ZoneID     types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Editable   types.Bool                     `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      customfield.List[types.String] `tfsdk:"value" json:"value,computed"`
}

func (m *OriginTLSComplianceModesDataSourceModel) toReadParams(_ context.Context) (params origin_tls_compliance_modes.OriginTLSComplianceModeGetParams, diags diag.Diagnostics) {
	params = origin_tls_compliance_modes.OriginTLSComplianceModeGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
