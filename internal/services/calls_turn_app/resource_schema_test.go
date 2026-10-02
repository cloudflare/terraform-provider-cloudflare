// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_turn_app_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/calls_turn_app"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCallsTURNAppModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*calls_turn_app.CallsTURNAppModel)(nil)
	schema := calls_turn_app.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	// key_id is intentionally marked Computed in schema.go (custom fix for
	// ESCALATION-10887): the API returns the resource identifier as `uid` in
	// the JSON body but uses `key_id` as the URL path parameter. The generated
	// model retains path:"key_id,optional" while the schema is computed_optional.
	errs.Ignore(t, ".@CallsTURNAppModel.key_id")
	errs.Report(t)
}
