package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareLogpushOwnershipChallengeSchema returns the legacy
// cloudflare_logpush_ownership_challenge schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
//
// Schema version: 0 (SDKv2 default - no explicit Version: field in v4 resource)
// Resource type: cloudflare_logpush_ownership_challenge
//
// This schema is minimal - it includes only the properties needed for state parsing:
// Required, Optional, Computed. Validators, PlanModifiers, and Descriptions are
// intentionally omitted.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_logpush_ownership_challenge.go
func SourceCloudflareLogpushOwnershipChallengeSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default - no explicit Version set in v4 provider
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"destination_conf": schema.StringAttribute{
				Required: true,
			},
			"ownership_challenge_filename": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
