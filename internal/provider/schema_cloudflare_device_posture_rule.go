package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareDevicePostureRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"serial_number", "file", "application", "gateway", "warp", "domain_joined", "os_version", "disk_encryption", "firewall", "workspace_one", "unique_client_id"}, false),
			Description:  fmt.Sprintf("The device posture rule type. %s", renderAvailableDocumentationValuesStringSlice([]string{"serial_number", "file", "application", "gateway", "warp", "domain_joined", "os_version", "disk_encryption", "firewall", "workspace_one", "unique_client_id"})),
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
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The Teams List id.",
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
						Description:  "The version comparison operator.",
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
						Description:  "The workspace one device compliance status.",
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
				},
			},
		},
	}
}
