// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_managed_networks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDeviceManagedNetworksListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_device_managed_networks.ZeroTrustDeviceManagedNetworksListDataSourceModel)(nil)
	schema := zero_trust_device_managed_networks.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
