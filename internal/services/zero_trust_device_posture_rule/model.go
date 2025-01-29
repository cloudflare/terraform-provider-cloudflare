// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureRuleResultEnvelope struct {
	Result ZeroTrustDevicePostureRuleModel `json:"result"`
}

type ZeroTrustDevicePostureRuleModel struct {
	ID          types.String                                                       `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String                                                       `tfsdk:"name" json:"name,required"`
	Type        types.String                                                       `tfsdk:"type" json:"type,required"`
	Description types.String                                                       `tfsdk:"description" json:"description,optional"`
	Expiration  types.String                                                       `tfsdk:"expiration" json:"expiration,optional"`
	Schedule    types.String                                                       `tfsdk:"schedule" json:"schedule,optional"`
	Input       customfield.NestedObject[ZeroTrustDevicePostureRuleInputModel]     `tfsdk:"input" json:"input,computed_optional"`
	Match       customfield.NestedObjectList[ZeroTrustDevicePostureRuleMatchModel] `tfsdk:"match" json:"match,computed_optional"`
}

func (m ZeroTrustDevicePostureRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDevicePostureRuleModel) MarshalJSONForUpdate(state ZeroTrustDevicePostureRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDevicePostureRuleInputModel struct {
	OperatingSystem  types.String                                                            `tfsdk:"operating_system" json:"operating_system,optional"`
	Path             types.String                                                            `tfsdk:"path" json:"path,optional"`
	Exists           types.Bool                                                              `tfsdk:"exists" json:"exists,optional"`
	Sha256           types.String                                                            `tfsdk:"sha256" json:"sha256,optional"`
	Thumbprint       types.String                                                            `tfsdk:"thumbprint" json:"thumbprint,optional"`
	ID               types.String                                                            `tfsdk:"id" json:"id,optional"`
	Domain           types.String                                                            `tfsdk:"domain" json:"domain,optional"`
	Operator         types.String                                                            `tfsdk:"operator" json:"operator,optional"`
	Version          types.String                                                            `tfsdk:"version" json:"version,optional"`
	OSDistroName     types.String                                                            `tfsdk:"os_distro_name" json:"os_distro_name,optional"`
	OSDistroRevision types.String                                                            `tfsdk:"os_distro_revision" json:"os_distro_revision,optional"`
	OSVersionExtra   types.String                                                            `tfsdk:"os_version_extra" json:"os_version_extra,optional"`
	Enabled          types.Bool                                                              `tfsdk:"enabled" json:"enabled,optional"`
	CheckDisks       *[]types.String                                                         `tfsdk:"check_disks" json:"checkDisks,optional"`
	RequireAll       types.Bool                                                              `tfsdk:"require_all" json:"requireAll,optional"`
	CertificateID    types.String                                                            `tfsdk:"certificate_id" json:"certificate_id,optional"`
	Cn               types.String                                                            `tfsdk:"cn" json:"cn,optional"`
	CheckPrivateKey  types.Bool                                                              `tfsdk:"check_private_key" json:"check_private_key,optional"`
	ExtendedKeyUsage *[]types.String                                                         `tfsdk:"extended_key_usage" json:"extended_key_usage,optional"`
	Locations        customfield.NestedObject[ZeroTrustDevicePostureRuleInputLocationsModel] `tfsdk:"locations" json:"locations,computed_optional"`
	ComplianceStatus types.String                                                            `tfsdk:"compliance_status" json:"compliance_status,optional"`
	ConnectionID     types.String                                                            `tfsdk:"connection_id" json:"connection_id,optional"`
	LastSeen         types.String                                                            `tfsdk:"last_seen" json:"last_seen,optional"`
	OS               types.String                                                            `tfsdk:"os" json:"os,optional"`
	Overall          types.String                                                            `tfsdk:"overall" json:"overall,optional"`
	SensorConfig     types.String                                                            `tfsdk:"sensor_config" json:"sensor_config,optional"`
	State            types.String                                                            `tfsdk:"state" json:"state,optional"`
	VersionOperator  types.String                                                            `tfsdk:"version_operator" json:"versionOperator,optional"`
	CountOperator    types.String                                                            `tfsdk:"count_operator" json:"countOperator,optional"`
	IssueCount       types.String                                                            `tfsdk:"issue_count" json:"issue_count,optional"`
	EidLastSeen      types.String                                                            `tfsdk:"eid_last_seen" json:"eid_last_seen,optional"`
	RiskLevel        types.String                                                            `tfsdk:"risk_level" json:"risk_level,optional"`
	ScoreOperator    types.String                                                            `tfsdk:"score_operator" json:"scoreOperator,optional"`
	TotalScore       types.Float64                                                           `tfsdk:"total_score" json:"total_score,optional"`
	ActiveThreats    types.Float64                                                           `tfsdk:"active_threats" json:"active_threats,optional"`
	Infected         types.Bool                                                              `tfsdk:"infected" json:"infected,optional"`
	IsActive         types.Bool                                                              `tfsdk:"is_active" json:"is_active,optional"`
	NetworkStatus    types.String                                                            `tfsdk:"network_status" json:"network_status,optional"`
	OperationalState types.String                                                            `tfsdk:"operational_state" json:"operational_state,optional"`
	Score            types.Float64                                                           `tfsdk:"score" json:"score,optional"`
}

type ZeroTrustDevicePostureRuleInputLocationsModel struct {
	Paths       *[]types.String `tfsdk:"paths" json:"paths,optional"`
	TrustStores *[]types.String `tfsdk:"trust_stores" json:"trust_stores,optional"`
}

type ZeroTrustDevicePostureRuleMatchModel struct {
	Platform types.String `tfsdk:"platform" json:"platform,optional"`
}
