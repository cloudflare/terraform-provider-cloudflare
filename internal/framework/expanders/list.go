package expanders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringList accepts a `types.List` and returns a slice of strings.
func StringList(ctx context.Context, in types.List) []string {
	results := []string{}
	_ = in.ElementsAs(ctx, &results, false)
	return results
}
