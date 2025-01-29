// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_secret"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersSecretModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_secret.WorkersSecretModel)(nil)
	schema := workers_secret.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
