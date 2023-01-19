package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resoureceCloudflareDevicesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"devices": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Device ID.",
					},
					"key": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device's public key.",
					},
					"device_type": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The type of the device.",
					},
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device name.",
					},
					"version": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The WARP client version.",
					},
					"model": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The device model name.",
					},
					"os_version": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The operating system version.",
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
					"ip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "IPv4 or IPv6 address.",
					},
					"last_seen": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was last seen.",
					},
					"created": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was created.",
					},
					"updated": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "When the device was updated.",
					},
					"user_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "User's ID.",
					},
					"user_email": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "User's email.",
					},
					"user_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "User's Name.",
					},
				},
			},
		},
	}
}
