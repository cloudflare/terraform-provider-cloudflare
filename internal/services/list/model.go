// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ListResultEnvelope struct {
	Result ListModel `json:"result"`
}

type ListModel struct {
	ID                    types.String  `tfsdk:"id" json:"id,computed"`
	AccountID             types.String  `tfsdk:"account_id" path:"account_id,required"`
	Kind                  types.String  `tfsdk:"kind" json:"kind,required"`
	Name                  types.String  `tfsdk:"name" json:"name,required"`
	Description           types.String  `tfsdk:"description" json:"description,optional"`
	CreatedOn             types.String  `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn            types.String  `tfsdk:"modified_on" json:"modified_on,computed"`
	NumItems              types.Float64 `tfsdk:"num_items" json:"num_items,computed"`
	NumReferencingFilters types.Float64 `tfsdk:"num_referencing_filters" json:"num_referencing_filters,computed"`
}

func (m ListModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ListModel) MarshalJSONForUpdate(state ListModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
