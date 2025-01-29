// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/turnstile_widget"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestTurnstileWidgetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*turnstile_widget.TurnstileWidgetModel)(nil)
	schema := turnstile_widget.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
