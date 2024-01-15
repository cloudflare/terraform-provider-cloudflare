package list_item

import "github.com/hashicorp/terraform-plugin-framework/types"

type ListItemModelV0 struct {
	AccountID types.String               `tfsdk:"account_id"`
	ListID    types.String               `tfsdk:"list_id"`
	ID        types.String               `tfsdk:"id"`
	IP        types.String               `tfsdk:"ip"`
	ASN       types.Int64                `tfsdk:"asn"`
	Hostname  []*ListItemHostnameModel   `tfsdk:"hostname"`
	Redirect  []*ListItemRedirectModelV0 `tfsdk:"redirect"`
	Comment   types.String               `tfsdk:"comment"`
}

type ListItemModelV1 struct {
	AccountID types.String               `tfsdk:"account_id"`
	ListID    types.String               `tfsdk:"list_id"`
	ID        types.String               `tfsdk:"id"`
	IP        types.String               `tfsdk:"ip"`
	ASN       types.Int64                `tfsdk:"asn"`
	Hostname  []*ListItemHostnameModel   `tfsdk:"hostname"`
	Redirect  []*ListItemRedirectModelV1 `tfsdk:"redirect"`
	Comment   types.String               `tfsdk:"comment"`
}

type ListItemHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname"`
}

type ListItemRedirectModelV1 struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string"`
}

type ListItemRedirectModelV0 struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	IncludeSubdomains   types.String `tfsdk:"include_subdomains"`
	SubpathMatching     types.String `tfsdk:"subpath_matching"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	PreservePathSuffix  types.String `tfsdk:"preserve_path_suffix"`
	PreserveQueryString types.String `tfsdk:"preserve_query_string"`
}
