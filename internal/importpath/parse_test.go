package importpath_test

import (
	"reflect"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
)

func TestParseID(t *testing.T) {
	a, b, c, d := false, int64(0), float64(0), ""
	diags := importpath.ParseImportID("true/1/1.1/hi", "///", &a, &b, &c, &d)
	results := []any{a, b, c, d}

	if diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if !reflect.DeepEqual(results, []any{true, int64(1), 1.1, "hi"}) {
		t.Fatalf("unexpected value: %v", results)
	}
}
