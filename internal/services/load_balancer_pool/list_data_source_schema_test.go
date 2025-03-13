// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLoadBalancerPoolsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*load_balancer_pool.LoadBalancerPoolsDataSourceModel)(nil)
  schema := load_balancer_pool.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
