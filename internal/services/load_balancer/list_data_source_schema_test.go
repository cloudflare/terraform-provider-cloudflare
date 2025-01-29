// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/load_balancer"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestLoadBalancersDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*load_balancer.LoadBalancersDataSourceModel)(nil)
	schema := load_balancer.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
