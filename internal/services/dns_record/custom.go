package dns_record

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

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
