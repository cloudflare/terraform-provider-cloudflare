// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[R2BucketsResultDataSourceModel] `json:"result,computed"`
}

type R2BucketsDataSourceModel struct {
	AccountID    types.String                                                 `tfsdk:"account_id" path:"account_id"`
	Direction    types.String                                                 `tfsdk:"direction" query:"direction"`
	NameContains types.String                                                 `tfsdk:"name_contains" query:"name_contains"`
	Order        types.String                                                 `tfsdk:"order" query:"order"`
	StartAfter   types.String                                                 `tfsdk:"start_after" query:"start_after"`
	MaxItems     types.Int64                                                  `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[R2BucketsResultDataSourceModel] `tfsdk:"result"`
}

func (m *R2BucketsDataSourceModel) toListParams() (params r2.BucketListParams, diags diag.Diagnostics) {
	params = r2.BucketListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(r2.BucketListParamsDirection(m.Direction.ValueString()))
	}
	if !m.NameContains.IsNull() {
		params.NameContains = cloudflare.F(m.NameContains.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(r2.BucketListParamsOrder(m.Order.ValueString()))
	}
	if !m.StartAfter.IsNull() {
		params.StartAfter = cloudflare.F(m.StartAfter.ValueString())
	}

	return
}

type R2BucketsResultDataSourceModel struct {
	CreationDate types.String `tfsdk:"creation_date" json:"creation_date"`
	Location     types.String `tfsdk:"location" json:"location"`
	Name         types.String `tfsdk:"name" json:"name"`
	StorageClass types.String `tfsdk:"storage_class" json:"storage_class,computed"`
}
