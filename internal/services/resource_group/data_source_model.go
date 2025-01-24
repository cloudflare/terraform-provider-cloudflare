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

type ResourceGroupDataSourceModel struct {
	AccountID       types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	ResourceGroupID types.String                                                    `tfsdk:"resource_group_id" path:"resource_group_id,required"`
	ID              types.String                                                    `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                    `tfsdk:"name" json:"name,computed"`
	Meta            customfield.NestedObject[ResourceGroupMetaDataSourceModel]      `tfsdk:"meta" json:"meta,computed"`
	Scope           customfield.NestedObjectList[ResourceGroupScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
}

func (m *ResourceGroupDataSourceModel) toReadParams(_ context.Context) (params iam.ResourceGroupGetParams, diags diag.Diagnostics) {
	params = iam.ResourceGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
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
