// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_rule

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureRuleResultDataSourceEnvelope struct {
	Result ZeroTrustDevicePostureRuleDataSourceModel `json:"result,computed"`
}

type ZeroTrustDevicePostureRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDevicePostureRuleDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDevicePostureRuleDataSourceModel struct {
	AccountID   types.String                                        `tfsdk:"account_id" path:"account_id"`
	RuleID      types.String                                        `tfsdk:"rule_id" path:"rule_id"`
	Description types.String                                        `tfsdk:"description" json:"description,computed_optional"`
	Expiration  types.String                                        `tfsdk:"expiration" json:"expiration,computed_optional"`
	ID          types.String                                        `tfsdk:"id" json:"id,computed_optional"`
	Name        types.String                                        `tfsdk:"name" json:"name,computed_optional"`
	Schedule    types.String                                        `tfsdk:"schedule" json:"schedule,computed_optional"`
	Type        types.String                                        `tfsdk:"type" json:"type,computed_optional"`
	Input       *ZeroTrustDevicePostureRuleInputDataSourceModel     `tfsdk:"input" json:"input,computed_optional"`
	Match       *[]*ZeroTrustDevicePostureRuleMatchDataSourceModel  `tfsdk:"match" json:"match,computed_optional"`
	Filter      *ZeroTrustDevicePostureRuleFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustDevicePostureRuleDataSourceModel) toReadParams() (params zero_trust.DevicePostureGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePostureGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDevicePostureRuleDataSourceModel) toListParams() (params zero_trust.DevicePostureListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePostureListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDevicePostureRuleInputDataSourceModel struct {
	OperatingSystem  types.String                                             `tfsdk:"operating_system" json:"operating_system,computed_optional"`
	Path             types.String                                             `tfsdk:"path" json:"path,computed_optional"`
	Exists           types.Bool                                               `tfsdk:"exists" json:"exists,computed_optional"`
	Sha256           types.String                                             `tfsdk:"sha256" json:"sha256,computed_optional"`
	Thumbprint       types.String                                             `tfsdk:"thumbprint" json:"thumbprint,computed_optional"`
	ID               types.String                                             `tfsdk:"id" json:"id,computed_optional"`
	Domain           types.String                                             `tfsdk:"domain" json:"domain,computed_optional"`
	Operator         types.String                                             `tfsdk:"operator" json:"operator,computed_optional"`
	Version          types.String                                             `tfsdk:"version" json:"version,computed_optional"`
	OSDistroName     types.String                                             `tfsdk:"os_distro_name" json:"os_distro_name,computed_optional"`
	OSDistroRevision types.String                                             `tfsdk:"os_distro_revision" json:"os_distro_revision,computed_optional"`
	OSVersionExtra   types.String                                             `tfsdk:"os_version_extra" json:"os_version_extra,computed_optional"`
	Enabled          types.Bool                                               `tfsdk:"enabled" json:"enabled,computed_optional"`
	CheckDisks       *[]types.String                                          `tfsdk:"check_disks" json:"checkDisks,computed_optional"`
	RequireAll       types.Bool                                               `tfsdk:"require_all" json:"requireAll,computed_optional"`
	CertificateID    types.String                                             `tfsdk:"certificate_id" json:"certificate_id,computed_optional"`
	Cn               types.String                                             `tfsdk:"cn" json:"cn,computed_optional"`
	CheckPrivateKey  types.Bool                                               `tfsdk:"check_private_key" json:"check_private_key,computed_optional"`
	ExtendedKeyUsage *[]types.String                                          `tfsdk:"extended_key_usage" json:"extended_key_usage,computed_optional"`
	Locations        *ZeroTrustDevicePostureRuleInputLocationsDataSourceModel `tfsdk:"locations" json:"locations,computed_optional"`
	ComplianceStatus types.String                                             `tfsdk:"compliance_status" json:"compliance_status,computed_optional"`
	ConnectionID     types.String                                             `tfsdk:"connection_id" json:"connection_id,computed_optional"`
	LastSeen         types.String                                             `tfsdk:"last_seen" json:"last_seen,computed_optional"`
	OS               types.String                                             `tfsdk:"os" json:"os,computed_optional"`
	Overall          types.String                                             `tfsdk:"overall" json:"overall,computed_optional"`
	SensorConfig     types.String                                             `tfsdk:"sensor_config" json:"sensor_config,computed_optional"`
	State            types.String                                             `tfsdk:"state" json:"state,computed_optional"`
	VersionOperator  types.String                                             `tfsdk:"version_operator" json:"versionOperator,computed_optional"`
	CountOperator    types.String                                             `tfsdk:"count_operator" json:"countOperator,computed_optional"`
	IssueCount       types.String                                             `tfsdk:"issue_count" json:"issue_count,computed_optional"`
	EidLastSeen      types.String                                             `tfsdk:"eid_last_seen" json:"eid_last_seen,computed_optional"`
	RiskLevel        types.String                                             `tfsdk:"risk_level" json:"risk_level,computed_optional"`
	ScoreOperator    types.String                                             `tfsdk:"score_operator" json:"scoreOperator,computed_optional"`
	TotalScore       types.Float64                                            `tfsdk:"total_score" json:"total_score,computed_optional"`
	ActiveThreats    types.Float64                                            `tfsdk:"active_threats" json:"active_threats,computed_optional"`
	Infected         types.Bool                                               `tfsdk:"infected" json:"infected,computed_optional"`
	IsActive         types.Bool                                               `tfsdk:"is_active" json:"is_active,computed_optional"`
	NetworkStatus    types.String                                             `tfsdk:"network_status" json:"network_status,computed_optional"`
}

type ZeroTrustDevicePostureRuleInputLocationsDataSourceModel struct {
	Paths       *[]types.String `tfsdk:"paths" json:"paths,computed_optional"`
	TrustStores *[]types.String `tfsdk:"trust_stores" json:"trust_stores,computed_optional"`
}

type ZeroTrustDevicePostureRuleMatchDataSourceModel struct {
	Platform types.String `tfsdk:"platform" json:"platform,computed_optional"`
}

type ZeroTrustDevicePostureRuleFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
