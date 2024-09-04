// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultDataSourceEnvelope struct {
	Result ListDataSourceModel `json:"result,computed"`
}

type ListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ListDataSourceModel] `json:"result,computed"`
}

type ListDataSourceModel struct {
	AccountID             types.String                  `tfsdk:"account_id" path:"account_id,optional"`
	ListID                types.String                  `tfsdk:"list_id" path:"list_id,optional"`
	CreatedOn             types.String                  `tfsdk:"created_on" json:"created_on,optional"`
	Description           types.String                  `tfsdk:"description" json:"description,optional"`
	ID                    types.String                  `tfsdk:"id" json:"id,optional"`
	Kind                  types.String                  `tfsdk:"kind" json:"kind,optional"`
	ModifiedOn            types.String                  `tfsdk:"modified_on" json:"modified_on,optional"`
	Name                  types.String                  `tfsdk:"name" json:"name,optional"`
	NumItems              types.Float64                 `tfsdk:"num_items" json:"num_items,optional"`
	NumReferencingFilters types.Float64                 `tfsdk:"num_referencing_filters" json:"num_referencing_filters,optional"`
	Filter                *ListFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ListDataSourceModel) toReadParams(_ context.Context) (params rules.ListGetParams, diags diag.Diagnostics) {
	params = rules.ListGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ListDataSourceModel) toListParams(_ context.Context) (params rules.ListListParams, diags diag.Diagnostics) {
	params = rules.ListListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ListFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
