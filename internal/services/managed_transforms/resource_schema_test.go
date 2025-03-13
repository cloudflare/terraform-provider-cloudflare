// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/managed_transforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestManagedTransformsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*managed_transforms.ManagedTransformsModel)(nil)
	schema := managed_transforms.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
