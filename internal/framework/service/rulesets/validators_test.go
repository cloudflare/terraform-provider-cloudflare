package rulesets

import (
	"context"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestEdgeTTLValidation(t *testing.T) {
	t.Parallel()

	var edgeValidator EdgeTTLValidator
	t.Run("when validating edge ttl", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		t.Run("override_origin mode is specified", func(t *testing.T) {
			t.Parallel()

			t.Run("errors when no ttl is passed", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructEdgeTTLObjectRequest("override_origin", nil)
				edgeValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: diag.Diagnostics{
						diag.NewAttributeErrorDiagnostic(path.Root("edge_ttl"), "invalid configuration", "using mode 'override_origin' requires setting a default for ttl"),
					},
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})

			t.Run("errors when invalid ttl is passed", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructEdgeTTLObjectRequest("override_origin", big.NewFloat(-1))
				edgeValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: diag.Diagnostics{
						diag.NewAttributeErrorDiagnostic(path.Root("edge_ttl"), "invalid configuration", "using mode 'override_origin' requires setting a default for ttl"),
					},
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})

			t.Run("passes valid ttl values", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructEdgeTTLObjectRequest("override_origin", big.NewFloat(10))
				edgeValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: nil,
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})
		})

		t.Run("respect_origin mode is specified", func(t *testing.T) {
			t.Parallel()

			t.Run("passes with ttl values", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructEdgeTTLObjectRequest("respect_origin", big.NewFloat(1))
				edgeValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: nil,
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})
		})
	})
}

func TestBrowserTTLValidation(t *testing.T) {
	t.Parallel()

	var browserValidator BrowserTTLValidator
	t.Run("when validating browser ttl", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		t.Run("override_origin mode is specified", func(t *testing.T) {
			t.Parallel()

			t.Run("errors when no ttl is passed", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructBrowserTTLObjectRequest("override_origin", nil)
				browserValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: diag.Diagnostics{
						diag.NewAttributeErrorDiagnostic(path.Root("browser_ttl"), "invalid configuration", "using mode 'override_origin' requires setting a default for ttl"),
					},
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})

			t.Run("errors when invalid ttl is passed", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructBrowserTTLObjectRequest("override_origin", big.NewFloat(-1))
				browserValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: diag.Diagnostics{
						diag.NewAttributeErrorDiagnostic(path.Root("browser_ttl"), "invalid configuration", "using mode 'override_origin' requires setting a default for ttl"),
					},
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})

			t.Run("passes valid ttl values", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructBrowserTTLObjectRequest("override_origin", big.NewFloat(10))
				browserValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: nil,
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})
		})

		t.Run("respect_origin mode is specified", func(t *testing.T) {
			t.Parallel()

			t.Run("passes with ttl values", func(t *testing.T) {
				t.Parallel()

				resp := &validator.ObjectResponse{}
				req := constructBrowserTTLObjectRequest("respect_origin", big.NewFloat(1))
				browserValidator.ValidateObject(ctx, req, resp)

				expected := &validator.ObjectResponse{
					Diagnostics: nil,
				}
				if diff := cmp.Diff(resp, expected); diff != "" {
					t.Errorf("unexpected difference: %s", diff)
				}
			})
		})
	})
}

func TestCacheKeyQueryStrings(t *testing.T) {
	t.Parallel()

	var cacheKeyValidator InvalidWildCardValidator
	t.Run("when including query strings in cache keys", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		t.Run("errors are thrown for full wildcards", func(t *testing.T) {
			t.Parallel()

			resp := &validator.SetResponse{}
			req := constructCacheKeyObjectRequest("include", []string{"1", "*", "foobarbaz"})
			cacheKeyValidator.ValidateSet(ctx, req, resp)

			expected := &validator.SetResponse{
				Diagnostics: diag.Diagnostics{
					diag.NewAttributeErrorDiagnostic(path.Root("include"), "invalid value", "full wildcards should use the ignore field instead, value: *"),
				},
			}
			if diff := cmp.Diff(resp, expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})

		t.Run("passes for every other value", func(t *testing.T) {
			t.Parallel()

			resp := &validator.SetResponse{}
			req := constructCacheKeyObjectRequest("include", []string{"1", "blah/*", "foobarbaz"})
			cacheKeyValidator.ValidateSet(ctx, req, resp)

			expected := &validator.SetResponse{
				Diagnostics: nil,
			}
			if diff := cmp.Diff(resp, expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	})

	t.Run("when excluding query strings in cache keys", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		t.Run("errors are thrown for full wildcards", func(t *testing.T) {
			t.Parallel()

			resp := &validator.SetResponse{}
			req := constructCacheKeyObjectRequest("exclude", []string{"1", "*", "foobarbaz"})
			cacheKeyValidator.ValidateSet(ctx, req, resp)

			expected := &validator.SetResponse{
				Diagnostics: diag.Diagnostics{
					diag.NewAttributeErrorDiagnostic(path.Root("exclude"), "invalid value", "full wildcards should use the ignore field instead, value: *"),
				},
			}
			if diff := cmp.Diff(resp, expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})

		t.Run("passes for every other value", func(t *testing.T) {
			t.Parallel()

			resp := &validator.SetResponse{}
			req := constructCacheKeyObjectRequest("exclude", []string{"1", "blah/*", "foobarbaz"})
			cacheKeyValidator.ValidateSet(ctx, req, resp)

			expected := &validator.SetResponse{
				Diagnostics: nil,
			}
			if diff := cmp.Diff(resp, expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	})
}

func constructStatusTTLObject() (tftypes.Value, attr.Value, types.ListType) {
	tftype := tftypes.NewValue(
		tftypes.List{
			ElementType: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"status_code": tftypes.Number,
					"value":       tftypes.Number,
				},
			},
		},
		[]tftypes.Value{tftypes.NewValue(
			tftypes.Object{AttributeTypes: map[string]tftypes.Type{
				"status_code": tftypes.Number,
				"value":       tftypes.Number,
			}},
			map[string]tftypes.Value{
				"status_code": tftypes.NewValue(
					tftypes.Number,
					100,
				),
				"value": tftypes.NewValue(
					tftypes.Number,
					100,
				),
			},
		)},
	)

	attributes, _ := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{"status_code": types.Int64Type, "value": types.Int64Type},
		},
		[]attr.Value{
			types.ObjectValueMust(
				map[string]attr.Type{"status_code": types.Int64Type, "value": types.Int64Type},
				map[string]attr.Value{"status_code": types.Int64Value(100), "value": types.Int64Value(100)},
			),
		},
	)

	listType := types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: map[string]attr.Type{"status_code": types.Int64Type, "value": types.Int64Type},
		},
	}
	return tftype, attributes, listType
}

func constructEdgeTTLObjectRequest(mode string, ttl *big.Float) validator.ObjectRequest {
	var objectValue basetypes.ObjectValue

	statusCodeTfTypes, statusCodeAttr, listType := constructStatusTTLObject()
	if ttl == nil {
		objectValue = types.ObjectValueMust(
			map[string]attr.Type{"mode": types.StringType, "status_code_ttl": listType, "default": types.Int64Type},
			map[string]attr.Value{"mode": types.StringValue(mode), "status_code_ttl": statusCodeAttr, "default": types.Int64Null()},
		)
	} else {
		intTTL, _ := ttl.Int64()
		objectValue = types.ObjectValueMust(
			map[string]attr.Type{"mode": types.StringType, "status_code_ttl": listType, "default": types.Int64Type},
			map[string]attr.Value{"mode": types.StringValue(mode), "status_code_ttl": statusCodeAttr, "default": types.Int64Value(intTTL)},
		)
	}

	return validator.ObjectRequest{
		Path:        path.Root("edge_ttl"),
		ConfigValue: objectValue,
		Config: tfsdk.Config{
			Raw: tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"mode":    tftypes.String,
						"default": tftypes.Number,
						"status_code_ttl": tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"status_code": tftypes.Number,
									"value":       tftypes.Number,
								},
							},
						},
					},
				},
				map[string]tftypes.Value{
					"mode": tftypes.NewValue(
						tftypes.String,
						mode,
					),
					"default": tftypes.NewValue(
						tftypes.Number,
						ttl,
					),
					"status_code_ttl": statusCodeTfTypes,
				},
			),
		},
	}
}

func constructBrowserTTLObjectRequest(mode string, ttl *big.Float) validator.ObjectRequest {
	var objectValue basetypes.ObjectValue

	if ttl == nil {
		objectValue = types.ObjectValueMust(
			map[string]attr.Type{"mode": types.StringType, "default": types.Int64Type},
			map[string]attr.Value{"mode": types.StringValue(mode), "default": types.Int64Null()},
		)
	} else {
		intTTL, _ := ttl.Int64()
		objectValue = types.ObjectValueMust(
			map[string]attr.Type{"mode": types.StringType, "default": types.Int64Type},
			map[string]attr.Value{"mode": types.StringValue(mode), "default": types.Int64Value(intTTL)},
		)
	}

	return validator.ObjectRequest{
		Path:        path.Root("browser_ttl"),
		ConfigValue: objectValue,
		Config: tfsdk.Config{
			Raw: tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"mode":    tftypes.String,
						"default": tftypes.Number,
					},
				},
				map[string]tftypes.Value{
					"mode": tftypes.NewValue(
						tftypes.String,
						mode,
					),
					"default": tftypes.NewValue(
						tftypes.Number,
						ttl,
					),
				},
			),
		},
	}
}

func constructCacheKeyObjectRequest(requestPath string, values []string) validator.SetRequest {
	var setValue basetypes.SetValue
	stringValues := make([]attr.Value, len(values))
	tfValues := make([]tftypes.Value, len(values))

	for i, v := range values {
		stringValues[i] = types.StringValue(v)
		tfValues[i] = tftypes.NewValue(tftypes.String, v)
	}

	setValue = types.SetValueMust(types.StringType, stringValues)

	return validator.SetRequest{
		Path:        path.Root(requestPath),
		ConfigValue: setValue,
		Config: tfsdk.Config{
			Raw: tftypes.NewValue(
				tftypes.Set{
					ElementType: tftypes.String,
				},
				tfValues,
			),
		},
	}
}
