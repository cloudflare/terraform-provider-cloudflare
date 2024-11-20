// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_discovery_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation_discovery"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAPIShieldOperationDiscoveryModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_operation_discovery.APIShieldOperationDiscoveryModel)(nil)
	schema := api_shield_operation_discovery.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
