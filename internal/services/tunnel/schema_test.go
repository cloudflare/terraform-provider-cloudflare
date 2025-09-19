package tunnel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTunnelSchema_MatchesV4(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)

	// Test that the schema has the expected attributes
	expectedAttributes := []string{
		"id",
		"account_id",
		"config_src",
		"name",
		"tunnel_secret",
		"account_tag",
		"conns_active_at",
		"conns_inactive_at",
		"created_at",
		"deleted_at",
		"remote_config",
		"status",
		"tun_type",
		"cname",
		"tunnel_token",
	}

	for _, attr := range expectedAttributes {
		assert.Contains(t, schema.Attributes, attr, "Schema should contain attribute: %s", attr)
	}

	// Test that required attributes are marked as required
	assert.True(t, schema.Attributes["account_id"].IsRequired(), "account_id should be required")
	assert.True(t, schema.Attributes["name"].IsRequired(), "name should be required")

	// Test that computed attributes are marked as computed
	assert.True(t, schema.Attributes["id"].IsComputed(), "id should be computed")
	assert.True(t, schema.Attributes["account_tag"].IsComputed(), "account_tag should be computed")
	assert.True(t, schema.Attributes["status"].IsComputed(), "status should be computed")
	assert.True(t, schema.Attributes["tun_type"].IsComputed(), "tun_type should be computed")
	assert.True(t, schema.Attributes["cname"].IsComputed(), "cname should be computed")
	assert.True(t, schema.Attributes["tunnel_token"].IsComputed(), "tunnel_token should be computed")

	// Test that sensitive attributes are marked as sensitive
	assert.True(t, schema.Attributes["tunnel_secret"].IsSensitive(), "tunnel_secret should be sensitive")
	assert.True(t, schema.Attributes["tunnel_token"].IsSensitive(), "tunnel_token should be sensitive")
}
