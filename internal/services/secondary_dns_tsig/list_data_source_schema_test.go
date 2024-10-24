// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_tsig_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secondary_dns_tsig"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecondaryDnstsiGsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secondary_dns_tsig.SecondaryDNSTSIGsDataSourceModel)(nil)
	schema := secondary_dns_tsig.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
