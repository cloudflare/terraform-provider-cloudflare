// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/dns_record"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestDNSRecordsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_record.DNSRecordsDataSourceModel)(nil)
	schema := dns_record.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
