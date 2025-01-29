// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestByoIPPrefixDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*byo_ip_prefix.ByoIPPrefixDataSourceModel)(nil)
	schema := byo_ip_prefix.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
