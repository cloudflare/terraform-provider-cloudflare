// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_config_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/token_validation_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestTokenValidationConfigModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*token_validation_config.TokenValidationConfigModel)(nil)
	schema := token_validation_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
