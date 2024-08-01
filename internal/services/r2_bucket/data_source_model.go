// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultDataSourceEnvelope struct {
	Result R2BucketDataSourceModel `json:"result,computed"`
}

type R2BucketResultListDataSourceEnvelope struct {
	Result *[]*R2BucketDataSourceModel `json:"result,computed"`
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

type R2BucketFindOneByDataSourceModel struct {
	AccountID    types.String  `tfsdk:"account_id" path:"account_id"`
	Cursor       types.String  `tfsdk:"cursor" query:"cursor"`
	Direction    types.String  `tfsdk:"direction" query:"direction"`
	NameContains types.String  `tfsdk:"name_contains" query:"name_contains"`
	Order        types.String  `tfsdk:"order" query:"order"`
	PerPage      types.Float64 `tfsdk:"per_page" query:"per_page"`
	StartAfter   types.String  `tfsdk:"start_after" query:"start_after"`
}
