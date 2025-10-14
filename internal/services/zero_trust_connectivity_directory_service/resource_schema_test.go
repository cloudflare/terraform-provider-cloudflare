// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_directory_service_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_connectivity_directory_service"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustConnectivityDirectoryServiceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_connectivity_directory_service.ZeroTrustConnectivityDirectoryServiceModel)(nil)
	schema := zero_trust_connectivity_directory_service.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
