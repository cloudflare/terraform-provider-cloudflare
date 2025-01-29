// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_deployment"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersDeploymentDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_deployment.WorkersDeploymentDataSourceModel)(nil)
	schema := workers_deployment.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
