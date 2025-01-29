// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicTransitSiteACLsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_transit_site_acl.MagicTransitSiteACLsDataSourceModel)(nil)
	schema := magic_transit_site_acl.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
