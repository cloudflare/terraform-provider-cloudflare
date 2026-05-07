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

type ZeroTrustDeviceDeploymentGroupsResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceDeploymentGroupsDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceDeploymentGroupsDataSourceModel struct {
	ID            types.String                                                                              `tfsdk:"id" path:"group_id,computed"`
	GroupID       types.String                                                                              `tfsdk:"group_id" path:"group_id,required"`
	AccountID     types.String                                                                              `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt     types.String                                                                              `tfsdk:"created_at" json:"created_at,computed"`
	Name          types.String                                                                              `tfsdk:"name" json:"name,computed"`
	UpdatedAt     types.String                                                                              `tfsdk:"updated_at" json:"updated_at,computed"`
	PolicyIDs     customfield.List[types.String]                                                            `tfsdk:"policy_ids" json:"policy_ids,computed"`
	VersionConfig customfield.NestedObjectList[ZeroTrustDeviceDeploymentGroupsVersionConfigDataSourceModel] `tfsdk:"version_config" json:"version_config,computed"`
}

func (m *ZeroTrustDeviceDeploymentGroupsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DeviceDeploymentGroupGetParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceDeploymentGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceDeploymentGroupsVersionConfigDataSourceModel struct {
	TargetEnvironment types.String `tfsdk:"target_environment" json:"target_environment,computed"`
	Version           types.String `tfsdk:"version" json:"version,computed"`
}
