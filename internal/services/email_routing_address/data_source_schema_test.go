// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_address"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailRoutingAddressDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_routing_address.EmailRoutingAddressDataSourceModel)(nil)
	schema := email_routing_address.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
