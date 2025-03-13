// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSRecordDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*dns_record.DNSRecordDataSourceModel)(nil)
  schema := dns_record.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
