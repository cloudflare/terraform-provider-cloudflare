// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

<<<<<<<< HEAD:internal/services/monetization_rules/data_source_schema_test.go
package monetization_rules_test
========
package mtls_certificate_test
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source_schema_test.go

import (
	"context"
	"testing"

<<<<<<<< HEAD:internal/services/monetization_rules/data_source_schema_test.go
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/monetization_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMonetizationRulesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*monetization_rules.MonetizationRulesDataSourceModel)(nil)
	schema := monetization_rules.DataSourceSchema(context.TODO())
========
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/mtls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMTLSCertificateDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*mtls_certificate.MTLSCertificateDataSourceModel)(nil)
	schema := mtls_certificate.DataSourceSchema(context.TODO())
>>>>>>>> db7d2bac00 (Apply custom code):internal/services/mtls_certificate/data_source_schema_test.go
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
