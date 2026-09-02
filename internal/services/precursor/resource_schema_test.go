// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/precursor"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPrecursorModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*precursor.PrecursorModel)(nil)
	schema := precursor.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
