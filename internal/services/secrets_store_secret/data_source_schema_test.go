// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_secret_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secrets_store_secret"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecretsStoreSecretDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secrets_store_secret.SecretsStoreSecretDataSourceModel)(nil)
	schema := secrets_store_secret.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
