// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security_custom_topics_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ai_security_custom_topics"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAISecurityCustomTopicsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ai_security_custom_topics.AISecurityCustomTopicsDataSourceModel)(nil)
	schema := ai_security_custom_topics.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
