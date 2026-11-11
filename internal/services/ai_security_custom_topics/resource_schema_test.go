// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security_custom_topics_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ai_security_custom_topics"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAISecurityCustomTopicsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ai_security_custom_topics.AISecurityCustomTopicsModel)(nil)
	schema := ai_security_custom_topics.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
