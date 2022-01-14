package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareTeamsAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"block_page": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: blockPageSchema,
			},
		},
		"fips": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: fipsSchema,
			},
		},
		"antivirus": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: antivirusSchema,
			},
		},
		"tls_decrypt_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"activity_log_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

var fipsSchema = map[string]*schema.Schema{
	"tls": {
		Type:     schema.TypeBool,
		Optional: true,
	},
}

var blockPageSchema = map[string]*schema.Schema{
	"enabled": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"footer_text": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"header_text": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"logo_path": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"background_color": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"name": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

var antivirusSchema = map[string]*schema.Schema{
	"enabled_download_phase": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"enabled_upload_phase": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"fail_closed": {
		Type:     schema.TypeBool,
		Required: true,
	},
}
