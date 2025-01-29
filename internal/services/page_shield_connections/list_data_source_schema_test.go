// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_connections_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/page_shield_connections"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestPageShieldConnectionsListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*page_shield_connections.PageShieldConnectionsListDataSourceModel)(nil)
	schema := page_shield_connections.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
