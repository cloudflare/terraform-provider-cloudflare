// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_acl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secondary_dns_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecondaryDnsacLsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secondary_dns_acl.SecondaryDNSACLsDataSourceModel)(nil)
	schema := secondary_dns_acl.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
