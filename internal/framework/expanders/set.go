package expanders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringSet accepts a `types.Set` and returns a slice of strings.
func StringSet(ctx context.Context, in types.Set) []string {
	results := []string{}
	_ = in.ElementsAs(ctx, &results, false)
	return results
}
