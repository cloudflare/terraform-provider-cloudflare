// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/email_security_impersonation_registry"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestEmailSecurityImpersonationRegistryDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_impersonation_registry.EmailSecurityImpersonationRegistryDataSourceModel)(nil)
	schema := email_security_impersonation_registry.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
