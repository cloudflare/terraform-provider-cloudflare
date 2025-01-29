// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestHyperdriveConfigDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*hyperdrive_config.HyperdriveConfigDataSourceModel)(nil)
	schema := hyperdrive_config.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
