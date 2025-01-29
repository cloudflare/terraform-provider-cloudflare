// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloudforce_one_request_asset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCloudforceOneRequestAssetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*cloudforce_one_request_asset.CloudforceOneRequestAssetDataSourceModel)(nil)
	schema := cloudforce_one_request_asset.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
