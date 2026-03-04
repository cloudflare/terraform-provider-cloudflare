// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificate_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/client_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestClientCertificateDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*client_certificate.ClientCertificateDataSourceModel)(nil)
	schema := client_certificate.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
