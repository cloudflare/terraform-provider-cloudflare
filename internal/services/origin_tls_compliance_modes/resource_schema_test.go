// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_compliance_modes_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_tls_compliance_modes"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOriginTLSComplianceModesModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*origin_tls_compliance_modes.OriginTLSComplianceModesModel)(nil)
	schema := origin_tls_compliance_modes.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
