// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationResultEnvelope struct {
	Result TeamsLocationModel `json:"result,computed"`
}

type TeamsLocationModel struct {
	ID                    types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                   `tfsdk:"account_id" path:"account_id"`
	Name                  types.String                   `tfsdk:"name" json:"name"`
	ClientDefault         types.Bool                     `tfsdk:"client_default" json:"client_default"`
	DNSDestinationIPsID   types.String                   `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id"`
	ECSSupport            types.Bool                     `tfsdk:"ecs_support" json:"ecs_support"`
	Networks              *[]*TeamsLocationNetworksModel `tfsdk:"networks" json:"networks"`
	CreatedAt             timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed"`
	DOHSubdomain          types.String                   `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	IP                    types.String                   `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination       types.String                   `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup types.String                   `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	UpdatedAt             timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed"`
}

type TeamsLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}
