// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/page_shield_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestPageShieldPolicyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*page_shield_policy.PageShieldPolicyDataSourceModel)(nil)
	schema := page_shield_policy.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
