package dns_record

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// normalizeSettings unconditionally materializes the settings object with false defaults
// for any sub-fields that are null. This handles two cases:
//  1. The API returns a settings object but omits some sub-fields (null sub-fields).
//  2. The API omits the settings object entirely (data.Settings.IsNull()), which it does
//     for non-proxied CNAME records where these fields are not applicable.
//
// In both cases we treat absent as false so that state is stable against a config of
// settings = { flatten_cname = false } (or ipv4_only / ipv6_only = false).
func normalizeSettings(ctx context.Context, data *DNSRecordModel, diags *diag.Diagnostics) {
	// settings sub-fields (ipv4_only, ipv6_only, flatten_cname) only apply to CNAME
	// records. For all other record types the API never returns a settings object, so
	// leave state as null to avoid injecting spurious settings into non-CNAME records.
	if data.Type.IsNull() || data.Type.IsUnknown() || data.Type.ValueString() != "CNAME" {
		return
	}

	var s DNSRecordSettingsModel
	if !data.Settings.IsNull() && !data.Settings.IsUnknown() {
		// Populate s from the existing object; ignore errors (s stays zero-value).
		data.Settings.As(ctx, &s, basetypes.ObjectAsOptions{})
	}
	// s sub-fields are null either because Settings was null (API omitted the object
	// entirely for non-proxied CNAME records) or because the API returned the object
	// without those particular sub-fields. Treat absent as false.
	if s.IPV4Only.IsNull() {
		s.IPV4Only = types.BoolValue(false)
	}
	if s.IPV6Only.IsNull() {
		s.IPV6Only = types.BoolValue(false)
	}
	if s.FlattenCNAME.IsNull() {
		s.FlattenCNAME = types.BoolValue(false)
	}
	normalized, d := customfield.NewObject[DNSRecordSettingsModel](ctx, &s)
	diags.Append(d...)
	if !diags.HasError() {
		data.Settings = normalized
	}
}

// Equal compares two DNSRecordDataModel instances for equality
func (d *DNSRecordDataModel) Equal(other *DNSRecordDataModel) bool {
	if d == nil && other == nil {
		return true
	}
	if d == nil || other == nil {
		return false
	}

	// Use semantic equality for Flags to handle Int64(0) == Float64(0.0)
	ctx := context.Background()
	flagsEqual := true
	if !d.Flags.IsNull() || !other.Flags.IsNull() {
		var diags diag.Diagnostics
		flagsEqual, diags = d.Flags.DynamicSemanticEquals(ctx, other.Flags)
		if diags.HasError() {
			flagsEqual = d.Flags.Equal(other.Flags)
		}
	}

	return flagsEqual &&
		d.Tag.Equal(other.Tag) &&
		d.Value.Equal(other.Value) &&
		d.Algorithm.Equal(other.Algorithm) &&
		d.Certificate.Equal(other.Certificate) &&
		d.KeyTag.Equal(other.KeyTag) &&
		d.Type.Equal(other.Type) &&
		d.Protocol.Equal(other.Protocol) &&
		d.PublicKey.Equal(other.PublicKey) &&
		d.Digest.Equal(other.Digest) &&
		d.DigestType.Equal(other.DigestType) &&
		d.Priority.Equal(other.Priority) &&
		d.Target.Equal(other.Target) &&
		d.Altitude.Equal(other.Altitude) &&
		d.LatDegrees.Equal(other.LatDegrees) &&
		d.LatDirection.Equal(other.LatDirection) &&
		d.LatMinutes.Equal(other.LatMinutes) &&
		d.LatSeconds.Equal(other.LatSeconds) &&
		d.LongDegrees.Equal(other.LongDegrees) &&
		d.LongDirection.Equal(other.LongDirection) &&
		d.LongMinutes.Equal(other.LongMinutes) &&
		d.LongSeconds.Equal(other.LongSeconds) &&
		d.PrecisionHorz.Equal(other.PrecisionHorz) &&
		d.PrecisionVert.Equal(other.PrecisionVert) &&
		d.Size.Equal(other.Size) &&
		d.Order.Equal(other.Order) &&
		d.Preference.Equal(other.Preference) &&
		d.Regex.Equal(other.Regex) &&
		d.Replacement.Equal(other.Replacement) &&
		d.Service.Equal(other.Service) &&
		d.MatchingType.Equal(other.MatchingType) &&
		d.Selector.Equal(other.Selector) &&
		d.Usage.Equal(other.Usage) &&
		d.Port.Equal(other.Port) &&
		d.Weight.Equal(other.Weight) &&
		d.Fingerprint.Equal(other.Fingerprint)
}
