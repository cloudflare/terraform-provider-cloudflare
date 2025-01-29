// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/custom_hostname_fallback_origin"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestCustomHostnameFallbackOriginModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_hostname_fallback_origin.CustomHostnameFallbackOriginModel)(nil)
	schema := custom_hostname_fallback_origin.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
