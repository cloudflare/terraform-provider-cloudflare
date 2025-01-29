// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/email_routing_address"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestEmailRoutingAddressModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_routing_address.EmailRoutingAddressModel)(nil)
	schema := email_routing_address.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
