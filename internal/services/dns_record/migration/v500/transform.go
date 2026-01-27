package v500

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

// Transform converts a source cloudflare_record state to target cloudflare_dns_record state.
func Transform(ctx context.Context, source SourceCloudflareRecordModel) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError("Missing required field", "zone_id is required for DNS record migration")
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError("Missing required field", "name is required for DNS record migration")
		return nil, diags
	}
	if source.Type.IsNull() || source.Type.IsUnknown() {
		diags.AddError("Missing required field", "type is required for DNS record migration")
		return nil, diags
	}

	target := &TargetDNSRecordModel{
		ID:      source.ID,
		ZoneID:  source.ZoneID,
		Name:    source.Name,
		Type:    source.Type,
		Comment: source.Comment,
		Proxied: source.Proxied,
	}

	// TTL: source is Int64, target is Float64
	// TTL defaults to 1 (auto) when not specified - this tells Cloudflare to use automatic TTL
	if !source.TTL.IsNull() && !source.TTL.IsUnknown() {
		target.TTL = types.Float64Value(float64(source.TTL.ValueInt64()))
	} else {
		target.TTL = types.Float64Value(1)
	}

	// Priority at root level: source is Int64, target is Float64
	if !source.Priority.IsNull() && !source.Priority.IsUnknown() {
		target.Priority = types.Float64Value(float64(source.Priority.ValueInt64()))
	}

	// Content: source uses "value" for user input, but API responses populate "content"
	// Prefer content if present (already in API format), otherwise use value
	if !source.Content.IsNull() && !source.Content.IsUnknown() && source.Content.ValueString() != "" {
		target.Content = source.Content
	} else if !source.Value.IsNull() && !source.Value.IsUnknown() {
		target.Content = source.Value
	}

	// Tags: source is Set[String], target is customfield.Set[String]
	if !source.Tags.IsNull() && !source.Tags.IsUnknown() {
		var tagValues []string
		diags.Append(source.Tags.ElementsAs(ctx, &tagValues, false)...)
		if !diags.HasError() && len(tagValues) > 0 {
			tagAttrs := make([]attr.Value, len(tagValues))
			for i, t := range tagValues {
				tagAttrs[i] = types.StringValue(t)
			}
			target.Tags = customfield.NewSetMust[types.String](ctx, tagAttrs)
		}
	}

	// Timestamps: source is String (RFC3339), target is timetypes.RFC3339
	if !source.CreatedOn.IsNull() && !source.CreatedOn.IsUnknown() {
		target.CreatedOn = timetypes.NewRFC3339ValueMust(source.CreatedOn.ValueString())
	}
	if !source.ModifiedOn.IsNull() && !source.ModifiedOn.IsUnknown() {
		target.ModifiedOn = timetypes.NewRFC3339ValueMust(source.ModifiedOn.ValueString())
	}

	// Data: source is []SourceDataModel (list block with MaxItems=1), target is *TargetDNSRecordDataModel (single nested object)
	if len(source.Data) > 0 {
		targetData, dataDiags := transformData(source.Data[0], source.Type.ValueString())
		diags.Append(dataDiags...)
		if !diags.HasError() {
			target.Data = targetData
		}
	}

	// Note: The following source fields are NOT migrated (removed or deprecated):
	// - allow_overwrite: Removed in v500
	// - hostname: Computed field, will be refreshed from API
	// - proxiable: Computed field, will be refreshed from API
	// - metadata: Computed field, will be refreshed from API

	return target, diags
}

// transformData converts source data block to target data object.
func transformData(sourceData SourceDataModel, recordType string) (*TargetDNSRecordDataModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	targetData := &TargetDNSRecordDataModel{}

	// Flags: source is String, target is DynamicValue (can be number or string)
	// The API typically returns flags as a number, so we attempt to parse numeric strings
	if !sourceData.Flags.IsNull() && !sourceData.Flags.IsUnknown() {
		flagsStr := sourceData.Flags.ValueString()
		if flagsStr != "" {
			// Try to parse as number first
			if flagsNum, err := strconv.ParseFloat(flagsStr, 64); err == nil {
				targetData.Flags = customfield.RawNormalizedDynamicValueFrom(types.Float64Value(flagsNum))
			} else {
				// Keep as string if not numeric
				targetData.Flags = customfield.RawNormalizedDynamicValueFrom(types.StringValue(flagsStr))
			}
		}
	}

	// Tag (CAA)
	targetData.Tag = sourceData.Tag

	// Value: For CAA records, source uses "content" field, target uses "value"
	// For other record types, the source already uses "value"
	if recordType == "CAA" && !sourceData.Content.IsNull() && !sourceData.Content.IsUnknown() {
		targetData.Value = sourceData.Content
	} else {
		targetData.Value = sourceData.Value
	}

	// SRV fields (strings, same type in source and target)
	targetData.Service = sourceData.Service
	targetData.Target = sourceData.Target

	// Numeric fields: source is Int64, target is Float64
	targetData.Priority = int64ToFloat64(sourceData.Priority)
	targetData.Weight = int64ToFloat64(sourceData.Weight)
	targetData.Port = int64ToFloat64(sourceData.Port)
	targetData.Algorithm = int64ToFloat64(sourceData.Algorithm)
	targetData.KeyTag = int64ToFloat64(sourceData.KeyTag)
	targetData.Type = int64ToFloat64(sourceData.Type)
	targetData.Protocol = int64ToFloat64(sourceData.Protocol)
	targetData.DigestType = int64ToFloat64(sourceData.DigestType)
	targetData.Usage = int64ToFloat64(sourceData.Usage)
	targetData.Selector = int64ToFloat64(sourceData.Selector)
	targetData.MatchingType = int64ToFloat64(sourceData.MatchingType)
	targetData.LatDegrees = int64ToFloat64(sourceData.LatDegrees)
	targetData.LatMinutes = int64ToFloat64(sourceData.LatMinutes)
	targetData.LongDegrees = int64ToFloat64(sourceData.LongDegrees)
	targetData.LongMinutes = int64ToFloat64(sourceData.LongMinutes)
	targetData.Order = int64ToFloat64(sourceData.Order)
	targetData.Preference = int64ToFloat64(sourceData.Preference)

	// Float64 fields (same type in v0 and target)
	targetData.Altitude = sourceData.Altitude
	targetData.LatSeconds = sourceData.LatSeconds
	targetData.LongSeconds = sourceData.LongSeconds
	targetData.PrecisionHorz = sourceData.PrecisionHorz
	targetData.PrecisionVert = sourceData.PrecisionVert
	targetData.Size = sourceData.Size

	// String fields (same type)
	targetData.LatDirection = sourceData.LatDirection
	targetData.LongDirection = sourceData.LongDirection
	targetData.PublicKey = sourceData.PublicKey
	targetData.Digest = sourceData.Digest
	targetData.Certificate = sourceData.Certificate
	targetData.Regex = sourceData.Regex
	targetData.Replacement = sourceData.Replacement
	targetData.Fingerprint = sourceData.Fingerprint

	// Note: The following source data fields are NOT migrated (removed in v500):
	// - proto: Removed from SRV data block
	// - name: Removed from SRV data block

	return targetData, diags
}

// int64ToFloat64 converts a types.Int64 to types.Float64.
// This conversion is safe for DNS record values, which are small integers (ports, priorities, etc.)
// and well within Float64's precision range (53 bits of mantissa).
func int64ToFloat64(v types.Int64) types.Float64 {
	if v.IsNull() {
		return types.Float64Null()
	}
	if v.IsUnknown() {
		return types.Float64Unknown()
	}
	return types.Float64Value(float64(v.ValueInt64()))
}
