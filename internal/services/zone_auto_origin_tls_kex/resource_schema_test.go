// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_auto_origin_tls_kex_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_auto_origin_tls_kex"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneAutoOriginTLSKexModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_auto_origin_tls_kex.ZoneAutoOriginTLSKexModel)(nil)
	schema := zone_auto_origin_tls_kex.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
