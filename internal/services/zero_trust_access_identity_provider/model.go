// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProviderResultEnvelope struct {
	Result ZeroTrustAccessIdentityProviderModel `json:"result"`
}

type ZeroTrustAccessIdentityProviderModel struct {
	ID         types.String                                    `tfsdk:"id" json:"id,computed_optional"`
	AccountID  types.String                                    `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID     types.String                                    `tfsdk:"zone_id" path:"zone_id,optional"`
	Name       types.String                                    `tfsdk:"name" json:"name,required"`
	Type       types.String                                    `tfsdk:"type" json:"type,required"`
	Config     *ZeroTrustAccessIdentityProviderConfigModel     `tfsdk:"config" json:"config,required"`
	SCIMConfig *ZeroTrustAccessIdentityProviderSCIMConfigModel `tfsdk:"scim_config" json:"scim_config,optional"`
}

type ZeroTrustAccessIdentityProviderConfigModel struct {
	Claims                   customfield.List[types.String] `tfsdk:"claims" json:"claims,computed_optional"`
	ClientID                 types.String                   `tfsdk:"client_id" json:"client_id,computed_optional"`
	ClientSecret             types.String                   `tfsdk:"client_secret" json:"client_secret,computed_optional"`
	ConditionalAccessEnabled types.Bool                     `tfsdk:"conditional_access_enabled" json:"conditional_access_enabled,computed_optional"`
	DirectoryID              types.String                   `tfsdk:"directory_id" json:"directory_id,computed_optional"`
	EmailClaimName           types.String                   `tfsdk:"email_claim_name" json:"email_claim_name,computed_optional"`
	Prompt                   types.String                   `tfsdk:"prompt" json:"prompt,computed_optional"`
	SupportGroups            types.Bool                     `tfsdk:"support_groups" json:"support_groups,computed_optional"`
}

type ZeroTrustAccessIdentityProviderSCIMConfigModel struct {
	Enabled                types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	GroupMemberDeprovision types.Bool   `tfsdk:"group_member_deprovision" json:"group_member_deprovision,computed_optional"`
	SeatDeprovision        types.Bool   `tfsdk:"seat_deprovision" json:"seat_deprovision,computed_optional"`
	Secret                 types.String `tfsdk:"secret" json:"secret,computed_optional"`
	UserDeprovision        types.Bool   `tfsdk:"user_deprovision" json:"user_deprovision,computed_optional"`
}
