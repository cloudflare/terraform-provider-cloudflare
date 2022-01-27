package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTeamsRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"precedence": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"action": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(cloudflare.TeamsRulesActionValues(), false),
			Required:     true,
		},
		"filters": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"traffic": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"identity": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"device_posture": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"rule_settings": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: teamsRuleSettings,
			},
		},
	}
}

var teamsRuleSettings = map[string]*schema.Schema{
	"block_page_enabled": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"block_page_reason": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"override_ips": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
	"override_host": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"l4override": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsL4OverrideSettings,
		},
	},
	"biso_admin_controls": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsBisoAdminControls,
		},
	},
	"check_session": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsCheckSessionSettings,
		},
	},
	"add_headers": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

var teamsL4OverrideSettings = map[string]*schema.Schema{
	"ip": {
		Type:     schema.TypeString,
		Required: true,
	},
	"port": {
		Type:     schema.TypeInt,
		Required: true,
	},
}

var teamsBisoAdminControls = map[string]*schema.Schema{
	"disable_printing": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"disable_copy_paste": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"disable_download": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"disable_keyboard": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"disable_upload": {
		Type:     schema.TypeBool,
		Optional: true,
	},
}

var teamsCheckSessionSettings = map[string]*schema.Schema{
	"enforce": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"duration": {
		Type:     schema.TypeString,
		Required: true,
	},
}
