package test_helpers

// This file is intentionally NOT code-generated.
// It provides ValidateMigrationModelSchemaIntegrity for use in migration parity tests.

import (
	"fmt"
	"log"
	"reflect"

	ds "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rs "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ValidateMigrationModelSchemaIntegrity checks that every attribute in the live
// resource schema has a corresponding tfsdk-tagged field in the migration target
// model, and vice versa — recursing into all nested attributes at every depth.
//
// Unlike ValidateResourceModelSchemaIntegrity, this function does NOT check Go
// type name identity for customfield.NestedObject[T] fields. Migration target
// structs use locally-defined types (e.g. customfield.NestedObject[TargetFooModel])
// while the live schema encodes the live package type
// (e.g. customfield.NestedObject[FooModel]). The names differ but the fields
// inside are structurally equivalent, so the type-name check would always produce
// false positives for migration structs.
func ValidateMigrationModelSchemaIntegrity(model any, schema rs.Schema) codingerrors {
	path := []string{}
	attributes := map[string]attrlike{}
	for name, attr := range schema.Attributes {
		attributes[name] = &attri{attr}
	}
	return migrationWalkAttributes(path, reflect.TypeOf(model), attributes)
}

func migrationWalkAttributes(path []string, model reflect.Type, attributes map[string]attrlike) (errs codingerrors) {
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
			errs = append(errs, migrationWalk(index, attr, field.Type)...)
		} else {
			errs = append(errs, &mismatch{path: index, expected: fmt.Sprintf("%T", attr.val()), actual: nil})
		}
	}
	return
}

func migrationWalk(path path, attribute attrlike, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	kind := model.Kind()

	switch a := attribute.val().(type) {
	case ds.BoolAttribute, rs.BoolAttribute, basetypes.BoolType:
		if kind != reflect.Bool && !model.Implements(reflect.TypeOf((*basetypes.BoolValuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "Bool", actual: model})
		}
	case ds.Int64Attribute, rs.Int64Attribute, basetypes.Int64Type:
		if kind != reflect.Int && kind != reflect.Int64 && !model.Implements(reflect.TypeOf((*basetypes.Int64Valuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "Int64", actual: model})
		}
	case ds.Float64Attribute, rs.Float64Attribute, basetypes.Float64Type:
		if kind != reflect.Float64 && !model.Implements(reflect.TypeOf((*basetypes.Float64Valuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "Float64", actual: model})
		}
	case ds.StringAttribute, rs.StringAttribute, basetypes.StringType:
		if kind != reflect.String && !model.Implements(reflect.TypeOf((*basetypes.StringValuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "String", actual: model})
		}
	case ds.ListAttribute, rs.ListAttribute:
		if kind != reflect.Slice && !model.Implements(reflect.TypeOf((*basetypes.ListValuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "List", actual: model})
		}
	case ds.SetAttribute, rs.SetAttribute:
		if kind != reflect.Slice && !model.Implements(reflect.TypeOf((*basetypes.SetValuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "Set", actual: model})
		}
	case ds.MapAttribute, rs.MapAttribute:
		if kind != reflect.Map && !model.Implements(reflect.TypeOf((*basetypes.MapValuable)(nil)).Elem()) {
			return append(errs, &mismatch{path: path, expected: "Map", actual: model})
		}
	case ds.SingleNestedAttribute, rs.SingleNestedAttribute:
		return append(errs, migrationWalkNested(path, a.(ds.NestedAttribute), kind, model)...)
	case ds.ListNestedAttribute, rs.ListNestedAttribute:
		return append(errs, migrationWalkNested(path, a.(ds.NestedAttribute), kind, model)...)
	case ds.SetNestedAttribute, rs.SetNestedAttribute:
		return append(errs, migrationWalkNested(path, a.(ds.NestedAttribute), kind, model)...)
	case ds.MapNestedAttribute, rs.MapNestedAttribute:
		return append(errs, migrationWalkNested(path, a.(ds.NestedAttribute), kind, model)...)
	default:
		log.Printf("migrationWalk: Unexpected attribute type at %q: %T", path.pathName(), a)
	}
	return
}

func migrationWalkNested(path path, attribute ds.NestedAttribute, kind reflect.Kind, model reflect.Type) (errs codingerrors) {
	model = deref(model)
	idx, exp := []string{}, "{...}"
	basetype, reflectype := reflect.TypeOf((*basetypes.ObjectValuable)(nil)), reflect.Struct

	switch attribute.(type) {
	case ds.ListNestedAttribute, rs.ListNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.ListValuable)(nil)), reflect.Slice
		idx, exp = []string{"[]"}, "List[...]"
	case ds.SetNestedAttribute, rs.SetNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.SetValuable)(nil)), reflect.Slice
		idx, exp = []string{"<>"}, "Set<...>"
	case ds.MapNestedAttribute, rs.MapNestedAttribute:
		basetype, reflectype = reflect.TypeOf((*basetypes.MapValuable)(nil)), reflect.Map
		idx, exp = []string{"{}"}, "Map{...}"
	}

	if kind == reflectype && reflectype != reflect.Struct {
		errs = append(errs, checkMapKeys(path, model)...)
		model = deref(model.Elem())
	} else if model.Implements(basetype.Elem()) {
		// Skip checkCustom — migration structs use local Target* types whose names
		// differ from the live package types, but the fields inside are equivalent.
		model = genericParam(model, 0)
	}

	if model == nil || model.Kind() != reflect.Struct {
		return append(errs, &mismatch{path: path, expected: exp, actual: model})
	}

	attributes := map[string]attrlike{}
	for name, nested := range attribute.GetNestedObject().GetAttributes() {
		attributes[name] = &attri{nested}
	}
	return append(errs, migrationWalkAttributes(append(path, idx...), model, attributes)...)
}
