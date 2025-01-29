// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/images"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImageVariantResultDataSourceEnvelope struct {
	Result ImageVariantDataSourceModel `json:"result,computed"`
}

type ImageVariantDataSourceModel struct {
	AccountID types.String                                                 `tfsdk:"account_id" path:"account_id,required"`
	VariantID types.String                                                 `tfsdk:"variant_id" path:"variant_id,required"`
	Variant   customfield.NestedObject[ImageVariantVariantDataSourceModel] `tfsdk:"variant" json:"variant,computed"`
}

func (m *ImageVariantDataSourceModel) toReadParams(_ context.Context) (params images.V1VariantGetParams, diags diag.Diagnostics) {
	params = images.V1VariantGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ImageVariantVariantDataSourceModel struct {
	ID                     types.String                                                        `tfsdk:"id" json:"id,computed"`
	Options                customfield.NestedObject[ImageVariantVariantOptionsDataSourceModel] `tfsdk:"options" json:"options,computed"`
	NeverRequireSignedURLs types.Bool                                                          `tfsdk:"never_require_signed_urls" json:"neverRequireSignedURLs,computed"`
}

type ImageVariantVariantOptionsDataSourceModel struct {
	Fit      types.String  `tfsdk:"fit" json:"fit,computed"`
	Height   types.Float64 `tfsdk:"height" json:"height,computed"`
	Metadata types.String  `tfsdk:"metadata" json:"metadata,computed"`
	Width    types.Float64 `tfsdk:"width" json:"width,computed"`
}
