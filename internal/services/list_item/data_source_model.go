// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListItemResultDataSourceEnvelope struct {
	Result ListItemDataSourceModel `json:"result,computed"`
}

type ListItemResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ListItemDataSourceModel] `json:"result,computed"`
}

type ListItemDataSourceModel struct {
	AccountIdentifier   types.String                      `tfsdk:"account_identifier" path:"account_identifier,optional"`
	ItemID              types.String                      `tfsdk:"item_id" path:"item_id,optional"`
	ListID              types.String                      `tfsdk:"list_id" path:"list_id,optional"`
	IncludeSubdomains   types.Bool                        `tfsdk:"include_subdomains" json:"include_subdomains,computed"`
	PreservePathSuffix  types.Bool                        `tfsdk:"preserve_path_suffix" json:"preserve_path_suffix,computed"`
	PreserveQueryString types.Bool                        `tfsdk:"preserve_query_string" json:"preserve_query_string,computed"`
	SourceURL           types.String                      `tfsdk:"source_url" json:"source_url,computed"`
	StatusCode          types.Int64                       `tfsdk:"status_code" json:"status_code,computed"`
	SubpathMatching     types.Bool                        `tfsdk:"subpath_matching" json:"subpath_matching,computed"`
	TargetURL           types.String                      `tfsdk:"target_url" json:"target_url,computed"`
	URLHostname         types.String                      `tfsdk:"url_hostname" json:"url_hostname,computed"`
	Filter              *ListItemFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ListItemDataSourceModel) toListParams(_ context.Context) (params rules.ListItemListParams, diags diag.Diagnostics) {
	params = rules.ListItemListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type ListItemFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ListID    types.String `tfsdk:"list_id" path:"list_id,required"`
	Search    types.String `tfsdk:"search" query:"search,optional"`
}
