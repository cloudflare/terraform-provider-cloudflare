// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_schema_validation_settings_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/api_shield_operation_schema_validation_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAPIShieldOperationSchemaValidationSettingsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_operation_schema_validation_settings.APIShieldOperationSchemaValidationSettingsModel)(nil)
	schema := api_shield_operation_schema_validation_settings.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
