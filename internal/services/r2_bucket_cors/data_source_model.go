// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_cors

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketCORSResultDataSourceEnvelope struct {
	Result R2BucketCORSDataSourceModel `json:"result,computed"`
}

type R2BucketCORSDataSourceModel struct {
	AccountID  types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	BucketName types.String                                                   `tfsdk:"bucket_name" path:"bucket_name,required"`
	Rules      customfield.NestedObjectList[R2BucketCORSRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *R2BucketCORSDataSourceModel) toReadParams(_ context.Context) (params r2.BucketCORSGetParams, diags diag.Diagnostics) {
	params = r2.BucketCORSGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type R2BucketCORSRulesDataSourceModel struct {
	Allowed       customfield.NestedObject[R2BucketCORSRulesAllowedDataSourceModel] `tfsdk:"allowed" json:"allowed,computed"`
	ID            types.String                                                      `tfsdk:"id" json:"id,computed"`
	ExposeHeaders customfield.List[types.String]                                    `tfsdk:"expose_headers" json:"exposeHeaders,computed"`
	MaxAgeSeconds types.Float64                                                     `tfsdk:"max_age_seconds" json:"maxAgeSeconds,computed"`
}

type R2BucketCORSRulesAllowedDataSourceModel struct {
	Methods customfield.List[types.String] `tfsdk:"methods" json:"methods,computed"`
	Origins customfield.List[types.String] `tfsdk:"origins" json:"origins,computed"`
	Headers customfield.List[types.String] `tfsdk:"headers" json:"headers,computed"`
}
