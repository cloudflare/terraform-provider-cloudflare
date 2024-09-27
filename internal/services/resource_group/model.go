// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceGroupResultEnvelope struct {
	Result ResourceGroupModel `json:"result"`
}

type ResourceGroupModel struct {
	ID        types.String             `tfsdk:"id" json:"id,computed"`
	AccountID types.String             `tfsdk:"account_id" path:"account_id,required"`
	Scope     *ResourceGroupScopeModel `tfsdk:"scope" json:"scope,required"`
	Meta      jsontypes.Normalized     `tfsdk:"meta" json:"meta,optional"`
	Name      types.String             `tfsdk:"name" json:"name,computed"`
}

type ResourceGroupScopeModel struct {
	Key     types.String                       `tfsdk:"key" json:"key,required"`
	Objects *[]*ResourceGroupScopeObjectsModel `tfsdk:"objects" json:"objects,required"`
}

type ResourceGroupScopeObjectsModel struct {
	Key types.String `tfsdk:"key" json:"key,required"`
}
