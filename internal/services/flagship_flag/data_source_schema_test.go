// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_flag_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/flagship_flag"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestFlagshipFlagDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*flagship_flag.FlagshipFlagDataSourceModel)(nil)
	schema := flagship_flag.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
