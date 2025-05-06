// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestAssetResultDataSourceEnvelope struct {
	Result CloudforceOneRequestAssetDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestAssetDataSourceModel struct {
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	AssetID     types.String      `tfsdk:"asset_id" path:"asset_id,required"`
	RequestID   types.String      `tfsdk:"request_id" path:"request_id,required"`
	Created     timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	FileType    types.String      `tfsdk:"file_type" json:"file_type,computed"`
	ID          types.Int64       `tfsdk:"id" json:"id,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
}

func (m *CloudforceOneRequestAssetDataSourceModel) toReadParams(_ context.Context) (params cloudforce_one.RequestAssetGetParams, diags diag.Diagnostics) {
	params = cloudforce_one.RequestAssetGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
