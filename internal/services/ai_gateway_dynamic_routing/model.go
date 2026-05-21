// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_dynamic_routing

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AIGatewayDynamicRoutingResultEnvelope struct {
	Result AIGatewayDynamicRoutingModel `json:"result"`
}

type AIGatewayDynamicRoutingModel struct {
	ID         types.String                                                     `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	GatewayID  types.String                                                     `tfsdk:"gateway_id" path:"gateway_id,required"`
	Elements   *[]*AIGatewayDynamicRoutingElementsModel                         `tfsdk:"elements" json:"elements,required"`
	Name       types.String                                                     `tfsdk:"name" json:"name,required"`
	CreatedAt  timetypes.RFC3339                                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339                                                `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Success    types.Bool                                                       `tfsdk:"success" json:"success,computed,no_refresh"`
	Deployment customfield.NestedObject[AIGatewayDynamicRoutingDeploymentModel] `tfsdk:"deployment" json:"deployment,computed"`
	Route      customfield.NestedObject[AIGatewayDynamicRoutingRouteModel]      `tfsdk:"route" json:"route,computed,no_refresh"`
	Version    customfield.NestedObject[AIGatewayDynamicRoutingVersionModel]    `tfsdk:"version" json:"version,computed"`
}

func (m AIGatewayDynamicRoutingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AIGatewayDynamicRoutingModel) MarshalJSONForUpdate(state AIGatewayDynamicRoutingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type AIGatewayDynamicRoutingElementsModel struct {
	ID         types.String                                    `tfsdk:"id" json:"id,required"`
	Outputs    *AIGatewayDynamicRoutingElementsOutputsModel    `tfsdk:"outputs" json:"outputs,required"`
	Type       types.String                                    `tfsdk:"type" json:"type,required"`
	Properties *AIGatewayDynamicRoutingElementsPropertiesModel `tfsdk:"properties" json:"properties,optional"`
}

type AIGatewayDynamicRoutingElementsOutputsModel struct {
	Next      *AIGatewayDynamicRoutingElementsOutputsNextModel     `tfsdk:"next" json:"next,optional"`
	False     *AIGatewayDynamicRoutingElementsOutputsFalseModel    `tfsdk:"false" json:"false,optional"`
	True      *AIGatewayDynamicRoutingElementsOutputsTrueModel     `tfsdk:"true" json:"true,optional"`
	ElementID types.String                                         `tfsdk:"element_id" json:"elementId,optional"`
	Fallback  *AIGatewayDynamicRoutingElementsOutputsFallbackModel `tfsdk:"fallback" json:"fallback,optional"`
	Success   *AIGatewayDynamicRoutingElementsOutputsSuccessModel  `tfsdk:"success" json:"success,optional"`
}

type AIGatewayDynamicRoutingElementsOutputsNextModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,required"`
}

type AIGatewayDynamicRoutingElementsOutputsFalseModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,required"`
}

type AIGatewayDynamicRoutingElementsOutputsTrueModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,required"`
}

type AIGatewayDynamicRoutingElementsOutputsFallbackModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,required"`
}

type AIGatewayDynamicRoutingElementsOutputsSuccessModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,required"`
}

type AIGatewayDynamicRoutingElementsPropertiesModel struct {
	Conditions                      jsontypes.Normalized `tfsdk:"conditions" json:"conditions,optional"`
	Key                             types.String         `tfsdk:"key" json:"key,optional"`
	Limit                           types.Float64        `tfsdk:"limit" json:"limit,optional"`
	LimitType                       types.String         `tfsdk:"limit_type" json:"limitType,optional"`
	Window                          types.Float64        `tfsdk:"window" json:"window,optional"`
	Model                           types.String         `tfsdk:"model" json:"model,optional"`
	AIGatewayDynamicRoutingProvider types.String         `tfsdk:"ai_gateway_dynamic_routing_provider" json:"provider,optional"`
	Retries                         types.Float64        `tfsdk:"retries" json:"retries,optional"`
	Timeout                         types.Float64        `tfsdk:"timeout" json:"timeout,optional"`
}

type AIGatewayDynamicRoutingDeploymentModel struct {
	CreatedAt    types.String `tfsdk:"created_at" json:"created_at,computed"`
	DeploymentID types.String `tfsdk:"deployment_id" json:"deployment_id,computed"`
	VersionID    types.String `tfsdk:"version_id" json:"version_id,computed"`
}

type AIGatewayDynamicRoutingRouteModel struct {
	ID         types.String                                                            `tfsdk:"id" json:"id,computed"`
	AccountTag types.String                                                            `tfsdk:"account_tag" json:"account_tag,computed"`
	CreatedAt  timetypes.RFC3339                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Deployment customfield.NestedObject[AIGatewayDynamicRoutingRouteDeploymentModel]   `tfsdk:"deployment" json:"deployment,computed"`
	Elements   customfield.NestedObjectList[AIGatewayDynamicRoutingRouteElementsModel] `tfsdk:"elements" json:"elements,computed"`
	GatewayID  types.String                                                            `tfsdk:"gateway_id" json:"gateway_id,computed"`
	ModifiedAt timetypes.RFC3339                                                       `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name       types.String                                                            `tfsdk:"name" json:"name,computed"`
	Version    customfield.NestedObject[AIGatewayDynamicRoutingRouteVersionModel]      `tfsdk:"version" json:"version,computed"`
}

type AIGatewayDynamicRoutingRouteDeploymentModel struct {
	CreatedAt    types.String `tfsdk:"created_at" json:"created_at,computed"`
	DeploymentID types.String `tfsdk:"deployment_id" json:"deployment_id,computed"`
	VersionID    types.String `tfsdk:"version_id" json:"version_id,computed"`
}

type AIGatewayDynamicRoutingRouteElementsModel struct {
	ID         types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Outputs    customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsOutputsModel]    `tfsdk:"outputs" json:"outputs,computed"`
	Type       types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Properties customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsPropertiesModel] `tfsdk:"properties" json:"properties,computed"`
}

type AIGatewayDynamicRoutingRouteElementsOutputsModel struct {
	Next      customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsOutputsNextModel]     `tfsdk:"next" json:"next,computed"`
	False     customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsOutputsFalseModel]    `tfsdk:"false" json:"false,computed"`
	True      customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsOutputsTrueModel]     `tfsdk:"true" json:"true,computed"`
	ElementID types.String                                                                       `tfsdk:"element_id" json:"elementId,computed"`
	Fallback  customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsOutputsFallbackModel] `tfsdk:"fallback" json:"fallback,computed"`
	Success   customfield.NestedObject[AIGatewayDynamicRoutingRouteElementsOutputsSuccessModel]  `tfsdk:"success" json:"success,computed"`
}

type AIGatewayDynamicRoutingRouteElementsOutputsNextModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingRouteElementsOutputsFalseModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingRouteElementsOutputsTrueModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingRouteElementsOutputsFallbackModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingRouteElementsOutputsSuccessModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingRouteElementsPropertiesModel struct {
	Conditions                      jsontypes.Normalized `tfsdk:"conditions" json:"conditions,computed"`
	Key                             types.String         `tfsdk:"key" json:"key,computed"`
	Limit                           types.Float64        `tfsdk:"limit" json:"limit,computed"`
	LimitType                       types.String         `tfsdk:"limit_type" json:"limitType,computed"`
	Window                          types.Float64        `tfsdk:"window" json:"window,computed"`
	Model                           types.String         `tfsdk:"model" json:"model,computed"`
	AIGatewayDynamicRoutingProvider types.String         `tfsdk:"ai_gateway_dynamic_routing_provider" json:"provider,computed"`
	Retries                         types.Float64        `tfsdk:"retries" json:"retries,computed"`
	Timeout                         types.Float64        `tfsdk:"timeout" json:"timeout,computed"`
}

type AIGatewayDynamicRoutingRouteVersionModel struct {
	Active    types.String `tfsdk:"active" json:"active,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Data      types.String `tfsdk:"data" json:"data,computed"`
	VersionID types.String `tfsdk:"version_id" json:"version_id,computed"`
}

type AIGatewayDynamicRoutingVersionModel struct {
	Active    types.String `tfsdk:"active" json:"active,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Data      types.String `tfsdk:"data" json:"data,computed"`
	VersionID types.String `tfsdk:"version_id" json:"version_id,computed"`
}
