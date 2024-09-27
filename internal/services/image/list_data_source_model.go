// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/images"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImagesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ImagesResultDataSourceModel] `json:"items,computed"`
}

type ImagesDataSourceModel struct {
	AccountID types.String                                              `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                               `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ImagesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ImagesDataSourceModel) toListParams(_ context.Context) (params images.V1ListParams, diags diag.Diagnostics) {
	params = images.V1ListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ImagesResultDataSourceModel struct {
	Images customfield.NestedObjectList[ImagesResultImagesDataSourceModel] `tfsdk:"images" json:"images,computed"`
}

type ImagesErrorsDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type ImagesMessagesDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type ImagesResultImagesDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	Filename          types.String                   `tfsdk:"filename" json:"filename,computed"`
	Meta              jsontypes.Normalized           `tfsdk:"meta" json:"meta,computed"`
	RequireSignedURLs types.Bool                     `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	Uploaded          timetypes.RFC3339              `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	Variants          customfield.List[types.String] `tfsdk:"variants" json:"variants,computed"`
}
