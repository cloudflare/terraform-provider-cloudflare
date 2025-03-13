// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/api_gateway"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[APIShieldOperationsResultDataSourceModel] `json:"result,computed"`
}

type APIShieldOperationsDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Direction types.String `tfsdk:"direction" query:"direction,optional"`
Endpoint types.String `tfsdk:"endpoint" query:"endpoint,optional"`
Order types.String `tfsdk:"order" query:"order,optional"`
Feature *[]types.String `tfsdk:"feature" query:"feature,optional"`
Host *[]types.String `tfsdk:"host" query:"host,optional"`
Method *[]types.String `tfsdk:"method" query:"method,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[APIShieldOperationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *APIShieldOperationsDataSourceModel) toListParams(_ context.Context) (params api_gateway.OperationListParams, diags diag.Diagnostics) {
  mFeature := []api_gateway.OperationListParamsFeature{}
  for _, item := range *m.Feature {
    mFeature = append(mFeature, api_gateway.OperationListParamsFeature(item.ValueString()))
  }
  mHost := []string{}
  for _, item := range *m.Host {
    mHost = append(mHost, item.ValueString())
  }
  mMethod := []string{}
  for _, item := range *m.Method {
    mMethod = append(mMethod, item.ValueString())
  }

  params = api_gateway.OperationListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
    Feature: cloudflare.F(mFeature),
    Host: cloudflare.F(mHost),
    Method: cloudflare.F(mMethod),
  }

  if !m.Direction.IsNull() {
    params.Direction = cloudflare.F(api_gateway.OperationListParamsDirection(m.Direction.ValueString()))
  }
  if !m.Endpoint.IsNull() {
    params.Endpoint = cloudflare.F(m.Endpoint.ValueString())
  }
  if !m.Order.IsNull() {
    params.Order = cloudflare.F(api_gateway.OperationListParamsOrder(m.Order.ValueString()))
  }

  return
}

type APIShieldOperationsResultDataSourceModel struct {
Endpoint types.String `tfsdk:"endpoint" json:"endpoint,computed"`
Host types.String `tfsdk:"host" json:"host,computed"`
LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
Method types.String `tfsdk:"method" json:"method,computed"`
OperationID types.String `tfsdk:"operation_id" json:"operation_id,computed"`
Features customfield.NestedObject[APIShieldOperationsFeaturesDataSourceModel] `tfsdk:"features" json:"features,computed"`
}

type APIShieldOperationsFeaturesDataSourceModel struct {
Thresholds customfield.NestedObject[APIShieldOperationsFeaturesThresholdsDataSourceModel] `tfsdk:"thresholds" json:"thresholds,computed"`
ParameterSchemas customfield.NestedObject[APIShieldOperationsFeaturesParameterSchemasDataSourceModel] `tfsdk:"parameter_schemas" json:"parameter_schemas,computed"`
APIRouting customfield.NestedObject[APIShieldOperationsFeaturesAPIRoutingDataSourceModel] `tfsdk:"api_routing" json:"api_routing,computed"`
ConfidenceIntervals customfield.NestedObject[APIShieldOperationsFeaturesConfidenceIntervalsDataSourceModel] `tfsdk:"confidence_intervals" json:"confidence_intervals,computed"`
SchemaInfo customfield.NestedObject[APIShieldOperationsFeaturesSchemaInfoDataSourceModel] `tfsdk:"schema_info" json:"schema_info,computed"`
}

type APIShieldOperationsFeaturesThresholdsDataSourceModel struct {
AuthIDTokens types.Int64 `tfsdk:"auth_id_tokens" json:"auth_id_tokens,computed"`
DataPoints types.Int64 `tfsdk:"data_points" json:"data_points,computed"`
LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
P50 types.Int64 `tfsdk:"p50" json:"p50,computed"`
P90 types.Int64 `tfsdk:"p90" json:"p90,computed"`
P99 types.Int64 `tfsdk:"p99" json:"p99,computed"`
PeriodSeconds types.Int64 `tfsdk:"period_seconds" json:"period_seconds,computed"`
Requests types.Int64 `tfsdk:"requests" json:"requests,computed"`
SuggestedThreshold types.Int64 `tfsdk:"suggested_threshold" json:"suggested_threshold,computed"`
}

type APIShieldOperationsFeaturesParameterSchemasDataSourceModel struct {
LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
ParameterSchemas customfield.NestedObject[APIShieldOperationsFeaturesParameterSchemasParameterSchemasDataSourceModel] `tfsdk:"parameter_schemas" json:"parameter_schemas,computed"`
}

type APIShieldOperationsFeaturesParameterSchemasParameterSchemasDataSourceModel struct {
Parameters customfield.List[jsontypes.Normalized] `tfsdk:"parameters" json:"parameters,computed"`
Responses jsontypes.Normalized `tfsdk:"responses" json:"responses,computed"`
}

type APIShieldOperationsFeaturesAPIRoutingDataSourceModel struct {
LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
Route types.String `tfsdk:"route" json:"route,computed"`
}

type APIShieldOperationsFeaturesConfidenceIntervalsDataSourceModel struct {
LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
SuggestedThreshold customfield.NestedObject[APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdDataSourceModel] `tfsdk:"suggested_threshold" json:"suggested_threshold,computed"`
}

type APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdDataSourceModel struct {
ConfidenceIntervals customfield.NestedObject[APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsDataSourceModel] `tfsdk:"confidence_intervals" json:"confidence_intervals,computed"`
Mean types.Float64 `tfsdk:"mean" json:"mean,computed"`
}

type APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsDataSourceModel struct {
P90 customfield.NestedObject[APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90DataSourceModel] `tfsdk:"p90" json:"p90,computed"`
P95 customfield.NestedObject[APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95DataSourceModel] `tfsdk:"p95" json:"p95,computed"`
P99 customfield.NestedObject[APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99DataSourceModel] `tfsdk:"p99" json:"p99,computed"`
}

type APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90DataSourceModel struct {
Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95DataSourceModel struct {
Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationsFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99DataSourceModel struct {
Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationsFeaturesSchemaInfoDataSourceModel struct {
ActiveSchema customfield.NestedObject[APIShieldOperationsFeaturesSchemaInfoActiveSchemaDataSourceModel] `tfsdk:"active_schema" json:"active_schema,computed"`
LearnedAvailable types.Bool `tfsdk:"learned_available" json:"learned_available,computed"`
MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action,computed"`
}

type APIShieldOperationsFeaturesSchemaInfoActiveSchemaDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
IsLearned types.Bool `tfsdk:"is_learned" json:"is_learned,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}
