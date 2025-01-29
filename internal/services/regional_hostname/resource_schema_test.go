// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/regional_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestRegionalHostnameModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*regional_hostname.RegionalHostnameModel)(nil)
	schema := regional_hostname.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
