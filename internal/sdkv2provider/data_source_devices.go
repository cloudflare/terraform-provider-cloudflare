package sdkv2provider

import (
	"context"
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareDevices() *schema.Resource {
	return &schema.Resource{
		Schema:      resoureceCloudflareDevicesSchema(),
		ReadContext: dataResourceCloudflareDevicesRead,
		Description: "Use this data source to lookup [Devices](https://api.cloudflare.com/#devices-list-devices).",
	}
}

func dataResourceCloudflareDevicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	d.SetId(accountID)

	devices, err := client.ListTeamsDevices(ctx, accountID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding devices in account %q: %w", accountID, err))
	}

	deviceDetails := make([]interface{}, 0)

	for _, device := range devices {
		deviceDetails = append(deviceDetails, map[string]interface{}{
			"id":                 device.ID,
			"key":                device.Key,
			"device_type":        device.DeviceType,
			"name":               device.Name,
			"version":            device.Version,
			"updated":            device.Updated,
			"created":            device.Created,
			"last_seen":          device.LastSeen,
			"revoked_at":         device.RevokedAt,
			"model":              device.Model,
			"os_version":         device.OSVersion,
			"os_distro_name":     device.OSDistroName,
			"os_distro_revision": device.OsDistroRevision,
			"ip":                 device.IP,
			"mac_address":        device.MacAddress,
			"serial_number":      device.SerialNumber,
			"user_id":            device.User.ID,
			"user_email":         device.User.Email,
			"user_name":          device.User.Name,
		})
	}

	if err = d.Set("devices", deviceDetails); err != nil {
		return diag.FromErr(fmt.Errorf("error setting device details: %w", err))
	}

	return nil
}
