// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_dataset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDatasetResultEnvelope struct {
	Result ZeroTrustDLPDatasetModel `json:"result"`
}

type ZeroTrustDLPDatasetModel struct {
	AccountID       types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	DatasetID       types.String                                                  `tfsdk:"dataset_id" path:"dataset_id,optional"`
	EncodingVersion types.Int64                                                   `tfsdk:"encoding_version" json:"encoding_version,optional"`
	Secret          types.Bool                                                    `tfsdk:"secret" json:"secret,optional"`
	Name            types.String                                                  `tfsdk:"name" json:"name,required"`
	Description     types.String                                                  `tfsdk:"description" json:"description,optional"`
	CreatedAt       timetypes.RFC3339                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID              types.String                                                  `tfsdk:"id" json:"id,computed"`
	MaxCells        types.Int64                                                   `tfsdk:"max_cells" json:"max_cells,computed"`
	NumCells        types.Int64                                                   `tfsdk:"num_cells" json:"num_cells,computed"`
	Status          types.String                                                  `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version         types.Int64                                                   `tfsdk:"version" json:"version,computed"`
	Columns         customfield.NestedObjectList[ZeroTrustDLPDatasetColumnsModel] `tfsdk:"columns" json:"columns,computed"`
	Dataset         customfield.NestedObject[ZeroTrustDLPDatasetDatasetModel]     `tfsdk:"dataset" json:"dataset,computed"`
	Uploads         customfield.NestedObjectList[ZeroTrustDLPDatasetUploadsModel] `tfsdk:"uploads" json:"uploads,computed"`
}

func (m ZeroTrustDLPDatasetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPDatasetModel) MarshalJSONForUpdate(state ZeroTrustDLPDatasetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPDatasetColumnsModel struct {
	EntryID      types.String `tfsdk:"entry_id" json:"entry_id,computed"`
	HeaderName   types.String `tfsdk:"header_name" json:"header_name,computed"`
	NumCells     types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	UploadStatus types.String `tfsdk:"upload_status" json:"upload_status,computed"`
}

type ZeroTrustDLPDatasetDatasetModel struct {
	ID              types.String                                                         `tfsdk:"id" json:"id,computed"`
	Columns         customfield.NestedObjectList[ZeroTrustDLPDatasetDatasetColumnsModel] `tfsdk:"columns" json:"columns,computed"`
	CreatedAt       timetypes.RFC3339                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	EncodingVersion types.Int64                                                          `tfsdk:"encoding_version" json:"encoding_version,computed"`
	Name            types.String                                                         `tfsdk:"name" json:"name,computed"`
	NumCells        types.Int64                                                          `tfsdk:"num_cells" json:"num_cells,computed"`
	Secret          types.Bool                                                           `tfsdk:"secret" json:"secret,computed"`
	Status          types.String                                                         `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Uploads         customfield.NestedObjectList[ZeroTrustDLPDatasetDatasetUploadsModel] `tfsdk:"uploads" json:"uploads,computed"`
	Description     types.String                                                         `tfsdk:"description" json:"description,computed"`
}

type ZeroTrustDLPDatasetDatasetColumnsModel struct {
	EntryID      types.String `tfsdk:"entry_id" json:"entry_id,computed"`
	HeaderName   types.String `tfsdk:"header_name" json:"header_name,computed"`
	NumCells     types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	UploadStatus types.String `tfsdk:"upload_status" json:"upload_status,computed"`
}

type ZeroTrustDLPDatasetDatasetUploadsModel struct {
	NumCells types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	Status   types.String `tfsdk:"status" json:"status,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed"`
}

type ZeroTrustDLPDatasetUploadsModel struct {
	NumCells types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	Status   types.String `tfsdk:"status" json:"status,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed"`
}
