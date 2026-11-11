// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_finding_type_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_casb_finding_type"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCasbFindingTypeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_casb_finding_type.ZeroTrustCasbFindingTypeModel)(nil)
	schema := zero_trust_casb_finding_type.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
