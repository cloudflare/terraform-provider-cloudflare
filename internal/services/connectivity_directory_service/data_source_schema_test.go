// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/connectivity_directory_service"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestConnectivityDirectoryServiceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*connectivity_directory_service.ConnectivityDirectoryServiceDataSourceModel)(nil)
	schema := connectivity_directory_service.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
