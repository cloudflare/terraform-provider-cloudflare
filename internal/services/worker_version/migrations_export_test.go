package worker_version

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ResourceSchemaV0ForTest exports resourceSchemaV0 for testing purposes
func ResourceSchemaV0ForTest(ctx context.Context) schema.Schema {
	return *resourceSchemaV0(ctx)
}

// ResourceModelV0ForTest is an alias for resourceModelV0 for testing purposes
type ResourceModelV0ForTest = resourceModelV0

