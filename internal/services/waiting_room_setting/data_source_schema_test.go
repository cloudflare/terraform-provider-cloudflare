// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_setting_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWaitingRoomSettingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*waiting_room_setting.WaitingRoomSettingDataSourceModel)(nil)
	schema := waiting_room_setting.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
