// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_cors

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketCORSResultEnvelope struct {
	Result R2BucketCORSModel `json:"result"`
}

type R2BucketCORSModel struct {
	AccountID    types.String               `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String               `tfsdk:"bucket_name" path:"bucket_name,required"`
	Jurisdiction types.String               `tfsdk:"jurisdiction" json:"-,computed_optional,no_refresh"`
	Rules        *[]*R2BucketCORSRulesModel `tfsdk:"rules" json:"rules,optional"`
}

func (m R2BucketCORSModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketCORSModel) MarshalJSONForUpdate(state R2BucketCORSModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketCORSRulesModel struct {
	Allowed       *R2BucketCORSRulesAllowedModel `tfsdk:"allowed" json:"allowed,required"`
	ID            types.String                   `tfsdk:"id" json:"id,optional"`
	ExposeHeaders *[]types.String                `tfsdk:"expose_headers" json:"exposeHeaders,optional"`
	MaxAgeSeconds types.Float64                  `tfsdk:"max_age_seconds" json:"maxAgeSeconds,optional"`
}

type R2BucketCORSRulesAllowedModel struct {
	Methods *[]types.String `tfsdk:"methods" json:"methods,required"`
	Origins *[]types.String `tfsdk:"origins" json:"origins,required"`
	Headers *[]types.String `tfsdk:"headers" json:"headers,optional"`
}
