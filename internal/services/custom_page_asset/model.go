// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_page_asset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPageAssetResultEnvelope struct {
	Result CustomPageAssetModel `json:"result"`
}

type CustomPageAssetModel struct {
	ID          types.String      `tfsdk:"id" json:"-,computed"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id,optional"`
	Description types.String      `tfsdk:"description" json:"description,required"`
	URL         types.String      `tfsdk:"url" json:"url,required"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	SizeBytes   types.Int64       `tfsdk:"size_bytes" json:"size_bytes,computed"`
}

func (m CustomPageAssetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomPageAssetModel) MarshalJSONForUpdate(state CustomPageAssetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
