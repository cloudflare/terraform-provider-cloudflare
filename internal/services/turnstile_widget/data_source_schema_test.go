// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/turnstile_widget"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestTurnstileWidgetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*turnstile_widget.TurnstileWidgetDataSourceModel)(nil)
	schema := turnstile_widget.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
