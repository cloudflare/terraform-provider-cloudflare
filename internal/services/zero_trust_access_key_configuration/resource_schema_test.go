// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_access_key_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustAccessKeyConfigurationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_key_configuration.ZeroTrustAccessKeyConfigurationModel)(nil)
	schema := zero_trust_access_key_configuration.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
