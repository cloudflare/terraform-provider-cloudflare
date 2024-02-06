package images_variant

import "github.com/hashicorp/terraform-plugin-framework/types"

type ImagesVariantModel struct {
	AccountID              types.String                 `tfsdk:"account_id"`
	ID                     types.String                 `tfsdk:"id"`
	NeverRequireSignedUrls types.Bool                   `tfsdk:"never_require_signed_urls"`
	Options                []*ImagesVariantOptionsModel `tfsdk:"options"`
}

type ImagesVariantOptionsModel struct {
	Metadata types.String `tfsdk:"metadata"`
	Fit      types.String `tfsdk:"fit"`
	Height   types.Int64  `tfsdk:"height"`
	Width    types.Int64  `tfsdk:"width"`
}
