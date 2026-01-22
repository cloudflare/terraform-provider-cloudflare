// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_hostname_certificate_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_hostname_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAuthenticatedOriginPullsHostnameCertificatesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*authenticated_origin_pulls_hostname_certificate.AuthenticatedOriginPullsHostnameCertificatesDataSourceModel)(nil)
	schema := authenticated_origin_pulls_hostname_certificate.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
