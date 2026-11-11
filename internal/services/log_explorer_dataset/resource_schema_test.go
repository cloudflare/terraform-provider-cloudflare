package log_explorer_dataset_test

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/log_explorer_dataset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
	"testing"
)

func TestLogExplorerDatasetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*log_explorer_dataset.LogExplorerDatasetModel)(nil)
	schema := log_explorer_dataset.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
