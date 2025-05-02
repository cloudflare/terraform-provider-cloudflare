// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_sippy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketSippyResultEnvelope struct {
	Result R2BucketSippyModel `json:"result"`
}

type R2BucketSippyModel struct {
	AccountID    types.String                   `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String                   `tfsdk:"bucket_name" path:"bucket_name,required"`
	Jurisdiction types.String                   `tfsdk:"jurisdiction" json:"-,computed_optional"`
	Destination  *R2BucketSippyDestinationModel `tfsdk:"destination" json:"destination,optional"`
	Source       *R2BucketSippySourceModel      `tfsdk:"source" json:"source,optional"`
	Enabled      types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
}

func (m R2BucketSippyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketSippyModel) MarshalJSONForUpdate(state R2BucketSippyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketSippyDestinationModel struct {
	AccessKeyID     types.String `tfsdk:"access_key_id" json:"accessKeyId,optional"`
	CloudProvider   types.String `tfsdk:"cloud_provider" json:"provider,optional"`
	SecretAccessKey types.String `tfsdk:"secret_access_key" json:"secretAccessKey,optional"`
}

type R2BucketSippySourceModel struct {
	AccessKeyID     types.String `tfsdk:"access_key_id" json:"accessKeyId,optional"`
	Bucket          types.String `tfsdk:"bucket" json:"bucket,optional"`
	CloudProvider   types.String `tfsdk:"cloud_provider" json:"provider,optional"`
	Region          types.String `tfsdk:"region" json:"region,optional"`
	SecretAccessKey types.String `tfsdk:"secret_access_key" json:"secretAccessKey,optional"`
	ClientEmail     types.String `tfsdk:"client_email" json:"clientEmail,optional"`
	PrivateKey      types.String `tfsdk:"private_key" json:"privateKey,optional"`
}
