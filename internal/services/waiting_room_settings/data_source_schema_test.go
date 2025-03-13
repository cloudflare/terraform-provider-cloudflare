// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_settings_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_settings"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWaitingRoomSettingsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*waiting_room_settings.WaitingRoomSettingsDataSourceModel)(nil)
  schema := waiting_room_settings.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
