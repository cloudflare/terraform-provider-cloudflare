package test_helpers

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	ds "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rs "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ codingerror = (*mismatch)(nil)
var _ codingerror = (*bug)(nil)
var _ error = (*codingerrors)(nil)
var _ attrlike = (*att)(nil)
var _ attrlike = (*attri)(nil)

type codingerror interface {
	error
	id() string
}

type attrlike interface {
	val() any
	tag() string
	custom([]string) any
}

type bug struct {
	path []string
	msg  string
}

type mismatch struct {
	path     []string
	expected string
	tagging  bool
	tags     string
	actual   reflect.Type
}

func pathName(path []string) string {
	return "." + strings.Join(path, ".")
}

func (e *bug) Error() string {
	return fmt.Sprintf("bug at %q: %q", e.id(), e.msg)
}
func (e *mismatch) Error() string {
	found := "<MISSING>"
	if e.tagging {
		found = e.tags
	} else if e.actual != nil {
		found = e.actual.Name()
		if found == "" {
			found = e.actual.Kind().String()
		}
	}
	return fmt.Sprintf("mismatch at %q: expected %q, received %q", e.id(), e.expected, found)
}

func (e *bug) id() string {
	return pathName(e.path)
}
func (e *mismatch) id() string {
	return pathName(e.path)
}

type codingerrors []codingerror

func (errs *codingerrors) Ignore(t *testing.T, paths ...string) {
	acc := map[string]any{}
	for _, path := range paths {
		acc[path] = nil
	}
	*errs = slices.DeleteFunc(*errs, func(err codingerror) bool {
		id := err.id()
		_, ok := acc[id]
		if ok {
			delete(acc, id)
		}
		return ok
	})
	if len(acc) > 0 {
		t.Errorf("Superfluous ignore paths: %v", acc)
	}
}

func (errs *codingerrors) Error() string {
	e := []string{}
	for _, err := range *errs {
		e = append(e, err.Error())
	}
	return strings.Join(e, "\n")
}

func (errs *codingerrors) Report(t *testing.T) {
	for _, err := range *errs {
		t.Error(err.Error())
	}
}

type att struct{ v attr.Type }
type attri struct{ v ds.Attribute }

func (a *att) val() any {
	return a.v
}
func (a *attri) val() any {
	return a.v
}

func (a *att) tag() string {
	return ""
}
func (a *attri) tag() string {
	if a.v.IsRequired() {
		return "required"
	}

	computed, optional := a.v.IsComputed(), a.v.IsOptional()
	if computed && optional {
		return "computed_optional"
	} else if computed {
		return "computed"
	} else if optional {
		return "optional"
	} else {
		return ""
	}
}

func (a *att) custom(_ []string) any {
	return nil
}
func (a *attri) custom(path []string) any {
	switch v := (a.v).(type) {
	case ds.BoolAttribute:
		return v.CustomType
	case rs.BoolAttribute:
		return v.CustomType

	case ds.Int64Attribute:
		return v.CustomType
	case rs.Int64Attribute:
		return v.CustomType

	case ds.Float64Attribute:
		return v.CustomType
	case rs.Float64Attribute:
		return v.CustomType

	case ds.NumberAttribute:
		return v.CustomType
	case rs.NumberAttribute:
		return v.CustomType

	case ds.StringAttribute:
		return v.CustomType
	case rs.StringAttribute:
		return v.CustomType

	case ds.ListAttribute:
		return v.CustomType
	case rs.ListAttribute:
		return v.CustomType

	case ds.SetAttribute:
		return v.CustomType
	case rs.SetAttribute:
		return v.CustomType

	case ds.MapAttribute:
		return v.CustomType
	case rs.MapAttribute:
		return v.CustomType

	case ds.ObjectAttribute:
		return v.CustomType
	case rs.ObjectAttribute:
		return v.CustomType

	case ds.SingleNestedAttribute:
		return v.CustomType
	case rs.SingleNestedAttribute:
		return v.CustomType

	case ds.ListNestedAttribute:
		return v.CustomType
	case rs.ListNestedAttribute:
		return v.CustomType

	case ds.SetNestedAttribute:
		return v.CustomType
	case rs.SetNestedAttribute:
		return v.CustomType

	case ds.MapNestedAttribute:
		return v.CustomType
	case rs.MapNestedAttribute:
		return v.CustomType

	default:
		log.Printf("custom: Unexpected attribute type at %q: %T", pathName(path), v)
		return nil
	}
}

func ValidateDataSourceModelSchemaIntegrity(model any, schema ds.Schema) codingerrors {
	path := []string{}
	attributes := map[string]attrlike{}
	for name, attr := range schema.Attributes {
		attributes[name] = &attri{attr}
	}
	return walkAttributes(path, reflect.TypeOf(model), attributes)
}

func ValidateResourceModelSchemaIntegrity(model any, schema rs.Schema) codingerrors {
	path := []string{}
	attributes := map[string]attrlike{}
	for name, attr := range schema.Attributes {
		attributes[name] = &attri{attr}
	}
	return walkAttributes(path, reflect.TypeOf(model), attributes)
}

func deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return deref(t.Elem())
	}
	return t
}

func checkTag(path []string, attr attrlike, field reflect.StructField) (errs codingerrors) {
	tag := attr.tag()
	if tag == "" {
		return
	}
	tagged, tags := false, []string{}

	for _, label := range []string{"path", "query", "json"} {
		if labelled, ok := field.Tag.Lookup(label); ok {
			tagged = true
			tags = append(tags, strings.Split(labelled, ",")...)
		}
	}

	if !tagged {
		return
	}
	for _, t := range tags {
		if t == tag {
			return
		}
	}
	return codingerrors{&mismatch{path: path, expected: tag, tagging: true, tags: strings.Join(tags, ",")}}
}

func checkCustom(path []string, attribute attrlike, model reflect.Type) (errs codingerrors) {
	custom := attribute.custom(path)
	if custom == nil {
		return
	}

	if a, ok := custom.(attr.Type); ok {
		custom = a.ValueType(context.TODO())
	}

	t := deref(reflect.TypeOf(custom))
	if model.ConvertibleTo(t) {
		return
	}

	return codingerrors{&mismatch{path: path, expected: t.Name(), actual: model}}
}

func walkAttributes(path []string, model reflect.Type, attributes map[string]attrlike) (errs codingerrors) {
	model = deref(model)
	if model.Kind() != reflect.Struct {
		return append(errs, &mismatch{path: path, expected: "Struct", actual: model})
	}

	path = append(path, "@"+model.Name())
	fields := map[string]reflect.StructField{}
	for i := 0; i < model.NumField(); i++ {
		field := model.Field(i)
		if tag, ok := field.Tag.Lookup("tfsdk"); ok {
			fields[tag] = field
		}
	}
	for name, attr := range attributes {
		index := append(path, name)
		if field, ok := fields[name]; ok {
			errs = append(append(errs, checkTag(index, attr, field)...), walk(index, attr, field.Type)...)
		} else {
			errs = append(errs, &mismatch{path: index, expected: fmt.Sprintf("%T", attr.val()), actual: nil})
		}
	}
	return
}

// workaround for lack of generic param in reflection
// we store an instance of the generics in struct itself
// https://github.com/golang/go/issues/54393
func genericParam(t reflect.Type, idx int) reflect.Type {
	t = deref(t)
	if t.Kind() != reflect.Struct {
		return nil
	}
	if t.NumField() <= idx {
		return nil
	}
	return t.Field(idx).Type
}

func checkMapKeys(path []string, model reflect.Type) (errs codingerrors) {
	if model.Kind() != reflect.Map {
		return
	}
	if deref(model.Key()).Kind() != reflect.String {
		return append(errs, &bug{path: path, msg: "checkMapKeys: Expected string keys"})
	}
	return
}

func walkCollection(path []string, collection attr.TypeWithElementType, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp, eltype := "[]", "List[T]", collection.ElementType()
	bstype, reflectype := reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice

	if eltype == nil {
		return append(errs, &bug{path: path, msg: "walkCollection: Unexpected nil ElementType"})
	}
	switch t := collection.(type) {
	case basetypes.ListType:
	case basetypes.SetType:
		idx, exp = "<>", "Set<T>"
		bstype = reflect.TypeOf((*basetypes.SetValuable)(nil))
	case basetypes.MapType:
		idx, exp, reflectype = "{}", "Map{T}", reflect.Map
		bstype = reflect.TypeOf((*basetypes.MapValuable)(nil))
	default:
		log.Printf("walkCollection: Unexpected attribute type at %q: %T", pathName(path), t)
		return
	}

	if kind == reflectype {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(bstype.Elem()) {
		model = genericParam(model, 0)
	} else {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}
	return append(errs, walk(append(path, idx), &att{eltype}, model)...)
}

func walkAttriCollection(path []string, attribute ds.Attribute, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp, eltype := "[]", "List[T]", (attr.Type)(nil)
	bstype, reflectype := reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice

	switch t := attribute.(type) {
	case ds.ListAttribute:
		eltype = t.ElementType
	case rs.ListAttribute:
		eltype = t.ElementType
	case ds.SetAttribute:
		idx, exp, eltype = "<>", "Set<T>", t.ElementType
		bstype = reflect.TypeOf((*basetypes.SetValuable)(nil))
	case rs.SetAttribute:
		idx, exp, eltype = "<>", "Set<T>", t.ElementType
		bstype = reflect.TypeOf((*basetypes.SetValuable)(nil))
	case ds.MapAttribute:
		idx, exp, eltype = "{}", "Map{T}", t.ElementType
		bstype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
	case rs.MapAttribute:
		idx, exp, eltype = "{}", "Map{T}", t.ElementType
		bstype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
	default:
		log.Printf("walkAttriCollection: Unexpected attribute type at %q: %T", pathName(path), t)
		return
	}

	if eltype == nil {
		return append(errs, &bug{path: path, msg: "walkAttriCollection: Unexpected nil ElementType"})
	}

	if kind == reflectype {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(bstype.Elem()) {
		for ty, t := range map[string]reflect.Type{
			"List": reflect.TypeOf(basetypes.ListValue{}),
			"Set":  reflect.TypeOf(basetypes.SetValue{}),
			"Map":  reflect.TypeOf(basetypes.MapValue{}),
		} {
			if model.AssignableTo(t) {
				return append(errs, &bug{path: path, msg: fmt.Sprintf("walkAttriCollection: Should use customfield.%s instead of %s", ty, t)})
			}
		}
		model = genericParam(model, 0)
	} else {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}
	return append(errs, walk(append(path, idx), &att{eltype}, model)...)
}

func walkNested(path []string, attribute ds.NestedAttribute, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp := []string{}, "{...}"
	basetype, reflectype := reflect.TypeOf((*basetypes.ObjectValuable)(nil)), reflect.Struct

	switch a := attribute.(type) {
	case ds.SingleNestedAttribute, rs.SingleNestedAttribute:
	case ds.ListNestedAttribute, rs.ListNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice
		idx, exp = []string{"[]"}, "List[...]"
	case ds.SetNestedAttribute, rs.SetNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.SetValuable)(nil)), reflect.Slice
		idx, exp = []string{"<>"}, "Set<...>"
	case ds.MapNestedAttribute, rs.MapNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
		idx, exp = []string{"{}"}, "Map{...}"
	default:
		log.Printf("walkNested: Unexpected attribute type %q: %T", pathName(path), a)
		return
	}

	if kind == reflectype && reflectype != reflect.Struct {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(basetype.Elem()) {
		errs = append(errs, checkCustom(append(path, idx...), &attri{attribute}, model)...)
		model = genericParam(model, 0)
	}

	if model.Kind() != reflect.Struct {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}

	attributes := map[string]attrlike{}
	for name, nested := range attribute.GetNestedObject().GetAttributes() {
		attributes[name] = &attri{nested}
	}
	return append(errs, walkAttributes(append(path, idx...), model, attributes)...)
}

func walk(path []string, attribute attrlike, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	kind := model.Kind()

	switch a := attribute.val().(type) {
	case ds.DynamicAttribute, rs.DynamicAttribute:
		if !model.ConvertibleTo(reflect.TypeOf(basetypes.DynamicValue{})) {
			return append(errs, &mismatch{path: path, expected: "Dynamic", actual: model})
		}
	case ds.BoolAttribute, rs.BoolAttribute, basetypes.BoolType:
		if model.Implements(reflect.TypeOf((*basetypes.BoolValuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.Bool {
			return append(errs, &mismatch{path: path, expected: "Bool", actual: model})
		}
	case ds.Int64Attribute, rs.Int64Attribute, basetypes.Int64Type:
		if model.Implements(reflect.TypeOf((*basetypes.Int64Valuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.Int && kind != reflect.Int64 {
			return append(errs, &mismatch{path: path, expected: "Int", actual: model})
		}
	case ds.Float64Attribute, rs.Float64Attribute, basetypes.Float64Type:
		if model.Implements(reflect.TypeOf((*basetypes.Float64Valuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.Float64 {
			return append(errs, &mismatch{path: path, expected: "Float", actual: model})
		}
	case ds.NumberAttribute, rs.NumberAttribute, basetypes.NumberType:
		if model.Implements(reflect.TypeOf((*basetypes.NumberValuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if model.ConvertibleTo(reflect.TypeOf(big.Int{})) || model.ConvertibleTo(reflect.TypeOf(big.Float{})) {
			return append(errs, &mismatch{path: path, expected: "Number", actual: model})
		}
	case timetypes.RFC3339Type:
		if !model.ConvertibleTo(reflect.TypeOf(timetypes.RFC3339{})) {
			return append(errs, &mismatch{path: path, expected: "Time", actual: model})
		}
	case jsontypes.NormalizedType:
		if !model.ConvertibleTo(reflect.TypeOf(jsontypes.Normalized{})) {
			return append(errs, &mismatch{path: path, expected: "JSON", actual: model})
		}
	case ds.StringAttribute, rs.StringAttribute, basetypes.StringType:
		if model.Implements(reflect.TypeOf((*basetypes.StringValuable)(nil)).Elem()) {
			return append(errs, checkCustom(path, attribute, model)...)
		} else if kind != reflect.String {
			return append(errs, &mismatch{path: path, expected: "String", actual: model})
		}
	case basetypes.ListType, basetypes.SetType, basetypes.MapType:
		return append(errs, walkCollection(path, a.(attr.TypeWithElementType), kind, model)...)
	case ds.ListAttribute, ds.SetAttribute, ds.MapAttribute, rs.ListAttribute, rs.SetAttribute, rs.MapAttribute:
		return append(errs, walkAttriCollection(path, a.(ds.Attribute), kind, model)...)
	case ds.ObjectAttribute, rs.ObjectAttribute, basetypes.ObjectType:
		if model.Implements(reflect.TypeOf((*basetypes.ObjectValuable)(nil)).Elem()) {
			errs = append(errs, checkCustom(path, attribute, model)...)
			model = genericParam(model, 0)
		} else if kind != reflect.Struct {
			return append(errs, &mismatch{path: path, expected: "Object", actual: model})
		}
		attributes := map[string]attrlike{}
		switch objt := a.(type) {
		case ds.ObjectAttribute:
			for name, nested := range objt.AttributeTypes {
				attributes[name] = &att{nested}
			}
		case rs.ObjectAttribute:
			for name, nested := range objt.AttributeTypes {
				attributes[name] = &att{nested}
			}
		case basetypes.ObjectType:
			for name, nested := range objt.AttrTypes {
				attributes[name] = &att{nested}
			}
		}
		return append(errs, walkAttributes(path, model, attributes)...)
	case ds.SingleNestedAttribute, ds.ListNestedAttribute, ds.MapNestedAttribute, ds.SetNestedAttribute, rs.SingleNestedAttribute, rs.SetNestedAttribute, rs.ListNestedAttribute, rs.MapNestedAttribute:
		return append(errs, walkNested(path, a.(ds.NestedAttribute), kind, model)...)
	default:
		log.Printf("walk: Unexpected attribute type at %q: %T", pathName(path), a)
	}

	return
}
