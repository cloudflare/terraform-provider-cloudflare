package expanders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Int64Set accepts a `types.Set` and returns a slice of int64.
func Int64Set(ctx context.Context, in types.Set) []int {
	results := []int{}
	_ = in.ElementsAs(ctx, &results, false)
	return results
}
