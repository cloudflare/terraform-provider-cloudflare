// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_dynamic_routing

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AIGatewayDynamicRoutingResultDataSourceEnvelope struct {
	Result AIGatewayDynamicRoutingDataSourceModel `json:"result,computed"`
}

type AIGatewayDynamicRoutingDataSourceModel struct {
	AccountID  types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	GatewayID  types.String                                                                 `tfsdk:"gateway_id" path:"gateway_id,required"`
	ID         types.String                                                                 `tfsdk:"id" path:"id,required"`
	CreatedAt  timetypes.RFC3339                                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339                                                            `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name       types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Deployment customfield.NestedObject[AIGatewayDynamicRoutingDeploymentDataSourceModel]   `tfsdk:"deployment" json:"deployment,computed"`
	Elements   customfield.NestedObjectList[AIGatewayDynamicRoutingElementsDataSourceModel] `tfsdk:"elements" json:"elements,computed"`
	Version    customfield.NestedObject[AIGatewayDynamicRoutingVersionDataSourceModel]      `tfsdk:"version" json:"version,computed"`
}

func (m *AIGatewayDynamicRoutingDataSourceModel) toReadParams(_ context.Context) (params ai_gateway.DynamicRoutingGetParams, diags diag.Diagnostics) {
	params = ai_gateway.DynamicRoutingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AIGatewayDynamicRoutingDeploymentDataSourceModel struct {
	CreatedAt    types.String `tfsdk:"created_at" json:"created_at,computed"`
	DeploymentID types.String `tfsdk:"deployment_id" json:"deployment_id,computed"`
	VersionID    types.String `tfsdk:"version_id" json:"version_id,computed"`
}

type AIGatewayDynamicRoutingElementsDataSourceModel struct {
	ID         types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Outputs    customfield.NestedObject[AIGatewayDynamicRoutingElementsOutputsDataSourceModel]    `tfsdk:"outputs" json:"outputs,computed"`
	Type       types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Properties customfield.NestedObject[AIGatewayDynamicRoutingElementsPropertiesDataSourceModel] `tfsdk:"properties" json:"properties,computed"`
}

type AIGatewayDynamicRoutingElementsOutputsDataSourceModel struct {
	Next      customfield.NestedObject[AIGatewayDynamicRoutingElementsOutputsNextDataSourceModel]     `tfsdk:"next" json:"next,computed"`
	False     customfield.NestedObject[AIGatewayDynamicRoutingElementsOutputsFalseDataSourceModel]    `tfsdk:"false" json:"false,computed"`
	True      customfield.NestedObject[AIGatewayDynamicRoutingElementsOutputsTrueDataSourceModel]     `tfsdk:"true" json:"true,computed"`
	ElementID types.String                                                                            `tfsdk:"element_id" json:"elementId,computed"`
	Fallback  customfield.NestedObject[AIGatewayDynamicRoutingElementsOutputsFallbackDataSourceModel] `tfsdk:"fallback" json:"fallback,computed"`
	Success   customfield.NestedObject[AIGatewayDynamicRoutingElementsOutputsSuccessDataSourceModel]  `tfsdk:"success" json:"success,computed"`
}

type AIGatewayDynamicRoutingElementsOutputsNextDataSourceModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingElementsOutputsFalseDataSourceModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingElementsOutputsTrueDataSourceModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingElementsOutputsFallbackDataSourceModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingElementsOutputsSuccessDataSourceModel struct {
	ElementID types.String `tfsdk:"element_id" json:"elementId,computed"`
}

type AIGatewayDynamicRoutingElementsPropertiesDataSourceModel struct {
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

type AIGatewayDynamicRoutingVersionDataSourceModel struct {
	Active    types.String `tfsdk:"active" json:"active,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Data      types.String `tfsdk:"data" json:"data,computed"`
	VersionID types.String `tfsdk:"version_id" json:"version_id,computed"`
}
