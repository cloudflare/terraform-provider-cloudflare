// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_dynamic_routing_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ai_gateway_dynamic_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAIGatewayDynamicRoutingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ai_gateway_dynamic_routing.AIGatewayDynamicRoutingModel)(nil)
	schema := ai_gateway_dynamic_routing.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
