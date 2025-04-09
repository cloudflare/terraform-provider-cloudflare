// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_logging

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayLoggingResultDataSourceEnvelope struct {
Result ZeroTrustGatewayLoggingDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayLoggingDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
RedactPii types.Bool `tfsdk:"redact_pii" json:"redact_pii,computed"`
SettingsByRuleType customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeDataSourceModel] `tfsdk:"settings_by_rule_type" json:"settings_by_rule_type,computed"`
}

func (m *ZeroTrustGatewayLoggingDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayLoggingGetParams, diags diag.Diagnostics) {
  params = zero_trust.GatewayLoggingGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeDataSourceModel struct {
DNS customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeDNSDataSourceModel] `tfsdk:"dns" json:"dns,computed"`
HTTP customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeHTTPDataSourceModel] `tfsdk:"http" json:"http,computed"`
L4 customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeL4DataSourceModel] `tfsdk:"l4" json:"l4,computed"`
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeDNSDataSourceModel struct {
LogAll types.Bool `tfsdk:"log_all" json:"log_all,computed"`
LogBlocks types.Bool `tfsdk:"log_blocks" json:"log_blocks,computed"`
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeHTTPDataSourceModel struct {
LogAll types.Bool `tfsdk:"log_all" json:"log_all,computed"`
LogBlocks types.Bool `tfsdk:"log_blocks" json:"log_blocks,computed"`
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeL4DataSourceModel struct {
LogAll types.Bool `tfsdk:"log_all" json:"log_all,computed"`
LogBlocks types.Bool `tfsdk:"log_blocks" json:"log_blocks,computed"`
}
