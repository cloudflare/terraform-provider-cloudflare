package dns_record

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithMoveState = (*DNSRecordResource)(nil)
var _ resource.ResourceWithUpgradeState = (*DNSRecordResource)(nil)

// MoveState handles moves from cloudflare_record (v4) to cloudflare_dns_record (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_record.example
//	    to   = cloudflare_dns_record.example
//	}
func (r *DNSRecordResource) MoveState(ctx context.Context) []resource.StateMover {
	v4Schema := V4CloudflareRecordSchema()
	return []resource.StateMover{
		{
			SourceSchema: &v4Schema,
			StateMover:   r.moveFromCloudflareRecord,
		},
	}
}

func (r *DNSRecordResource) moveFromCloudflareRecord(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Verify source is cloudflare_record from cloudflare provider
	if req.SourceTypeName != "cloudflare_record" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_record to cloudflare_dns_record",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the v4 state using the v4 schema
	var v4State V4CloudflareRecordModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 state to v5 state
	v5State, diags := transformV4ToV5(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the v5 state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5State)...)

	tflog.Info(ctx, "State move from cloudflare_record to cloudflare_dns_record completed successfully")
}

// transformV4ToV5 converts a v4 cloudflare_record state to v5 cloudflare_dns_record state.
func transformV4ToV5(ctx context.Context, v4 V4CloudflareRecordModel) (*DNSRecordModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &DNSRecordModel{
		ID:      v4.ID,
		ZoneID:  v4.ZoneID,
		Name:    v4.Name,
		Type:    v4.Type,
		Comment: v4.Comment,
		Proxied: v4.Proxied,
	}

	// TTL: v4 is Int64, v5 is Float64
	if !v4.TTL.IsNull() && !v4.TTL.IsUnknown() {
		v5.TTL = types.Float64Value(float64(v4.TTL.ValueInt64()))
	} else {
		v5.TTL = types.Float64Value(1) // Default TTL
	}

	// Priority at root level: v4 is Int64, v5 is Float64
	if !v4.Priority.IsNull() && !v4.Priority.IsUnknown() {
		v5.Priority = types.Float64Value(float64(v4.Priority.ValueInt64()))
	}

	// Content: v4 uses "value" for simple records, v5 uses "content"
	// Check if v4 has content already (from API), otherwise use value
	if !v4.Content.IsNull() && !v4.Content.IsUnknown() && v4.Content.ValueString() != "" {
		v5.Content = v4.Content
	} else if !v4.Value.IsNull() && !v4.Value.IsUnknown() {
		v5.Content = v4.Value
	}

	// Tags: v4 is Set[String], v5 is customfield.Set[String]
	if !v4.Tags.IsNull() && !v4.Tags.IsUnknown() {
		var tagValues []string
		diags.Append(v4.Tags.ElementsAs(ctx, &tagValues, false)...)
		if !diags.HasError() && len(tagValues) > 0 {
			tagAttrs := make([]attr.Value, len(tagValues))
			for i, t := range tagValues {
				tagAttrs[i] = types.StringValue(t)
			}
			v5.Tags = customfield.NewSetMust[types.String](ctx, tagAttrs)
		}
	}

	// Timestamps: v4 is String, v5 is timetypes.RFC3339
	if !v4.CreatedOn.IsNull() && !v4.CreatedOn.IsUnknown() {
		v5.CreatedOn = timetypes.NewRFC3339ValueMust(v4.CreatedOn.ValueString())
	}
	if !v4.ModifiedOn.IsNull() && !v4.ModifiedOn.IsUnknown() {
		v5.ModifiedOn = timetypes.NewRFC3339ValueMust(v4.ModifiedOn.ValueString())
	}

	// Data: v4 is []V4DataModel (list block), v5 is *DNSRecordDataModel (nested object)
	if len(v4.Data) > 0 {
		v5Data, dataDiags := transformV4DataToV5(ctx, v4.Data[0], v4.Type.ValueString())
		diags.Append(dataDiags...)
		if !diags.HasError() {
			v5.Data = v5Data
		}
	}

	// Note: The following v4 fields are NOT mapped to v5 (removed/deprecated):
	// - allow_overwrite
	// - hostname
	// - proxiable (computed, will be refreshed)
	// - metadata (computed, will be refreshed)

	return v5, diags
}

// transformV4DataToV5 converts v4 data block to v5 data object.
func transformV4DataToV5(ctx context.Context, v4Data V4DataModel, recordType string) (*DNSRecordDataModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5Data := &DNSRecordDataModel{}

	// Flags: v4 is String, v5 is DynamicValue
	// The API returns flags as a number, so we convert string "0" to number 0
	if !v4Data.Flags.IsNull() && !v4Data.Flags.IsUnknown() {
		flagsStr := v4Data.Flags.ValueString()
		if flagsStr != "" {
			// Try to parse as number
			if flagsNum, err := strconv.ParseFloat(flagsStr, 64); err == nil {
				v5Data.Flags = customfield.RawNormalizedDynamicValueFrom(types.Float64Value(flagsNum))
			} else {
				// Keep as string if not a number
				v5Data.Flags = customfield.RawNormalizedDynamicValueFrom(types.StringValue(flagsStr))
			}
		}
	}

	// Tag (CAA)
	v5Data.Tag = v4Data.Tag

	// Value: v4 CAA uses "content", v5 uses "value"
	if recordType == "CAA" && !v4Data.Content.IsNull() && !v4Data.Content.IsUnknown() {
		v5Data.Value = v4Data.Content
	} else {
		v5Data.Value = v4Data.Value
	}

	// SRV fields
	v5Data.Service = v4Data.Service
	v5Data.Target = v4Data.Target

	// Numeric fields: v4 is Int64, v5 is Float64
	v5Data.Priority = int64ToFloat64(v4Data.Priority)
	v5Data.Weight = int64ToFloat64(v4Data.Weight)
	v5Data.Port = int64ToFloat64(v4Data.Port)
	v5Data.Algorithm = int64ToFloat64(v4Data.Algorithm)
	v5Data.KeyTag = int64ToFloat64(v4Data.KeyTag)
	v5Data.Type = int64ToFloat64(v4Data.Type)
	v5Data.Protocol = int64ToFloat64(v4Data.Protocol)
	v5Data.DigestType = int64ToFloat64(v4Data.DigestType)
	v5Data.Usage = int64ToFloat64(v4Data.Usage)
	v5Data.Selector = int64ToFloat64(v4Data.Selector)
	v5Data.MatchingType = int64ToFloat64(v4Data.MatchingType)
	v5Data.LatDegrees = int64ToFloat64(v4Data.LatDegrees)
	v5Data.LatMinutes = int64ToFloat64(v4Data.LatMinutes)
	v5Data.LongDegrees = int64ToFloat64(v4Data.LongDegrees)
	v5Data.LongMinutes = int64ToFloat64(v4Data.LongMinutes)
	v5Data.Order = int64ToFloat64(v4Data.Order)
	v5Data.Preference = int64ToFloat64(v4Data.Preference)

	// Float64 fields (same type in v4 and v5)
	v5Data.Altitude = v4Data.Altitude
	v5Data.LatSeconds = v4Data.LatSeconds
	v5Data.LongSeconds = v4Data.LongSeconds
	v5Data.PrecisionHorz = v4Data.PrecisionHorz
	v5Data.PrecisionVert = v4Data.PrecisionVert
	v5Data.Size = v4Data.Size

	// String fields (same type)
	v5Data.LatDirection = v4Data.LatDirection
	v5Data.LongDirection = v4Data.LongDirection
	v5Data.PublicKey = v4Data.PublicKey
	v5Data.Digest = v4Data.Digest
	v5Data.Certificate = v4Data.Certificate
	v5Data.Regex = v4Data.Regex
	v5Data.Replacement = v4Data.Replacement
	v5Data.Fingerprint = v4Data.Fingerprint

	// Note: v4 "proto" and "name" in data block are NOT in v5 schema

	return v5Data, diags
}

// int64ToFloat64 converts a types.Int64 to types.Float64.
func int64ToFloat64(v types.Int64) types.Float64 {
	if v.IsNull() {
		return types.Float64Null()
	}
	if v.IsUnknown() {
		return types.Float64Unknown()
	}
	return types.Float64Value(float64(v.ValueInt64()))
}

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}

// UpgradeState handles schema version upgrades for cloudflare_dns_record.
// This is triggered when users manually run `terraform state mv` (Terraform < 1.8).
func (r *DNSRecordResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := V4CloudflareRecordSchema()
	v5Schema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from earlier v5 versions (no schema changes, just version bump)
		0: {
			PriorSchema:   &v5Schema,
			StateUpgrader: r.upgradeFromV5,
		},
		// Handle state moved from cloudflare_record (v4 provider, schema_version=3)
		// When users run `terraform state mv cloudflare_record.x cloudflare_dns_record.x`,
		// the schema_version=3 is preserved, triggering this upgrader.
		3: {
			PriorSchema:   &v4Schema,
			StateUpgrader: r.upgradeFromV4,
		},
	}
}
func (r *DNSRecordResource) upgradeFromV5(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	// No-op upgrade: schema is compatible, just copy state through
	var state DNSRecordModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *DNSRecordResource) upgradeFromV4(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading DNS record state from v4 cloudflare_record format (schema version 3)")

	// Parse the v4 state
	var v4State V4CloudflareRecordModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to v5
	v5State, diags := transformV4ToV5(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)

	tflog.Info(ctx, "State upgrade from v4 cloudflare_record completed successfully")
}
