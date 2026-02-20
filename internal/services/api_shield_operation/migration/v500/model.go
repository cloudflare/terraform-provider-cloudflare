package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareAPIShieldOperationModel represents the source cloudflare_api_shield_operation state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromV4 to parse legacy state.
//
// Note: This resource was added late to v4 and has a very simple schema with no nested structures.
type SourceCloudflareAPIShieldOperationModel struct {
	ID       types.String `tfsdk:"id"`
	ZoneID   types.String `tfsdk:"zone_id"`
	Method   types.String `tfsdk:"method"`
	Host     types.String `tfsdk:"host"`
	Endpoint types.String `tfsdk:"endpoint"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAPIShieldOperationModel represents the target cloudflare_api_shield_operation state structure (v500).
// This matches the structure in the parent package's model.go file.
type TargetAPIShieldOperationModel struct {
	ID          types.String                                                          `tfsdk:"id"`
	OperationID types.String                                                          `tfsdk:"operation_id"`
	ZoneID      types.String                                                          `tfsdk:"zone_id"`
	Endpoint    types.String                                                          `tfsdk:"endpoint"`
	Host        types.String                                                          `tfsdk:"host"`
	Method      types.String                                                          `tfsdk:"method"`
	LastUpdated timetypes.RFC3339                                                     `tfsdk:"last_updated"`
	Features    customfield.NestedObject[TargetAPIShieldOperationFeaturesModel]       `tfsdk:"features"`
}

// TargetAPIShieldOperationFeaturesModel represents the features nested structure.
// All fields are computed and populated by the API.
type TargetAPIShieldOperationFeaturesModel struct {
	Thresholds          customfield.NestedObject[TargetAPIShieldOperationFeaturesThresholdsModel]          `tfsdk:"thresholds"`
	ParameterSchemas    customfield.NestedObject[TargetAPIShieldOperationFeaturesParameterSchemasModel]    `tfsdk:"parameter_schemas"`
	APIRouting          customfield.NestedObject[TargetAPIShieldOperationFeaturesAPIRoutingModel]          `tfsdk:"api_routing"`
	ConfidenceIntervals customfield.NestedObject[TargetAPIShieldOperationFeaturesConfidenceIntervalsModel] `tfsdk:"confidence_intervals"`
	SchemaInfo          customfield.NestedObject[TargetAPIShieldOperationFeaturesSchemaInfoModel]          `tfsdk:"schema_info"`
}

type TargetAPIShieldOperationFeaturesThresholdsModel struct {
	AuthIDTokens       types.Int64       `tfsdk:"auth_id_tokens"`
	DataPoints         types.Int64       `tfsdk:"data_points"`
	LastUpdated        timetypes.RFC3339 `tfsdk:"last_updated"`
	P50                types.Int64       `tfsdk:"p50"`
	P90                types.Int64       `tfsdk:"p90"`
	P99                types.Int64       `tfsdk:"p99"`
	PeriodSeconds      types.Int64       `tfsdk:"period_seconds"`
	Requests           types.Int64       `tfsdk:"requests"`
	SuggestedThreshold types.Int64       `tfsdk:"suggested_threshold"`
}

type TargetAPIShieldOperationFeaturesParameterSchemasModel struct {
	LastUpdated      timetypes.RFC3339                                                                                     `tfsdk:"last_updated"`
	ParameterSchemas customfield.NestedObject[TargetAPIShieldOperationFeaturesParameterSchemasParameterSchemasModel]       `tfsdk:"parameter_schemas"`
}

type TargetAPIShieldOperationFeaturesParameterSchemasParameterSchemasModel struct {
	Parameters customfield.List[jsontypes.Normalized] `tfsdk:"parameters"`
	Responses  jsontypes.Normalized                   `tfsdk:"responses"`
}

type TargetAPIShieldOperationFeaturesAPIRoutingModel struct {
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated"`
	Route       types.String      `tfsdk:"route"`
}

type TargetAPIShieldOperationFeaturesConfidenceIntervalsModel struct {
	LastUpdated        timetypes.RFC3339                                                                                          `tfsdk:"last_updated"`
	SuggestedThreshold customfield.NestedObject[TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdModel]      `tfsdk:"suggested_threshold"`
}

type TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdModel struct {
	ConfidenceIntervals customfield.NestedObject[TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsModel] `tfsdk:"confidence_intervals"`
	Mean                types.Float64                                                                                                            `tfsdk:"mean"`
}

type TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsModel struct {
	P90 customfield.NestedObject[TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90Model] `tfsdk:"p90"`
	P95 customfield.NestedObject[TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95Model] `tfsdk:"p95"`
	P99 customfield.NestedObject[TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99Model] `tfsdk:"p99"`
}

type TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90Model struct {
	Lower types.Float64 `tfsdk:"lower"`
	Upper types.Float64 `tfsdk:"upper"`
}

type TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95Model struct {
	Lower types.Float64 `tfsdk:"lower"`
	Upper types.Float64 `tfsdk:"upper"`
}

type TargetAPIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99Model struct {
	Lower types.Float64 `tfsdk:"lower"`
	Upper types.Float64 `tfsdk:"upper"`
}

type TargetAPIShieldOperationFeaturesSchemaInfoModel struct {
	ActiveSchema     customfield.NestedObject[TargetAPIShieldOperationFeaturesSchemaInfoActiveSchemaModel] `tfsdk:"active_schema"`
	LearnedAvailable types.Bool                                                                            `tfsdk:"learned_available"`
	MitigationAction types.String                                                                          `tfsdk:"mitigation_action"`
}

type TargetAPIShieldOperationFeaturesSchemaInfoActiveSchemaModel struct {
	ID        types.String      `tfsdk:"id"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at"`
	IsLearned types.Bool        `tfsdk:"is_learned"`
	Name      types.String      `tfsdk:"name"`
}
