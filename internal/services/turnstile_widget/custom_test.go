package turnstile_widget

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func domain(s string) types.String {
	return types.StringValue(s)
}

func domainSlice(ss ...string) *[]types.String {
	ds := make([]types.String, len(ss))
	for i, s := range ss {
		ds[i] = types.StringValue(s)
	}
	return &ds
}

func domainStrings(domains *[]types.String) []string {
	out := make([]string, 0, len(*domains))
	for _, d := range *domains {
		out = append(out, d.ValueString())
	}
	return out
}

func stringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestReorderDomainsMatchesPlannedOrder is the core regression test: the API
// returns domains sorted alphabetically, but they must be reordered back to
// the order the user planned so Terraform sees no diff.
func TestReorderDomainsMatchesPlannedOrder(t *testing.T) {
	// User planned non-alphabetical order.
	planned := domainSlice("zebra.example.com", "alpha.example.com", "middle.example.com")
	snap := snapshotDomainOrder(planned)

	// API returns alphabetical order.
	apiResp := domainSlice("alpha.example.com", "middle.example.com", "zebra.example.com")
	reorderDomainsFromSnapshot(apiResp, snap)

	want := []string{"zebra.example.com", "alpha.example.com", "middle.example.com"}
	if got := domainStrings(apiResp); !stringsEqual(got, want) {
		t.Fatalf("reordered domains = %v, want %v", got, want)
	}
}

// TestReorderDomainsNoSnapshotCanonicalSort verifies that with no prior order
// (e.g. import), domains fall back to a deterministic alphabetical sort.
func TestReorderDomainsNoSnapshotCanonicalSort(t *testing.T) {
	apiResp := domainSlice("zebra.example.com", "alpha.example.com", "middle.example.com")
	reorderDomainsFromSnapshot(apiResp, nil)

	want := []string{"alpha.example.com", "middle.example.com", "zebra.example.com"}
	if got := domainStrings(apiResp); !stringsEqual(got, want) {
		t.Fatalf("canonical sort = %v, want %v", got, want)
	}
}

// TestReorderDomainsAddedDomainSortsAfterMatched verifies that a domain
// present in the API response but absent from the prior plan is placed after
// the matched domains, in alphabetical order.
func TestReorderDomainsAddedDomainSortsAfterMatched(t *testing.T) {
	planned := domainSlice("zebra.example.com", "alpha.example.com")
	snap := snapshotDomainOrder(planned)

	apiResp := domainSlice("alpha.example.com", "new.example.com", "zebra.example.com")
	reorderDomainsFromSnapshot(apiResp, snap)

	want := []string{"zebra.example.com", "alpha.example.com", "new.example.com"}
	if got := domainStrings(apiResp); !stringsEqual(got, want) {
		t.Fatalf("reordered domains = %v, want %v", got, want)
	}
}

// TestReorderDomainsRemovedDomainHandled verifies that a domain present in
// the prior plan but absent from the API response does not cause issues.
func TestReorderDomainsRemovedDomainHandled(t *testing.T) {
	planned := domainSlice("zebra.example.com", "alpha.example.com", "removed.example.com")
	snap := snapshotDomainOrder(planned)

	apiResp := domainSlice("alpha.example.com", "zebra.example.com")
	reorderDomainsFromSnapshot(apiResp, snap)

	want := []string{"zebra.example.com", "alpha.example.com"}
	if got := domainStrings(apiResp); !stringsEqual(got, want) {
		t.Fatalf("reordered domains = %v, want %v", got, want)
	}
}

// TestReorderDomainsEmptySlice verifies that nil and empty slices are handled
// without panics.
func TestReorderDomainsEmptySlice(t *testing.T) {
	// nil domains
	reorderDomainsFromSnapshot(nil, nil)

	// empty domains
	empty := domainSlice()
	reorderDomainsFromSnapshot(empty, nil)

	if len(*empty) != 0 {
		t.Fatalf("expected empty slice, got %v", domainStrings(empty))
	}
}

// TestReorderDomainsSingleDomain verifies the trivial single-domain case.
func TestReorderDomainsSingleDomain(t *testing.T) {
	planned := domainSlice("example.com")
	snap := snapshotDomainOrder(planned)

	apiResp := domainSlice("example.com")
	reorderDomainsFromSnapshot(apiResp, snap)

	want := []string{"example.com"}
	if got := domainStrings(apiResp); !stringsEqual(got, want) {
		t.Fatalf("reordered domains = %v, want %v", got, want)
	}
}

// TestReorderDomainsNumericSortIssue reproduces the exact scenario from
// GitHub #7028: staging-10 sorts between staging-1 and staging-2
// alphabetically, but the user wants numeric order.
func TestReorderDomainsNumericSortIssue(t *testing.T) {
	// User planned numeric-friendly order.
	planned := domainSlice(
		"staging-1.example.com",
		"staging-2.example.com",
		"staging-10.example.com",
		"staging-11.example.com",
	)
	snap := snapshotDomainOrder(planned)

	// API returns alphabetical order (10 sorts between 1 and 2).
	apiResp := domainSlice(
		"staging-1.example.com",
		"staging-10.example.com",
		"staging-11.example.com",
		"staging-2.example.com",
	)
	reorderDomainsFromSnapshot(apiResp, snap)

	want := []string{
		"staging-1.example.com",
		"staging-2.example.com",
		"staging-10.example.com",
		"staging-11.example.com",
	}
	if got := domainStrings(apiResp); !stringsEqual(got, want) {
		t.Fatalf("reordered domains = %v, want %v", got, want)
	}
}
