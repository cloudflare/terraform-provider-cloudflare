// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package field_extractor_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/field_extractor"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestFieldExtractorDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*field_extractor.FieldExtractorDataSourceModel)(nil)
	schema := field_extractor.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
