package expanders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringSet accepts a `types.Set` and returns a slice of strings.
func StringSet(in types.Set) []string {
	results := []string{}
	_ = in.ElementsAs(context.Background(), &results, false)
	return results
}
