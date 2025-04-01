// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_sippy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketSippyResultEnvelope struct {
	Result R2BucketSippyModel `json:"result"`
}

type R2BucketSippyModel struct {
	AccountID    types.String                                            `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String                                            `tfsdk:"bucket_name" path:"bucket_name,required"`
	Jurisdiction types.String                                            `tfsdk:"jurisdiction" json:"-,computed_optional"`
	Destination  customfield.NestedObject[R2BucketSippyDestinationModel] `tfsdk:"destination" json:"destination,computed_optional"`
	Source       customfield.NestedObject[R2BucketSippySourceModel]      `tfsdk:"source" json:"source,computed_optional"`
	Enabled      types.Bool                                              `tfsdk:"enabled" json:"enabled,computed"`
}

func (m R2BucketSippyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketSippyModel) MarshalJSONForUpdate(state R2BucketSippyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketSippyDestinationModel struct {
	AccessKeyID     types.String         `tfsdk:"access_key_id" json:"accessKeyId,optional"`
	Provider        jsontypes.Normalized `tfsdk:"provider" json:"provider,optional"`
	SecretAccessKey types.String         `tfsdk:"secret_access_key" json:"secretAccessKey,optional"`
}

type R2BucketSippySourceModel struct {
	AccessKeyID     types.String `tfsdk:"access_key_id" json:"accessKeyId,optional"`
	Bucket          types.String `tfsdk:"bucket" json:"bucket,optional"`
	Provider        types.String `tfsdk:"provider" json:"provider,optional"`
	Region          types.String `tfsdk:"region" json:"region,optional"`
	SecretAccessKey types.String `tfsdk:"secret_access_key" json:"secretAccessKey,optional"`
	ClientEmail     types.String `tfsdk:"client_email" json:"clientEmail,optional"`
	PrivateKey      types.String `tfsdk:"private_key" json:"privateKey,optional"`
}
