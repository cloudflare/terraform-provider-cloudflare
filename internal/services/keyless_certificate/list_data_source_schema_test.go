// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/keyless_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestKeylessCertificatesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*keyless_certificate.KeylessCertificatesDataSourceModel)(nil)
	schema := keyless_certificate.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
