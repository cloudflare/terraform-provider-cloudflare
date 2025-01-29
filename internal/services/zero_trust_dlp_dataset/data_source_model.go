// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_dataset

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDatasetResultDataSourceEnvelope struct {
	Result ZeroTrustDLPDatasetDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPDatasetDataSourceModel struct {
	AccountID       types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	DatasetID       types.String                                                            `tfsdk:"dataset_id" path:"dataset_id,required"`
	CreatedAt       timetypes.RFC3339                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description     types.String                                                            `tfsdk:"description" json:"description,computed"`
	EncodingVersion types.Int64                                                             `tfsdk:"encoding_version" json:"encoding_version,computed"`
	ID              types.String                                                            `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                            `tfsdk:"name" json:"name,computed"`
	NumCells        types.Int64                                                             `tfsdk:"num_cells" json:"num_cells,computed"`
	Secret          types.Bool                                                              `tfsdk:"secret" json:"secret,computed"`
	Status          types.String                                                            `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Columns         customfield.NestedObjectList[ZeroTrustDLPDatasetColumnsDataSourceModel] `tfsdk:"columns" json:"columns,computed"`
	Uploads         customfield.NestedObjectList[ZeroTrustDLPDatasetUploadsDataSourceModel] `tfsdk:"uploads" json:"uploads,computed"`
}

func (m *ZeroTrustDLPDatasetDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPDatasetGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPDatasetGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPDatasetColumnsDataSourceModel struct {
	EntryID      types.String `tfsdk:"entry_id" json:"entry_id,computed"`
	HeaderName   types.String `tfsdk:"header_name" json:"header_name,computed"`
	NumCells     types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	UploadStatus types.String `tfsdk:"upload_status" json:"upload_status,computed"`
}

type ZeroTrustDLPDatasetUploadsDataSourceModel struct {
	NumCells types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	Status   types.String `tfsdk:"status" json:"status,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed"`
}
