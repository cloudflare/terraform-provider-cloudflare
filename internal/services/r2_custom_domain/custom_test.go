/* Unit tests for custom domain false drift when API does not return certain fields due to intermittent errors */
package r2_custom_domain

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// makeStatus creates a NestedObject[R2CustomDomainStatusModel] from the given values.
func makeStatus(t *testing.T, ssl, ownership string) customfield.NestedObject[R2CustomDomainStatusModel] {
	t.Helper()
	ctx := context.Background()
	status := &R2CustomDomainStatusModel{
		SSL:       types.StringValue(ssl),
		Ownership: types.StringValue(ownership),
	}
	obj, diags := customfield.NewObject(ctx, status)
	if diags.HasError() {
		t.Fatalf("failed to create status object: %v", diags)
	}
	return obj
}

// makeNullStatus creates a null NestedObject[R2CustomDomainStatusModel].
func makeNullStatus(t *testing.T) customfield.NestedObject[R2CustomDomainStatusModel] {
	t.Helper()
	ctx := context.Background()
	return customfield.NullObject[R2CustomDomainStatusModel](ctx)
}

// getStatusValues extracts ssl and ownership strings from a status NestedObject.
func getStatusValues(t *testing.T, status customfield.NestedObject[R2CustomDomainStatusModel]) (ssl, ownership string) {
	t.Helper()
	ctx := context.Background()
	s, diags := status.Value(ctx)
	if diags.HasError() {
		t.Fatalf("failed to get status value: %v", diags)
	}
	if s == nil {
		t.Fatal("status value is nil")
	}
	return s.SSL.ValueString(), s.Ownership.ValueString()
}

func TestPreserveStateOnDegradedResponse_StatusPreserved(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	ciphers := []types.String{
		types.StringValue("ECDHE-RSA-AES128-GCM-SHA256"),
		types.StringValue("ECDHE-RSA-AES256-GCM-SHA384"),
	}

	// Simulate: previous state had active/active, API returned unknown/unknown
	data := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneID:   types.StringValue("zone123"),
		ZoneName: types.StringNull(), // API omitted this
		MinTLS:   types.StringNull(), // API omitted this
		Ciphers:  nil,                // API omitted this
		Status:   makeStatus(t, "unknown", "unknown"),
	}

	previousState := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneID:   types.StringValue("zone123"),
		ZoneName: types.StringValue("example.com"),
		MinTLS:   types.StringValue("1.2"),
		Ciphers:  &ciphers,
		Status:   makeStatus(t, "active", "active"),
	}

	preserveStateOnDegradedResponse(ctx, data, previousState)

	// Status should be restored to active/active
	ssl, ownership := getStatusValues(t, data.Status)
	if ssl != "active" {
		t.Errorf("expected status.ssl to be preserved as 'active', got %q", ssl)
	}
	if ownership != "active" {
		t.Errorf("expected status.ownership to be preserved as 'active', got %q", ownership)
	}

	// zone_name, min_tls, and ciphers should be restored
	if data.ZoneName.ValueString() != "example.com" {
		t.Errorf("expected zone_name to be preserved as 'example.com', got %q", data.ZoneName.ValueString())
	}
	if data.MinTLS.ValueString() != "1.2" {
		t.Errorf("expected min_tls to be preserved as '1.2', got %q", data.MinTLS.ValueString())
	}
	if data.Ciphers == nil {
		t.Error("expected ciphers to be preserved, got nil")
	} else if len(*data.Ciphers) != 2 {
		t.Errorf("expected 2 ciphers to be preserved, got %d", len(*data.Ciphers))
	}
}

func TestPreserveStateOnDegradedResponse_NonDegradedNotChanged(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Simulate: API returned a legitimate status change (active -> pending)
	data := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneID:   types.StringValue("zone123"),
		ZoneName: types.StringValue("example.com"),
		MinTLS:   types.StringValue("1.2"),
		Status:   makeStatus(t, "pending", "active"),
	}

	previousState := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneID:   types.StringValue("zone123"),
		ZoneName: types.StringValue("example.com"),
		MinTLS:   types.StringValue("1.2"),
		Status:   makeStatus(t, "active", "active"),
	}

	preserveStateOnDegradedResponse(ctx, data, previousState)

	// Status should NOT be overridden — this is a real status change
	ssl, ownership := getStatusValues(t, data.Status)
	if ssl != "pending" {
		t.Errorf("expected status.ssl to remain 'pending' (real change), got %q", ssl)
	}
	if ownership != "active" {
		t.Errorf("expected status.ownership to remain 'active', got %q", ownership)
	}
}

func TestPreserveStateOnDegradedResponse_NilPreviousState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Simulate: no previous state (first read, e.g. after import)
	data := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneID:   types.StringValue("zone123"),
		ZoneName: types.StringNull(),
		MinTLS:   types.StringNull(),
		Status:   makeStatus(t, "unknown", "unknown"),
	}

	// Should not panic with nil previousState
	preserveStateOnDegradedResponse(ctx, data, nil)

	// Values should remain unchanged (no previous state to restore from)
	ssl, ownership := getStatusValues(t, data.Status)
	if ssl != "unknown" {
		t.Errorf("expected status.ssl to remain 'unknown' (no prior state), got %q", ssl)
	}
	if ownership != "unknown" {
		t.Errorf("expected status.ownership to remain 'unknown' (no prior state), got %q", ownership)
	}
}

func TestPreserveStateOnDegradedResponse_NilData(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	previousState := &R2CustomDomainModel{
		Status: makeStatus(t, "active", "active"),
	}

	// Should not panic with nil data
	preserveStateOnDegradedResponse(ctx, nil, previousState)
}

func TestPreserveStateOnDegradedResponse_PreviousAlsoUnknown(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Simulate: previous state was already unknown (e.g. consecutive failures)
	data := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneName: types.StringNull(),
		MinTLS:   types.StringNull(),
		Status:   makeStatus(t, "unknown", "unknown"),
	}

	previousState := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneName: types.StringNull(),
		MinTLS:   types.StringNull(),
		Status:   makeStatus(t, "unknown", "unknown"),
	}

	preserveStateOnDegradedResponse(ctx, data, previousState)

	// Nothing useful to restore — status stays unknown
	ssl, ownership := getStatusValues(t, data.Status)
	if ssl != "unknown" {
		t.Errorf("expected status.ssl to remain 'unknown' (no better value), got %q", ssl)
	}
	if ownership != "unknown" {
		t.Errorf("expected status.ownership to remain 'unknown' (no better value), got %q", ownership)
	}
}

func TestPreserveStateOnDegradedResponse_OnlyZoneNameOmitted(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Simulate: non-degraded response but zone_name omitted for some reason
	// (e.g. domain in initializing state — status would be error, not unknown)
	data := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneName: types.StringNull(),
		MinTLS:   types.StringValue("1.2"),
		Status:   makeStatus(t, "error", "error"),
	}

	previousState := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneName: types.StringValue("example.com"),
		MinTLS:   types.StringValue("1.2"),
		Status:   makeStatus(t, "active", "active"),
	}

	preserveStateOnDegradedResponse(ctx, data, previousState)

	// Status is not "unknown"/"unknown", so this is NOT a degraded response.
	// zone_name should NOT be restored — the API returned a real (non-degraded) response.
	ssl, ownership := getStatusValues(t, data.Status)
	if ssl != "error" {
		t.Errorf("expected status.ssl to remain 'error', got %q", ssl)
	}
	if ownership != "error" {
		t.Errorf("expected status.ownership to remain 'error', got %q", ownership)
	}
	if !data.ZoneName.IsNull() {
		t.Errorf("expected zone_name to remain null (non-degraded response), got %q", data.ZoneName.ValueString())
	}
}

func TestPreserveStateOnDegradedResponse_NullPreviousStatus(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Simulate: previous state had null status (shouldn't happen normally)
	data := &R2CustomDomainModel{
		Domain: types.StringValue("cdn.example.com"),
		Status: makeStatus(t, "unknown", "unknown"),
	}

	previousState := &R2CustomDomainModel{
		Domain: types.StringValue("cdn.example.com"),
		Status: makeNullStatus(t),
	}

	preserveStateOnDegradedResponse(ctx, data, previousState)

	// Can't restore from null previous status — values stay as-is
	ssl, ownership := getStatusValues(t, data.Status)
	if ssl != "unknown" {
		t.Errorf("expected status.ssl to remain 'unknown' (null previous), got %q", ssl)
	}
	if ownership != "unknown" {
		t.Errorf("expected status.ownership to remain 'unknown' (null previous), got %q", ownership)
	}
}

func TestIsDegradedStatusResponse(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name     string
		ssl      string
		own      string
		expected bool
	}{
		{"both unknown", "unknown", "unknown", true},
		{"ssl unknown only", "unknown", "active", false},
		{"ownership unknown only", "active", "unknown", false},
		{"both active", "active", "active", false},
		{"both error", "error", "error", false},
		{"ssl pending", "pending", "active", false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			data := &R2CustomDomainModel{
				Status: makeStatus(t, tt.ssl, tt.own),
			}
			got := isDegradedStatusResponse(ctx, data)
			if got != tt.expected {
				t.Errorf("isDegradedStatusResponse(ssl=%q, ownership=%q) = %v, want %v",
					tt.ssl, tt.own, got, tt.expected)
			}
		})
	}
}

func TestIsDegradedStatusResponse_NullStatus(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	data := &R2CustomDomainModel{
		Status: makeNullStatus(t),
	}
	if isDegradedStatusResponse(ctx, data) {
		t.Error("expected null status to not be considered degraded")
	}
}

func TestSnapshotState(t *testing.T) {
	t.Parallel()

	original := &R2CustomDomainModel{
		Domain:   types.StringValue("cdn.example.com"),
		ZoneID:   types.StringValue("zone123"),
		ZoneName: types.StringValue("example.com"),
		MinTLS:   types.StringValue("1.2"),
	}

	snapshot := snapshotState(original)

	// Modify original — snapshot should be unaffected
	original.ZoneName = types.StringNull()
	original.MinTLS = types.StringNull()

	if snapshot.ZoneName.ValueString() != "example.com" {
		t.Errorf("snapshot zone_name was affected by original mutation, got %q", snapshot.ZoneName.ValueString())
	}
	if snapshot.MinTLS.ValueString() != "1.2" {
		t.Errorf("snapshot min_tls was affected by original mutation, got %q", snapshot.MinTLS.ValueString())
	}
}

func TestSnapshotState_Nil(t *testing.T) {
	t.Parallel()

	snapshot := snapshotState(nil)
	if snapshot != nil {
		t.Error("expected nil input to return nil snapshot")
	}
}

func TestPreserveOmittedFields(t *testing.T) {
	t.Parallel()

	ciphers1 := []types.String{
		types.StringValue("ECDHE-RSA-AES128-GCM-SHA256"),
		types.StringValue("ECDHE-RSA-AES256-GCM-SHA384"),
	}
	ciphers2 := []types.String{
		types.StringValue("ECDHE-RSA-AES128-SHA256"),
	}

	tests := []struct {
		name             string
		dataZoneName     types.String
		dataMinTLS       types.String
		dataCiphers      *[]types.String
		prevZoneName     types.String
		prevMinTLS       types.String
		prevCiphers      *[]types.String
		expectedZoneName types.String
		expectedMinTLS   types.String
		expectedCiphers  *[]types.String
	}{
		{
			name:             "all null, all restored",
			dataZoneName:     types.StringNull(),
			dataMinTLS:       types.StringNull(),
			dataCiphers:      nil,
			prevZoneName:     types.StringValue("example.com"),
			prevMinTLS:       types.StringValue("1.2"),
			prevCiphers:      &ciphers1,
			expectedZoneName: types.StringValue("example.com"),
			expectedMinTLS:   types.StringValue("1.2"),
			expectedCiphers:  &ciphers1,
		},
		{
			name:             "zone_name null, min_tls present, ciphers nil",
			dataZoneName:     types.StringNull(),
			dataMinTLS:       types.StringValue("1.3"),
			dataCiphers:      nil,
			prevZoneName:     types.StringValue("example.com"),
			prevMinTLS:       types.StringValue("1.2"),
			prevCiphers:      &ciphers1,
			expectedZoneName: types.StringValue("example.com"),
			expectedMinTLS:   types.StringValue("1.3"), // not overwritten, was present
			expectedCiphers:  &ciphers1,
		},
		{
			name:             "all present, nothing changes",
			dataZoneName:     types.StringValue("example.com"),
			dataMinTLS:       types.StringValue("1.2"),
			dataCiphers:      &ciphers2,
			prevZoneName:     types.StringValue("old.com"),
			prevMinTLS:       types.StringValue("1.0"),
			prevCiphers:      &ciphers1,
			expectedZoneName: types.StringValue("example.com"),
			expectedMinTLS:   types.StringValue("1.2"),
			expectedCiphers:  &ciphers2, // not overwritten, was present
		},
		{
			name:             "previous also null, nothing to restore",
			dataZoneName:     types.StringNull(),
			dataMinTLS:       types.StringNull(),
			dataCiphers:      nil,
			prevZoneName:     types.StringNull(),
			prevMinTLS:       types.StringNull(),
			prevCiphers:      nil,
			expectedZoneName: types.StringNull(),
			expectedMinTLS:   types.StringNull(),
			expectedCiphers:  nil,
		},
		{
			name:             "ciphers nil, should restore",
			dataZoneName:     types.StringValue("example.com"),
			dataMinTLS:       types.StringValue("1.2"),
			dataCiphers:      nil,
			prevZoneName:     types.StringValue("example.com"),
			prevMinTLS:       types.StringValue("1.2"),
			prevCiphers:      &ciphers1,
			expectedZoneName: types.StringValue("example.com"),
			expectedMinTLS:   types.StringValue("1.2"),
			expectedCiphers:  &ciphers1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data := &R2CustomDomainModel{
				ZoneName: tt.dataZoneName,
				MinTLS:   tt.dataMinTLS,
				Ciphers:  tt.dataCiphers,
			}
			prev := &R2CustomDomainModel{
				ZoneName: tt.prevZoneName,
				MinTLS:   tt.prevMinTLS,
				Ciphers:  tt.prevCiphers,
			}

			preserveOmittedFields(data, prev)

			if !data.ZoneName.Equal(tt.expectedZoneName) {
				t.Errorf("zone_name: got %v, want %v", data.ZoneName, tt.expectedZoneName)
			}
			if !data.MinTLS.Equal(tt.expectedMinTLS) {
				t.Errorf("min_tls: got %v, want %v", data.MinTLS, tt.expectedMinTLS)
			}
			if (data.Ciphers == nil) != (tt.expectedCiphers == nil) {
				t.Errorf("ciphers nil mismatch: got %v, want %v", data.Ciphers == nil, tt.expectedCiphers == nil)
			}
			if data.Ciphers != nil && tt.expectedCiphers != nil {
				if len(*data.Ciphers) != len(*tt.expectedCiphers) {
					t.Errorf("ciphers length: got %d, want %d", len(*data.Ciphers), len(*tt.expectedCiphers))
				}
			}
		})
	}
}
