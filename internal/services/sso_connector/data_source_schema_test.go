// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/sso_connector"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSSOConnectorDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*sso_connector.SSOConnectorDataSourceModel)(nil)
	schema := sso_connector.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
