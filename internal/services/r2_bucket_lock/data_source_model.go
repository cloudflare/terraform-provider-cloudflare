// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lock

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketLockResultDataSourceEnvelope struct {
	Result R2BucketLockDataSourceModel `json:"result,computed"`
}

type R2BucketLockDataSourceModel struct {
	AccountID  types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	BucketName types.String                                                   `tfsdk:"bucket_name" path:"bucket_name,required"`
	Rules      customfield.NestedObjectList[R2BucketLockRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *R2BucketLockDataSourceModel) toReadParams(_ context.Context) (params r2.BucketLockGetParams, diags diag.Diagnostics) {
	params = r2.BucketLockGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2BucketLockRulesDataSourceModel struct {
	ID        types.String                                                        `tfsdk:"id" json:"id,computed"`
	Condition customfield.NestedObject[R2BucketLockRulesConditionDataSourceModel] `tfsdk:"condition" json:"condition,computed"`
	Enabled   types.Bool                                                          `tfsdk:"enabled" json:"enabled,computed"`
	Prefix    types.String                                                        `tfsdk:"prefix" json:"prefix,computed"`
}

type R2BucketLockRulesConditionDataSourceModel struct {
	MaxAgeSeconds types.Int64       `tfsdk:"max_age_seconds" json:"maxAgeSeconds,computed"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	Date          timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
}
