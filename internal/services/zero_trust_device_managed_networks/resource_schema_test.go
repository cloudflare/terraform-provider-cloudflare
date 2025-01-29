// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_device_managed_networks"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustDeviceManagedNetworksModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_device_managed_networks.ZeroTrustDeviceManagedNetworksModel)(nil)
	schema := zero_trust_device_managed_networks.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
