package log_explorer_dataset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogExplorerDatasetResultEnvelope struct {
	Result LogExplorerDatasetModel `json:"result"`
}

type LogExplorerDatasetModel struct {
	ID                  types.String                                                `tfsdk:"id" json:"-,computed"`
	DatasetID           types.String                                                `tfsdk:"dataset_id" json:"dataset_id,computed"`
	AccountID           types.String                                                `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID              types.String                                                `tfsdk:"zone_id" path:"zone_id,optional"`
	Dataset             types.String                                                `tfsdk:"dataset" json:"dataset,required"`
	DeletionProtection  types.Bool                                                  `tfsdk:"deletion_protection" json:"deletion_protection,computed_optional"`
	Enabled             types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Filter              types.String                                                `tfsdk:"filter" json:"filter,computed_optional"`
	FilterAttackTraffic types.Bool                                                  `tfsdk:"filter_attack_traffic" json:"filter_attack_traffic,computed_optional"`
	Fields              customfield.NestedObjectList[LogExplorerDatasetFieldsModel] `tfsdk:"fields" json:"fields,computed_optional"`
	CreatedAt           timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ObjectID            types.String                                                `tfsdk:"object_id" json:"object_id,computed"`
	ObjectType          types.String                                                `tfsdk:"object_type" json:"object_type,computed"`
	UpdatedAt           timetypes.RFC3339                                           `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type LogExplorerDatasetFieldsModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	Name    types.String `tfsdk:"name" json:"name,required"`
}
