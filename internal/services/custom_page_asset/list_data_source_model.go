// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_page_asset

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPageAssetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomPageAssetsResultDataSourceModel] `json:"result,computed"`
}

type CustomPageAssetsDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                        `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CustomPageAssetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomPageAssetsDataSourceModel) toListParams(_ context.Context) (params custom_pages.AssetListParams, diags diag.Diagnostics) {
	params = custom_pages.AssetListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type CustomPageAssetsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"name,computed"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	SizeBytes   types.Int64       `tfsdk:"size_bytes" json:"size_bytes,computed"`
	URL         types.String      `tfsdk:"url" json:"url,computed"`
}
