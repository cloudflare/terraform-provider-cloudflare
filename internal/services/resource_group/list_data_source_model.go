// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ResourceGroupsResultDataSourceModel] `json:"result,computed"`
}

type ResourceGroupsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	ID        types.String                                                      `tfsdk:"id" query:"id,optional"`
	Name      types.String                                                      `tfsdk:"name" query:"name,optional"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ResourceGroupsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ResourceGroupsDataSourceModel) toListParams(_ context.Context) (params iam.ResourceGroupListParams, diags diag.Diagnostics) {
	params = iam.ResourceGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}

	return
}

type ResourceGroupsResultDataSourceModel struct {
	ID    types.String                                                     `tfsdk:"id" json:"id,computed"`
	Scope customfield.NestedObjectList[ResourceGroupsScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Meta  customfield.NestedObject[ResourceGroupsMetaDataSourceModel]      `tfsdk:"meta" json:"meta,computed"`
	Name  types.String                                                     `tfsdk:"name" json:"name,computed"`
}

type ResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                            `tfsdk:"key" json:"key,computed"`
	Objects customfield.NestedObjectList[ResourceGroupsScopeObjectsDataSourceModel] `tfsdk:"objects" json:"objects,computed"`
}

type ResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type ResourceGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
