package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"account_id"},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"precedence": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"decision": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "non_identity", "bypass"}, false),
		},
		"require": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     AccessGroupOptionSchemaElement,
		},
		"exclude": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     AccessGroupOptionSchemaElement,
		},
		"include": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     AccessGroupOptionSchemaElement,
		},
		"purpose_justification_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"purpose_justification_prompt": {
			Type:         schema.TypeString,
			Optional:     true,
			RequiredWith: []string{"purpose_justification_required"},
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
		},
		"approvals_needed": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntAtLeast(0),
		},
	},
}
