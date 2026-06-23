package load_balancer_pool

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func origin(name, address string) *LoadBalancerPoolOriginsModel {
	return &LoadBalancerPoolOriginsModel{
		Name:    types.StringValue(name),
		Address: types.StringValue(address),
	}
}

func names(origins *[]*LoadBalancerPoolOriginsModel) []string {
	out := make([]string, 0, len(*origins))
	for _, o := range *origins {
		out = append(out, o.Name.ValueString())
	}
	return out
}

func eq(a, b []string) bool {
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

// TestReorderOriginsMatchesPlannedOrder is the core LB-5712 regression: the API
// returns origins sorted ascending by name, but they must be reordered back to
// the order the user planned so Terraform sees no diff.
func TestReorderOriginsMatchesPlannedOrder(t *testing.T) {
	// User planned descending order.
	planned := &[]*LoadBalancerPoolOriginsModel{
		origin("clusterB", "192.0.2.2"),
		origin("clusterA", "192.0.2.1"),
	}
	snap := snapshotOriginOrder(planned)

	// API returns ascending-by-name order.
	apiResp := &[]*LoadBalancerPoolOriginsModel{
		origin("clusterA", "192.0.2.1"),
		origin("clusterB", "192.0.2.2"),
	}
	reorderOriginsFromSnapshot(apiResp, snap)

	want := []string{"clusterB", "clusterA"}
	if got := names(apiResp); !eq(got, want) {
		t.Fatalf("reordered origins = %v, want %v", got, want)
	}
}

// TestReorderOriginsNoSnapshotCanonicalSort verifies that with no prior order
// (e.g. import), origins fall back to a deterministic canonical sort by name.
func TestReorderOriginsNoSnapshotCanonicalSort(t *testing.T) {
	apiResp := &[]*LoadBalancerPoolOriginsModel{
		origin("clusterB", "192.0.2.2"),
		origin("clusterA", "192.0.2.1"),
	}
	reorderOriginsFromSnapshot(apiResp, nil)

	want := []string{"clusterA", "clusterB"}
	if got := names(apiResp); !eq(got, want) {
		t.Fatalf("canonical sort = %v, want %v", got, want)
	}
}

// TestReorderOriginsAddedOriginSortsAfterMatched verifies that an origin present
// in the API response but absent from the prior plan is placed after the matched
// origins, in canonical name order, rather than dropped or causing instability.
func TestReorderOriginsAddedOriginSortsAfterMatched(t *testing.T) {
	planned := &[]*LoadBalancerPoolOriginsModel{
		origin("clusterB", "192.0.2.2"),
		origin("clusterA", "192.0.2.1"),
	}
	snap := snapshotOriginOrder(planned)

	apiResp := &[]*LoadBalancerPoolOriginsModel{
		origin("clusterA", "192.0.2.1"),
		origin("clusterZ", "192.0.2.9"), // not in plan
		origin("clusterB", "192.0.2.2"),
	}
	reorderOriginsFromSnapshot(apiResp, snap)

	want := []string{"clusterB", "clusterA", "clusterZ"}
	if got := names(apiResp); !eq(got, want) {
		t.Fatalf("reordered origins = %v, want %v", got, want)
	}
}

// TestReorderOriginsSameNameDifferentAddress verifies that identity uses both
// name and address, so origins are matched precisely when names collide.
func TestReorderOriginsSameNameDifferentAddress(t *testing.T) {
	planned := &[]*LoadBalancerPoolOriginsModel{
		origin("dup", "192.0.2.9"),
		origin("dup", "192.0.2.1"),
	}
	snap := snapshotOriginOrder(planned)

	apiResp := &[]*LoadBalancerPoolOriginsModel{
		origin("dup", "192.0.2.1"),
		origin("dup", "192.0.2.9"),
	}
	reorderOriginsFromSnapshot(apiResp, snap)

	wantAddrs := []string{"192.0.2.9", "192.0.2.1"}
	got := make([]string, 0, 2)
	for _, o := range *apiResp {
		got = append(got, o.Address.ValueString())
	}
	if !eq(got, wantAddrs) {
		t.Fatalf("reordered addresses = %v, want %v", got, wantAddrs)
	}
}
