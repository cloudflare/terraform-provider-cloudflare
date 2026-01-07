// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/leaked_credential_checks"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LeakedCredentialCheckRuleResultDataSourceEnvelope struct {
	Result LeakedCredentialCheckRuleDataSourceModel `json:"result,computed"`
}

type LeakedCredentialCheckRuleDataSourceModel struct {
	ID          types.String `tfsdk:"id" path:"detection_id,computed"`
	DetectionID types.String `tfsdk:"detection_id" path:"detection_id,required"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Password    types.String `tfsdk:"password" json:"password,computed"`
	Username    types.String `tfsdk:"username" json:"username,computed"`
}

func (m *LeakedCredentialCheckRuleDataSourceModel) toReadParams(_ context.Context) (params leaked_credential_checks.DetectionGetParams, diags diag.Diagnostics) {
	params = leaked_credential_checks.DetectionGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
