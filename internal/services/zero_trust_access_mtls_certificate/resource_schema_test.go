// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_certificate_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_access_mtls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustAccessMTLSCertificateModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_mtls_certificate.ZeroTrustAccessMTLSCertificateModel)(nil)
	schema := zero_trust_access_mtls_certificate.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
