// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_compliance_modes_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_tls_compliance_modes"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOriginTLSComplianceModesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*origin_tls_compliance_modes.OriginTLSComplianceModesDataSourceModel)(nil)
	schema := origin_tls_compliance_modes.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
