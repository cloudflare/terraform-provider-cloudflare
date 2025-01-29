// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_discovery_operation_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/api_shield_discovery_operation"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAPIShieldDiscoveryOperationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_discovery_operation.APIShieldDiscoveryOperationModel)(nil)
	schema := api_shield_discovery_operation.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
