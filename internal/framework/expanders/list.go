package expanders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringList accepts a `types.List` and returns a slice of strings.
func StringList(in types.List) []string {
	results := []string{}
	_ = in.ElementsAs(context.Background(), &results, false)
	return results
}
