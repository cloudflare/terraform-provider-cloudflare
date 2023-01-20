package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var waitingRoomRulesActionValues = []string{
	"bypass_waiting_room",
}

func resourceCloudflareWaitingRoomRulesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"waiting_room_id": {
			Description: "The Waiting Room ID the rules should apply to.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"rules": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of rules to apply to the ruleset.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Unique rule identifier.",
					},
					"version": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Version of the waiting room rule.",
					},
					"status": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
						Description:  fmt.Sprintf("Whether the rule is enabled or disabled. %s", renderAvailableDocumentationValuesStringSlice([]string{"enabled", "disabled"})),
					},
					"action": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice(waitingRoomRulesActionValues, false),
						Description:  fmt.Sprintf("Action to perform in the ruleset rule. %s", renderAvailableDocumentationValuesStringSlice(waitingRoomRulesActionValues)),
					},
					"expression": {
						Description: "Criteria for an HTTP request to trigger the waiting room rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Waiting Room Rules Docs](https://developers.cloudflare.com/waiting-room/additional-options/waiting-room-rules/bypass-rules/)",
						Type:        schema.TypeString,
						Required:    true,
					},
					"description": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Brief summary of the waiting room rule and its intended use.",
					},
				},
			},
		},
	}
}
