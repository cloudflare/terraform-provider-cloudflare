package access_policy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessPolicySchema_MatchesV4(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)

	// Test that the schema has the expected attributes
	expectedAttributes := []string{
		"id",
		"account_id",
		"zone_id",
		"application_id",
		"decision",
		"name",
		"precedence",
		"approval_required",
		"isolation_required",
		"purpose_justification_prompt",
		"purpose_justification_required",
		"session_duration",
		"exclude",
		"include",
		"require",
	}

	for _, attr := range expectedAttributes {
		assert.Contains(t, schema.Attributes, attr, "Schema should contain attribute: %s", attr)
	}

	// Test that required attributes are marked as required
	assert.True(t, schema.Attributes["application_id"].IsRequired(), "application_id should be required")
	assert.True(t, schema.Attributes["decision"].IsRequired(), "decision should be required")
	assert.True(t, schema.Attributes["name"].IsRequired(), "name should be required")
	assert.True(t, schema.Attributes["precedence"].IsRequired(), "precedence should be required")
	assert.True(t, schema.Attributes["include"].IsRequired(), "include should be required")

	// Test that computed attributes are marked as computed
	assert.True(t, schema.Attributes["id"].IsComputed(), "id should be computed")
	assert.True(t, schema.Attributes["session_duration"].IsComputed(), "session_duration should be computed")

	// Test that optional attributes are marked as optional
	assert.True(t, schema.Attributes["account_id"].IsOptional(), "account_id should be optional")
	assert.True(t, schema.Attributes["zone_id"].IsOptional(), "zone_id should be optional")
	assert.True(t, schema.Attributes["approval_required"].IsOptional(), "approval_required should be optional")
	assert.True(t, schema.Attributes["isolation_required"].IsOptional(), "isolation_required should be optional")
	assert.True(t, schema.Attributes["purpose_justification_prompt"].IsOptional(), "purpose_justification_prompt should be optional")
	assert.True(t, schema.Attributes["purpose_justification_required"].IsOptional(), "purpose_justification_required should be optional")
	assert.True(t, schema.Attributes["exclude"].IsOptional(), "exclude should be optional")
	assert.True(t, schema.Attributes["require"].IsOptional(), "require should be optional")
}
