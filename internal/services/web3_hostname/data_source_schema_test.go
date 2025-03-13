// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/web3_hostname"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWeb3HostnameDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*web3_hostname.Web3HostnameDataSourceModel)(nil)
  schema := web3_hostname.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
