// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_custom_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2CustomDomainResultEnvelope struct {
	Result R2CustomDomainModel `json:"result"`
}

type R2CustomDomainModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	BucketName types.String `tfsdk:"bucket_name" path:"bucket_name,required"`
	Domain     types.String `tfsdk:"domain" json:"domain,required"`
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	ZoneID     types.String `tfsdk:"zone_id" json:"zoneId,required"`
	MinTLS     types.String `tfsdk:"min_tls" json:"minTLS,optional"`
}

func (m R2CustomDomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2CustomDomainModel) MarshalJSONForUpdate(state R2CustomDomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
