// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/healthcheck"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestHealthcheckModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*healthcheck.HealthcheckModel)(nil)
	schema := healthcheck.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
