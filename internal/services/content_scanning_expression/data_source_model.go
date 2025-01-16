// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning_expression

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/content_scanning"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ContentScanningExpressionResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ContentScanningExpressionDataSourceModel] `json:"result,computed"`
}

type ContentScanningExpressionDataSourceModel struct {
	ID      types.String                                       `tfsdk:"id" json:"id,computed"`
	Payload types.String                                       `tfsdk:"payload" json:"payload,computed"`
	Filter  *ContentScanningExpressionFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ContentScanningExpressionDataSourceModel) toListParams(_ context.Context) (params content_scanning.PayloadListParams, diags diag.Diagnostics) {
	params = content_scanning.PayloadListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type ContentScanningExpressionFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}
