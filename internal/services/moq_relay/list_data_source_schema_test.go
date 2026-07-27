// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay_test

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/moq_relay"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
	"testing"
)

func TestMoQRelaysDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*moq_relay.MoQRelaysDataSourceModel)(nil)
	schema := moq_relay.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
