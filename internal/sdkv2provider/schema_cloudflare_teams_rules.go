package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTeamsRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the teams rule.",
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The description of the teams rule.",
		},
		"precedence": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The evaluation precedence of the teams rule.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicator of rule enablement.",
		},
		"action": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(cloudflare.TeamsRulesActionValues(), false),
			Required:     true,
			Description:  fmt.Sprintf("The action executed by matched teams rule. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.TeamsRulesActionValues())),
		},
		"filters": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "The protocol or layer to evaluate the traffic and identity expressions.",
		},
		"traffic": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The wirefilter expression to be used for traffic matching.",
		},
		"identity": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The wirefilter expression to be used for identity matching.",
		},
		"device_posture": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The wirefilter expression to be used for device_posture check matching.",
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
			Description: "Additional rule settings.",
		},
	}
}

var teamsRuleSettings = map[string]*schema.Schema{
	"block_page_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Indicator of block page enablement.",
	},
	"block_page_reason": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "The displayed reason for a user being blocked.",
	},
	"override_ips": {
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "The IPs to override matching DNS queries with.",
	},
	"override_host": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "The host to override matching DNS queries with.",
	},
	"ip_categories": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Turns on IP category based filter on dns if the rule contains dns category checks.",
	},
	"allow_child_bypass": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Allow parent MSP accounts to enable bypass their children's rules.",
	},
	"bypass_parent_rule": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Allow child MSP accounts to bypass their parent's rule.",
	},
	"audit_ssh": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsAuditSSHSettings,
		},
		Description: "Settings for auditing SSH usage.",
	},
	"l4override": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsL4OverrideSettings,
		},
		Description: "Settings to forward layer 4 traffic.",
	},
	"biso_admin_controls": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsBisoAdminControls,
		},
		Description: "Configure how browser isolation behaves.",
	},
	"check_session": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsCheckSessionSettings,
		},
		Description: "Configure how session check behaves.",
	},
	"add_headers": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "Add custom headers to allowed requests in the form of key-value pairs.",
	},
	"insecure_disable_dnssec_validation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Disable DNSSEC validation (must be Allow rule).",
	},
	"egress": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: egressSettings,
		},
		Description: "Configure how Proxy traffic egresses. Can be set for rules with Egress action and Egress filter. Can be omitted to indicate local egress via Warp IPs.",
	},
	"untrusted_cert": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: untrustedCertSettings,
		},
		Description: "Configure untrusted certificate settings for this rule.",
	},
	"payload_log": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: payloadLogSettings,
		},
		Description: "Configure DLP Payload Logging settings for this rule.",
	},
}

var payloadLogSettings = map[string]*schema.Schema{
	"enabled": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Enable or disable DLP Payload Logging for this rule.",
	},
}

var untrustedCertSettings = map[string]*schema.Schema{
	"action": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice(cloudflare.TeamsRulesUntrustedCertActionValues(), false),
		Optional:     true,
		Description:  fmt.Sprintf("Action to be taken when the SSL certificate of upstream is invalid. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.TeamsRulesUntrustedCertActionValues())),
	},
}

var egressSettings = map[string]*schema.Schema{
	"ipv6": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The IPv6 range to be used for egress.",
	},
	"ipv4": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The IPv4 address to be used for egress.",
	},
	"ipv4_fallback": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "The IPv4 address to be used for egress in the event of an error egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egreass via Warp IPs.",
	},
}

var teamsL4OverrideSettings = map[string]*schema.Schema{
	"ip": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Override IP to forward traffic to.",
	},
	"port": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Override Port to forward traffic to.",
	},
}

var teamsAuditSSHSettings = map[string]*schema.Schema{
	"command_logging": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Log all SSH commands.",
	},
}

var teamsBisoAdminControls = map[string]*schema.Schema{
	"disable_printing": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Disable printing.",
	},
	"disable_copy_paste": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Disable copy-paste.",
	},
	"disable_download": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Disable download.",
	},
	"disable_keyboard": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Disable keyboard usage.",
	},
	"disable_upload": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Disable upload.",
	},
}

var teamsCheckSessionSettings = map[string]*schema.Schema{
	"enforce": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Enable session enforcement for this rule.",
	},
	"duration": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Configure how fresh the session needs to be to be considered valid.",
	},
}
