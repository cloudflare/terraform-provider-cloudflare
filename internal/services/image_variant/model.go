// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImageVariantResultEnvelope struct {
	Result ImageVariantModel `json:"result"`
}

type ImageVariantModel struct {
	ID                     types.String                                       `tfsdk:"id" json:"id,required"`
	AccountID              types.String                                       `tfsdk:"account_id" path:"account_id,required"`
	Options                *ImageVariantOptionsModel                          `tfsdk:"options" json:"options,required,no_refresh"`
	NeverRequireSignedURLs types.Bool                                         `tfsdk:"never_require_signed_urls" json:"neverRequireSignedURLs,computed_optional,no_refresh"`
	Variant                customfield.NestedObject[ImageVariantVariantModel] `tfsdk:"variant" json:"variant,computed"`
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

type ImageVariantVariantModel struct {
	ID                     types.String                                              `tfsdk:"id" json:"id,computed"`
	Options                customfield.NestedObject[ImageVariantVariantOptionsModel] `tfsdk:"options" json:"options,computed"`
	NeverRequireSignedURLs types.Bool                                                `tfsdk:"never_require_signed_urls" json:"neverRequireSignedURLs,computed"`
}

type ImageVariantVariantOptionsModel struct {
	Fit      types.String  `tfsdk:"fit" json:"fit,computed"`
	Height   types.Float64 `tfsdk:"height" json:"height,computed"`
	Metadata types.String  `tfsdk:"metadata" json:"metadata,computed"`
	Width    types.Float64 `tfsdk:"width" json:"width,computed"`
}
