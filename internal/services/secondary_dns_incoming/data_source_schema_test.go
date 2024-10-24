// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_incoming_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secondary_dns_incoming"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecondaryDNSIncomingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secondary_dns_incoming.SecondaryDNSIncomingDataSourceModel)(nil)
	schema := secondary_dns_incoming.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
