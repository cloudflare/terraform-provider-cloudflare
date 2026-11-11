package log_explorer_dataset

import (
	"context"
	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/logs"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogExplorerDatasetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LogExplorerDatasetsResultDataSourceModel] `json:"result,computed"`
}

type LogExplorerDatasetsDataSourceModel struct {
	AccountID    types.String                                                           `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID       types.String                                                           `tfsdk:"zone_id" path:"zone_id,optional"`
	IncludeZones types.Bool                                                             `tfsdk:"include_zones" query:"include_zones,optional"`
	MaxItems     types.Int64                                                            `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[LogExplorerDatasetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *LogExplorerDatasetsDataSourceModel) toListParams(_ context.Context) (params logs.LogExplorerDatasetListParams, diags diag.Diagnostics) {
	params = logs.LogExplorerDatasetListParams{}

	if !m.IncludeZones.IsNull() {
		params.IncludeZones = cloudflare.F(m.IncludeZones.ValueBool())
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type LogExplorerDatasetsResultDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" json:"dataset_id,computed"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Dataset            types.String      `tfsdk:"dataset" json:"dataset,computed"`
	DatasetID          types.String      `tfsdk:"dataset_id" json:"dataset_id,computed"`
	DeletionProtection types.Bool        `tfsdk:"deletion_protection" json:"deletion_protection,computed"`
	Enabled            types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ObjectID           types.String      `tfsdk:"object_id" json:"object_id,computed"`
	ObjectType         types.String      `tfsdk:"object_type" json:"object_type,computed"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
