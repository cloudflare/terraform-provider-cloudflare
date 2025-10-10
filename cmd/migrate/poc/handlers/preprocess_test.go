package handlers_test

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/handlers"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
)

// MockResourceTransformer is a test double for ResourceTransformer
type MockResourceTransformer struct {
	resourceType    string
	preprocessCalls int
	preprocessFunc  func(content string) string
}

func (m *MockResourceTransformer) CanHandle(resourceType string) bool {
	return resourceType == m.resourceType
}

func (m *MockResourceTransformer) GetResourceType() string {
	return m.resourceType
}

func (m *MockResourceTransformer) TransformConfig(block *hclwrite.Block) (*interfaces.TransformResult, error) {
	return nil, nil
}

func (m *MockResourceTransformer) TransformState(json gjson.Result, resourcePath string) (string, error) {
	return "", nil
}

func (m *MockResourceTransformer) Preprocess(content string) string {
	m.preprocessCalls++
	if m.preprocessFunc != nil {
		return m.preprocessFunc(content)
	}
	return content
}

func TestPreprocessHandler(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		resources      []MockResourceTransformer
		expectedOutput string
	}{
		{
			name:           "No registered resources",
			input:          `resource "cloudflare_record" "test" {}`,
			resources:      []MockResourceTransformer{},
			expectedOutput: `resource "cloudflare_record" "test" {}`,
		},
		{
			name:  "Single resource preprocessor",
			input: `resource "old_resource" "test" {}`,
			resources: []MockResourceTransformer{
				{
					resourceType: "old_resource",
					preprocessFunc: func(content string) string {
						return "resource \"new_resource\" \"test\" {}"
					},
				},
			},
			expectedOutput: `resource "new_resource" "test" {}`,
		},
		{
			name:  "Multiple resource preprocessors applied in order",
			input: `resource "resource_a" "test" {} resource "resource_b" "test2" {}`,
			resources: []MockResourceTransformer{
				{
					resourceType: "resource_a",
					preprocessFunc: func(content string) string {
						// First preprocessor changes resource_a to resource_x
						return "resource \"resource_x\" \"test\" {} resource \"resource_b\" \"test2\" {}"
					},
				},
				{
					resourceType: "resource_b",
					preprocessFunc: func(content string) string {
						// Second preprocessor changes resource_b to resource_y
						return "resource \"resource_x\" \"test\" {} resource \"resource_y\" \"test2\" {}"
					},
				},
			},
			expectedOutput: `resource "resource_x" "test" {} resource "resource_y" "test2" {}`,
		},
		{
			name:  "Preprocessor that returns content unchanged",
			input: `resource "some_resource" "test" {}`,
			resources: []MockResourceTransformer{
				{
					resourceType: "other_resource",
					preprocessFunc: func(content string) string {
						// This preprocessor doesn't find anything to change
						return content
					},
				},
			},
			expectedOutput: `resource "some_resource" "test" {}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create registry and register mock resources
			reg := registry.NewStrategyRegistry()
			for i := range tt.resources {
				// Need to get pointer to avoid loop variable issues
				resource := &tt.resources[i]
				reg.Register(resource)
			}

			// Create handler
			handler := handlers.NewPreprocessHandler(reg)

			// Create context
			ctx := &interfaces.TransformContext{
				Content: []byte(tt.input),
			}

			// Process
			result, err := handler.Handle(ctx)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check output
			if string(result.Content) != tt.expectedOutput {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tt.expectedOutput, string(result.Content))
			}

			// Verify all preprocessors were called
			for i, resource := range tt.resources {
				if resource.preprocessCalls != 1 {
					t.Errorf("Resource %d preprocessor called %d times, expected 1", i, resource.preprocessCalls)
				}
			}
		})
	}
}

func TestPreprocessHandlerCallsAllRegisteredPreprocessors(t *testing.T) {
	// This test verifies that ALL registered preprocessors are called,
	// not just those whose resource types are found in the content

	callOrder := []string{}

	resource1 := &MockResourceTransformer{
		resourceType: "resource_type_1",
		preprocessFunc: func(content string) string {
			callOrder = append(callOrder, "resource_1")
			return content
		},
	}

	resource2 := &MockResourceTransformer{
		resourceType: "resource_type_2",
		preprocessFunc: func(content string) string {
			callOrder = append(callOrder, "resource_2")
			return content
		},
	}

	resource3 := &MockResourceTransformer{
		resourceType: "resource_type_3",
		preprocessFunc: func(content string) string {
			callOrder = append(callOrder, "resource_3")
			return content
		},
	}

	// Create registry and register resources
	reg := registry.NewStrategyRegistry()
	reg.Register(resource1)
	reg.Register(resource2)
	reg.Register(resource3)

	// Create handler
	handler := handlers.NewPreprocessHandler(reg)

	// Content that doesn't mention any of these resource types
	ctx := &interfaces.TransformContext{
		Content: []byte(`resource "completely_different" "test" {}`),
	}

	// Process
	_, err := handler.Handle(ctx)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify all preprocessors were called in registration order
	expectedOrder := []string{"resource_1", "resource_2", "resource_3"}
	if len(callOrder) != len(expectedOrder) {
		t.Errorf("Expected %d preprocessor calls, got %d", len(expectedOrder), len(callOrder))
	}

	for i, expected := range expectedOrder {
		if i >= len(callOrder) || callOrder[i] != expected {
			t.Errorf("Expected call order[%d] to be %s, got %v", i, expected, callOrder)
		}
	}
}
