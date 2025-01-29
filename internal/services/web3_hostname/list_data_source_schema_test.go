// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/web3_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWeb3HostnamesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*web3_hostname.Web3HostnamesDataSourceModel)(nil)
	schema := web3_hostname.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
