// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemResultEnvelope struct {
	Result ListItemModel `json:"result"`
}

type ListItemModel struct {
	ListID      types.String                                    `tfsdk:"list_id" path:"list_id,required"`
	AccountID   types.String                                    `tfsdk:"account_id" path:"account_id,optional"`
	ID          types.String                                    `tfsdk:"id" path:"item_id,computed"`
	ASN         types.Int64                                     `tfsdk:"asn" json:"asn,optional"`
	Comment     types.String                                    `tfsdk:"comment" json:"comment,optional"`
	IP          types.String                                    `tfsdk:"ip" json:"ip,optional"`
	Hostname    customfield.NestedObject[ListItemHostnameModel] `tfsdk:"hostname" json:"hostname,optional"`
	Redirect    customfield.NestedObject[ListItemRedirectModel] `tfsdk:"redirect" json:"redirect,optional"`
	OperationID types.String                                    `tfsdk:"operation_id" json:"operation_id,computed"`
	ModifiedOn  types.String                                    `tfsdk:"modified_on" json:"modified_on,computed"`
	CreatedOn   types.String                                    `tfsdk:"created_on" json:"created_on,computed"`
}

func (m ListItemModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

// MarshalSingleToCollectionJSON takes a single `ListItemModel` and wraps
// it as a collection in order to satisfy the API requirements.
//
// The endpoint is designed for bulk creation however, in the provider,
// we only manage a single resource at a time.
func (m ListItemModel) MarshalSingleToCollectionJSON(existingJSON []byte) (data []byte) {
	var output []byte
	output = append(output, []byte("[")...)
	output = append(output, existingJSON...)
	output = append(output, []byte("]")...)
	return output
}

func (m ListItemModel) MarshalJSONForUpdate(state ListItemModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ListItemHostnameModel struct {
	URLHostname types.String `tfsdk:"url_hostname" json:"url_hostname,required"`
}

type ListItemRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url" json:"source_url,required"`
	TargetURL           types.String `tfsdk:"target_url" json:"target_url,required"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains" json:"include_subdomains,computed_optional"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix,computed_optional"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string" json:"preserve_query_string,computed_optional"`
	StatusCode          types.Int64  `tfsdk:"status_code" json:"status_code,computed_optional"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching" json:"subpath_matching,computed_optional"`
}
