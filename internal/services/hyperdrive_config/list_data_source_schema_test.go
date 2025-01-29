// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/hyperdrive_config"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestHyperdriveConfigsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*hyperdrive_config.HyperdriveConfigsDataSourceModel)(nil)
	schema := hyperdrive_config.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
