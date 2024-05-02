// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DevicePostureRuleResultEnvelope struct {
	Result DevicePostureRuleModel `json:"result,computed"`
}

type DevicePostureRuleModel struct {
	ID          types.String                    `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                    `tfsdk:"account_id" path:"account_id"`
	Name        types.String                    `tfsdk:"name" json:"name"`
	Type        types.String                    `tfsdk:"type" json:"type"`
	Description types.String                    `tfsdk:"description" json:"description"`
	Expiration  types.String                    `tfsdk:"expiration" json:"expiration"`
	Input       *DevicePostureRuleInputModel    `tfsdk:"input" json:"input"`
	Match       *[]*DevicePostureRuleMatchModel `tfsdk:"match" json:"match"`
	Schedule    types.String                    `tfsdk:"schedule" json:"schedule"`
}

type DevicePostureRuleInputModel struct {
	OperatingSystem  types.String  `tfsdk:"operating_system" json:"operating_system"`
	Path             types.String  `tfsdk:"path" json:"path"`
	Exists           types.Bool    `tfsdk:"exists" json:"exists"`
	Sha256           types.String  `tfsdk:"sha256" json:"sha256"`
	Thumbprint       types.String  `tfsdk:"thumbprint" json:"thumbprint"`
	ID               types.String  `tfsdk:"id" json:"id"`
	Domain           types.String  `tfsdk:"domain" json:"domain"`
	Operator         types.String  `tfsdk:"operator" json:"operator"`
	Version          types.String  `tfsdk:"version" json:"version"`
	OSDistroName     types.String  `tfsdk:"os_distro_name" json:"os_distro_name"`
	OSDistroRevision types.String  `tfsdk:"os_distro_revision" json:"os_distro_revision"`
	OSVersionExtra   types.String  `tfsdk:"os_version_extra" json:"os_version_extra"`
	Enabled          types.Bool    `tfsdk:"enabled" json:"enabled"`
	CheckDisks       types.String  `tfsdk:"checkdisks" json:"checkDisks"`
	RequireAll       types.Bool    `tfsdk:"requireall" json:"requireAll"`
	CertificateID    types.String  `tfsdk:"certificate_id" json:"certificate_id"`
	Cn               types.String  `tfsdk:"cn" json:"cn"`
	ComplianceStatus types.String  `tfsdk:"compliance_status" json:"compliance_status"`
	ConnectionID     types.String  `tfsdk:"connection_id" json:"connection_id"`
	LastSeen         types.String  `tfsdk:"last_seen" json:"last_seen"`
	OS               types.String  `tfsdk:"os" json:"os"`
	Overall          types.String  `tfsdk:"overall" json:"overall"`
	SensorConfig     types.String  `tfsdk:"sensor_config" json:"sensor_config"`
	State            types.String  `tfsdk:"state" json:"state"`
	VersionOperator  types.String  `tfsdk:"versionoperator" json:"versionOperator"`
	CountOperator    types.String  `tfsdk:"countoperator" json:"countOperator"`
	IssueCount       types.String  `tfsdk:"issue_count" json:"issue_count"`
	EidLastSeen      types.String  `tfsdk:"eid_last_seen" json:"eid_last_seen"`
	RiskLevel        types.String  `tfsdk:"risk_level" json:"risk_level"`
	ScoreOperator    types.String  `tfsdk:"scoreoperator" json:"scoreOperator"`
	TotalScore       types.Float64 `tfsdk:"total_score" json:"total_score"`
	ActiveThreats    types.Float64 `tfsdk:"active_threats" json:"active_threats"`
	Infected         types.Bool    `tfsdk:"infected" json:"infected"`
	IsActive         types.Bool    `tfsdk:"is_active" json:"is_active"`
	NetworkStatus    types.String  `tfsdk:"network_status" json:"network_status"`
}

type DevicePostureRuleMatchModel struct {
	Platform types.String `tfsdk:"platform" json:"platform"`
}
