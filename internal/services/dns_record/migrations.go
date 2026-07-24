package dns_record

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record/migration/v501"
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
	v500Schema := sourceSchemaV500(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle state written by earlier v5 releases that used schema version 0
		// with the same shape as v500.
		0: {
			PriorSchema:   &v500Schema,
			StateUpgrader: v501.UpgradeFromV0,
		},
		// Handle state moved from legacy cloudflare_record (schema_version=3 from the SDKv2 provider)
		// When users run `terraform state mv cloudflare_record.x cloudflare_dns_record.x`,
		// the schema_version=3 is preserved, triggering this upgrader.
		3: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV3,
		},
		500: {
			PriorSchema:   &v500Schema,
			StateUpgrader: v501.UpgradeFromV500,
		},
	}
}

// sourceSchemaV500 freezes the only structural difference between v500 and
// v501: priority was configurable at the resource root. All other attributes
// retain the same state representation.
func sourceSchemaV500(ctx context.Context) schema.Schema {
	prior := ResourceSchema(ctx)
	prior.Version = 500
	prior.Attributes["priority"] = schema.Float64Attribute{Optional: true}
	return prior
}
