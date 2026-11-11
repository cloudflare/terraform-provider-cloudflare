// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_cws_tenant_structure"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCWSTenantStructureDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_cws_tenant_structure.ZeroTrustCWSTenantStructureDataSourceModel)(nil)
	schema := zero_trust_cws_tenant_structure.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
