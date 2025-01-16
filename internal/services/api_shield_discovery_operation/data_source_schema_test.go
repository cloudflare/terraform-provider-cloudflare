// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_discovery_operation_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_discovery_operation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAPIShieldDiscoveryOperationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_discovery_operation.APIShieldDiscoveryOperationDataSourceModel)(nil)
	schema := api_shield_discovery_operation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
