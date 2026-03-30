// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_rule"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_rule/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestPageRuleMigrationModelSchemaParity verifies that TargetV5PageRuleModel
// (used in UpgradeFromVersion0, UpgradeFromV4 → resp.State.Set) stays in sync with the live ResourceSchema.
func TestPageRuleMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetV5PageRuleModel)(nil)
	schema := page_rule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
