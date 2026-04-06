package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/leaked_credential_check_rule/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTransform_EmptyID verifies that Transform preserves an empty ID from v4
// state. This is the critical case: the v4 provider had a bug where the
// detection_id was not stored in state (id = ""). The migration must pass it
// through unchanged so that Read() in the v5 provider can detect it and remove
// the resource from state rather than calling the API with an empty ID.
func TestTransform_EmptyID(t *testing.T) {
	ctx := context.Background()

	source := v500.SourceLeakedCredentialCheckRuleModel{
		ID:       types.StringValue(""), // v4 bug: detection_id not stored
		ZoneID:   types.StringValue("023e105f4ecef8ad9ca31a8372d0c353"),
		Username: types.StringValue(`lookup_json_string(http.request.body.raw, "user")`),
		Password: types.StringValue(`lookup_json_string(http.request.body.raw, "secret")`),
	}

	target, diags := v500.Transform(ctx, source)

	require.False(t, diags.HasError(), "Transform should not error on empty ID: %v", diags)
	require.NotNil(t, target)

	// Empty ID must be preserved — Read() uses this to detect the v4 bug
	// and remove the resource from state instead of calling the API.
	assert.Equal(t, "", target.ID.ValueString(),
		"empty ID must be preserved through migration so Read() can handle it")
	assert.Equal(t, source.ZoneID.ValueString(), target.ZoneID.ValueString())
	assert.Equal(t, source.Username.ValueString(), target.Username.ValueString())
	assert.Equal(t, source.Password.ValueString(), target.Password.ValueString())
}

// TestTransform_PopulatedID verifies that a resource with a valid ID migrates
// correctly — the normal (non-buggy) v4 state path.
func TestTransform_PopulatedID(t *testing.T) {
	ctx := context.Background()

	source := v500.SourceLeakedCredentialCheckRuleModel{
		ID:       types.StringValue("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		ZoneID:   types.StringValue("023e105f4ecef8ad9ca31a8372d0c353"),
		Username: types.StringValue(`lookup_json_string(http.request.body.raw, "user")`),
		Password: types.StringValue(`lookup_json_string(http.request.body.raw, "secret")`),
	}

	target, diags := v500.Transform(ctx, source)

	require.False(t, diags.HasError())
	require.NotNil(t, target)
	assert.Equal(t, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", target.ID.ValueString())
	assert.Equal(t, source.ZoneID.ValueString(), target.ZoneID.ValueString())
}
