// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_deployment_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDeploymentGroupsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDeviceDeploymentGroupsListResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDeviceDeploymentGroupsListDataSourceModel struct {
	AccountID types.String                                                                           `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDeviceDeploymentGroupsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDeviceDeploymentGroupsListDataSourceModel) toListParams(_ context.Context) (params zero_trust.DeviceDeploymentGroupListParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceDeploymentGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceDeploymentGroupsListResultDataSourceModel struct {
	ID            types.String                                                                                  `tfsdk:"id" json:"id,computed"`
	CreatedAt     types.String                                                                                  `tfsdk:"created_at" json:"created_at,computed"`
	Name          types.String                                                                                  `tfsdk:"name" json:"name,computed"`
	UpdatedAt     types.String                                                                                  `tfsdk:"updated_at" json:"updated_at,computed"`
	VersionConfig customfield.NestedObjectList[ZeroTrustDeviceDeploymentGroupsListVersionConfigDataSourceModel] `tfsdk:"version_config" json:"version_config,computed"`
	PolicyIDs     customfield.List[types.String]                                                                `tfsdk:"policy_ids" json:"policy_ids,computed"`
}

type ZeroTrustDeviceDeploymentGroupsListVersionConfigDataSourceModel struct {
	TargetEnvironment types.String `tfsdk:"target_environment" json:"target_environment,computed"`
	Version           types.String `tfsdk:"version" json:"version,computed"`
}
