// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_identity_providers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustIdentityProvidersResultEnvelope struct {
	Result ZeroTrustIdentityProvidersModel `json:"result,computed"`
}

type ZeroTrustIdentityProvidersModel struct {
	ID         types.String                               `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                               `tfsdk:"account_id" path:"account_id"`
	ZoneID     types.String                               `tfsdk:"zone_id" path:"zone_id"`
	Config     *ZeroTrustIdentityProvidersConfigModel     `tfsdk:"config" json:"config"`
	Name       types.String                               `tfsdk:"name" json:"name"`
	Type       types.String                               `tfsdk:"type" json:"type"`
	ScimConfig *ZeroTrustIdentityProvidersScimConfigModel `tfsdk:"scim_config" json:"scim_config"`
}

type ZeroTrustIdentityProvidersConfigModel struct {
	Claims                   *[]types.String `tfsdk:"claims" json:"claims"`
	ClientID                 types.String    `tfsdk:"client_id" json:"client_id"`
	ClientSecret             types.String    `tfsdk:"client_secret" json:"client_secret"`
	ConditionalAccessEnabled types.Bool      `tfsdk:"conditional_access_enabled" json:"conditional_access_enabled"`
	DirectoryID              types.String    `tfsdk:"directory_id" json:"directory_id"`
	EmailClaimName           types.String    `tfsdk:"email_claim_name" json:"email_claim_name"`
	Prompt                   types.String    `tfsdk:"prompt" json:"prompt"`
	SupportGroups            types.Bool      `tfsdk:"support_groups" json:"support_groups"`
}

type ZeroTrustIdentityProvidersScimConfigModel struct {
	Enabled                types.Bool   `tfsdk:"enabled" json:"enabled"`
	GroupMemberDeprovision types.Bool   `tfsdk:"group_member_deprovision" json:"group_member_deprovision"`
	SeatDeprovision        types.Bool   `tfsdk:"seat_deprovision" json:"seat_deprovision"`
	Secret                 types.String `tfsdk:"secret" json:"secret"`
	UserDeprovision        types.Bool   `tfsdk:"user_deprovision" json:"user_deprovision"`
}
