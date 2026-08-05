package v500_test

import (
	"context"
	"testing"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransform_Basic(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	source := v500.SourceCloudflareHyperdriveConfigModel{
		ID:        types.StringValue("test-id-123"),
		AccountID: types.StringValue("account-id-456"),
		Name:      types.StringValue("my-hyperdrive"),
		Origin: &v500.SourceCloudflareHyperdriveConfigOriginModel{
			Database: types.StringValue("mydb"),
			Password: types.StringValue("secret"),
			Host:     types.StringValue("db.example.com"),
			Port:     types.Int64Value(5432),
			Scheme:   types.StringValue("postgres"),
			User:     types.StringValue("admin"),
		},
		Caching: types.ObjectNull(map[string]attr.Type{
			"disabled":               types.BoolType,
			"max_age":                types.Int64Type,
			"stale_while_revalidate": types.Int64Type,
		}),
	}

	target, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	// Verify direct copies
	assertStringValue(t, "id", target.ID, "test-id-123")
	assertStringValue(t, "account_id", target.AccountID, "account-id-456")
	assertStringValue(t, "name", target.Name, "my-hyperdrive")

	// Verify origin fields
	if target.Origin == nil {
		t.Fatal("expected origin to be non-nil")
	}
	assertStringValue(t, "origin.database", target.Origin.Database, "mydb")
	assertStringValue(t, "origin.password", target.Origin.Password, "secret")
	assertStringValue(t, "origin.host", target.Origin.Host, "db.example.com")
	assertInt64Value(t, "origin.port", target.Origin.Port, 5432)
	assertStringValue(t, "origin.scheme", target.Origin.Scheme, "postgres")
	assertStringValue(t, "origin.user", target.Origin.User, "admin")

	// Verify new origin field is null
	assertNull(t, "origin.service_id", target.Origin.ServiceID)

	// Verify new fields are null
	assertNull(t, "origin_connection_limit", target.OriginConnectionLimit)
	if target.MTLS != nil {
		t.Error("expected mtls to be nil")
	}
	if target.Caching != nil {
		t.Error("expected caching to be nil when source is null")
	}
	assertRFC3339Null(t, "created_on", target.CreatedOn)
	assertRFC3339Null(t, "modified_on", target.ModifiedOn)
}

func TestTransform_WithCaching(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	cachingAttrTypes := map[string]attr.Type{
		"disabled":               types.BoolType,
		"max_age":                types.Int64Type,
		"stale_while_revalidate": types.Int64Type,
	}

	source := v500.SourceCloudflareHyperdriveConfigModel{
		ID:        types.StringValue("test-id-123"),
		AccountID: types.StringValue("account-id-456"),
		Name:      types.StringValue("my-hyperdrive"),
		Origin: &v500.SourceCloudflareHyperdriveConfigOriginModel{
			Database: types.StringValue("mydb"),
			Password: types.StringValue("secret"),
			Host:     types.StringValue("db.example.com"),
			Port:     types.Int64Value(5432),
			Scheme:   types.StringValue("postgres"),
			User:     types.StringValue("admin"),
		},
		Caching: types.ObjectValueMust(cachingAttrTypes, map[string]attr.Value{
			"disabled":               types.BoolValue(false),
			"max_age":                types.Int64Value(300),
			"stale_while_revalidate": types.Int64Value(60),
		}),
	}

	target, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	// Verify caching was converted from types.Object to struct pointer
	if target.Caching == nil {
		t.Fatal("expected caching to be non-nil")
	}
	assertBoolValue(t, "caching.disabled", target.Caching.Disabled, false)
	assertInt64Value(t, "caching.max_age", target.Caching.MaxAge, 300)
	assertInt64Value(t, "caching.stale_while_revalidate", target.Caching.StaleWhileRevalidate, 60)
}

func TestTransform_WithAccessOrigin(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	source := v500.SourceCloudflareHyperdriveConfigModel{
		ID:        types.StringValue("test-id-123"),
		AccountID: types.StringValue("account-id-456"),
		Name:      types.StringValue("my-hyperdrive-access"),
		Origin: &v500.SourceCloudflareHyperdriveConfigOriginModel{
			Database:           types.StringValue("mydb"),
			Password:           types.StringValue("secret"),
			Host:               types.StringValue("db.internal.example.com"),
			Port:               types.Int64Null(),
			Scheme:             types.StringValue("postgres"),
			User:               types.StringValue("admin"),
			AccessClientID:     types.StringValue("client-id-abc"),
			AccessClientSecret: types.StringValue("client-secret-xyz"),
		},
		Caching: types.ObjectNull(map[string]attr.Type{
			"disabled":               types.BoolType,
			"max_age":                types.Int64Type,
			"stale_while_revalidate": types.Int64Type,
		}),
	}

	target, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	// Verify access fields are preserved
	assertStringValue(t, "origin.access_client_id", target.Origin.AccessClientID, "client-id-abc")
	assertStringValue(t, "origin.access_client_secret", target.Origin.AccessClientSecret, "client-secret-xyz")
	assertNull(t, "origin.port", target.Origin.Port)
	assertNull(t, "origin.service_id", target.Origin.ServiceID)
}

func TestTransform_MissingAccountID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	source := v500.SourceCloudflareHyperdriveConfigModel{
		ID:        types.StringValue("test-id"),
		AccountID: types.StringNull(),
		Name:      types.StringValue("test"),
	}

	_, diags := v500.Transform(ctx, source)
	if !diags.HasError() {
		t.Fatal("expected error for missing account_id")
	}
}

func TestTransform_MissingName(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	source := v500.SourceCloudflareHyperdriveConfigModel{
		ID:        types.StringValue("test-id"),
		AccountID: types.StringValue("account-id"),
		Name:      types.StringNull(),
	}

	_, diags := v500.Transform(ctx, source)
	if !diags.HasError() {
		t.Fatal("expected error for missing name")
	}
}

func TestTransform_CachingDisabledOnly(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	cachingAttrTypes := map[string]attr.Type{
		"disabled":               types.BoolType,
		"max_age":                types.Int64Type,
		"stale_while_revalidate": types.Int64Type,
	}

	source := v500.SourceCloudflareHyperdriveConfigModel{
		ID:        types.StringValue("test-id"),
		AccountID: types.StringValue("account-id"),
		Name:      types.StringValue("test"),
		Origin: &v500.SourceCloudflareHyperdriveConfigOriginModel{
			Database: types.StringValue("mydb"),
			Password: types.StringValue("secret"),
			Host:     types.StringValue("db.example.com"),
			Scheme:   types.StringValue("postgres"),
			User:     types.StringValue("admin"),
		},
		Caching: types.ObjectValueMust(cachingAttrTypes, map[string]attr.Value{
			"disabled":               types.BoolValue(true),
			"max_age":                types.Int64Null(),
			"stale_while_revalidate": types.Int64Null(),
		}),
	}

	target, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if target.Caching == nil {
		t.Fatal("expected caching to be non-nil")
	}
	assertBoolValue(t, "caching.disabled", target.Caching.Disabled, true)
	assertNull(t, "caching.max_age", target.Caching.MaxAge)
	assertNull(t, "caching.stale_while_revalidate", target.Caching.StaleWhileRevalidate)
}

// Test helpers

func assertStringValue(t *testing.T, field string, got types.String, want string) {
	t.Helper()
	if got.IsNull() || got.IsUnknown() {
		t.Errorf("%s: expected %q, got null/unknown", field, want)
		return
	}
	if got.ValueString() != want {
		t.Errorf("%s: expected %q, got %q", field, want, got.ValueString())
	}
}

func assertInt64Value(t *testing.T, field string, got types.Int64, want int64) {
	t.Helper()
	if got.IsNull() || got.IsUnknown() {
		t.Errorf("%s: expected %d, got null/unknown", field, want)
		return
	}
	if got.ValueInt64() != want {
		t.Errorf("%s: expected %d, got %d", field, want, got.ValueInt64())
	}
}

func assertBoolValue(t *testing.T, field string, got types.Bool, want bool) {
	t.Helper()
	if got.IsNull() || got.IsUnknown() {
		t.Errorf("%s: expected %v, got null/unknown", field, want)
		return
	}
	if got.ValueBool() != want {
		t.Errorf("%s: expected %v, got %v", field, want, got.ValueBool())
	}
}

func assertNull(t *testing.T, field string, got attr.Value) {
	t.Helper()
	if !got.IsNull() {
		t.Errorf("%s: expected null, got %v", field, got)
	}
}

func assertRFC3339Null(t *testing.T, field string, got timetypes.RFC3339) {
	t.Helper()
	if !got.IsNull() {
		t.Errorf("%s: expected null, got %v", field, got)
	}
}
