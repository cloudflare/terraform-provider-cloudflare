// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketsResultListDataSourceEnvelope struct {
	Result *[]*R2BucketsResultDataSourceModel `json:"result,computed"`
}

type R2BucketsDataSourceModel struct {
	AccountID    types.String                       `tfsdk:"account_id" path:"account_id"`
	Cursor       types.String                       `tfsdk:"cursor" query:"cursor"`
	Direction    types.String                       `tfsdk:"direction" query:"direction"`
	NameContains types.String                       `tfsdk:"name_contains" query:"name_contains"`
	Order        types.String                       `tfsdk:"order" query:"order"`
	PerPage      types.Float64                      `tfsdk:"per_page" query:"per_page"`
	StartAfter   types.String                       `tfsdk:"start_after" query:"start_after"`
	MaxItems     types.Int64                        `tfsdk:"max_items"`
	Result       *[]*R2BucketsResultDataSourceModel `tfsdk:"result"`
}

type R2BucketsResultDataSourceModel struct {
	CreationDate types.String `tfsdk:"creation_date" json:"creation_date"`
	Location     types.String `tfsdk:"location" json:"location"`
	Name         types.String `tfsdk:"name" json:"name"`
	StorageClass types.String `tfsdk:"storage_class" json:"storage_class,computed"`
}
