// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationResultEnvelope struct {
	Result APIShieldOperationModel `json:"result"`
}

type APIShieldOperationModel struct {
	ID          types.String                                              `tfsdk:"id" json:"-,computed"`
	OperationID types.String                                              `tfsdk:"operation_id" json:"operation_id,computed"`
	ZoneID      types.String                                              `tfsdk:"zone_id" path:"zone_id,required"`
	Endpoint    types.String                                              `tfsdk:"endpoint" json:"endpoint,required"`
	Host        types.String                                              `tfsdk:"host" json:"host,required"`
	Method      types.String                                              `tfsdk:"method" json:"method,required"`
	LastUpdated timetypes.RFC3339                                         `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Features    customfield.NestedObject[APIShieldOperationFeaturesModel] `tfsdk:"features" json:"features,computed"`
}

func (m APIShieldOperationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIShieldOperationModel) MarshalJSONForUpdate(state APIShieldOperationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type APIShieldOperationFeaturesModel struct {
	Thresholds          customfield.NestedObject[APIShieldOperationFeaturesThresholdsModel]          `tfsdk:"thresholds" json:"thresholds,computed"`
	ParameterSchemas    customfield.NestedObject[APIShieldOperationFeaturesParameterSchemasModel]    `tfsdk:"parameter_schemas" json:"parameter_schemas,computed"`
	APIRouting          customfield.NestedObject[APIShieldOperationFeaturesAPIRoutingModel]          `tfsdk:"api_routing" json:"api_routing,computed"`
	ConfidenceIntervals customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsModel] `tfsdk:"confidence_intervals" json:"confidence_intervals,computed"`
	SchemaInfo          customfield.NestedObject[APIShieldOperationFeaturesSchemaInfoModel]          `tfsdk:"schema_info" json:"schema_info,computed"`
}

type APIShieldOperationFeaturesThresholdsModel struct {
	AuthIDTokens       types.Int64       `tfsdk:"auth_id_tokens" json:"auth_id_tokens,computed"`
	DataPoints         types.Int64       `tfsdk:"data_points" json:"data_points,computed"`
	LastUpdated        timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	P50                types.Int64       `tfsdk:"p50" json:"p50,computed"`
	P90                types.Int64       `tfsdk:"p90" json:"p90,computed"`
	P99                types.Int64       `tfsdk:"p99" json:"p99,computed"`
	PeriodSeconds      types.Int64       `tfsdk:"period_seconds" json:"period_seconds,computed"`
	Requests           types.Int64       `tfsdk:"requests" json:"requests,computed"`
	SuggestedThreshold types.Int64       `tfsdk:"suggested_threshold" json:"suggested_threshold,computed"`
}

type APIShieldOperationFeaturesParameterSchemasModel struct {
	LastUpdated      timetypes.RFC3339                                                                         `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	ParameterSchemas customfield.NestedObject[APIShieldOperationFeaturesParameterSchemasParameterSchemasModel] `tfsdk:"parameter_schemas" json:"parameter_schemas,computed"`
}

type APIShieldOperationFeaturesParameterSchemasParameterSchemasModel struct {
	Parameters customfield.List[jsontypes.Normalized] `tfsdk:"parameters" json:"parameters,computed"`
	Responses  jsontypes.Normalized                   `tfsdk:"responses" json:"responses,computed"`
}

type APIShieldOperationFeaturesAPIRoutingModel struct {
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Route       types.String      `tfsdk:"route" json:"route,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsModel struct {
	LastUpdated        timetypes.RFC3339                                                                              `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	SuggestedThreshold customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdModel] `tfsdk:"suggested_threshold" json:"suggested_threshold,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdModel struct {
	ConfidenceIntervals customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsModel] `tfsdk:"confidence_intervals" json:"confidence_intervals,computed"`
	Mean                types.Float64                                                                                                     `tfsdk:"mean" json:"mean,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsModel struct {
	P90 customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90Model] `tfsdk:"p90" json:"p90,computed"`
	P95 customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95Model] `tfsdk:"p95" json:"p95,computed"`
	P99 customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99Model] `tfsdk:"p99" json:"p99,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90Model struct {
	Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
	Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95Model struct {
	Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
	Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99Model struct {
	Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
	Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationFeaturesSchemaInfoModel struct {
	ActiveSchema     customfield.NestedObject[APIShieldOperationFeaturesSchemaInfoActiveSchemaModel] `tfsdk:"active_schema" json:"active_schema,computed"`
	LearnedAvailable types.Bool                                                                      `tfsdk:"learned_available" json:"learned_available,computed"`
	MitigationAction types.String                                                                    `tfsdk:"mitigation_action" json:"mitigation_action,computed"`
}

type APIShieldOperationFeaturesSchemaInfoActiveSchemaModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsLearned types.Bool        `tfsdk:"is_learned" json:"is_learned,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
}
