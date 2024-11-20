// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_discovery_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation_discovery"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAPIShieldOperationDiscoveryDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_operation_discovery.APIShieldOperationDiscoveryDataSourceModel)(nil)
	schema := api_shield_operation_discovery.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
