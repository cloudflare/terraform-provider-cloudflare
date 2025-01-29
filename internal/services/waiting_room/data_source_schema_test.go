// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/waiting_room"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWaitingRoomDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*waiting_room.WaitingRoomDataSourceModel)(nil)
	schema := waiting_room.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
