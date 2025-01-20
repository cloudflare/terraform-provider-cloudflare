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

type APIShieldOperationResultDataSourceEnvelope struct {
	Result APIShieldOperationDataSourceModel `json:"result,computed"`
}

type APIShieldOperationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APIShieldOperationDataSourceModel] `json:"result,computed"`
}

type APIShieldOperationDataSourceModel struct {
	ZoneID      types.String                                                        `tfsdk:"zone_id" path:"zone_id,optional"`
	OperationID types.String                                                        `tfsdk:"operation_id" path:"operation_id,computed_optional"`
	Feature     *[]types.String                                                     `tfsdk:"feature" query:"feature,optional"`
	Endpoint    types.String                                                        `tfsdk:"endpoint" json:"endpoint,computed"`
	Host        types.String                                                        `tfsdk:"host" json:"host,computed"`
	LastUpdated timetypes.RFC3339                                                   `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Method      types.String                                                        `tfsdk:"method" json:"method,computed"`
	Features    customfield.NestedObject[APIShieldOperationFeaturesDataSourceModel] `tfsdk:"features" json:"features,computed"`
	Filter      *APIShieldOperationFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *APIShieldOperationDataSourceModel) toReadParams(_ context.Context) (params api_gateway.OperationGetParams, diags diag.Diagnostics) {
	params = api_gateway.OperationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *APIShieldOperationDataSourceModel) toListParams(_ context.Context) (params api_gateway.OperationListParams, diags diag.Diagnostics) {
	mFilterFeature := []api_gateway.OperationListParamsFeature{}
	for _, item := range *m.Filter.Feature {
		mFilterFeature = append(mFilterFeature, api_gateway.OperationListParamsFeature(item.ValueString()))
	}
	mFilterHost := []string{}
	for _, item := range *m.Filter.Host {
		mFilterHost = append(mFilterHost, item.ValueString())
	}
	mFilterMethod := []string{}
	for _, item := range *m.Filter.Method {
		mFilterMethod = append(mFilterMethod, item.ValueString())
	}

	params = api_gateway.OperationListParams{
		ZoneID:  cloudflare.F(m.Filter.ZoneID.ValueString()),
		Feature: cloudflare.F(mFilterFeature),
		Host:    cloudflare.F(mFilterHost),
		Method:  cloudflare.F(mFilterMethod),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(api_gateway.OperationListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Endpoint.IsNull() {
		params.Endpoint = cloudflare.F(m.Filter.Endpoint.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(api_gateway.OperationListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type APIShieldOperationFeaturesDataSourceModel struct {
	Thresholds          customfield.NestedObject[APIShieldOperationFeaturesThresholdsDataSourceModel]          `tfsdk:"thresholds" json:"thresholds,computed"`
	ParameterSchemas    customfield.NestedObject[APIShieldOperationFeaturesParameterSchemasDataSourceModel]    `tfsdk:"parameter_schemas" json:"parameter_schemas,computed"`
	APIRouting          customfield.NestedObject[APIShieldOperationFeaturesAPIRoutingDataSourceModel]          `tfsdk:"api_routing" json:"api_routing,computed"`
	ConfidenceIntervals customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsDataSourceModel] `tfsdk:"confidence_intervals" json:"confidence_intervals,computed"`
	SchemaInfo          customfield.NestedObject[APIShieldOperationFeaturesSchemaInfoDataSourceModel]          `tfsdk:"schema_info" json:"schema_info,computed"`
}

type APIShieldOperationFeaturesThresholdsDataSourceModel struct {
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

type APIShieldOperationFeaturesParameterSchemasDataSourceModel struct {
	LastUpdated      timetypes.RFC3339                                                                                   `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	ParameterSchemas customfield.NestedObject[APIShieldOperationFeaturesParameterSchemasParameterSchemasDataSourceModel] `tfsdk:"parameter_schemas" json:"parameter_schemas,computed"`
}

type APIShieldOperationFeaturesParameterSchemasParameterSchemasDataSourceModel struct {
	Parameters customfield.List[jsontypes.Normalized] `tfsdk:"parameters" json:"parameters,computed"`
	Responses  jsontypes.Normalized                   `tfsdk:"responses" json:"responses,computed"`
}

type APIShieldOperationFeaturesAPIRoutingDataSourceModel struct {
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Route       types.String      `tfsdk:"route" json:"route,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsDataSourceModel struct {
	LastUpdated        timetypes.RFC3339                                                                                        `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	SuggestedThreshold customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdDataSourceModel] `tfsdk:"suggested_threshold" json:"suggested_threshold,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdDataSourceModel struct {
	ConfidenceIntervals customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsDataSourceModel] `tfsdk:"confidence_intervals" json:"confidence_intervals,computed"`
	Mean                types.Float64                                                                                                               `tfsdk:"mean" json:"mean,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsDataSourceModel struct {
	P90 customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90DataSourceModel] `tfsdk:"p90" json:"p90,computed"`
	P95 customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95DataSourceModel] `tfsdk:"p95" json:"p95,computed"`
	P99 customfield.NestedObject[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99DataSourceModel] `tfsdk:"p99" json:"p99,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90DataSourceModel struct {
	Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
	Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95DataSourceModel struct {
	Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
	Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99DataSourceModel struct {
	Lower types.Float64 `tfsdk:"lower" json:"lower,computed"`
	Upper types.Float64 `tfsdk:"upper" json:"upper,computed"`
}

type APIShieldOperationFeaturesSchemaInfoDataSourceModel struct {
	ActiveSchema     customfield.NestedObject[APIShieldOperationFeaturesSchemaInfoActiveSchemaDataSourceModel] `tfsdk:"active_schema" json:"active_schema,computed"`
	LearnedAvailable types.Bool                                                                                `tfsdk:"learned_available" json:"learned_available,computed"`
	MitigationAction types.String                                                                              `tfsdk:"mitigation_action" json:"mitigation_action,computed"`
}

type APIShieldOperationFeaturesSchemaInfoActiveSchemaDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsLearned types.Bool        `tfsdk:"is_learned" json:"is_learned,computed"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
}

type APIShieldOperationFindOneByDataSourceModel struct {
	ZoneID    types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	Direction types.String    `tfsdk:"direction" query:"direction,optional"`
	Endpoint  types.String    `tfsdk:"endpoint" query:"endpoint,optional"`
	Feature   *[]types.String `tfsdk:"feature" query:"feature,optional"`
	Host      *[]types.String `tfsdk:"host" query:"host,optional"`
	Method    *[]types.String `tfsdk:"method" query:"method,optional"`
	Order     types.String    `tfsdk:"order" query:"order,optional"`
}
