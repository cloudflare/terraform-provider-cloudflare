// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/flagship_app"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestFlagshipAppModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*flagship_app.FlagshipAppModel)(nil)
	schema := flagship_app.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
