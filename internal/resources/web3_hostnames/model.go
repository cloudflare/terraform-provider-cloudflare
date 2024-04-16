// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostnames

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnamesResultEnvelope struct {
	Result Web3HostnamesModel `json:"result,computed"`
}

type Web3HostnamesModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	Target         types.String `tfsdk:"target" json:"target"`
	Description    types.String `tfsdk:"description" json:"description"`
	Dnslink        types.String `tfsdk:"dnslink" json:"dnslink"`
}
