// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/sso_connector"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSSOConnectorsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*sso_connector.SSOConnectorsDataSourceModel)(nil)
	schema := sso_connector.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
