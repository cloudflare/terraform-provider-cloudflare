package access_application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessApplicationSchema_MatchesV4(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)

	// Test that the schema has the expected attributes
	expectedAttributes := []string{
		"id",
		"account_id",
		"zone_id",
		"name",
		"domain",
		"type",
		"session_duration",
		"auto_redirect_to_identity",
		"enable_binding_cookie",
		"allowed_idps",
		"custom_deny_message",
		"custom_deny_url",
		"custom_non_identity_deny_url",
		"logo_url",
		"created_at",
		"updated_at",
	}

	for _, attr := range expectedAttributes {
		assert.Contains(t, schema.Attributes, attr, "Schema should contain attribute: %s", attr)
	}

	// Test that required attributes are marked as required
	assert.True(t, schema.Attributes["name"].IsRequired(), "name should be required")

	// Test that computed attributes are marked as computed
	assert.True(t, schema.Attributes["id"].IsComputed(), "id should be computed")
	assert.True(t, schema.Attributes["domain"].IsComputed(), "domain should be computed")
	assert.True(t, schema.Attributes["type"].IsComputed(), "type should be computed")
	assert.True(t, schema.Attributes["session_duration"].IsComputed(), "session_duration should be computed")
	assert.True(t, schema.Attributes["auto_redirect_to_identity"].IsComputed(), "auto_redirect_to_identity should be computed")
	assert.True(t, schema.Attributes["enable_binding_cookie"].IsComputed(), "enable_binding_cookie should be computed")
	assert.True(t, schema.Attributes["created_at"].IsComputed(), "created_at should be computed")
	assert.True(t, schema.Attributes["updated_at"].IsComputed(), "updated_at should be computed")

	// Test that optional attributes are marked as optional
	assert.True(t, schema.Attributes["account_id"].IsOptional(), "account_id should be optional")
	assert.True(t, schema.Attributes["zone_id"].IsOptional(), "zone_id should be optional")
	assert.True(t, schema.Attributes["domain"].IsOptional(), "domain should be optional")
	assert.True(t, schema.Attributes["type"].IsOptional(), "type should be optional")
	assert.True(t, schema.Attributes["session_duration"].IsOptional(), "session_duration should be optional")
	assert.True(t, schema.Attributes["auto_redirect_to_identity"].IsOptional(), "auto_redirect_to_identity should be optional")
	assert.True(t, schema.Attributes["enable_binding_cookie"].IsOptional(), "enable_binding_cookie should be optional")
	assert.True(t, schema.Attributes["allowed_idps"].IsOptional(), "allowed_idps should be optional")
	assert.True(t, schema.Attributes["custom_deny_message"].IsOptional(), "custom_deny_message should be optional")
	assert.True(t, schema.Attributes["custom_deny_url"].IsOptional(), "custom_deny_url should be optional")
	assert.True(t, schema.Attributes["custom_non_identity_deny_url"].IsOptional(), "custom_non_identity_deny_url should be optional")
	assert.True(t, schema.Attributes["logo_url"].IsOptional(), "logo_url should be optional")
}
