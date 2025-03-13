// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_event"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWaitingRoomEventModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*waiting_room_event.WaitingRoomEventModel)(nil)
  schema := waiting_room_event.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
