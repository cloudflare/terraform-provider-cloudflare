package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resoureceCloudflareDevicesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"devices": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"created": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was created.",
					},
					"deleted": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether the device has been deleted.",
					},
					"device_type": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The type of the device.",
					},
					"id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Device ID.",
					},
					"ip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "IPv4 or IPv6 address.",
					},
					"key": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device's public key.",
					},
					"last_seen": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was last seen.",
					},
					"mac_address": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device's MAC address.",
					},
					"manufacturer": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device manufacturer's name.",
					},
					"model": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device model name.",
					},
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device name.",
					},
					"os_distro_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The Linux distribution name.",
					},
					"os_distro_revision": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The Linux distribution revision.",
					},
					"os_version": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The operating system version.",
					},
					"revoked_at": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was revoked.",
					},
					"serial_number": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device's serial number.",
					},
					"updated": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was updated.",
					},
					"user_email": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "User's email.",
					},
					"user_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "User's ID.",
					},
					"user_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "User's Name.",
					},
					"version": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The WARP client version.",
					},
				},
			},
		},
	}
}
