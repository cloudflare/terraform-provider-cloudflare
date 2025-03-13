// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWaitingRoomsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*waiting_room.WaitingRoomsDataSourceModel)(nil)
  schema := waiting_room.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
