// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_rego_metadata_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_cws_rego_metadata"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCWSRegoMetadataDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_cws_rego_metadata.ZeroTrustCWSRegoMetadataDataSourceModel)(nil)
	schema := zero_trust_cws_rego_metadata.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
