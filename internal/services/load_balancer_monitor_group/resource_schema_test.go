// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_group_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLoadBalancerMonitorGroupModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*load_balancer_monitor_group.LoadBalancerMonitorGroupModel)(nil)
	schema := load_balancer_monitor_group.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
