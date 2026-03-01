package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SourceListModel represents the combined cloudflare_list state structure
// that can hold both v4 (SDKv2) and v5 (Framework) state data.
// v4 state uses Item blocks; v5 state uses Items attribute.
type SourceListModel struct {
	ID                    types.String      `tfsdk:"id"`
	AccountID             types.String      `tfsdk:"account_id"`
	Kind                  types.String      `tfsdk:"kind"`
	Name                  types.String      `tfsdk:"name"`
	Description           types.String      `tfsdk:"description"`
	NumItems              types.Float64     `tfsdk:"num_items"`
	Item                  []SourceItemModel `tfsdk:"item"`
	// v5 fields (null when deserializing v4 state)
	Items                 basetypes.SetValue `tfsdk:"items"`
	CreatedOn             types.String       `tfsdk:"created_on"`
	ModifiedOn            types.String       `tfsdk:"modified_on"`
	NumReferencingFilters types.Float64      `tfsdk:"num_referencing_filters"`
}

// SourceItemModel represents a v4 item block.
type SourceItemModel struct {
	Comment types.String      `tfsdk:"comment"`
	Value   []SourceValueModel `tfsdk:"value"`
}

// SourceValueModel represents a v4 value block within an item block.
type SourceValueModel struct {
	IP       types.String           `tfsdk:"ip"`
	ASN      types.Int64            `tfsdk:"asn"`
	Hostname []SourceHostnameModel  `tfsdk:"hostname"`
	Redirect []SourceRedirectModel  `tfsdk:"redirect"`
}

// SourceHostnameModel represents a v4 hostname block.
type SourceHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname"`
}

// SourceRedirectModel represents a v4 redirect block.
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

// TargetListModel represents the v5 cloudflare_list state structure.
type TargetListModel struct {
	ID                    types.String                                       `tfsdk:"id"`
	AccountID             types.String                                       `tfsdk:"account_id"`
	Kind                  types.String                                       `tfsdk:"kind"`
	Name                  types.String                                       `tfsdk:"name"`
	Description           types.String                                       `tfsdk:"description"`
	CreatedOn             types.String                                       `tfsdk:"created_on"`
	ModifiedOn            types.String                                       `tfsdk:"modified_on"`
	NumItems              types.Float64                                      `tfsdk:"num_items"`
	NumReferencingFilters types.Float64                                      `tfsdk:"num_referencing_filters"`
	Items                 customfield.NestedObjectSet[TargetListItemModel]   `tfsdk:"items"`
}

// TargetListItemModel represents a v5 item within the items set.
type TargetListItemModel struct {
	ASN      types.Int64                                      `tfsdk:"asn"`
	Comment  types.String                                     `tfsdk:"comment"`
	IP       types.String                                     `tfsdk:"ip"`
	Hostname customfield.NestedObject[TargetHostnameModel]    `tfsdk:"hostname"`
	Redirect customfield.NestedObject[TargetRedirectModel]    `tfsdk:"redirect"`
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
