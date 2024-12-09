// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_smart_routing_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_smart_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestArgoSmartRoutingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*argo_smart_routing.ArgoSmartRoutingModel)(nil)
	schema := argo_smart_routing.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}