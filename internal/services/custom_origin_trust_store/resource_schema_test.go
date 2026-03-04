// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_origin_trust_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomOriginTrustStoreModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_origin_trust_store.CustomOriginTrustStoreModel)(nil)
	schema := custom_origin_trust_store.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
