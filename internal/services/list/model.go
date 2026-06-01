// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultEnvelope struct {
	Result ListModel `json:"result"`
}

type ListModel struct {
	ID                    types.String                               `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                               `tfsdk:"account_id" path:"account_id,required"`
	Kind                  types.String                               `tfsdk:"kind" json:"kind,required"`
	Name                  types.String                               `tfsdk:"name" json:"name,required"`
	Description           types.String                               `tfsdk:"description" json:"description,optional"`
	CreatedOn             types.String                               `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn            types.String                               `tfsdk:"modified_on" json:"modified_on,computed"`
	NumItems              types.Float64                              `tfsdk:"num_items" json:"num_items,computed"`
	NumReferencingFilters types.Float64                              `tfsdk:"num_referencing_filters" json:"num_referencing_filters,computed"`
	Items                 customfield.NestedObjectSet[ListItemModel] `tfsdk:"items" json:"items,optional"`
}

func (m ListModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ListModel) MarshalJSONForUpdate(state ListModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ListItemModel struct {
	ASN      types.Int64                                     `tfsdk:"asn" json:"asn,optional"`
	Comment  types.String                                    `tfsdk:"comment" json:"comment,optional"`
	IP       types.String                                    `tfsdk:"ip" json:"ip,optional"`
	Hostname customfield.NestedObject[ListItemHostnameModel] `tfsdk:"hostname" json:"hostname,optional"`
	Redirect customfield.NestedObject[ListItemRedirectModel] `tfsdk:"redirect" json:"redirect,optional"`
}

func (m ListItemModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

type ListItemHostnameModel struct {
	URLHostname          types.String `tfsdk:"url_hostname" json:"url_hostname,required"`
	ExcludeExactHostname types.Bool   `tfsdk:"exclude_exact_hostname" json:"exclude_exact_hostname,optional"`
}

type ListItemRedirectModel struct {
	SourceURL           types.String `tfsdk:"source_url" json:"source_url,required"`
	TargetURL           types.String `tfsdk:"target_url" json:"target_url,required"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains" json:"include_subdomains,optional"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix,optional"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string" json:"preserve_query_string,optional"`
	StatusCode          types.Int64  `tfsdk:"status_code" json:"status_code,optional"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching" json:"subpath_matching,optional"`
}
