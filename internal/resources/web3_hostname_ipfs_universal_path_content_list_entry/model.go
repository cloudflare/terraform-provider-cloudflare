// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname_ipfs_universal_path_content_list_entry

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnameIPFSUniversalPathContentListEntryResultEnvelope struct {
	Result Web3HostnameIPFSUniversalPathContentListEntryModel `json:"result,computed"`
}

type Web3HostnameIPFSUniversalPathContentListEntryModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	Identifier     types.String `tfsdk:"identifier" path:"identifier"`
	Content        types.String `tfsdk:"content" json:"content"`
	Type           types.String `tfsdk:"type" json:"type"`
	Description    types.String `tfsdk:"description" json:"description"`
}
