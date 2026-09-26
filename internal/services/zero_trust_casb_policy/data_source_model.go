// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbPolicyResultDataSourceEnvelope struct {
	Result ZeroTrustCasbPolicyDataSourceModel `json:"result,computed"`
}

type ZeroTrustCasbPolicyDataSourceModel struct {
	ID                       types.String                                                        `tfsdk:"id" path:"policy_id,computed"`
	PolicyID                 types.String                                                        `tfsdk:"policy_id" path:"policy_id,required"`
	AccountID                types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	AppliesToAllIntegrations types.Bool                                                          `tfsdk:"applies_to_all_integrations" json:"applies_to_all_integrations,computed"`
	CreatedAt                timetypes.RFC3339                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description              types.String                                                        `tfsdk:"description" json:"description,computed"`
	DisabledAt               timetypes.RFC3339                                                   `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	DisplayName              types.String                                                        `tfsdk:"display_name" json:"display_name,computed"`
	Enabled                  types.Bool                                                          `tfsdk:"enabled" json:"enabled,computed"`
	FindingTypeID            types.String                                                        `tfsdk:"finding_type_id" json:"finding_type_id,computed"`
	LastTriggeredAt          timetypes.RFC3339                                                   `tfsdk:"last_triggered_at" json:"last_triggered_at,computed" format:"date-time"`
	UpdatedAt                timetypes.RFC3339                                                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	IntegrationIDs           customfield.List[types.String]                                      `tfsdk:"integration_ids" json:"integration_ids,computed"`
	Actions                  customfield.NestedObject[ZeroTrustCasbPolicyActionsDataSourceModel] `tfsdk:"actions" json:"actions,computed"`
}

func (m *ZeroTrustCasbPolicyDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CasbPosturePolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.CasbPosturePolicyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustCasbPolicyActionsDataSourceModel struct {
	RemediationTypes customfield.NestedObjectList[ZeroTrustCasbPolicyActionsRemediationTypesDataSourceModel] `tfsdk:"remediation_types" json:"remediation_types,computed"`
	WebhookConfigs   customfield.NestedObjectList[ZeroTrustCasbPolicyActionsWebhookConfigsDataSourceModel]   `tfsdk:"webhook_configs" json:"webhook_configs,computed"`
}

type ZeroTrustCasbPolicyActionsRemediationTypesDataSourceModel struct {
	DisplayName       types.String `tfsdk:"display_name" json:"display_name,computed"`
	RemediationType   types.String `tfsdk:"remediation_type" json:"remediation_type,computed"`
	RemediationTypeID types.String `tfsdk:"remediation_type_id" json:"remediation_type_id,computed"`
}

type ZeroTrustCasbPolicyActionsWebhookConfigsDataSourceModel struct {
	DisplayName     types.String `tfsdk:"display_name" json:"display_name,computed"`
	WebhookConfigID types.String `tfsdk:"webhook_config_id" json:"webhook_config_id,computed"`
}
