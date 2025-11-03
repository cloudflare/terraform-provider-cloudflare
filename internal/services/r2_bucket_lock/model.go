// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lock

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketLockResultEnvelope struct {
	Result R2BucketLockModel `json:"result"`
}

type R2BucketLockModel struct {
	AccountID    types.String               `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String               `tfsdk:"bucket_name" path:"bucket_name,required"`
	Jurisdiction types.String               `tfsdk:"jurisdiction" json:"-,computed_optional,no_refresh"`
	Rules        *[]*R2BucketLockRulesModel `tfsdk:"rules" json:"rules,optional"`
}

func (m R2BucketLockModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketLockModel) MarshalJSONForUpdate(state R2BucketLockModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketLockRulesModel struct {
	ID        types.String                     `tfsdk:"id" json:"id,required"`
	Condition *R2BucketLockRulesConditionModel `tfsdk:"condition" json:"condition,required"`
	Enabled   types.Bool                       `tfsdk:"enabled" json:"enabled,required"`
	Prefix    types.String                     `tfsdk:"prefix" json:"prefix,optional"`
}

type R2BucketLockRulesConditionModel struct {
	MaxAgeSeconds types.Int64       `tfsdk:"max_age_seconds" json:"maxAgeSeconds,optional"`
	Type          types.String      `tfsdk:"type" json:"type,required"`
	Date          timetypes.RFC3339 `tfsdk:"date" json:"date,optional" format:"date-time"`
}
