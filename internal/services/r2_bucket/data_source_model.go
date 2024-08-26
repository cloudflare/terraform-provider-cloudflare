// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultDataSourceEnvelope struct {
	Result R2BucketDataSourceModel `json:"result,computed"`
}

type R2BucketResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[R2BucketDataSourceModel] `json:"result,computed"`
}

type R2BucketDataSourceModel struct {
	AccountID    types.String                      `tfsdk:"account_id" path:"account_id"`
	BucketName   types.String                      `tfsdk:"bucket_name" path:"bucket_name"`
	StorageClass types.String                      `tfsdk:"storage_class" json:"storage_class,computed"`
	CreationDate types.String                      `tfsdk:"creation_date" json:"creation_date"`
	Location     types.String                      `tfsdk:"location" json:"location"`
	Name         types.String                      `tfsdk:"name" json:"name"`
	Filter       *R2BucketFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *R2BucketDataSourceModel) toReadParams() (params r2.BucketGetParams, diags diag.Diagnostics) {
	params = r2.BucketGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *R2BucketDataSourceModel) toListParams() (params r2.BucketListParams, diags diag.Diagnostics) {
	params = r2.BucketListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(r2.BucketListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.NameContains.IsNull() {
		params.NameContains = cloudflare.F(m.Filter.NameContains.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(r2.BucketListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.StartAfter.IsNull() {
		params.StartAfter = cloudflare.F(m.Filter.StartAfter.ValueString())
	}

	return
}

type R2BucketFindOneByDataSourceModel struct {
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	Direction    types.String `tfsdk:"direction" query:"direction"`
	NameContains types.String `tfsdk:"name_contains" query:"name_contains"`
	Order        types.String `tfsdk:"order" query:"order"`
	StartAfter   types.String `tfsdk:"start_after" query:"start_after"`
}
