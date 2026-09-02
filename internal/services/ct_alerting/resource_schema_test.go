// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ct_alerting_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ct_alerting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCTAlertingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ct_alerting.CTAlertingModel)(nil)
	schema := ct_alerting.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
