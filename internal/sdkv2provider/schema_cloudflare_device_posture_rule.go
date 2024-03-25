package sdkv2provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

var devicePostureRuleTypes = []string{"serial_number", "file", "application", "gateway", "warp", "domain_joined", "os_version", "disk_encryption", "firewall", "client_certificate", "workspace_one", "unique_client_id", "crowdstrike_s2s", "sentinelone", "kolide", "tanium_s2s", "intune", "sentinelone_s2s"}

func resourceCloudflareDevicePostureRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(devicePostureRuleTypes, false),
			Description:  fmt.Sprintf("The device posture rule type. %s", renderAvailableDocumentationValuesStringSlice(devicePostureRuleTypes)),
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the device posture rule.",
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"schedule": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Tells the client when to run the device posture check. Must be in the format `1h` or `30m`. Valid units are `h` and `m`.",
		},
		"expiration": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Expire posture results after the specified amount of time. Must be in the format `1h` or `30m`. Valid units are `h` and `m`.",
		},
		"match": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The conditions that the client must match to run the rule.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"platform": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"windows", "mac", "linux", "android", "ios", "chromeos"}, false),
						Description:  fmt.Sprintf("The platform of the device. %s", renderAvailableDocumentationValuesStringSlice([]string{"windows", "mac", "linux", "android", "ios", "chromeos"})),
					},
				},
			},
		},
		"input": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			Description: "Required for all rule types except `warp`, `gateway`, and `tanium`.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The Teams List id. Required for `serial_number` and `unique_client_id` rule types.",
					},
					"path": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The path to the file.",
					},
					"exists": {
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
						Description: "Checks if the file should exist.",
					},
					"thumbprint": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The thumbprint of the file certificate.",
					},
					"sha256": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The sha256 hash of the file.",
					},
					"running": {
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
						Description: "Checks if the application should be running",
					},
					"require_all": {
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
						Description: "True if all drives must be encrypted.",
					},
					"check_disks": {
						Type: schema.TypeSet,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional:    true,
						Description: "Specific volume(s) to check for encryption.",
					},
					"enabled": {
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
						Description: "True if the firewall must be enabled.",
					},
					"version": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The operating system semantic version.",
					},
					"operator": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{">", ">=", "<", "<=", "=="}, true),
						Description:  fmt.Sprintf("The version comparison operator. %s", renderAvailableDocumentationValuesStringSlice([]string{">", ">=", "<", "<=", "=="})),
					},
					"domain": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The domain that the client must join.",
					},
					"connection_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The workspace one connection id.",
					},
					"compliance_status": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"compliant", "noncompliant"}, true),
						Description:  fmt.Sprintf("The workspace one device compliance status. %s", renderAvailableDocumentationValuesStringSlice([]string{"compliant", "noncompliant"})),
					},
					"os_distro_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The operating system excluding version information.",
					},
					"os_distro_revision": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The operating system version excluding OS name information or release name.",
					},
					"os": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "OS signal score from Crowdstrike. Value must be between 1 and 100.",
					},
					"overall": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Overall ZTA score from Crowdstrike. Value must be between 1 and 100.",
					},
					"sensor_config": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Sensor signal score from Crowdstrike. Value must be between 1 and 100.",
					},
					"version_operator": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{">", ">=", "<", "<=", "=="}, true),
						Description:  fmt.Sprintf("The version comparison operator for crowdstrike. %s", renderAvailableDocumentationValuesStringSlice([]string{">", ">=", "<", "<=", "=="})),
					},
					"last_seen": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The duration of time that the host was last seen from Crowdstrike. Must be in the format `1h` or `30m`. Valid units are `d`, `h` and `m`.",
					},
					"state": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"online", "offline", "unknown"}, true),
						Description:  fmt.Sprintf("The host’s current online status from Crowdstrike. %s", renderAvailableDocumentationValuesStringSlice([]string{"online", "offline", "unknown"})),
					},
					"count_operator": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{">", ">=", "<", "<=", "=="}, true),
						Description:  fmt.Sprintf("The count comparison operator for kolide. %s", renderAvailableDocumentationValuesStringSlice([]string{">", ">=", "<", "<=", "=="})),
					},
					"issue_count": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The number of issues for kolide.",
					},
					"certificate_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The UUID of a Cloudflare managed certificate.",
					},
					"cn": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The common name for a certificate.",
					},
					"active_threats": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "The number of active threats from SentinelOne.",
					},
					"network_status": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"connected", "disconnected", "disconnecting", "connecting"}, true),
						Description:  fmt.Sprintf("The network status from SentinelOne. %s", renderAvailableDocumentationValuesStringSlice([]string{"connected", "disconnected", "disconnecting", "connecting"})),
					},
					"infected": {
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
						Description: "True if SentinelOne device is infected.",
					},
					"is_active": {
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
						Description: "True if SentinelOne device is active.",
					},
					"eid_last_seen": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The datetime a device last seen in RFC 3339 format from Tanium.",
					},
					"risk_level": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high", "critical"}, true),
						Description:  fmt.Sprintf("The risk level from Tanium. %s", renderAvailableDocumentationValuesStringSlice([]string{"low", "medium", "high", "critical"})),
					},
					"total_score": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "The total score from Tanium.",
					},
				},
			},
		},
	}
}
