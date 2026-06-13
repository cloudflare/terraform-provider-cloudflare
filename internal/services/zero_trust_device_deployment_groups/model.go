// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_deployment_groups

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDeploymentGroupsResultEnvelope struct {
	Result ZeroTrustDeviceDeploymentGroupsModel `json:"result"`
}

type ZeroTrustDeviceDeploymentGroupsModel struct {
	ID            types.String                                          `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                          `tfsdk:"account_id" path:"account_id,required"`
	Name          types.String                                          `tfsdk:"name" json:"name,required"`
	VersionConfig *[]*ZeroTrustDeviceDeploymentGroupsVersionConfigModel `tfsdk:"version_config" json:"version_config,required"`
	PolicyIDs     *[]types.String                                       `tfsdk:"policy_ids" json:"policy_ids,optional"`
	CreatedAt     types.String                                          `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt     types.String                                          `tfsdk:"updated_at" json:"updated_at,computed"`
}

func (m ZeroTrustDeviceDeploymentGroupsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceDeploymentGroupsModel) MarshalJSONForUpdate(state ZeroTrustDeviceDeploymentGroupsModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ZeroTrustDeviceDeploymentGroupsVersionConfigModel struct {
	TargetEnvironment types.String `tfsdk:"target_environment" json:"target_environment,required"`
	Version           types.String `tfsdk:"version" json:"version,required"`
}
