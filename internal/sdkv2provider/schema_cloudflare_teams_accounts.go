package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTeamsAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"block_page": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Configuration for a custom block page.",
			Elem: &schema.Resource{
				Schema: blockPageSchema,
			},
		},
		"body_scanning": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Configuration for body scanning.",
			Elem: &schema.Resource{
				Schema: bodyScanningSchema,
			},
		},
		"fips": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: fipsSchema,
			},
			Description: "Configure compliance with Federal Information Processing Standards.",
		},
		"antivirus": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: antivirusSchema,
			},
			Description: "Configuration block for antivirus traffic scanning.",
		},
		"tls_decrypt_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicator that decryption of TLS traffic is enabled.",
		},
		"protocol_detection_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicator that protocol detection is enabled.",
		},
		"activity_log_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to enable the activity log.",
		},
		"url_browser_isolation_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Safely browse websites in Browser Isolation through a URL.",
		},
		"non_identity_browser_isolation_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enable non-identity onramp for Browser Isolation.",
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
			Description: "Configuration block for specifying which protocols are proxied.",
		},
		"ssh_session_log": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: sshSessionLogSchema,
			},
			Description: "Configuration for SSH Session Logging.",
		},
		"payload_log": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: payloadLogSchema,
			},
			Description: "Configuration for DLP Payload Logging.",
		},
	}
}

var fipsSchema = map[string]*schema.Schema{
	"tls": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Only allow FIPS-compliant TLS configuration.",
	},
}

var blockPageSchema = map[string]*schema.Schema{
	"enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Indicator of enablement.",
	},
	"footer_text": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Block page footer text.",
	},
	"header_text": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Block page header text.",
	},
	"logo_path": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL of block page logo.",
	},
	"background_color": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Hex code of block page background color.",
	},
	"name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of block page configuration.",
	},
	"mailto_address": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Admin email for users to contact.",
	},
	"mailto_subject": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Subject line for emails created from block page.",
	},
}

var inspectionModeOptions = []string{"deep", "shallow"}
var bodyScanningSchema = map[string]*schema.Schema{
	"inspection_mode": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(inspectionModeOptions, false),
		Description: fmt.Sprintf(
			"Body scanning inspection mode. %s",
			renderAvailableDocumentationValuesStringSlice(inspectionModeOptions),
		),
	},
}

var antivirusSchema = map[string]*schema.Schema{
	"enabled_download_phase": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Scan on file download.",
	},
	"enabled_upload_phase": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Scan on file upload.",
	},
	"fail_closed": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Block requests for files that cannot be scanned.",
	},
}

var proxySchema = map[string]*schema.Schema{
	"tcp": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Whether gateway proxy is enabled on gateway devices for TCP traffic.",
	},
	"udp": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Whether gateway proxy is enabled on gateway devices for UDP traffic.",
	},
	"root_ca": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Whether root ca is enabled account wide for ZT clients.",
	},
}

var loggingSchema = map[string]*schema.Schema{
	"settings_by_rule_type": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Required:    true,
		Description: "Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dns": {
					Type:        schema.TypeList,
					MaxItems:    1,
					Required:    true,
					Description: "Logging configuration for DNS requests.",
					Elem: &schema.Resource{
						Schema: loggingEnabledSchema,
					},
				},
				"http": {
					Type:        schema.TypeList,
					MaxItems:    1,
					Required:    true,
					Description: "Logging configuration for HTTP requests.",
					Elem: &schema.Resource{
						Schema: loggingEnabledSchema,
					},
				},
				"l4": {
					Type:        schema.TypeList,
					MaxItems:    1,
					Required:    true,
					Description: "Logging configuration for layer 4 requests.",
					Elem: &schema.Resource{
						Schema: loggingEnabledSchema,
					},
				},
			},
		},
	},
	"redact_pii": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).",
	},
}

var loggingEnabledSchema = map[string]*schema.Schema{
	"log_all": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Whether to log all activity.",
	},
	"log_blocks": {
		Type:     schema.TypeBool,
		Required: true,
	},
}

var sshSessionLogSchema = map[string]*schema.Schema{
	"public_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Public key used to encrypt ssh session.",
	},
}

var payloadLogSchema = map[string]*schema.Schema{
	"public_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Public key used to encrypt matched payloads.",
	},
}
