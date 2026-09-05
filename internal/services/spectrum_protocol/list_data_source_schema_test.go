// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_protocol_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/spectrum_protocol"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSpectrumProtocolsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*spectrum_protocol.SpectrumProtocolsDataSourceModel)(nil)
	schema := spectrum_protocol.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
