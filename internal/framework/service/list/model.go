package list

import "github.com/hashicorp/terraform-plugin-framework/types"

type ListModel struct {
	ID          types.String    `tfsdk:"id"`
	AccountID   types.String    `tfsdk:"account_id"`
	Name        types.String    `tfsdk:"name"`
	Description types.String    `tfsdk:"description"`
	Kind        types.String    `tfsdk:"kind"`
	Items       []ListItemModel `tfsdk:"item"`
}

type ListItemModel struct {
	IP       types.String             `tfsdk:"ip"`
	ASN      types.Int64              `tfsdk:"asn"`
	Hostname []*ListItemHostnameModel `tfsdk:"hostname"`
	Redirect []*ListItemRedirectModel `tfsdk:"redirect"`
	Comment  types.String             `tfsdk:"comment"`
}

type ListItemHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname"`
}

type ListItemRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	IncludeSubdomains   types.String `tfsdk:"include_subdomains"`
	SubpathMatching     types.String `tfsdk:"subpath_matching"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	PreservePathSuffix  types.String `tfsdk:"preserve_path_suffix"`
	PreserveQueryString types.String `tfsdk:"preserve_query_string"`
}
