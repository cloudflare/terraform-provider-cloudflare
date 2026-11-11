package list

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ListItemsSet replaces customfield.NestedObjectSet[ListItemModel] for
// performance only. These tests check that it behaves exactly like the
// original types.

func ipItem(ctx context.Context, ip string, comment *string) ListItemModel {
	c := types.StringNull()
	if comment != nil {
		c = types.StringValue(*comment)
	}
	return ListItemModel{
		ASN:      types.Int64Null(),
		Comment:  c,
		IP:       types.StringValue(ip),
		Hostname: customfield.NullObject[ListItemHostnameModel](ctx),
		Redirect: customfield.NullObject[ListItemRedirectModel](ctx),
	}
}

func strPtr(s string) *string { return &s }

func testItems(ctx context.Context) (a, b, c, unknownComment, hostname, redirect, asn0, asnNull ListItemModel) {
	a = ipItem(ctx, "192.0.2.1", strPtr("a"))
	b = ipItem(ctx, "192.0.2.2", nil)
	c = ipItem(ctx, "192.0.2.3", strPtr(""))
	unknownComment = ipItem(ctx, "192.0.2.4", nil)
	unknownComment.Comment = types.StringUnknown()
	hostname = ListItemModel{
		ASN:     types.Int64Null(),
		Comment: types.StringNull(),
		IP:      types.StringNull(),
		Hostname: customfield.NewObjectMust(ctx, &ListItemHostnameModel{
			URLHostname:          types.StringValue("*.example.com"),
			ExcludeExactHostname: types.BoolValue(true),
		}),
		Redirect: customfield.NullObject[ListItemRedirectModel](ctx),
	}
	redirect = ListItemModel{
		ASN:      types.Int64Null(),
		Comment:  types.StringNull(),
		IP:       types.StringNull(),
		Hostname: customfield.NullObject[ListItemHostnameModel](ctx),
		Redirect: customfield.NewObjectMust(ctx, &ListItemRedirectModel{
			SourceURL:           types.StringValue("example.com/a"),
			TargetURL:           types.StringValue("https://example.net/a"),
			IncludeSubdomains:   types.BoolNull(),
			PreservePathSuffix:  types.BoolValue(false),
			PreserveQueryString: types.BoolNull(),
			StatusCode:          types.Int64Value(301),
			SubpathMatching:     types.BoolNull(),
		}),
	}
	asn0 = ListItemModel{
		ASN:      types.Int64Value(0),
		Comment:  types.StringNull(),
		IP:       types.StringNull(),
		Hostname: customfield.NullObject[ListItemHostnameModel](ctx),
		Redirect: customfield.NullObject[ListItemRedirectModel](ctx),
	}
	asnNull = asn0
	asnNull.ASN = types.Int64Null()
	return
}

func listItemElements(t testing.TB, ctx context.Context, items []ListItemModel) ([]attr.Value, []tftypes.Value) {
	t.Helper()
	elemType, d := listItemElemType(ctx)
	if d.HasError() {
		t.Fatal(d)
	}
	attrs := make([]attr.Value, 0, len(items))
	tfs := make([]tftypes.Value, 0, len(items))
	for _, item := range items {
		obj, d := basetypes.NewObjectValueFrom(ctx, elemType.AttrTypes, item)
		if d.HasError() {
			t.Fatal(d)
		}
		tv, err := obj.ToTerraformValue(ctx)
		if err != nil {
			t.Fatal(err)
		}
		attrs = append(attrs, obj)
		tfs = append(tfs, tv)
	}
	return attrs, tfs
}

// rawListItemsSet builds a set without any duplicate check, so that invalid
// states can be compared as well.
func rawListItemsSet(t testing.TB, ctx context.Context, items []ListItemModel) ListItemsSet {
	t.Helper()
	if items == nil {
		return NullListItemsSet(ctx)
	}
	elems, _ := listItemElements(t, ctx, items)
	elemType, _ := listItemElemType(ctx)
	v, d := basetypes.NewSetValue(elemType, elems)
	if d.HasError() {
		t.Fatal(d)
	}
	return ListItemsSet{SetValue: v}
}

func TestListItemsSetEqualMatchesSetValueEqual(t *testing.T) {
	ctx := context.Background()
	a, b, c, unk, host, redir, asn0, asnNull := testItems(ctx)
	unknown := UnknownListItemsSet(ctx)
	null := NullListItemsSet(ctx)

	sets := map[string]ListItemsSet{
		"null":                null,
		"unknown":             unknown,
		"zero":                {},
		"empty":               rawListItemsSet(t, ctx, []ListItemModel{}),
		"a":                   rawListItemsSet(t, ctx, []ListItemModel{a}),
		"abc":                 rawListItemsSet(t, ctx, []ListItemModel{a, b, c}),
		"cba":                 rawListItemsSet(t, ctx, []ListItemModel{c, b, a}),
		"abd":                 rawListItemsSet(t, ctx, []ListItemModel{a, b, ipItem(ctx, "192.0.2.3", nil)}),
		"ab":                  rawListItemsSet(t, ctx, []ListItemModel{a, b}),
		"aab":                 rawListItemsSet(t, ctx, []ListItemModel{a, a, b}),
		"abb":                 rawListItemsSet(t, ctx, []ListItemModel{a, b, b}),
		"aa":                  rawListItemsSet(t, ctx, []ListItemModel{a, a}),
		"a,unknown comment":   rawListItemsSet(t, ctx, []ListItemModel{a, unk}),
		"unknown comment,a":   rawListItemsSet(t, ctx, []ListItemModel{unk, a}),
		"mixed kinds":         rawListItemsSet(t, ctx, []ListItemModel{host, redir, asn0, a}),
		"mixed kinds shuffle": rawListItemsSet(t, ctx, []ListItemModel{asn0, a, redir, host}),
		"asn null":            rawListItemsSet(t, ctx, []ListItemModel{host, redir, asnNull, a}),
	}

	for an, av := range sets {
		for bn, bv := range sets {
			want := av.SetValue.Equal(bv.SetValue)
			if got := av.Equal(bv); got != want {
				t.Errorf("%s.Equal(%s) = %v, want %v", an, bn, got, want)
			}
		}
		// Like customfield.NestedObjectSet, other value types are never equal.
		ref := customfield.NestedObjectSet[ListItemModel]{SetValue: av.SetValue}
		if av.Equal(ref) || av.Equal(av.SetValue) {
			t.Errorf("%s.Equal(other type) = true, want false", an)
		}
	}
}

func TestListItemsSetSemanticEqualsOnlyForIdenticalElements(t *testing.T) {
	ctx := context.Background()
	a, b, c, unk, host, redir, _, _ := testItems(ctx)

	for _, tc := range []struct {
		name            string
		proposed, prior ListItemsSet
		want            bool
	}{
		{"same order", rawListItemsSet(t, ctx, []ListItemModel{a, b, c}), rawListItemsSet(t, ctx, []ListItemModel{a, b, c}), true},
		{"reordered", rawListItemsSet(t, ctx, []ListItemModel{c, host, a, redir}), rawListItemsSet(t, ctx, []ListItemModel{redir, a, host, c}), true},
		{"empty", rawListItemsSet(t, ctx, []ListItemModel{}), rawListItemsSet(t, ctx, []ListItemModel{}), true},
		{"element changed", rawListItemsSet(t, ctx, []ListItemModel{a, b}), rawListItemsSet(t, ctx, []ListItemModel{a, c}), false},
		{"element added", rawListItemsSet(t, ctx, []ListItemModel{a, b, c}), rawListItemsSet(t, ctx, []ListItemModel{a, b}), false},
		{"element removed", rawListItemsSet(t, ctx, []ListItemModel{a}), rawListItemsSet(t, ctx, []ListItemModel{a, b}), false},
		{"duplicates in proposed", rawListItemsSet(t, ctx, []ListItemModel{a, a, b}), rawListItemsSet(t, ctx, []ListItemModel{a, b, b}), false},
		{"duplicates in prior", rawListItemsSet(t, ctx, []ListItemModel{a, b}), rawListItemsSet(t, ctx, []ListItemModel{a, a}), false},
		{"unknown element value", rawListItemsSet(t, ctx, []ListItemModel{a, unk}), rawListItemsSet(t, ctx, []ListItemModel{a, unk}), false},
		{"null prior", rawListItemsSet(t, ctx, []ListItemModel{a}), NullListItemsSet(ctx), false},
		{"unknown proposed", UnknownListItemsSet(ctx), rawListItemsSet(t, ctx, []ListItemModel{a}), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, diags := tc.proposed.SetSemanticEquals(ctx, tc.prior)
			if diags.HasError() {
				t.Fatal(diags)
			}
			if got != tc.want {
				t.Errorf("SetSemanticEquals() = %v, want %v", got, tc.want)
			}
			if got && !tc.proposed.SetValue.Equal(tc.prior.SetValue) {
				t.Errorf("SetSemanticEquals() = true for values that are not equal")
			}
		})
	}

	// Other value types are never semantically equal.
	v := rawListItemsSet(t, ctx, []ListItemModel{a})
	if got, _ := v.SetSemanticEquals(ctx, v.SetValue); got {
		t.Errorf("SetSemanticEquals(basetypes.SetValue) = true, want false")
	}
}

func TestListItemsSetTypeValidateMatchesSetType(t *testing.T) {
	ctx := context.Background()
	a, b, c, unk, host, redir, asn0, _ := testItems(ctx)
	typ := NewListItemsSetType(ctx)
	ref := customfield.NewNestedObjectSetType[ListItemModel](ctx)
	setTFType := typ.TerraformType(ctx)

	set := func(items ...ListItemModel) tftypes.Value {
		_, tfs := listItemElements(t, ctx, items)
		return tftypes.NewValue(setTFType, tfs)
	}

	for _, tc := range []struct {
		name string
		in   tftypes.Value
	}{
		{"distinct", set(a, b, c, host, redir, asn0)},
		{"empty", set()},
		{"null", tftypes.NewValue(setTFType, nil)},
		{"unknown", tftypes.NewValue(setTFType, tftypes.UnknownValue)},
		{"no type", tftypes.Value{}},
		{"wrong type", tftypes.NewValue(tftypes.String, "x")},
		{"one duplicate", set(a, b, a, c)},
		{"many duplicates", set(a, b, a, c, a, b, host, host)},
		{"duplicate with unknown elements", set(unk, a, unk, a)},
		{"unknown elements only", set(unk, unk)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := path.Root("items")
			want := ref.Validate(ctx, tc.in, p)
			got := typ.Validate(ctx, tc.in, p)
			if !got.Equal(want) {
				t.Errorf("Validate() diagnostics differ\n got: %v\nwant: %v", got, want)
			}
		})
	}
}

func TestNewListItemsSetMatchesNewObjectSet(t *testing.T) {
	ctx := context.Background()
	a, b, c, unk, host, redir, asn0, asnNull := testItems(ctx)

	for _, tc := range []struct {
		name  string
		items []ListItemModel
	}{
		{"nil", nil},
		{"empty", []ListItemModel{}},
		{"single", []ListItemModel{a}},
		{"mixed kinds", []ListItemModel{c, host, a, redir, asn0, b, asnNull}},
		{"unknown element value", []ListItemModel{a, unk}},
		{"duplicates", []ListItemModel{a, b, a, c, b}},
		{"zero value model", []ListItemModel{{}, a}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			want, wantDiags := customfield.NewObjectSet[ListItemModel](ctx, tc.items)
			got, gotDiags := NewListItemsSet(ctx, tc.items)
			if !gotDiags.Equal(wantDiags) {
				t.Fatalf("diagnostics differ\n got: %v\nwant: %v", gotDiags, wantDiags)
			}
			assertSameSetValue(t, ctx, got, want.SetValue)
		})
	}

	if got, want := NullListItemsSet(ctx), customfield.NullObjectSet[ListItemModel](ctx); !got.SetValue.Equal(want.SetValue) {
		t.Errorf("NullListItemsSet() = %v, want %v", got, want)
	}
	if got, want := UnknownListItemsSet(ctx), customfield.UnknownObjectSet[ListItemModel](ctx); !got.SetValue.Equal(want.SetValue) {
		t.Errorf("UnknownListItemsSet() = %v, want %v", got, want)
	}
}

func assertSameSetValue(t *testing.T, ctx context.Context, got ListItemsSet, want basetypes.SetValue) {
	t.Helper()
	gotTF, err := got.ToTerraformValue(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wantTF, err := customfield.NestedObjectSet[ListItemModel]{SetValue: want}.ToTerraformValue(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// String() includes element order, which must also be preserved.
	if !gotTF.Equal(wantTF) || gotTF.String() != wantTF.String() || got.SetValue.String() != want.String() {
		t.Errorf("value differs\n got: %s\nwant: %s", gotTF, wantTF)
	}
	if !got.SetValue.Equal(want) {
		t.Errorf("SetValue differs\n got: %v\nwant: %v", got.SetValue, want)
	}
}

func TestListItemsSetTypeMatchesNestedObjectSetType(t *testing.T) {
	ctx := context.Background()
	a, b, _, _, host, _, _, _ := testItems(ctx)
	typ := NewListItemsSetType(ctx)
	ref := customfield.NewNestedObjectSetType[ListItemModel](ctx)

	if typ.String() != ref.String() {
		t.Errorf("String() = %q, want %q", typ.String(), ref.String())
	}
	if !typ.TerraformType(ctx).Equal(ref.TerraformType(ctx)) {
		t.Errorf("TerraformType() = %s, want %s", typ.TerraformType(ctx), ref.TerraformType(ctx))
	}
	if !typ.ElementType().Equal(ref.ElementType()) {
		t.Errorf("ElementType() = %s, want %s", typ.ElementType(), ref.ElementType())
	}
	if !typ.Equal(NewListItemsSetType(ctx)) || typ.Equal(ref) {
		t.Errorf("Equal() must only match ListItemsSetType")
	}

	setTFType := typ.TerraformType(ctx)
	_, tfs := listItemElements(t, ctx, []ListItemModel{b, host, a})
	for _, in := range []tftypes.Value{
		tftypes.NewValue(setTFType, tfs),
		tftypes.NewValue(setTFType, []tftypes.Value{}),
		tftypes.NewValue(setTFType, nil),
		tftypes.NewValue(setTFType, tftypes.UnknownValue),
	} {
		got, err := typ.ValueFromTerraform(ctx, in)
		if err != nil {
			t.Fatal(err)
		}
		want, err := ref.ValueFromTerraform(ctx, in)
		if err != nil {
			t.Fatal(err)
		}
		gotSet, ok := got.(ListItemsSet)
		if !ok {
			t.Fatalf("ValueFromTerraform() returned %T", got)
		}
		assertSameSetValue(t, ctx, gotSet, mustAs[customfield.NestedObjectSet[ListItemModel]](t, want).SetValue)
		if !gotSet.Type(ctx).Equal(typ) {
			t.Errorf("Type() = %s, want %s", gotSet.Type(ctx), typ)
		}

		gotStructs, gotDiags := gotSet.AsStructSliceT(ctx)
		wantStructs, wantDiags := mustAs[customfield.NestedObjectSet[ListItemModel]](t, want).AsStructSliceT(ctx)
		if !gotDiags.Equal(wantDiags) || fmt.Sprint(gotStructs) != fmt.Sprint(wantStructs) {
			t.Errorf("AsStructSliceT() = %v, %v, want %v, %v", gotStructs, gotDiags, wantStructs, wantDiags)
		}
	}

	if _, err := typ.ValueFromTerraform(ctx, tftypes.NewValue(tftypes.String, "x")); err == nil {
		t.Errorf("ValueFromTerraform(string) succeeded, want error")
	}
}

func TestListItemsSetAPIJSONMatchesNestedObjectSet(t *testing.T) {
	type refModel struct {
		Items customfield.NestedObjectSet[ListItemModel] `json:"items,optional"`
	}
	type newModel struct {
		Items ListItemsSet `json:"items,optional"`
	}
	ctx := context.Background()

	for _, data := range []string{
		`{"items":[{"ip":"192.0.2.1","comment":"a"},{"ip":"192.0.2.0/24"},{"hostname":{"url_hostname":"*.example.com","exclude_exact_hostname":false}},{"redirect":{"source_url":"example.com/a","target_url":"https://example.net/a","status_code":301}},{"asn":13335}]}`,
		`{"items":[]}`,
		`{"items":null}`,
		`{}`,
	} {
		t.Run(data, func(t *testing.T) {
			var ref refModel
			var got newModel
			if err := apijson.Unmarshal([]byte(data), &ref); err != nil {
				t.Fatal(err)
			}
			if err := apijson.Unmarshal([]byte(data), &got); err != nil {
				t.Fatal(err)
			}
			assertSameSetValue(t, ctx, got.Items, ref.Items.SetValue)

			refJSON, refErr := apijson.MarshalRoot(ref)
			gotJSON, gotErr := apijson.MarshalRoot(got)
			if fmt.Sprint(refErr) != fmt.Sprint(gotErr) || string(refJSON) != string(gotJSON) {
				t.Errorf("MarshalRoot() = %s, %v, want %s, %v", gotJSON, gotErr, refJSON, refErr)
			}
		})
	}
}

// The tests above would also pass if ListItemsSet always used the original,
// quadratic implementations. The tests below check that typical inputs take
// the indexed paths, and that the public methods scale accordingly.

// manyListItems returns n distinct items of all kinds.
func manyListItems(ctx context.Context, n int) []ListItemModel {
	_, _, _, _, host, redir, asn0, _ := testItems(ctx)
	items := []ListItemModel{host, redir, asn0}
	for i := 0; len(items) < n; i++ {
		var comment *string
		switch i % 3 {
		case 0:
			comment = strPtr(fmt.Sprintf("comment %d", i))
		case 1:
			comment = strPtr("")
		}
		items = append(items, ipItem(ctx, fmt.Sprintf("10.%d.%d.%d", i>>16&255, i>>8&255, i&255), comment))
	}
	return items[:n]
}

func reversedListItems(items []ListItemModel) []ListItemModel {
	out := make([]ListItemModel, len(items))
	for i, item := range items {
		out[len(items)-1-i] = item
	}
	return out
}

func withChangedComment(items []ListItemModel) []ListItemModel {
	out := append([]ListItemModel(nil), items...)
	out[len(out)/2].Comment = types.StringValue("changed")
	return out
}

func TestListItemsSetFastPathsAreTaken(t *testing.T) {
	ctx := context.Background()
	elemType, _ := listItemElemType(ctx)
	items := manyListItems(ctx, 50)
	set := rawListItemsSet(t, ctx, items)
	reordered := rawListItemsSet(t, ctx, reversedListItems(items))
	changed := rawListItemsSet(t, ctx, withChangedComment(items))
	tv, err := set.ToTerraformValue(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !listItemsSetDistinct(elemType, tv) {
		t.Errorf("Validate fast path not taken for distinct items")
	}
	if _, ok := newListItemsSetFast(ctx, items); !ok {
		t.Errorf("NewListItemsSet fast path not taken for distinct items")
	}
	for name, other := range map[string]ListItemsSet{"unchanged": rawListItemsSet(t, ctx, items), "reordered": reordered, "one changed": changed} {
		wantEqual := name != "one changed"
		if eq, ok := listItemsSetEqualIndexed(set.SetValue, other.SetValue); !ok || eq != wantEqual {
			t.Errorf("Equal fast path (%s) = %v, %v, want %v, true", name, eq, ok, wantEqual)
		}
		if got := listItemsSetSameElements(ctx, other.SetValue, set.SetValue); got != wantEqual {
			t.Errorf("SetSemanticEquals fast path (%s) = %v, want %v", name, got, wantEqual)
		}
	}

	// Inputs the fast paths cannot decide are left to the original code.
	dups := append(manyListItems(ctx, 10), items[5])
	dupsTV, _ := rawListItemsSet(t, ctx, dups).ToTerraformValue(ctx)
	if listItemsSetDistinct(elemType, dupsTV) {
		t.Errorf("Validate fast path taken for duplicates")
	}
	if _, ok := newListItemsSetFast(ctx, dups); ok {
		t.Errorf("NewListItemsSet fast path taken for duplicates")
	}
	_, _, _, unk, _, _, _, _ := testItems(ctx)
	withUnknown := rawListItemsSet(t, ctx, append(manyListItems(ctx, 10), unk))
	if _, ok := listItemsSetEqualIndexed(withUnknown.SetValue, withUnknown.SetValue); ok {
		t.Errorf("Equal fast path taken for unknown values")
	}
	if listItemsSetSameElements(ctx, withUnknown.SetValue, withUnknown.SetValue) {
		t.Errorf("SetSemanticEquals fast path taken for unknown values")
	}
}

// countingObject counts calls to Equal, to detect pairwise comparisons.
type countingObject struct {
	basetypes.ObjectValue
	calls *int
}

func (o countingObject) Equal(other attr.Value) bool {
	*o.calls++
	c, ok := other.(countingObject)
	return ok && o.ObjectValue.Equal(c.ObjectValue)
}

func countingListItemsSet(t testing.TB, ctx context.Context, items []ListItemModel, calls *int) ListItemsSet {
	t.Helper()
	elems, _ := listItemElements(t, ctx, items)
	for i, e := range elems {
		elems[i] = countingObject{ObjectValue: mustAs[basetypes.ObjectValue](t, e), calls: calls}
	}
	elemType, _ := listItemElemType(ctx)
	v, d := basetypes.NewSetValue(elemType, elems)
	if d.HasError() {
		t.Fatal(d)
	}
	return ListItemsSet{SetValue: v}
}

func TestListItemsSetScalesLinearly(t *testing.T) {
	ctx := context.Background()
	const n = 800

	// Equal and SetSemanticEquals: the original implementations compare every
	// pair of elements (about n*n/2 calls for reordered sets); the indexed
	// ones compare each element with its matching candidate only.
	items := manyListItems(ctx, n)
	var calls int
	set := countingListItemsSet(t, ctx, items, &calls)
	reordered := countingListItemsSet(t, ctx, reversedListItems(items), &calls)
	if !set.Equal(reordered) {
		t.Fatal("Equal(reordered) = false")
	}
	if calls > 2*n {
		t.Errorf("Equal compared elements %d times for %d items, want at most %d", calls, n, 2*n)
	}
	calls = 0
	if eq, _ := reordered.SetSemanticEquals(ctx, set); !eq {
		t.Fatal("SetSemanticEquals(reordered) = false")
	}
	if calls > 2*n {
		t.Errorf("SetSemanticEquals compared elements %d times for %d items, want at most %d", calls, n, 2*n)
	}

	// Validate and NewListItemsSet: the original duplicate check allocates for
	// every pair of elements, so 4x the items means about 16x the allocations.
	// The indexed check is about 4x.
	typ := NewListItemsSetType(ctx)
	allocs := func(n int) (validate, construct float64) {
		items := manyListItems(ctx, n)
		tv, err := rawListItemsSet(t, ctx, items).ToTerraformValue(ctx)
		if err != nil {
			t.Fatal(err)
		}
		validate = testing.AllocsPerRun(2, func() { typ.Validate(ctx, tv, path.Root("items")) })
		construct = testing.AllocsPerRun(2, func() { NewListItemsSet(ctx, items) })
		return
	}
	smallValidate, smallConstruct := allocs(n / 4)
	largeValidate, largeConstruct := allocs(n)
	if r := largeValidate / smallValidate; r > 8 {
		t.Errorf("Validate allocations grew %.1fx for 4x the items, want under 8x", r)
	}
	if r := largeConstruct / smallConstruct; r > 8 {
		t.Errorf("NewListItemsSet allocations grew %.1fx for 4x the items, want under 8x", r)
	}
}

// Run with:
//
//	go test ./internal/services/list/ -run '^$' -bench 'ListItemsSet' -benchmem
func BenchmarkListItemsSet(b *testing.B) {
	ctx := context.Background()
	typ := NewListItemsSetType(ctx)
	for _, n := range []int{100, 1000, 5000} {
		items := manyListItems(ctx, n)
		set := rawListItemsSet(b, ctx, items)
		tv, err := set.ToTerraformValue(ctx)
		if err != nil {
			b.Fatal(err)
		}
		others := []struct {
			name string
			set  ListItemsSet
		}{
			{"unchanged", rawListItemsSet(b, ctx, items)},
			{"reordered", rawListItemsSet(b, ctx, reversedListItems(items))},
			{"one_changed", rawListItemsSet(b, ctx, withChangedComment(items))},
		}

		b.Run(fmt.Sprintf("Validate/%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				typ.Validate(ctx, tv, path.Root("items"))
			}
		})
		for _, o := range others {
			b.Run(fmt.Sprintf("Equal/%s/%d", o.name, n), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					set.Equal(o.set)
				}
			})
			b.Run(fmt.Sprintf("SetSemanticEquals/%s/%d", o.name, n), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = o.set.SetSemanticEquals(ctx, set)
				}
			})
		}
		b.Run(fmt.Sprintf("NewListItemsSet/%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = NewListItemsSet(ctx, items)
			}
		})
	}
}
