// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/address_map"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAddressMapModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*address_map.AddressMapModel)(nil)
	schema := address_map.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
