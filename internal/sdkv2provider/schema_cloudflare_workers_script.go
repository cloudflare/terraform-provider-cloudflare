package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var kvNamespaceBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"namespace_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ID of the KV namespace you want to use.",
		},
	},
}

var plainTextBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"text": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The plain text you want to store.",
		},
	},
}

var secretTextBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"text": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "The secret text you want to store.",
		},
	},
}

var webAssemblyBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"module": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The base64 encoded wasm module you want to store.",
		},
	},
}

var serviceBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"service": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the Worker to bind to.",
		},
		"environment": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the Worker environment to bind to.",
		},
	},
}

var r2BucketBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"bucket_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the Bucket to bind to.",
		},
	},
}

var analyticsEngineBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The global variable for the binding in your Worker code.",
		},
		"dataset": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the Analytics Engine dataset to write to.",
		},
	},
}

var queueBindingResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"binding": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the global variable for the binding in your Worker code.",
		},
		"queue": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the queue you want to use.",
		},
	},
}

func resourceCloudflareWorkerScriptSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The name for the script.",
		},
		"content": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The script content.",
		},
		"module": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to upload Worker as a module.",
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
		"service_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     serviceBindingResource,
		},
		"r2_bucket_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     r2BucketBindingResource,
		},
		"analytics_engine_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     analyticsEngineBindingResource,
		},
		"queue_binding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     queueBindingResource,
		},
	}
}
