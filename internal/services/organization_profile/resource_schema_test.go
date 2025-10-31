// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_profile_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/organization_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOrganizationProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*organization_profile.OrganizationProfileModel)(nil)
	schema := organization_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
