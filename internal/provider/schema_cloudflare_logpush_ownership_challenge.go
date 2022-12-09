package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareLogpushOwnershipChallengeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:  "The account identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"zone_id": {
			Description:  "The zone identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"destination_conf": {
			Type:     schema.TypeString,
			ForceNew: true,
			Required: true,
		},
		"ownership_challenge_filename": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
