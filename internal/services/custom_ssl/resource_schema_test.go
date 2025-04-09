// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomSSLModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_ssl.CustomSSLModel)(nil)
	schema := custom_ssl.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
