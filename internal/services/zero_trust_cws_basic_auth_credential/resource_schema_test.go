// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_credential_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_cws_basic_auth_credential"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCWSBasicAuthCredentialModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_cws_basic_auth_credential.ZeroTrustCWSBasicAuthCredentialModel)(nil)
	schema := zero_trust_cws_basic_auth_credential.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
