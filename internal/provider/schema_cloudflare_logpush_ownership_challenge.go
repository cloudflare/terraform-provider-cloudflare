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
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#destination).",
		},
		"ownership_challenge_filename": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The filename of the ownership challenge which	contains the contents required for Logpush Job creation.",
		},
	}
}
