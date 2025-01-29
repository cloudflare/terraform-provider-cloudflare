// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_managed_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2ManagedDomainResultEnvelope struct {
	Result R2ManagedDomainModel `json:"result"`
}

type R2ManagedDomainModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	BucketName types.String `tfsdk:"bucket_name" path:"bucket_name,required"`
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	BucketID   types.String `tfsdk:"bucket_id" json:"bucketId,computed"`
	Domain     types.String `tfsdk:"domain" json:"domain,computed"`
}

func (m R2ManagedDomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2ManagedDomainModel) MarshalJSONForUpdate(state R2ManagedDomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
