// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/authenticated_origin_pulls"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAuthenticatedOriginPullsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*authenticated_origin_pulls.AuthenticatedOriginPullsModel)(nil)
	schema := authenticated_origin_pulls.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
