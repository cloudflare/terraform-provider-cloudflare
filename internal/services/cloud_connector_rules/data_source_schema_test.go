// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloud_connector_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCloudConnectorRulesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*cloud_connector_rules.CloudConnectorRulesDataSourceModel)(nil)
	schema := cloud_connector_rules.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
