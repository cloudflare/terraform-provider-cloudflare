// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/cloudforce_one_request_asset"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestCloudforceOneRequestAssetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*cloudforce_one_request_asset.CloudforceOneRequestAssetModel)(nil)
	schema := cloudforce_one_request_asset.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
