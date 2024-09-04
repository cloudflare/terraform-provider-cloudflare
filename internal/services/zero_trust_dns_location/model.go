// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationResultEnvelope struct {
	Result ZeroTrustDNSLocationModel `json:"result"`
}

type ZeroTrustDNSLocationModel struct {
	ID                    types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	ClientDefault         types.Bool                                                      `tfsdk:"client_default" json:"client_default,computed_optional"`
	DNSDestinationIPsID   types.String                                                    `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id,computed_optional"`
	ECSSupport            types.Bool                                                      `tfsdk:"ecs_support" json:"ecs_support,computed_optional"`
	Name                  types.String                                                    `tfsdk:"name" json:"name,computed_optional"`
	Endpoints             customfield.NestedObject[ZeroTrustDNSLocationEndpointsModel]    `tfsdk:"endpoints" json:"endpoints,computed_optional"`
	Networks              customfield.NestedObjectList[ZeroTrustDNSLocationNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
	CreatedAt             timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DOHSubdomain          types.String                                                    `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	IP                    types.String                                                    `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination       types.String                                                    `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup types.String                                                    `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	UpdatedAt             timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustDNSLocationEndpointsModel struct {
	DOH  customfield.NestedObject[ZeroTrustDNSLocationEndpointsDOHModel]  `tfsdk:"doh" json:"doh,computed_optional"`
	DOT  customfield.NestedObject[ZeroTrustDNSLocationEndpointsDOTModel]  `tfsdk:"dot" json:"dot,computed_optional"`
	IPV4 customfield.NestedObject[ZeroTrustDNSLocationEndpointsIPV4Model] `tfsdk:"ipv4" json:"ipv4,computed_optional"`
	IPV6 customfield.NestedObject[ZeroTrustDNSLocationEndpointsIPV6Model] `tfsdk:"ipv6" json:"ipv6,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOHModel struct {
	Enabled      types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks     customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsDOHNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
	RequireToken types.Bool                                                                  `tfsdk:"require_token" json:"require_token,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOHNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOTModel struct {
	Enabled  types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsDOTNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOTNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsIPV4Model struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsIPV6Model struct {
	Enabled  types.Bool                                                                   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsIPV6NetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsIPV6NetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,computed_optional"`
}

type ZeroTrustDNSLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,computed_optional"`
}
