// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultDataSourceEnvelope struct {
	Result ListDataSourceModel `json:"result,computed"`
}

type ListDataSourceModel struct {
	AccountID             types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	ListID                types.String                                         `tfsdk:"list_id" path:"list_id,required"`
	ID                    types.String                                         `tfsdk:"id" path:"list_id,computed"`
	CreatedOn             types.String                                         `tfsdk:"created_on" json:"created_on,computed"`
	Description           types.String                                         `tfsdk:"description" json:"description,computed"`
	Kind                  types.String                                         `tfsdk:"kind" json:"kind,computed"`
	ModifiedOn            types.String                                         `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                  types.String                                         `tfsdk:"name" json:"name,computed"`
	NumItems              types.Float64                                        `tfsdk:"num_items" json:"num_items,computed"`
	NumReferencingFilters types.Float64                                        `tfsdk:"num_referencing_filters" json:"num_referencing_filters,computed"`
	Search                types.String                                         `tfsdk:"search" json:"search,optional"`
	Items                 customfield.NestedObjectSet[ListItemDataSourceModel] `tfsdk:"items" json:"items,computed"`
}

func (m *ListDataSourceModel) toReadParams(_ context.Context) (params rules.ListGetParams, diags diag.Diagnostics) {
	params = rules.ListGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ListItemDataSourceModel struct {
	ASN      types.Int64                                               `tfsdk:"asn" json:"asn,computed"`
	Comment  types.String                                              `tfsdk:"comment" json:"comment,computed"`
	IP       types.String                                              `tfsdk:"ip" json:"ip,computed"`
	Hostname customfield.NestedObject[ListItemHostnameDataSourceModel] `tfsdk:"hostname" json:"hostname,computed"`
	Redirect customfield.NestedObject[ListItemRedirectDataSourceModel] `tfsdk:"redirect" json:"redirect,computed"`
}

func (m ListItemDataSourceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

type ListItemHostnameDataSourceModel struct {
	URLHostname          types.String `tfsdk:"url_hostname" json:"url_hostname,computed"`
	ExcludeExactHostname types.Bool   `tfsdk:"exclude_exact_hostname" json:"exclude_exact_hostname,computed"`
}

type ListItemRedirectDataSourceModel struct {
	SourceURL           types.String `tfsdk:"source_url" json:"source_url,computed"`
	TargetURL           types.String `tfsdk:"target_url" json:"target_url,computed"`
	IncludeSubdomains   types.Bool   `tfsdk:"include_subdomains" json:"include_subdomains,computed"`
	PreservePathSuffix  types.Bool   `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix,computed"`
	PreserveQueryString types.Bool   `tfsdk:"preserve_query_string" json:"preserve_query_string,computed"`
	StatusCode          types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
	SubpathMatching     types.Bool   `tfsdk:"subpath_matching" json:"subpath_matching,computed"`
}
