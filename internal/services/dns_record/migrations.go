// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudflare/tf-migrate/pkg/state"

	"github.com/tidwall/gjson"
)

var _ resource.ResourceWithUpgradeState = (*DNSRecordResource)(nil)
var _ resource.ResourceWithMoveState = (*DNSRecordResource)(nil)

// MoveState handles moves from cloudflare_record (v4) to cloudflare_dns_record (v5)
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_record.example
//	    to   = cloudflare_dns_record.example
//	}
func (r *DNSRecordResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			// No SourceSchema needed - we use raw JSON transformation
			StateMover: r.moveFromCloudflareRecord,
		},
	}
}

func (r *DNSRecordResource) moveFromCloudflareRecord(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	tflog.Info(ctx, "Starting state move from cloudflare_record to cloudflare_dns_record",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Get raw state as JSON for transformation
	if req.SourceRawState == nil || len(req.SourceRawState.JSON) == 0 {
		resp.Diagnostics.AddError("Failed to read source state", "source raw state is empty")
		return
	}
	sourceJSON := req.SourceRawState.JSON

	tflog.Debug(ctx, "Source state JSON", map[string]interface{}{
		"json": string(sourceJSON),
	})

	// Transform using the same logic as tf-migrate
	transformedJSON, err := state.TransformDNSRecordState(gjson.ParseBytes(sourceJSON))
	if err != nil {
		resp.Diagnostics.AddError("Failed to transform state", err.Error())
		return
	}

	tflog.Debug(ctx, "Transformed state JSON", map[string]interface{}{
		"json": transformedJSON,
	})

	// Get the schema type for parsing JSON
	schemaType := resp.TargetState.Schema.Type().TerraformType(ctx)

	// Parse JSON directly into tftypes.Value using the schema
	val, err := tftypes.ValueFromJSON([]byte(transformedJSON), schemaType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse transformed state", err.Error())
		return
	}

	// Set the raw state directly
	resp.TargetState.Raw = val

	tflog.Info(ctx, "State move from cloudflare_record to cloudflare_dns_record completed successfully")
}

// UpgradeState handles state upgrades for cloudflare_dns_record
// This is triggered when:
//  1. Existing v5 users upgrade (schema_version 0 -> 4, no-op)
//  2. A user manually runs `terraform state mv cloudflare_record.x cloudflare_dns_record.x`
//     (the moved state preserves schema_version=3 from cloudflare_record)
func (r *DNSRecordResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Existing v5 users have schema_version=0 - no transformation needed
		0: {
			StateUpgrader: r.upgradeFromV5,
		},
		// Handle state moved from cloudflare_record (v4 provider, schema_version=3)
		3: {
			StateUpgrader: r.upgradeFromV4,
		},
	}
}

// upgradeFromV5 handles upgrades for existing v5 users (schema_version 0, 1, 2 -> 4)
// This is a no-op since the state is already in v5 format
func (r *DNSRecordResource) upgradeFromV5(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading DNS record state from v5 format (no-op)")

	// Get raw state as JSON
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		resp.Diagnostics.AddError("Failed to read source state", "raw state is empty")
		return
	}

	// Get the schema type for parsing JSON
	schemaType := resp.State.Schema.Type().TerraformType(ctx)

	// Parse JSON directly into tftypes.Value using the schema
	val, err := tftypes.ValueFromJSON(req.RawState.JSON, schemaType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse state", err.Error())
		return
	}

	// Set the raw state directly
	resp.State.Raw = val
}

// upgradeFromV4 handles upgrades from v4 cloudflare_record (schema_version 3 -> 4)
// This is triggered when users manually run `terraform state mv`
func (r *DNSRecordResource) upgradeFromV4(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading DNS record state from v4 cloudflare_record format")

	// Get raw state as JSON
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		resp.Diagnostics.AddError("Failed to read source state", "raw state is empty")
		return
	}
	sourceJSON := req.RawState.JSON

	// Transform using the same logic as tf-migrate
	transformedJSON, err := state.TransformDNSRecordState(gjson.ParseBytes(sourceJSON))
	if err != nil {
		resp.Diagnostics.AddError("Failed to transform state", err.Error())
		return
	}

	// Get the schema type for parsing JSON
	schemaType := resp.State.Schema.Type().TerraformType(ctx)

	// Parse JSON directly into tftypes.Value using the schema
	val, err := tftypes.ValueFromJSON([]byte(transformedJSON), schemaType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse transformed state", err.Error())
		return
	}

	// Set the raw state directly
	resp.State.Raw = val
}
