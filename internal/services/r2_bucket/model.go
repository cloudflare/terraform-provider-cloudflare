// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultEnvelope struct {
	Result R2BucketModel `json:"result"`
}

type R2BucketModel struct {
	ID           types.String `tfsdk:"id" json:"-,computed"`
	Name         types.String `tfsdk:"name" json:"name,required"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id,required"`
	Location     types.String `tfsdk:"location" json:"location,computed_optional"`
	StorageClass types.String `tfsdk:"storage_class" json:"storage_class,computed_optional"`
	CreationDate types.String `tfsdk:"creation_date" json:"creation_date,computed"`
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"jurisdiction,computed_optional,no_refresh"`
}

type CreateR2BucketModel struct {
	ID           types.String `tfsdk:"id" json:"-,computed"`
	Name         types.String `tfsdk:"name" json:"name,required"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id,required"`
	Location     types.String `tfsdk:"location" json:"locationHint,computed_optional"`
	Jurisdiction types.String `tfsdk:"jurisdiction" json:"-,computed_optional,no_refresh"`
	StorageClass types.String `tfsdk:"storage_class" json:"storageClass,computed_optional"`
	CreationDate types.String `tfsdk:"creation_date" json:"-,computed"`
}

func (m CreateR2BucketModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}
