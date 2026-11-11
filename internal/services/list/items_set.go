package list

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ListItemsSetType and ListItemsSet are drop-in replacements for
// customfield.NestedObjectSetType[ListItemModel] and
// customfield.NestedObjectSet[ListItemModel] used by cloudflare_list.items.
//
// The schema, state encoding and diagnostics are unchanged. The difference is
// that whole-set comparisons (Validate duplicate detection, Equal and semantic
// equality) index elements by an exact fingerprint instead of comparing every
// pair of elements, which is quadratic and dominates plan/refresh time for
// lists with thousands of items. Whenever the fast path cannot prove its
// result (for example duplicates are present or a value has an unexpected
// shape), the original framework implementation is used so that results and
// diagnostics stay identical.

var (
	_ basetypes.SetTypable                    = ListItemsSetType{}
	_ xattr.TypeWithValidate                  = ListItemsSetType{}
	_ basetypes.SetValuable                   = ListItemsSet{}
	_ basetypes.SetValuableWithSemanticEquals = ListItemsSet{}
	_ customfield.NestedObjectListLike        = ListItemsSet{}
)

type ListItemsSetType struct {
	basetypes.SetType
}

func NewListItemsSetType(ctx context.Context) ListItemsSetType {
	return ListItemsSetType{SetType: customfield.NewNestedObjectSetType[ListItemModel](ctx).SetType}
}

func (t ListItemsSetType) Equal(o attr.Type) bool {
	other, ok := o.(ListItemsSetType)
	if !ok {
		return false
	}

	return t.SetType.Equal(other.SetType)
}

func (t ListItemsSetType) String() string {
	var zero ListItemModel
	return fmt.Sprintf("SetType[%T]", zero)
}

func (t ListItemsSetType) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NullListItemsSet(ctx), diags
	}
	if in.IsUnknown() {
		return UnknownListItemsSet(ctx), diags
	}

	ty, d := listItemElemType(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return UnknownListItemsSet(ctx), diags
	}

	v, d := basetypes.NewSetValue(ty, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return UnknownListItemsSet(ctx), diags
	}

	return ListItemsSet{SetValue: v}, diags
}

func (t ListItemsSetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	setValue, ok := attrValue.(basetypes.SetValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	setValuable, diags := t.ValueFromSet(ctx, setValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return setValuable, nil
}

func (t ListItemsSetType) ValueType(ctx context.Context) attr.Value {
	return ListItemsSet{}
}

func (t ListItemsSetType) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	return NullListItemsSet(ctx), diags
}

// Validate returns the same diagnostics as basetypes.SetType.Validate. When all
// fully known elements are provably distinct, there is nothing to report and
// the pairwise duplicate scan is skipped. Otherwise the original
// implementation runs and produces the diagnostics.
func (t ListItemsSetType) Validate(ctx context.Context, in tftypes.Value, p path.Path) diag.Diagnostics {
	if listItemsSetDistinct(t.ElementType(), in) {
		return nil
	}
	return t.SetType.Validate(ctx, in, p)
}

// listItemsSetDistinct reports whether in is a known, non-null set whose fully
// known elements all have distinct fingerprints, and whose element type does
// not require per-element validation. In that case
// basetypes.SetType.Validate returns no diagnostics.
func listItemsSetDistinct(elemType attr.Type, in tftypes.Value) bool {
	if elemType == nil || in.Type() == nil || !in.Type().Is(tftypes.Set{}) {
		return false
	}
	if !in.IsKnown() || in.IsNull() {
		return false
	}
	// basetypes.SetType.Validate also validates elements whose type implements
	// the (deprecated) xattr.TypeWithValidate; leave that case to it.
	if _, ok := elemType.(xattr.TypeWithValidate); ok {
		return false
	}

	var elems []tftypes.Value
	if err := in.As(&elems); err != nil {
		return false
	}

	seen := make(map[string]struct{}, len(elems))
	var buf []byte
	for _, elem := range elems {
		if !elem.IsFullyKnown() {
			continue
		}
		var ok bool
		buf, ok = appendListItemFingerprint(buf[:0], elem)
		if !ok {
			return false
		}
		key := string(buf)
		if _, dup := seen[key]; dup {
			return false
		}
		seen[key] = struct{}{}
	}
	return true
}

// ListItemsSet is the value type of ListItemsSetType.
type ListItemsSet struct {
	// placeholder must stay the first field: apijson and the schema/model
	// integrity checks read the element model type via reflection on Field(0).
	//lint:ignore U1000 the placeholder is for easy reflection-based-access
	placeholder ListItemModel
	basetypes.SetValue
}

func (v ListItemsSet) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	tv := v.SetValue
	if tv.ElementType(ctx) == nil {
		tv = NullListItemsSet(ctx).SetValue
	}
	return tv.ToTerraformValue(ctx)
}

func (v ListItemsSet) NullValue(ctx context.Context) customfield.NestedObjectListLike {
	return NullListItemsSet(ctx)
}

func (v ListItemsSet) UnknownValue(ctx context.Context) customfield.NestedObjectListLike {
	return UnknownListItemsSet(ctx)
}

func (v ListItemsSet) KnownValue(ctx context.Context, anyValues any) customfield.NestedObjectListLike {
	r, diags := NewListItemsSet(ctx, anyValues.([]ListItemModel))
	if diags.HasError() {
		panic(fmt.Errorf("unexpected error creating NestedObjectSet: %v", diags))
	}
	return r
}

func (v ListItemsSet) IsNullOrUnknown() bool {
	return v.IsNull() || v.IsUnknown()
}

// Equal has the same result as basetypes.SetValue.Equal.
func (v ListItemsSet) Equal(o attr.Value) bool {
	other, ok := o.(ListItemsSet)
	if !ok {
		return false
	}

	if eq, ok := listItemsSetEqualIndexed(v.SetValue, other.SetValue); ok {
		return eq
	}
	return v.SetValue.Equal(other.SetValue)
}

// SetSemanticEquals reports whether the proposed value (v) is exactly equal to
// the prior value. The framework then keeps the prior value without walking
// every pair of elements. It returns false in every other case, which makes
// the framework run its original element-by-element semantic equality.
//
// Known limitation: when the API and prior state differ by even one item
// (drift), the framework's element walk in
// fwschemadata.ValueSemanticEqualitySetElements still compares every pair of
// elements, so a drifted refresh stays quadratic (about 10s for 829 items and
// 150s for 3,251). The provider cannot skip that walk; it needs the framework
// change in hashicorp/terraform-plugin-framework#1315.
func (v ListItemsSet) SetSemanticEquals(ctx context.Context, other basetypes.SetValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	prior, ok := other.(ListItemsSet)
	if !ok {
		return false, diags
	}

	return listItemsSetSameElements(ctx, v.SetValue, prior.SetValue), diags
}

func (v ListItemsSet) Type(ctx context.Context) attr.Type {
	return NewListItemsSetType(ctx)
}

func (v ListItemsSet) AsStructSlice(ctx context.Context) (any, diag.Diagnostics) {
	return v.AsStructSliceT(ctx)
}

func (v ListItemsSet) AsStructSliceT(ctx context.Context) ([]ListItemModel, diag.Diagnostics) {
	ts := []ListItemModel{}
	if len(v.SetValue.Elements()) == 0 {
		return ts, nil
	}
	diags := v.SetValue.ElementsAs(ctx, &ts, true)
	if diags.HasError() {
		return nil, diags
	}
	return ts, nil
}

func (v ListItemsSet) Value(ctx context.Context) ([]attr.Value, diag.Diagnostics) {
	return v.SetValue.Elements(), nil
}

func NullListItemsSet(ctx context.Context) ListItemsSet {
	return ListItemsSet{SetValue: customfield.NullObjectSet[ListItemModel](ctx).SetValue}
}

func UnknownListItemsSet(ctx context.Context) ListItemsSet {
	return ListItemsSet{SetValue: customfield.UnknownObjectSet[ListItemModel](ctx).SetValue}
}

// NewListItemsSet builds the same value as customfield.NewObjectSet. It
// mirrors basetypes.NewSetValueFrom, but uses the indexed duplicate check. If
// anything at all is reported, the original constructor is used instead so
// the returned value and diagnostics are identical.
func NewListItemsSet(ctx context.Context, values []ListItemModel) (ListItemsSet, diag.Diagnostics) {
	if v, ok := newListItemsSetFast(ctx, values); ok {
		return v, nil
	}

	r, diags := customfield.NewObjectSet[ListItemModel](ctx, values)
	return ListItemsSet{SetValue: r.SetValue}, diags
}

func newListItemsSetFast(ctx context.Context, values []ListItemModel) (ListItemsSet, bool) {
	if values == nil {
		return ListItemsSet{}, false
	}

	elemType, d := listItemElemType(ctx)
	if d.HasError() {
		return ListItemsSet{}, false
	}
	setType := basetypes.SetType{ElemType: elemType}
	tfType := setType.TerraformType(ctx)

	// Leave element types with their own validation to the original constructor.
	if _, ok := attr.Type(elemType).(xattr.TypeWithValidate); ok {
		return ListItemsSet{}, false
	}

	tfElems := make([]tftypes.Value, 0, len(values))
	for _, item := range values {
		obj, d := basetypes.NewObjectValueFrom(ctx, elemType.AttrTypes, item)
		if len(d) > 0 {
			return ListItemsSet{}, false
		}
		if _, ok := attr.Value(obj).(xattr.ValidateableAttribute); ok {
			return ListItemsSet{}, false
		}
		tfVal, err := obj.ToTerraformValue(ctx)
		if err != nil {
			return ListItemsSet{}, false
		}
		tfElems = append(tfElems, tfVal)
	}

	if err := tftypes.ValidateValue(tfType, tfElems); err != nil {
		return ListItemsSet{}, false
	}
	tfVal := tftypes.NewValue(tfType, tfElems)

	attrVal, err := setType.ValueFromTerraform(ctx, tfVal)
	if err != nil {
		return ListItemsSet{}, false
	}
	setValue, ok := attrVal.(basetypes.SetValue)
	if !ok {
		return ListItemsSet{}, false
	}
	if _, ok := attrVal.(xattr.ValidateableAttribute); ok {
		return ListItemsSet{}, false
	}
	if !listItemsSetDistinct(elemType, tfVal) {
		return ListItemsSet{}, false
	}

	return ListItemsSet{SetValue: setValue}, true
}

func listItemElemType(ctx context.Context) (types.ObjectType, diag.Diagnostics) {
	var diags diag.Diagnostics
	attrTypes, d := customfield.StructToAttributes[ListItemModel](ctx)
	diags.Append(d...)
	if diags.HasError() {
		return types.ObjectType{}, diags
	}
	return types.ObjectType{AttrTypes: attrTypes}, diags
}

// listItemsSetEqualIndexed computes basetypes.SetValue.Equal using a
// fingerprint index. ok is false when the fast path does not apply, and the
// caller must then use the original implementation.
func listItemsSetEqualIndexed(a, b basetypes.SetValue) (eq bool, ok bool) {
	ctx := context.Background()

	// Everything except the known/known element comparison is cheap in the
	// original implementation.
	if a.IsNull() || a.IsUnknown() || b.IsNull() || b.IsUnknown() {
		return false, false
	}
	aElems, bElems := a.Elements(), b.Elements()
	if len(aElems) != len(bElems) || len(aElems) < 2 {
		return false, false
	}
	at, bt := a.ElementType(ctx), b.ElementType(ctx)
	if at == nil || bt == nil || !at.Equal(bt) {
		return false, false
	}

	index, ok := listItemsFingerprintIndex(ctx, bElems)
	if !ok {
		return false, false
	}

	var buf []byte
	for _, elem := range aElems {
		tv, err := elem.ToTerraformValue(ctx)
		if err != nil {
			return false, false
		}
		buf, ok = appendListItemFingerprint(buf[:0], tv)
		if !ok {
			return false, false
		}
		found := false
		for _, idx := range index[string(buf)] {
			if elem.Equal(bElems[idx]) {
				found = true
				break
			}
		}
		if !found {
			return false, true
		}
	}
	return true, true
}

// listItemsSetSameElements reports whether proposed and prior are known,
// fully known sets that contain exactly the same distinct elements. That is
// the only case in which the framework's element-by-element semantic
// equality walk would leave the value equal to the prior value.
func listItemsSetSameElements(ctx context.Context, proposed, prior basetypes.SetValue) bool {
	if proposed.IsNull() || proposed.IsUnknown() || prior.IsNull() || prior.IsUnknown() {
		return false
	}
	pt, rt := proposed.ElementType(ctx), prior.ElementType(ctx)
	if pt == nil || rt == nil || !pt.Equal(rt) {
		return false
	}
	pElems, rElems := proposed.Elements(), prior.Elements()
	if len(pElems) != len(rElems) {
		return false
	}

	index, ok := listItemsFingerprintIndex(ctx, rElems)
	if !ok || len(index) != len(rElems) {
		// Unsupported values or duplicate prior elements.
		return false
	}

	seen := make(map[string]struct{}, len(pElems))
	var buf []byte
	for _, elem := range pElems {
		tv, err := elem.ToTerraformValue(ctx)
		if err != nil || !tv.IsFullyKnown() {
			return false
		}
		buf, ok = appendListItemFingerprint(buf[:0], tv)
		if !ok {
			return false
		}
		key := string(buf)
		if _, dup := seen[key]; dup {
			return false
		}
		seen[key] = struct{}{}
		idxs := index[key]
		if len(idxs) != 1 || !elem.Equal(rElems[idxs[0]]) {
			return false
		}
	}
	return true
}

// listItemsFingerprintIndex maps element fingerprints to element indexes.
// It only accepts fully known elements.
func listItemsFingerprintIndex(ctx context.Context, elems []attr.Value) (map[string][]int, bool) {
	index := make(map[string][]int, len(elems))
	var buf []byte
	for i, elem := range elems {
		tv, err := elem.ToTerraformValue(ctx)
		if err != nil || !tv.IsFullyKnown() {
			return nil, false
		}
		var ok bool
		buf, ok = appendListItemFingerprint(buf[:0], tv)
		if !ok {
			return nil, false
		}
		key := string(buf)
		index[key] = append(index[key], i)
	}
	return index, true
}

// appendListItemFingerprint appends an encoding of v to b such that two values
// of the same type have the same encoding if and only if they are equal under
// tftypes.Value.Equal (and so under the corresponding attr.Value Equal for the
// list item types). It returns false for values it does not support; callers
// then fall back to the original comparison.
func appendListItemFingerprint(b []byte, v tftypes.Value) ([]byte, bool) {
	if !v.IsKnown() {
		return append(b, 'U'), true
	}
	if v.IsNull() {
		return append(b, 'N'), true
	}

	typ := v.Type()
	switch {
	case typ.Is(tftypes.String):
		var s string
		if err := v.As(&s); err != nil {
			return b, false
		}
		b = append(b, 'S')
		b = strconv.AppendInt(b, int64(len(s)), 10)
		b = append(b, ':')
		return append(b, s...), true

	case typ.Is(tftypes.Bool):
		var x bool
		if err := v.As(&x); err != nil {
			return b, false
		}
		if x {
			return append(b, 'T'), true
		}
		return append(b, 'F'), true

	case typ.Is(tftypes.Number):
		f := new(big.Float)
		if err := v.As(&f); err != nil || f == nil {
			return b, false
		}
		// Only exact integers are supported (asn and status_code are Int64).
		// -0 and 0 are equal and both encode as 0.
		if !f.IsInt() {
			return b, false
		}
		i, acc := f.Int64()
		if acc != big.Exact {
			return b, false
		}
		b = append(b, 'I')
		b = strconv.AppendInt(b, i, 10)
		return append(b, ';'), true

	case typ.Is(tftypes.Object{}):
		var attrs map[string]tftypes.Value
		if err := v.As(&attrs); err != nil {
			return b, false
		}
		names := make([]string, 0, len(attrs))
		for name := range attrs {
			names = append(names, name)
		}
		sort.Strings(names)
		b = append(b, 'O')
		b = strconv.AppendInt(b, int64(len(names)), 10)
		b = append(b, '{')
		for _, name := range names {
			b = strconv.AppendInt(b, int64(len(name)), 10)
			b = append(b, ':')
			b = append(b, name...)
			var ok bool
			b, ok = appendListItemFingerprint(b, attrs[name])
			if !ok {
				return b, false
			}
		}
		return append(b, '}'), true
	}

	return b, false
}
