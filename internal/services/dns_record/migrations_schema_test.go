package dns_record

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestPrioritySchemaMigrationBoundary(t *testing.T) {
	t.Parallel()

	current := ResourceSchema(context.Background())
	if current.Version != 501 {
		t.Fatalf("current schema version = %d, want 501", current.Version)
	}
	currentPriority := current.Attributes["priority"].(schema.Float64Attribute)
	if !currentPriority.Computed || currentPriority.Optional {
		t.Fatalf("current priority flags = computed:%t optional:%t, want computed-only", currentPriority.Computed, currentPriority.Optional)
	}

	prior := sourceSchemaV500(context.Background())
	if prior.Version != 500 {
		t.Fatalf("prior schema version = %d, want 500", prior.Version)
	}
	priorPriority := prior.Attributes["priority"].(schema.Float64Attribute)
	if !priorPriority.Optional || priorPriority.Computed {
		t.Fatalf("prior priority flags = computed:%t optional:%t, want optional-only", priorPriority.Computed, priorPriority.Optional)
	}
}
