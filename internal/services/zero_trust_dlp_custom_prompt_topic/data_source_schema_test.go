// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_prompt_topic_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_custom_prompt_topic"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDLPCustomPromptTopicDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dlp_custom_prompt_topic.ZeroTrustDLPCustomPromptTopicDataSourceModel)(nil)
	schema := zero_trust_dlp_custom_prompt_topic.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
