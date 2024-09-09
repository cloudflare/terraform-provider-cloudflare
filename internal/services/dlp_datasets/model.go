// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_datasets

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLPDatasetsResultEnvelope struct {
	Result DLPDatasetsModel `json:"result"`
}

type DLPDatasetsModel struct {
	AccountID       types.String                                          `tfsdk:"account_id" path:"account_id,required"`
	DatasetID       types.String                                          `tfsdk:"dataset_id" path:"dataset_id,optional"`
	EncodingVersion types.Int64                                           `tfsdk:"encoding_version" json:"encoding_version,computed_optional"`
	Secret          types.Bool                                            `tfsdk:"secret" json:"secret,computed_optional"`
	Name            types.String                                          `tfsdk:"name" json:"name,required"`
	Description     types.String                                          `tfsdk:"description" json:"description,computed_optional"`
	CreatedAt       timetypes.RFC3339                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID              types.String                                          `tfsdk:"id" json:"id,computed"`
	MaxCells        types.Int64                                           `tfsdk:"max_cells" json:"max_cells,computed"`
	NumCells        types.Int64                                           `tfsdk:"num_cells" json:"num_cells,computed"`
	Status          types.String                                          `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                     `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version         types.Int64                                           `tfsdk:"version" json:"version,computed"`
	Columns         customfield.NestedObjectList[DLPDatasetsColumnsModel] `tfsdk:"columns" json:"columns,computed"`
	Dataset         customfield.NestedObject[DLPDatasetsDatasetModel]     `tfsdk:"dataset" json:"dataset,computed"`
	Uploads         customfield.NestedObjectList[DLPDatasetsUploadsModel] `tfsdk:"uploads" json:"uploads,computed"`
}

type DLPDatasetsColumnsModel struct {
	EntryID      types.String `tfsdk:"entry_id" json:"entry_id,computed"`
	HeaderName   types.String `tfsdk:"header_name" json:"header_name,computed"`
	NumCells     types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	UploadStatus types.String `tfsdk:"upload_status" json:"upload_status,computed"`
}

type DLPDatasetsDatasetModel struct {
	ID              types.String                                                 `tfsdk:"id" json:"id,computed"`
	Columns         customfield.NestedObjectList[DLPDatasetsDatasetColumnsModel] `tfsdk:"columns" json:"columns,computed"`
	CreatedAt       timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	EncodingVersion types.Int64                                                  `tfsdk:"encoding_version" json:"encoding_version,computed"`
	Name            types.String                                                 `tfsdk:"name" json:"name,computed"`
	NumCells        types.Int64                                                  `tfsdk:"num_cells" json:"num_cells,computed"`
	Secret          types.Bool                                                   `tfsdk:"secret" json:"secret,computed"`
	Status          types.String                                                 `tfsdk:"status" json:"status,computed"`
	UpdatedAt       timetypes.RFC3339                                            `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Uploads         customfield.NestedObjectList[DLPDatasetsDatasetUploadsModel] `tfsdk:"uploads" json:"uploads,computed"`
	Description     types.String                                                 `tfsdk:"description" json:"description,computed"`
}

type DLPDatasetsDatasetColumnsModel struct {
	EntryID      types.String `tfsdk:"entry_id" json:"entry_id,computed"`
	HeaderName   types.String `tfsdk:"header_name" json:"header_name,computed"`
	NumCells     types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	UploadStatus types.String `tfsdk:"upload_status" json:"upload_status,computed"`
}

type DLPDatasetsDatasetUploadsModel struct {
	NumCells types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	Status   types.String `tfsdk:"status" json:"status,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed"`
}

type DLPDatasetsUploadsModel struct {
	NumCells types.Int64  `tfsdk:"num_cells" json:"num_cells,computed"`
	Status   types.String `tfsdk:"status" json:"status,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed"`
}
