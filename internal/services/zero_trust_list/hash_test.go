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
			name: "null vs empty string value",
			items1: []*ZeroTrustListItemsModel{
				{Value: types.StringNull()},
			},
			items2: []*ZeroTrustListItemsModel{
				{Value: types.StringValue("")},
			},
			shouldMatch: true, // both result in empty string
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
