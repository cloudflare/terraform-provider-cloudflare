// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultEnvelope struct {
	Result R2BucketModel `json:"result,computed"`
}

type R2BucketModel struct {
	ID           types.String `tfsdk:"id" json:"-,computed"`
	Name         types.String `tfsdk:"name" json:"name"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	Location     types.String `tfsdk:"location" json:"locationHint"`
	StorageClass types.String `tfsdk:"storage_class" json:"storageClass"`
	CreationDate types.String `tfsdk:"creation_date" json:"creation_date,computed"`
}
