// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_tiered_caching"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestArgoTieredCachingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*argo_tiered_caching.ArgoTieredCachingDataSourceModel)(nil)
	schema := argo_tiered_caching.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
