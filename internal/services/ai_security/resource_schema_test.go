// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ai_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAISecurityModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ai_security.AISecurityModel)(nil)
	schema := ai_security.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
