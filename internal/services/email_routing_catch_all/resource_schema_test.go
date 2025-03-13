// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_catch_all"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailRoutingCatchAllModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*email_routing_catch_all.EmailRoutingCatchAllModel)(nil)
  schema := email_routing_catch_all.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
