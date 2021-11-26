package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var kvNamespaceBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"namespace_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var plainTextBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"text": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var secretTextBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"text": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
		},
	},
}

var webAssemblyBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"module": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

func resourceCloudflareWorkerScriptSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"content": {
			Type:     schema.TypeString,
			Required: true,
		},
		"plain_text_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     plainTextBindingResource,
		},
		"secret_text_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     secretTextBindingResource,
		},
		"kv_namespace_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     kvNamespaceBindingResource,
		},
		"webassembly_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     webAssemblyBindingResource,
		},
	}
}
