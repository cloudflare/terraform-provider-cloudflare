package log_explorer_dataset_test

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
	"testing"
)

func TestLogExplorerDatasetsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*log_explorer_dataset.LogExplorerDatasetsDataSourceModel)(nil)
	schema := log_explorer_dataset.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
