// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/r2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultDataSourceEnvelope struct {
	Result R2BucketDataSourceModel `json:"result,computed"`
}

type R2BucketDataSourceModel struct {
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	BucketName   types.String `tfsdk:"bucket_name" path:"bucket_name"`
	CreationDate types.String `tfsdk:"creation_date" json:"creation_date"`
	Location     types.String `tfsdk:"location" json:"location"`
	Name         types.String `tfsdk:"name" json:"name"`
	StorageClass types.String `tfsdk:"storage_class" json:"storage_class,computed_optional"`
}

func (m *R2BucketDataSourceModel) toReadParams() (params r2.BucketGetParams, diags diag.Diagnostics) {
	params = r2.BucketGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
