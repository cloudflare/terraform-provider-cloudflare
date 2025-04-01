// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_sippy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketSippyResultDataSourceEnvelope struct {
	Result R2BucketSippyDataSourceModel `json:"result,computed"`
}

type R2BucketSippyDataSourceModel struct {
	AccountID   types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	BucketName  types.String                                                      `tfsdk:"bucket_name" path:"bucket_name,required"`
	Enabled     types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed"`
	Destination customfield.NestedObject[R2BucketSippyDestinationDataSourceModel] `tfsdk:"destination" json:"destination,computed"`
	Source      customfield.NestedObject[R2BucketSippySourceDataSourceModel]      `tfsdk:"source" json:"source,computed"`
}

func (m *R2BucketSippyDataSourceModel) toReadParams(_ context.Context) (params r2.BucketSippyGetParams, diags diag.Diagnostics) {
	params = r2.BucketSippyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2BucketSippyDestinationDataSourceModel struct {
	AccessKeyID types.String         `tfsdk:"access_key_id" json:"accessKeyId,computed"`
	Account     types.String         `tfsdk:"account" json:"account,computed"`
	Bucket      types.String         `tfsdk:"bucket" json:"bucket,computed"`
	Provider    jsontypes.Normalized `tfsdk:"provider" json:"provider,computed"`
}

type R2BucketSippySourceDataSourceModel struct {
	Bucket   types.String `tfsdk:"bucket" json:"bucket,computed"`
	Provider types.String `tfsdk:"provider" json:"provider,computed"`
	Region   types.String `tfsdk:"region" json:"region,computed"`
}
