// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemResultEnvelope struct {
	Result ListItemModel `json:"result"`
}

type ListItemModel struct {
	ListID            types.String           `tfsdk:"list_id" path:"list_id"`
	AccountID         types.String           `tfsdk:"account_id" path:"account_id"`
	AccountIdentifier types.String           `tfsdk:"account_identifier" path:"account_identifier"`
	ItemID            types.String           `tfsdk:"item_id" path:"item_id"`
	ASN               types.Int64            `tfsdk:"asn" json:"asn"`
	Comment           types.String           `tfsdk:"comment" json:"comment"`
	IP                types.String           `tfsdk:"ip" json:"ip"`
	Hostname          *ListItemHostnameModel `tfsdk:"hostname" json:"hostname"`
	Redirect          *ListItemRedirectModel `tfsdk:"redirect" json:"redirect"`
	OperationID       types.String           `tfsdk:"operation_id" json:"operation_id,computed"`
}

type ListItemHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname" json:"url_hostname"`
}

type ListItemRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url" json:"source_url"`
	TargetURL           types.String `tfsdk:"target_url" json:"target_url"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains" json:"include_subdomains"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string" json:"preserve_query_string"`
	StatusCode          types.Int64  `tfsdk:"status_code" json:"status_code"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching" json:"subpath_matching"`
}
