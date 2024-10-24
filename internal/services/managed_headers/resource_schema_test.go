// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/managed_headers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestManagedHeadersModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*managed_headers.ManagedHeadersModel)(nil)
	schema := managed_headers.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
