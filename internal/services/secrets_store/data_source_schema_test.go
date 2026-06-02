// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecretsStoreDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secrets_store.SecretsStoreDataSourceModel)(nil)
	schema := secrets_store.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
