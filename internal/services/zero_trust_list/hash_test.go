package zero_trust_list

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestComputeItemsHash verifies the hash function produces consistent,
// order-independent hashes for list items.
func TestComputeItemsHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		items1      []*ZeroTrustListItemsModel
		items2      []*ZeroTrustListItemsModel
		shouldMatch bool
	}{
		{
			name:        "empty lists match",
			items1:      []*ZeroTrustListItemsModel{},
			items2:      []*ZeroTrustListItemsModel{},
			shouldMatch: true,
		},
		{
			name: "same items same order match",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
				{Value: types.StringValue("10.0.0.2")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
				{Value: types.StringValue("10.0.0.2")},
			},
			shouldMatch: true,
		},
		{
			name: "same items different order match (set semantics)",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
				{Value: types.StringValue("10.0.0.2")},
				{Value: types.StringValue("10.0.0.3")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.3")},
				{Value: types.StringValue("10.0.0.1")},
				{Value: types.StringValue("10.0.0.2")},
			},
			shouldMatch: true,
		},
		{
			name: "different items do not match",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.2")},
			},
			shouldMatch: false,
		},
		{
			name: "different count does not match",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
				{Value: types.StringValue("10.0.0.2")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
			},
			shouldMatch: false,
		},
		{
			name: "same value different description does not match",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringValue("server1")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringValue("server2")},
			},
			shouldMatch: false,
		},
		{
			name: "same value and description match",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringValue("server1")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringValue("server1")},
			},
			shouldMatch: true,
		},
		{
			name: "nil items handled gracefully",
			items1: []*ZeroTrustListItemsModel{
				nil,
				{Value: types.StringValue("10.0.0.1")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1")},
				nil,
			},
			shouldMatch: true,
		},
		{
			name: "null description does not match empty string description",
			// The hash explicitly distinguishes null from "" to avoid suppressing
			// a diff when a user changes description from null to "" or vice versa.
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringNull()},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringValue("")},
			},
			shouldMatch: false,
		},
		{
			name: "null value does not match empty string value",
			// The hash explicitly distinguishes null from "" for value as well.
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringNull()},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("")},
			},
			shouldMatch: false,
		},
		{
			name: "nil item is skipped — matches empty slice",
			// nil pointers are skipped during hashing (consistent with nil-stripping
			// in suppressItemsDiff), so [nil] hashes the same as [].
			items1: []*ZeroTrustListItemsModel{
				nil,
			},
			items2:      []*ZeroTrustListItemsModel{},
			shouldMatch: true,
		},
		{
			name: "value containing null byte does not collide with value+description pair",
			// Regression: with the old "\x00" intra-item separator, {value="a", desc="b"}
			// and {value="a\x00b", desc=""} both encoded as "a\x00b", producing a false match.
			// Length-prefixed encoding with type tags eliminates this.
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("a"), Description: types.StringValue("b")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("a\x00b"), Description: types.StringValue("")},
			},
			shouldMatch: false,
		},
		{
			name: "two entries do not collide with one entry whose value looks like concatenated encodings",
			// Checks that a single item whose value equals the raw concatenation of two
			// items' encodings does not produce a hash match.
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("a")},
				{Value: types.StringValue("b")},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("v(1):a/d:null\x00v(1):b/d:null")},
			},
			shouldMatch: false,
		},
		{
			name: "unknown value does not match null value",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringUnknown()},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringNull()},
			},
			shouldMatch: false,
		},
		{
			name: "unknown description does not match null description",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringUnknown()},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("10.0.0.1"), Description: types.StringNull()},
			},
			shouldMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := computeItemsHash(tt.items1)
			hash2 := computeItemsHash(tt.items2)

			if tt.shouldMatch && hash1 != hash2 {
				t.Errorf("expected hashes to match, got %x != %x", hash1, hash2)
			}
			if !tt.shouldMatch && hash1 == hash2 {
				t.Errorf("expected hashes to differ, got %x == %x", hash1, hash2)
			}
		})
	}
}

// TestComputeItemsHashDeterministic verifies the hash is deterministic
// (same input always produces same output).
func TestComputeItemsHashDeterministic(t *testing.T) {
	t.Parallel()

	items := []*ZeroTrustListItemsModel{
		{Value: types.StringValue("10.0.0.1"), Description: types.StringValue("test")},
		{Value: types.StringValue("10.0.0.2")},
		{Value: types.StringValue("192.168.1.1")},
	}

	hash1 := computeItemsHash(items)
	hash2 := computeItemsHash(items)
	hash3 := computeItemsHash(items)

	if hash1 != hash2 || hash2 != hash3 {
		t.Errorf("hash should be deterministic: %x, %x, %x", hash1, hash2, hash3)
	}
}

// TestComputeItemsHashLargeList verifies performance with large lists.
func TestComputeItemsHashLargeList(t *testing.T) {
	t.Parallel()

	// Create a list with 5000 items
	items := make([]*ZeroTrustListItemsModel, 5000)
	for i := 0; i < 5000; i++ {
		items[i] = &ZeroTrustListItemsModel{
			Value: types.StringValue(fmt.Sprintf("10.%d.%d.1", i/256, i%256)),
		}
	}

	// Should complete without timeout (test has default timeout)
	hash := computeItemsHash(items)
	if hash == [32]byte{} {
		t.Error("hash should not be zero")
	}
}

// BenchmarkComputeItemsHash measures hash performance for various list sizes.
func BenchmarkComputeItemsHash(b *testing.B) {
	sizes := []int{100, 1000, 2000, 5000}

	for _, size := range sizes {
		items := make([]*ZeroTrustListItemsModel, size)
		for i := 0; i < size; i++ {
			items[i] = &ZeroTrustListItemsModel{
				Value: types.StringValue(fmt.Sprintf("10.%d.%d.1", i/256, i%256)),
			}
		}

		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				computeItemsHash(items)
			}
		})
	}
}
