// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_job_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_job"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLogpushJobModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*logpush_job.LogpushJobModel)(nil)
	schema := logpush_job.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
