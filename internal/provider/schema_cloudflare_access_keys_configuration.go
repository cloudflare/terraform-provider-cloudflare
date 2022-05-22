package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAccessKeysConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"key_rotation_interval_days": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}
}
