// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImageVariantWriteResultEnvelope struct {
	Result ImageVariantWriteResult `json:"result"`
}

type ImageVariantWriteResult struct {
	Variant ImageVariantModel `json:"variant"`
}

type ImageVariantModel struct {
	ID                     types.String              `tfsdk:"id" json:"id,required"`
	AccountID              types.String              `tfsdk:"account_id" path:"account_id,required"`
	Options                *ImageVariantOptionsModel `tfsdk:"options" json:"options,required"`
	NeverRequireSignedURLs types.Bool                `tfsdk:"never_require_signed_urls" json:"neverRequireSignedURLs,computed_optional"`
}

func (m ImageVariantModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ImageVariantModel) MarshalJSONForUpdate(state ImageVariantModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ImageVariantOptionsModel struct {
	Fit      types.String  `tfsdk:"fit" json:"fit,required"`
	Height   types.Float64 `tfsdk:"height" json:"height,required"`
	Metadata types.String  `tfsdk:"metadata" json:"metadata,required"`
	Width    types.Float64 `tfsdk:"width" json:"width,required"`
}
