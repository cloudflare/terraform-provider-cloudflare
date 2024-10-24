// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_outgoing_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secondary_dns_outgoing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecondaryDNSOutgoingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secondary_dns_outgoing.SecondaryDNSOutgoingModel)(nil)
	schema := secondary_dns_outgoing.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
