// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/r2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultDataSourceEnvelope struct {
	Result R2BucketDataSourceModel `json:"result,computed"`
}

type R2BucketDataSourceModel struct {
	AccountID    types.String `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String `tfsdk:"bucket_name" path:"bucket_name,required"`
	CreationDate types.String `tfsdk:"creation_date" json:"creation_date,optional"`
	Location     types.String `tfsdk:"location" json:"location,optional"`
	Name         types.String `tfsdk:"name" json:"name,optional"`
	StorageClass types.String `tfsdk:"storage_class" json:"storage_class,computed_optional"`
}

func (m *R2BucketDataSourceModel) toReadParams(_ context.Context) (params r2.BucketGetParams, diags diag.Diagnostics) {
	params = r2.BucketGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
