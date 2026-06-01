package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy v4 SDKv2 Provider)
// ============================================================================

// SourceDevicePostureRuleModel represents the legacy cloudflare_device_posture_rule state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_device_posture_rule
type SourceDevicePostureRuleModel struct {
	ID          types.String       `tfsdk:"id"`
	AccountID   types.String       `tfsdk:"account_id"`
	Name        types.String       `tfsdk:"name"`
	Type        types.String       `tfsdk:"type"`
	Description types.String       `tfsdk:"description"`
	Expiration  types.String       `tfsdk:"expiration"`
	Schedule    types.String       `tfsdk:"schedule"`
	Input       []SourceInputModel `tfsdk:"input"` // v4 block stored as list (MaxItems:1)
	Match       []SourceMatchModel `tfsdk:"match"` // v4 multiple blocks stored as list
}

// SourceInputModel represents the input block from v4.x provider.
type SourceInputModel struct {
	OperatingSystem         types.String             `tfsdk:"operating_system"`
	Path                    types.String             `tfsdk:"path"`
	Exists                  types.Bool               `tfsdk:"exists"`
	Sha256                  types.String             `tfsdk:"sha256"`
	Thumbprint              types.String             `tfsdk:"thumbprint"`
	ID                      types.String             `tfsdk:"id"`
	Domain                  types.String             `tfsdk:"domain"`
	Operator                types.String             `tfsdk:"operator"`
	Version                 types.String             `tfsdk:"version"`
	OSDistroName            types.String             `tfsdk:"os_distro_name"`
	OSDistroRevision        types.String             `tfsdk:"os_distro_revision"`
	OSVersionExtra          types.String             `tfsdk:"os_version_extra"`
	Enabled                 types.Bool               `tfsdk:"enabled"`
	CheckDisks              types.Set                `tfsdk:"check_disks"` // Set in v4, List in v5
	RequireAll              types.Bool               `tfsdk:"require_all"`
	CertificateID           types.String             `tfsdk:"certificate_id"`
	Cn                      types.String             `tfsdk:"cn"`
	CheckPrivateKey         types.Bool               `tfsdk:"check_private_key"`
	ExtendedKeyUsage        types.List               `tfsdk:"extended_key_usage"`
	Locations               []SourceLocationsModel   `tfsdk:"locations"` // v4 nested block as list
	SubjectAlternativeNames types.List               `tfsdk:"subject_alternative_names"`
	UpdateWindowDays        types.Float64            `tfsdk:"update_window_days"`
	ComplianceStatus        types.String             `tfsdk:"compliance_status"`
	ConnectionID            types.String             `tfsdk:"connection_id"`
	LastSeen                types.String             `tfsdk:"last_seen"`
	OS                      types.String             `tfsdk:"os"`
	Overall                 types.String             `tfsdk:"overall"`
	SensorConfig            types.String             `tfsdk:"sensor_config"`
	State                   types.String             `tfsdk:"state"`
	VersionOperator         types.String             `tfsdk:"version_operator"`
	CountOperator           types.String             `tfsdk:"count_operator"`
	IssueCount              types.String             `tfsdk:"issue_count"`
	EidLastSeen             types.String             `tfsdk:"eid_last_seen"`
	RiskLevel               types.String             `tfsdk:"risk_level"`
	ScoreOperator           types.String             `tfsdk:"score_operator"`
	TotalScore              types.Float64            `tfsdk:"total_score"`
	ActiveThreats           types.Float64            `tfsdk:"active_threats"`
	Infected                types.Bool               `tfsdk:"infected"`
	IsActive                types.Bool               `tfsdk:"is_active"`
	NetworkStatus           types.String             `tfsdk:"network_status"`
	OperationalState        types.String             `tfsdk:"operational_state"`
	Score                   types.Float64            `tfsdk:"score"`
	Running                 types.Bool               `tfsdk:"running"` // Removed in v5
}

// SourceLocationsModel represents the nested locations block from v4.x provider.
type SourceLocationsModel struct {
	Paths       types.List `tfsdk:"paths"`
	TrustStores types.List `tfsdk:"trust_stores"`
}

// SourceMatchModel represents the match block from v4.x provider.
type SourceMatchModel struct {
	Platform types.String `tfsdk:"platform"`
}

// ============================================================================
// Target Models (Current v5 Provider)
// ============================================================================

// TargetDevicePostureRuleModel represents the current cloudflare_zero_trust_device_posture_rule state.
// Schema version: 500
// Resource type: cloudflare_zero_trust_device_posture_rule
type TargetDevicePostureRuleModel struct {
	ID          types.String          `tfsdk:"id"`
	AccountID   types.String          `tfsdk:"account_id"`
	Name        types.String          `tfsdk:"name"`
	Type        types.String          `tfsdk:"type"`
	Description types.String          `tfsdk:"description"`
	Expiration  types.String          `tfsdk:"expiration"`
	Schedule    types.String          `tfsdk:"schedule"`
	Input       *TargetInputModel     `tfsdk:"input"`
	Match       *[]*TargetMatchModel  `tfsdk:"match"`
}

// TargetInputModel represents the input nested attribute in v5.
type TargetInputModel struct {
	OperatingSystem         types.String          `tfsdk:"operating_system"`
	Path                    types.String          `tfsdk:"path"`
	Exists                  types.Bool            `tfsdk:"exists"`
	Sha256                  types.String          `tfsdk:"sha256"`
	Thumbprint              types.String          `tfsdk:"thumbprint"`
	ID                      types.String          `tfsdk:"id"`
	Domain                  types.String          `tfsdk:"domain"`
	Operator                types.String          `tfsdk:"operator"`
	Version                 types.String          `tfsdk:"version"`
	OSDistroName            types.String          `tfsdk:"os_distro_name"`
	OSDistroRevision        types.String          `tfsdk:"os_distro_revision"`
	OSVersionExtra          types.String          `tfsdk:"os_version_extra"`
	Enabled                 types.Bool            `tfsdk:"enabled"`
	CheckDisks              *[]types.String       `tfsdk:"check_disks"`
	RequireAll              types.Bool            `tfsdk:"require_all"`
	CertificateID           types.String          `tfsdk:"certificate_id"`
	Cn                      types.String          `tfsdk:"cn"`
	CheckPrivateKey         types.Bool            `tfsdk:"check_private_key"`
	ExtendedKeyUsage        *[]types.String       `tfsdk:"extended_key_usage"`
	Locations               *TargetLocationsModel `tfsdk:"locations"`
	SubjectAlternativeNames *[]types.String       `tfsdk:"subject_alternative_names"`
	UpdateWindowDays        types.Float64         `tfsdk:"update_window_days"`
	ComplianceStatus        types.String          `tfsdk:"compliance_status"`
	ConnectionID            types.String          `tfsdk:"connection_id"`
	LastSeen                types.String          `tfsdk:"last_seen"`
	OS                      types.String          `tfsdk:"os"`
	Overall                 types.String          `tfsdk:"overall"`
	SensorConfig            types.String          `tfsdk:"sensor_config"`
	State                   types.String          `tfsdk:"state"`
	VersionOperator         types.String          `tfsdk:"version_operator"`
	CountOperator           types.String          `tfsdk:"count_operator"`
	IssueCount              types.String          `tfsdk:"issue_count"`
	EidLastSeen             types.String          `tfsdk:"eid_last_seen"`
	RiskLevel               types.String          `tfsdk:"risk_level"`
	ScoreOperator           types.String          `tfsdk:"score_operator"`
	TotalScore              types.Float64         `tfsdk:"total_score"`
	ActiveThreats           types.Float64         `tfsdk:"active_threats"`
	Infected                types.Bool            `tfsdk:"infected"`
	IsActive                types.Bool            `tfsdk:"is_active"`
	NetworkStatus           types.String          `tfsdk:"network_status"`
	OperationalState        types.String          `tfsdk:"operational_state"`
	Score                   types.Float64         `tfsdk:"score"`
	// Note: Running field from v4 is intentionally not included (removed in v5)
}

// TargetLocationsModel represents the nested locations attribute in v5.
type TargetLocationsModel struct {
	Paths       *[]types.String `tfsdk:"paths"`
	TrustStores *[]types.String `tfsdk:"trust_stores"`
}

// TargetMatchModel represents each match entry in v5.
type TargetMatchModel struct {
	Platform types.String `tfsdk:"platform"`
}
