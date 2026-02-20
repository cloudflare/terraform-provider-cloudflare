package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceListItemModel represents the v4 (SDKv2) cloudflare_list_item state structure.
// In v4, hostname and redirect were TypeList with MaxItems: 1 (stored as arrays in state).
type SourceListItemModel struct {
	AccountID types.String             `tfsdk:"account_id"`
	ListID    types.String             `tfsdk:"list_id"`
	ID        types.String             `tfsdk:"id"`
	IP        types.String             `tfsdk:"ip"`
	ASN       types.Int64              `tfsdk:"asn"`
	Comment   types.String             `tfsdk:"comment"`
	Hostname  []SourceHostnameModel    `tfsdk:"hostname"`
	Redirect  []SourceRedirectModel    `tfsdk:"redirect"`
	CreatedOn types.String             `tfsdk:"created_on"`
}

// SourceHostnameModel represents the v4 hostname block structure.
type SourceHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname"`
}

// SourceRedirectModel represents the v4 redirect block structure.
// Boolean fields are stored as strings ("enabled"/"disabled") in v4.
type SourceRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	IncludeSubdomains   types.String `tfsdk:"include_subdomains"`
	SubpathMatching     types.String `tfsdk:"subpath_matching"`
	PreserveQueryString types.String `tfsdk:"preserve_query_string"`
	PreservePathSuffix  types.String `tfsdk:"preserve_path_suffix"`
}

// TargetListItemModel represents the v5 cloudflare_list_item state structure.
type TargetListItemModel struct {
	AccountID   types.String                                    `tfsdk:"account_id"`
	ListID      types.String                                    `tfsdk:"list_id"`
	ID          types.String                                    `tfsdk:"id"`
	ASN         types.Int64                                     `tfsdk:"asn"`
	Comment     types.String                                    `tfsdk:"comment"`
	IP          types.String                                    `tfsdk:"ip"`
	Hostname    customfield.NestedObject[TargetHostnameModel]   `tfsdk:"hostname"`
	Redirect    customfield.NestedObject[TargetRedirectModel]   `tfsdk:"redirect"`
	OperationID types.String                                    `tfsdk:"operation_id"`
	ModifiedOn  types.String                                    `tfsdk:"modified_on"`
	CreatedOn   types.String                                    `tfsdk:"created_on"`
}

// SourceListItemV1Model represents the v4.52.5 framework cloudflare_list_item state (schema_version=1).
// In v4.52.5, hostname and redirect are ListNestedBlocks.
// Redirect boolean fields are actual bools (v4.52.5's 0→1 upgrader converted from strings).
type SourceListItemV1Model struct {
	AccountID types.String            `tfsdk:"account_id"`
	ListID    types.String            `tfsdk:"list_id"`
	ID        types.String            `tfsdk:"id"`
	IP        types.String            `tfsdk:"ip"`
	ASN       types.Int64             `tfsdk:"asn"`
	Comment   types.String            `tfsdk:"comment"`
	Hostname  []SourceHostnameModel   `tfsdk:"hostname"`
	Redirect  []SourceRedirectV1Model `tfsdk:"redirect"`
}

// SourceRedirectV1Model represents the v4.52.5 redirect block (schema_version=1).
// Boolean fields are actual bools (already converted from "enabled"/"disabled" by v4.52.5's upgrader).
type SourceRedirectV1Model struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix"`
}

// TargetHostnameModel represents the v5 hostname nested object.
type TargetHostnameModel struct {
	URLHostname          types.String `tfsdk:"url_hostname"`
	ExcludeExactHostname types.Bool   `tfsdk:"exclude_exact_hostname"`
}

// TargetRedirectModel represents the v5 redirect nested object.
// Boolean fields are actual booleans in v5 (not "enabled"/"disabled" strings).
type TargetRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url"`
	TargetURL           types.String `tfsdk:"target_url"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string"`
	StatusCode          types.Int64  `tfsdk:"status_code"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching"`
}
