package sdkv2provider

import (
	"fmt"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			RequiredWith: []string{"precedence"},
			Deprecated: "This field is deprecated. Policies can now be standalone and reusable by multiple applications." +
				"Please use `cloudflare_access_application.policies` to associate policies with applications.",
			Description: "The ID of the application the policy is associated with.",
		},
		consts.AccountIDSchemaKey: {
			Description:  consts.AccountIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ExactlyOneOf: []string{consts.AccountIDSchemaKey, consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Friendly name of the Access Policy.",
		},
		"precedence": {
			Type:         schema.TypeInt,
			Optional:     true,
			RequiredWith: []string{"application_id"},
			Deprecated: "This field is deprecated. Access policies can now be reusable by multiple applications." +
				" Please use `cloudflare_access_application.policies` to link policies to an application with" +
				" ascending order of precedence.",
			Description: "The unique precedence for policies on a single application.",
		},
		"decision": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "non_identity", "bypass"}, false),
			Description:  fmt.Sprintf("Defines the action Access will take if the policy matches the user. %s", renderAvailableDocumentationValuesStringSlice([]string{"allow", "deny", "non_identity", "bypass"})),
		},
		"require": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        AccessGroupOptionSchemaElement,
			Description: "A series of access conditions, see [Access Groups](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).",
		},
		"exclude": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        AccessGroupOptionSchemaElement,
			Description: "A series of access conditions, see [Access Groups](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).",
		},
		"include": {
			Type:        schema.TypeList,
			Required:    true,
			Elem:        AccessGroupOptionSchemaElement,
			Description: "A series of access conditions, see [Access Groups](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).",
		},
		"isolation_required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Require this application to be served in an isolated browser for users matching this policy.",
		},
		"purpose_justification_required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to prompt the user for a justification for accessing the resource.",
		},
		"purpose_justification_prompt": {
			Type:         schema.TypeString,
			Optional:     true,
			RequiredWith: []string{"purpose_justification_required"},
			Description:  "The prompt to display to the user for a justification for accessing the resource.",
		},
		"approval_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"approval_group": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     AccessPolicyApprovalGroupElement,
		},
		"session_duration": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
				v := val.(string)
				_, err := time.ParseDuration(v)
				if err != nil {
					errs = append(errs, fmt.Errorf(`%q only supports "ns", "us" (or "µs"), "ms", "s", "m", or "h" as valid units`, key))
				}
				return
			},
			Description: "How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`",
		},
	}
}

var AccessPolicyApprovalGroupElement = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"email_list_uuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"email_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of emails to request approval from.",
		},
		"approvals_needed": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntAtLeast(0),
			Description:  "Number of approvals needed.",
		},
	},
}
