// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DevicePostureRulesResultListDataSourceEnvelope struct {
	Result *[]*DevicePostureRulesItemsDataSourceModel `json:"result,computed"`
}

type DevicePostureRulesDataSourceModel struct {
	AccountID types.String                               `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                `tfsdk:"max_items"`
	Items     *[]*DevicePostureRulesItemsDataSourceModel `tfsdk:"items"`
}

type DevicePostureRulesItemsDataSourceModel struct {
	ID          types.String                                    `tfsdk:"id" json:"id,computed"`
	Description types.String                                    `tfsdk:"description" json:"description,computed"`
	Expiration  types.String                                    `tfsdk:"expiration" json:"expiration,computed"`
	Match       *[]*DevicePostureRulesItemsMatchDataSourceModel `tfsdk:"match" json:"match,computed"`
	Name        types.String                                    `tfsdk:"name" json:"name,computed"`
	Schedule    types.String                                    `tfsdk:"schedule" json:"schedule,computed"`
	Type        types.String                                    `tfsdk:"type" json:"type,computed"`
}

type DevicePostureRulesItemsInputDataSourceModel struct {
	OperatingSystem  types.String  `tfsdk:"operating_system" json:"operating_system,computed"`
	Path             types.String  `tfsdk:"path" json:"path,computed"`
	Exists           types.Bool    `tfsdk:"exists" json:"exists,computed"`
	Sha256           types.String  `tfsdk:"sha256" json:"sha256,computed"`
	Thumbprint       types.String  `tfsdk:"thumbprint" json:"thumbprint,computed"`
	ID               types.String  `tfsdk:"id" json:"id,computed"`
	Domain           types.String  `tfsdk:"domain" json:"domain,computed"`
	Operator         types.String  `tfsdk:"operator" json:"operator,computed"`
	Version          types.String  `tfsdk:"version" json:"version,computed"`
	OSDistroName     types.String  `tfsdk:"os_distro_name" json:"os_distro_name,computed"`
	OSDistroRevision types.String  `tfsdk:"os_distro_revision" json:"os_distro_revision,computed"`
	OSVersionExtra   types.String  `tfsdk:"os_version_extra" json:"os_version_extra,computed"`
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled,computed"`
	CheckDisks       types.String  `tfsdk:"check_disks" json:"checkDisks,computed"`
	RequireAll       types.Bool    `tfsdk:"require_all" json:"requireAll,computed"`
	CertificateID    types.String  `tfsdk:"certificate_id" json:"certificate_id,computed"`
	Cn               types.String  `tfsdk:"cn" json:"cn,computed"`
	ComplianceStatus types.String  `tfsdk:"compliance_status" json:"compliance_status,computed"`
	ConnectionID     types.String  `tfsdk:"connection_id" json:"connection_id,computed"`
	LastSeen         types.String  `tfsdk:"last_seen" json:"last_seen,computed"`
	OS               types.String  `tfsdk:"os" json:"os,computed"`
	Overall          types.String  `tfsdk:"overall" json:"overall,computed"`
	SensorConfig     types.String  `tfsdk:"sensor_config" json:"sensor_config,computed"`
	State            types.String  `tfsdk:"state" json:"state,computed"`
	VersionOperator  types.String  `tfsdk:"version_operator" json:"versionOperator,computed"`
	CountOperator    types.String  `tfsdk:"count_operator" json:"countOperator,computed"`
	IssueCount       types.String  `tfsdk:"issue_count" json:"issue_count,computed"`
	EidLastSeen      types.String  `tfsdk:"eid_last_seen" json:"eid_last_seen,computed"`
	RiskLevel        types.String  `tfsdk:"risk_level" json:"risk_level,computed"`
	ScoreOperator    types.String  `tfsdk:"score_operator" json:"scoreOperator,computed"`
	TotalScore       types.Float64 `tfsdk:"total_score" json:"total_score,computed"`
	ActiveThreats    types.Float64 `tfsdk:"active_threats" json:"active_threats,computed"`
	Infected         types.Bool    `tfsdk:"infected" json:"infected,computed"`
	IsActive         types.Bool    `tfsdk:"is_active" json:"is_active,computed"`
	NetworkStatus    types.String  `tfsdk:"network_status" json:"network_status,computed"`
}

type DevicePostureRulesItemsMatchDataSourceModel struct {
	Platform types.String `tfsdk:"platform" json:"platform,computed"`
}
