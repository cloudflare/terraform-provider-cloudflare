// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_authorities_hostname_associations_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/certificate_authorities_hostname_associations"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCertificateAuthoritiesHostnameAssociationsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*certificate_authorities_hostname_associations.CertificateAuthoritiesHostnameAssociationsModel)(nil)
	schema := certificate_authorities_hostname_associations.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
