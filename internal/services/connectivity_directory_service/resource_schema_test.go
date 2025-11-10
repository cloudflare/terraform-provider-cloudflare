// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/connectivity_directory_service"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestConnectivityDirectoryServiceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*connectivity_directory_service.ConnectivityDirectoryServiceModel)(nil)
	schema := connectivity_directory_service.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
