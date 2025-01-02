// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/leaked_credential_checks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LeakedCredentialCheckRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LeakedCredentialCheckRuleDataSourceModel] `json:"result,computed"`
}

type LeakedCredentialCheckRuleDataSourceModel struct {
	ID       types.String                                       `tfsdk:"id" json:"id,computed"`
	Password types.String                                       `tfsdk:"password" json:"password,computed"`
	Username types.String                                       `tfsdk:"username" json:"username,computed"`
	Filter   *LeakedCredentialCheckRuleFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *LeakedCredentialCheckRuleDataSourceModel) toListParams(_ context.Context) (params leaked_credential_checks.DetectionListParams, diags diag.Diagnostics) {
	params = leaked_credential_checks.DetectionListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type LeakedCredentialCheckRuleFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}
