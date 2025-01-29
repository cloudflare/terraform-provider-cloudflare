// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/load_balancer_monitor"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestLoadBalancerMonitorModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*load_balancer_monitor.LoadBalancerMonitorModel)(nil)
	schema := load_balancer_monitor.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
