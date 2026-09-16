package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformCKQueryStringExcludeWildcard(t *testing.T) {
	t.Parallel()

	result, diags := transformCKQueryString(context.Background(), &SourceV4QueryStringModel{
		Exclude: types.SetValueMust(types.StringType, []attr.Value{types.StringValue("*")}),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if result.Exclude == nil {
		t.Fatal("expected exclude to be set")
	}
	if !result.Exclude.All.ValueBool() {
		t.Errorf("expected exclude wildcard to set all=true, got %v", result.Exclude.All)
	}
	if result.Exclude.List != nil {
		t.Errorf("expected exclude wildcard list to be nil, got %v", result.Exclude.List)
	}
}

func TestTransformCKQueryStringExcludePreservesListSemantics(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		exclude     types.Set
		wantExclude bool
		wantAll     bool
		wantList    []string
	}{
		{
			name:        "ordinary value",
			exclude:     stringSet("parameter"),
			wantExclude: true,
			wantAll:     false,
			wantList:    []string{"parameter"},
		},
		{
			name:        "wildcard with another value",
			exclude:     stringSet("*", "parameter"),
			wantExclude: true,
			wantAll:     false,
			wantList:    []string{"*", "parameter"},
		},
		{
			name:        "unknown",
			exclude:     types.SetUnknown(types.StringType),
			wantExclude: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, diags := transformCKQueryString(context.Background(), &SourceV4QueryStringModel{Exclude: tc.exclude})
			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
			if (result.Exclude != nil) != tc.wantExclude {
				t.Fatalf("exclude presence = %t, want %t", result.Exclude != nil, tc.wantExclude)
			}
			if !tc.wantExclude {
				return
			}
			if result.Exclude.All.ValueBool() != tc.wantAll {
				t.Errorf("all = %t, want %t", result.Exclude.All.ValueBool(), tc.wantAll)
			}
			if len(result.Exclude.List) != len(tc.wantList) {
				t.Fatalf("list length = %d, want %d", len(result.Exclude.List), len(tc.wantList))
			}
			gotList := make(map[string]bool, len(result.Exclude.List))
			for _, value := range result.Exclude.List {
				gotList[value.ValueString()] = true
			}
			for _, want := range tc.wantList {
				if !gotList[want] {
					t.Errorf("list = %v, want it to contain %q", gotList, want)
				}
			}
		})
	}
}

func stringSet(values ...string) types.Set {
	attributes := make([]attr.Value, 0, len(values))
	for _, value := range values {
		attributes = append(attributes, types.StringValue(value))
	}
	return types.SetValueMust(types.StringType, attributes)
}
