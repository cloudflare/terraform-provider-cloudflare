// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkerVersionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*worker_version.WorkerVersionModel)(nil)
	schema := worker_version.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	// Schema must be Computed+Optional for proper Terraform behavior, but model uses "optional" (not "computed_optional")
	// because apijson.UnmarshalComputed panics on null dynamic values. Default is set via setRunWorkerFirstDefault().
	errs.Ignore(t, ".@WorkerVersionModel.assets.@WorkerVersionAssetsModel.config.@WorkerVersionAssetsConfigModel.run_worker_first")
	errs.Report(t)
}
