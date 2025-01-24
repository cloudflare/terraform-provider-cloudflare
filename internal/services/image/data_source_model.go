// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/images"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImageResultDataSourceEnvelope struct {
	Result ImageDataSourceModel `json:"result,computed"`
}

type ImageDataSourceModel struct {
	AccountID         types.String                   `tfsdk:"account_id" path:"account_id,required"`
	ImageID           types.String                   `tfsdk:"image_id" path:"image_id,required"`
	Filename          types.String                   `tfsdk:"filename" json:"filename,computed"`
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	RequireSignedURLs types.Bool                     `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	Uploaded          timetypes.RFC3339              `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	Variants          customfield.List[types.String] `tfsdk:"variants" json:"variants,computed"`
	Meta              jsontypes.Normalized           `tfsdk:"meta" json:"meta,computed"`
}

func (m *ImageDataSourceModel) toReadParams(_ context.Context) (params images.V1GetParams, diags diag.Diagnostics) {
	params = images.V1GetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
