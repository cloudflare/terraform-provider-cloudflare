// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemResultEnvelope struct {
	Result *[]*ListItemBodyModel `json:"result"`
}

type ListItemModel struct {
	ListID              types.String          `tfsdk:"list_id" path:"list_id,required"`
	AccountID           types.String          `tfsdk:"account_id" path:"account_id,optional"`
	AccountIdentifier   types.String          `tfsdk:"account_identifier" path:"account_identifier,optional"`
	ItemID              types.String          `tfsdk:"item_id" path:"item_id,optional"`
	Body                *[]*ListItemBodyModel `tfsdk:"body" json:"body,required"`
	IncludeSubdomains   types.Bool            `tfsdk:"include_subdomains" json:"include_subdomains,computed"`
	OperationID         types.String          `tfsdk:"operation_id" json:"operation_id,computed"`
	PreservePathSuffix  types.Bool            `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix,computed"`
	PreserveQueryString types.Bool            `tfsdk:"preserve_query_string" json:"preserve_query_string,computed"`
	SourceURL           types.String          `tfsdk:"source_url" json:"source_url,computed"`
	StatusCode          types.Int64           `tfsdk:"status_code" json:"status_code,computed"`
	SubpathMatching     types.Bool            `tfsdk:"subpath_matching" json:"subpath_matching,computed"`
	TargetURL           types.String          `tfsdk:"target_url" json:"target_url,computed"`
	URLHostname         types.String          `tfsdk:"url_hostname" json:"url_hostname,computed"`
}

func (m ListItemModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m ListItemModel) MarshalJSONForUpdate(state ListItemModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}

type ListItemBodyModel struct {
	ASN      types.Int64                `tfsdk:"asn" json:"asn,optional"`
	Comment  types.String               `tfsdk:"comment" json:"comment,optional"`
	Hostname *ListItemBodyHostnameModel `tfsdk:"hostname" json:"hostname,optional"`
	IP       types.String               `tfsdk:"ip" json:"ip,optional"`
	Redirect *ListItemBodyRedirectModel `tfsdk:"redirect" json:"redirect,optional"`
}

type ListItemBodyHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname" json:"url_hostname,required"`
}

type ListItemBodyRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url" json:"source_url,required"`
	TargetURL           types.String `tfsdk:"target_url" json:"target_url,required"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains" json:"include_subdomains,computed_optional"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix,computed_optional"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string" json:"preserve_query_string,computed_optional"`
	StatusCode          types.Int64  `tfsdk:"status_code" json:"status_code,computed_optional"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching" json:"subpath_matching,computed_optional"`
}
