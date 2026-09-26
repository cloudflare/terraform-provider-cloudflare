// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbPolicyResultEnvelope struct {
	Result ZeroTrustCasbPolicyModel `json:"result"`
}

type ZeroTrustCasbPolicyModel struct {
	ID                       types.String                     `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                     `tfsdk:"account_id" path:"account_id,required"`
	FindingTypeID            types.String                     `tfsdk:"finding_type_id" json:"finding_type_id,required"`
	AppliesToAllIntegrations types.Bool                       `tfsdk:"applies_to_all_integrations" json:"applies_to_all_integrations,required"`
	DisplayName              types.String                     `tfsdk:"display_name" json:"display_name,required"`
	Enabled                  types.Bool                       `tfsdk:"enabled" json:"enabled,required"`
	Actions                  *ZeroTrustCasbPolicyActionsModel `tfsdk:"actions" json:"actions,required"`
	Description              types.String                     `tfsdk:"description" json:"description,computed_optional"`
	IntegrationIDs           customfield.List[types.String]   `tfsdk:"integration_ids" json:"integration_ids,computed_optional"`
	CreatedAt                timetypes.RFC3339                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DisabledAt               timetypes.RFC3339                `tfsdk:"disabled_at" json:"disabled_at,computed" format:"date-time"`
	LastTriggeredAt          timetypes.RFC3339                `tfsdk:"last_triggered_at" json:"last_triggered_at,computed" format:"date-time"`
	UpdatedAt                timetypes.RFC3339                `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCasbPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCasbPolicyModel) MarshalJSONForUpdate(state ZeroTrustCasbPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustCasbPolicyActionsModel struct {
	RemediationTypes *[]*ZeroTrustCasbPolicyActionsRemediationTypesModel `tfsdk:"remediation_types" json:"remediation_types,optional"`
	WebhookConfigs   *[]*ZeroTrustCasbPolicyActionsWebhookConfigsModel   `tfsdk:"webhook_configs" json:"webhook_configs,optional"`
}

type ZeroTrustCasbPolicyActionsRemediationTypesModel struct {
	RemediationTypeID types.String `tfsdk:"remediation_type_id" json:"remediation_type_id,required"`
}

type ZeroTrustCasbPolicyActionsWebhookConfigsModel struct {
	WebhookConfigID types.String `tfsdk:"webhook_config_id" json:"webhook_config_id,required"`
}
