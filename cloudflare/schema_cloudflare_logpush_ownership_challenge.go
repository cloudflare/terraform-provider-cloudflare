package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareLogpushOwnershipChallengeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
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
