package provider

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
		"url_browser_isolation_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"logging": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: loggingSchema,
			},
		},
		"proxy": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: proxySchema,
			},
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

var proxySchema = map[string]*schema.Schema{
	"tcp": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"udp": {
		Type:     schema.TypeBool,
		Required: true,
	},
}

var loggingSchema = map[string]*schema.Schema{
	"settings_by_rule_type": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dns": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Required: true,
					Elem: &schema.Resource{
						Schema: loggingEnabledSchema,
					},
				},
				"http": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Required: true,
					Elem: &schema.Resource{
						Schema: loggingEnabledSchema,
					},
				},
				"l4": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Required: true,
					Elem: &schema.Resource{
						Schema: loggingEnabledSchema,
					},
				},
			},
		},
	},
	"redact_pii": {
		Type:     schema.TypeBool,
		Required: true,
	},
}

var loggingEnabledSchema = map[string]*schema.Schema{
	"log_all": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"log_blocks": {
		Type:     schema.TypeBool,
		Required: true,
	},
}
