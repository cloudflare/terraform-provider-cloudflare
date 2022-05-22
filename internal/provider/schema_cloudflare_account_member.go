package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAccountMemberSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_address": {
			Type:     schema.TypeString,
			Required: true,
		},

		"role_ids": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}
