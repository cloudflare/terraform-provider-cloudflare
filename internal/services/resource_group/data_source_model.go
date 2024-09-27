// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceGroupResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ResourceGroupDataSourceModel] `json:"result,computed"`
}

type ResourceGroupDataSourceModel struct {
	AccountID       types.String                           `tfsdk:"account_id" path:"account_id,optional"`
	ResourceGroupID types.String                           `tfsdk:"resource_group_id" path:"resource_group_id,optional"`
	ID              types.String                           `tfsdk:"id" json:"id,optional"`
	Name            types.String                           `tfsdk:"name" json:"name,optional"`
	Meta            *ResourceGroupMetaDataSourceModel      `tfsdk:"meta" json:"meta,optional"`
	Scope           *[]*ResourceGroupScopeDataSourceModel  `tfsdk:"scope" json:"scope,optional"`
	Filter          *ResourceGroupFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ResourceGroupDataSourceModel) toReadParams(_ context.Context) (params iam.ResourceGroupGetParams, diags diag.Diagnostics) {
	params = iam.ResourceGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ResourceGroupDataSourceModel) toListParams(_ context.Context) (params iam.ResourceGroupListParams, diags diag.Diagnostics) {
	params = iam.ResourceGroupListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.ID.IsNull() {
		params.ID = cloudflare.F(m.Filter.ID.ValueString())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}

	return
}

type ResourceGroupMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type ResourceGroupScopeDataSourceModel struct {
	Key     types.String                                                           `tfsdk:"key" json:"key,computed"`
	Objects customfield.NestedObjectList[ResourceGroupScopeObjectsDataSourceModel] `tfsdk:"objects" json:"objects,computed"`
}

type ResourceGroupScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type ResourceGroupFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ID        types.String `tfsdk:"id" query:"id,optional"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
}
