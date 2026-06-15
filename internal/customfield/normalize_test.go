package customfield

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type testHostnameModel struct {
	URLHostname          types.String `tfsdk:"url_hostname"`
	ExcludeExactHostname types.Bool   `tfsdk:"exclude_exact_hostname"`
}

type testEmptyMarkerModel struct{}

type testSingleFieldModel struct {
	Name types.String `tfsdk:"name"`
}

func TestNullifyEmptyObject_AllNullFields(t *testing.T) {
	ctx := context.Background()

	// Create a known object where all fields are null (simulates API returning {})
	obj := NewObjectMust(ctx, &testHostnameModel{
		URLHostname:          types.StringNull(),
		ExcludeExactHostname: types.BoolNull(),
	})

	if obj.IsNull() {
		t.Fatal("precondition: object should be known (not null) before normalization")
	}

	result := NullifyEmptyObject(ctx, obj)

	if !result.IsNull() {
		t.Error("expected NullifyEmptyObject to return null when all attributes are null")
	}
}

func TestNullifyEmptyObject_SomeFieldsSet(t *testing.T) {
	ctx := context.Background()

	obj := NewObjectMust(ctx, &testHostnameModel{
		URLHostname:          types.StringValue("example.com"),
		ExcludeExactHostname: types.BoolNull(),
	})

	result := NullifyEmptyObject(ctx, obj)

	if result.IsNull() {
		t.Error("expected NullifyEmptyObject to preserve object when some attributes are set")
	}
}

func TestNullifyEmptyObject_AllFieldsSet(t *testing.T) {
	ctx := context.Background()

	obj := NewObjectMust(ctx, &testHostnameModel{
		URLHostname:          types.StringValue("example.com"),
		ExcludeExactHostname: types.BoolValue(true),
	})

	result := NullifyEmptyObject(ctx, obj)

	if result.IsNull() {
		t.Error("expected NullifyEmptyObject to preserve object when all attributes are set")
	}
}

func TestNullifyEmptyObject_AlreadyNull(t *testing.T) {
	ctx := context.Background()

	obj := NullObject[testHostnameModel](ctx)

	result := NullifyEmptyObject(ctx, obj)

	if !result.IsNull() {
		t.Error("expected NullifyEmptyObject to return null for already-null input")
	}
}

func TestNullifyEmptyObject_AlreadyUnknown(t *testing.T) {
	ctx := context.Background()

	obj := UnknownObject[testHostnameModel](ctx)

	result := NullifyEmptyObject(ctx, obj)

	if !result.IsUnknown() {
		t.Error("expected NullifyEmptyObject to return unknown for already-unknown input")
	}
}

func TestNullifyEmptyObject_EmptyMarkerStruct(t *testing.T) {
	ctx := context.Background()

	// Empty marker struct (zero fields) -- presence IS the value.
	// NullifyEmptyObject must NOT convert this to null.
	obj := NewObjectMust(ctx, &testEmptyMarkerModel{})

	if obj.IsNull() {
		t.Fatal("precondition: empty marker object should be known (not null)")
	}

	result := NullifyEmptyObject(ctx, obj)

	if result.IsNull() {
		t.Error("expected NullifyEmptyObject to preserve empty marker struct (zero fields)")
	}
}

func TestNullifyEmptyObject_SingleFieldNull(t *testing.T) {
	ctx := context.Background()

	obj := NewObjectMust(ctx, &testSingleFieldModel{
		Name: types.StringNull(),
	})

	result := NullifyEmptyObject(ctx, obj)

	if !result.IsNull() {
		t.Error("expected NullifyEmptyObject to return null when single field is null")
	}
}

func TestNullifyEmptyObject_SingleFieldSet(t *testing.T) {
	ctx := context.Background()

	obj := NewObjectMust(ctx, &testSingleFieldModel{
		Name: types.StringValue("test"),
	})

	result := NullifyEmptyObject(ctx, obj)

	if result.IsNull() {
		t.Error("expected NullifyEmptyObject to preserve object when single field is set")
	}
}
