package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareNotificationPolicySchema() map[string]*schema.Schema {
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
			Optional: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"alert_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"filters": notificationPolicyFilterSchema(),
		"created": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"modified": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"email_integration": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     mechanismData,
		},
		"webhooks_integration": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     mechanismData,
		},
		"pagerduty_integration": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     mechanismData,
		},
	}
}

var mechanismData = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

func notificationPolicyFilterSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"status": {
					Type:     schema.TypeSet,
					Elem:     &schema.Schema{Type: schema.TypeString},
					Optional: true,
				},
				"health_check_id": {
					Type:         schema.TypeSet,
					Elem:         &schema.Schema{Type: schema.TypeString},
					Optional:     true,
					RequiredWith: []string{"filters.0.status"},
				},
				"zones": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"services": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"product": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"limit": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"enabled": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"pool_id": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:     true,
					RequiredWith: []string{"filters.0.enabled"},
				},
				"slo": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
			},
		},
	}
}
