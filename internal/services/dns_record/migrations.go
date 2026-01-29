package dns_record

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record/migration/v500"
)

var _ resource.ResourceWithMoveState = (*DNSRecordResource)(nil)
var _ resource.ResourceWithUpgradeState = (*DNSRecordResource)(nil)

// MoveState handles moves from cloudflare_record (v0) to cloudflare_dns_record (v500).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_record.example
//	    to   = cloudflare_dns_record.example
//	}
func (r *DNSRecordResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareRecordSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_dns_record.
// This is triggered when users manually run `terraform state mv` (Terraform < 1.8).
func (r *DNSRecordResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareRecordSchema()
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from earlier v500 versions (no schema changes, just version bump)
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle state moved from legacy cloudflare_record (schema_version=3 from the SDKv2 provider)
		// When users run `terraform state mv cloudflare_record.x cloudflare_dns_record.x`,
		// the schema_version=3 is preserved, triggering this upgrader.
		3: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV3,
		},
		// Handle upgrades from v5.x.x releases that used schema_version=4
		// All v5 releases up to 5.16.0 used Version: 4, we changed it to Version: 500
		4: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}
