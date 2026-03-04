// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate_associations_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/mtls_certificate_associations"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMTLSCertificateAssociationsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*mtls_certificate_associations.MTLSCertificateAssociationsDataSourceModel)(nil)
	schema := mtls_certificate_associations.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
