// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_profile_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/organization_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOrganizationProfileDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*organization_profile.OrganizationProfileDataSourceModel)(nil)
	schema := organization_profile.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
