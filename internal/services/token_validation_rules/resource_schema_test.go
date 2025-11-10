// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_rules_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/token_validation_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestTokenValidationRulesModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*token_validation_rules.TokenValidationRulesModel)(nil)
	schema := token_validation_rules.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
