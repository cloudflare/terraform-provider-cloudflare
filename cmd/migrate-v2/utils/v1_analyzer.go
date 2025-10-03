package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// MigrationAnalysis contains the analyzed patterns from v1 migration code
type MigrationAnalysis struct {
	ResourceType      string                 `yaml:"resource_type"`
	SourceVersion     string                 `yaml:"source_version"`
	TargetVersion     string                 `yaml:"target_version"`
	Description       string                 `yaml:"description"`
	Config           ConfigAnalysis         `yaml:"config"`
	State            StateAnalysis          `yaml:"state"`
	CustomLogicNotes []string               `yaml:"custom_logic_notes,omitempty"`
}

// ConfigAnalysis contains config transformation patterns
type ConfigAnalysis struct {
	AttributeRenames    map[string]string      `yaml:"attribute_renames,omitempty"`
	AttributeRemovals   []string               `yaml:"attribute_removals,omitempty"`
	TypeConversions     []TypeConversion       `yaml:"type_conversions,omitempty"`
	BlocksToLists       []string               `yaml:"blocks_to_lists,omitempty"`
	ListsToBlocks       []string               `yaml:"lists_to_blocks,omitempty"`
	DefaultValues       map[string]interface{} `yaml:"default_values,omitempty"`
	StructuralChanges   []StructuralChange     `yaml:"structural_changes,omitempty"`
	ConditionalRemovals []ConditionalRemoval   `yaml:"conditional_removals,omitempty"`
}

// StateAnalysis contains state transformation patterns
type StateAnalysis struct {
	AttributeRenames  map[string]string    `yaml:"attribute_renames,omitempty"`
	TypeConversions   []TypeConversion     `yaml:"type_conversions,omitempty"`
	ArrayToObject     []string             `yaml:"array_to_object,omitempty"`
	StructuralChanges []StructuralChange   `yaml:"structural_changes,omitempty"`
	SchemaVersion     int                  `yaml:"schema_version,omitempty"`
}

// TypeConversion defines type conversion patterns
type TypeConversion struct {
	Attribute string `yaml:"attribute"`
	FromType  string `yaml:"from_type"`
	ToType    string `yaml:"to_type"`
	Pattern   string `yaml:"pattern,omitempty"`
}

// StructuralChange defines complex structural transformations
type StructuralChange struct {
	Type       string                 `yaml:"type"`
	Source     string                 `yaml:"source"`
	Target     string                 `yaml:"target,omitempty"`
	Transform  string                 `yaml:"transform,omitempty"`
	Parameters map[string]interface{} `yaml:"parameters,omitempty"`
}

// ConditionalRemoval defines conditional attribute removal
type ConditionalRemoval struct {
	Attribute string                 `yaml:"attribute"`
	Condition map[string]interface{} `yaml:"condition"`
}

// AnalyzeV1Migration analyzes a v1 migration file and extracts patterns
func AnalyzeV1Migration(resourceName string, configPath string, statePath string) (*MigrationAnalysis, error) {
	analysis := &MigrationAnalysis{
		ResourceType:  fmt.Sprintf("cloudflare_%s", resourceName),
		SourceVersion: "v4",
		TargetVersion: "v5",
		Description:   fmt.Sprintf("Migrate %s from v4 to v5", resourceName),
		Config: ConfigAnalysis{
			AttributeRenames:  make(map[string]string),
			AttributeRemovals: []string{},
			DefaultValues:     make(map[string]interface{}),
		},
		State: StateAnalysis{
			AttributeRenames: make(map[string]string),
		},
	}

	// Analyze config transformations
	if configPath != "" {
		if err := analyzeConfigFile(configPath, analysis); err != nil {
			return nil, fmt.Errorf("failed to analyze config: %w", err)
		}
	}

	// Analyze state transformations
	if statePath != "" {
		if err := analyzeStateFile(statePath, resourceName, analysis); err != nil {
			return nil, fmt.Errorf("failed to analyze state: %w", err)
		}
	}

	return analysis, nil
}

// analyzeConfigFile analyzes v1 config transformation patterns
func analyzeConfigFile(path string, analysis *MigrationAnalysis) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, content, parser.ParseComments)
	if err != nil {
		return err
	}

	// Walk the AST to find transformation patterns
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.CallExpr:
			extractCallPatterns(node, analysis)
		case *ast.FuncDecl:
			extractFunctionPatterns(node, analysis)
		}
		return true
	})

	return nil
}

// extractCallPatterns extracts patterns from function calls
func extractCallPatterns(call *ast.CallExpr, analysis *MigrationAnalysis) {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		switch sel.Sel.Name {
		case "RemoveAttribute":
			// Extract attribute removal
			if len(call.Args) > 0 {
				if lit, ok := call.Args[0].(*ast.BasicLit); ok {
					attrName := strings.Trim(lit.Value, `"`)
					analysis.Config.AttributeRemovals = append(analysis.Config.AttributeRemovals, attrName)
				}
			}

		case "SetAttributeValue":
			// Extract default value setting
			if len(call.Args) >= 2 {
				if nameLit, ok := call.Args[0].(*ast.BasicLit); ok {
					attrName := strings.Trim(nameLit.Value, `"`)
					
					// Check if it's setting a default value
					if valCall, ok := call.Args[1].(*ast.CallExpr); ok {
						if sel, ok := valCall.Fun.(*ast.SelectorExpr); ok {
							switch sel.Sel.Name {
							case "StringVal":
								if len(valCall.Args) > 0 {
									if lit, ok := valCall.Args[0].(*ast.BasicLit); ok {
										value := strings.Trim(lit.Value, `"`)
										analysis.Config.DefaultValues[attrName] = value
									}
								}
							case "BoolVal":
								if len(valCall.Args) > 0 {
									if ident, ok := valCall.Args[0].(*ast.Ident); ok {
										analysis.Config.DefaultValues[attrName] = ident.Name == "true"
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// extractFunctionPatterns extracts patterns from function declarations
func extractFunctionPatterns(fn *ast.FuncDecl, analysis *MigrationAnalysis) {
	funcName := fn.Name.Name
	
	// Look for specific transformation functions
	switch {
	case strings.Contains(funcName, "transformSetToList"):
		// Found set to list conversion
		analysis.CustomLogicNotes = append(analysis.CustomLogicNotes, 
			"Uses set to list conversion for some attributes")
		
	case strings.Contains(funcName, "transformPolicies"):
		// Found policies transformation
		analysis.Config.StructuralChanges = append(analysis.Config.StructuralChanges, StructuralChange{
			Transform: "string_list_to_object_list",
			Source:    "policies",
			Parameters: map[string]interface{}{
				"object_key": "id",
			},
		})
		
	case strings.Contains(funcName, "convertDestinations"):
		// Found destinations block to list conversion
		analysis.Config.BlocksToLists = append(analysis.Config.BlocksToLists, "destinations")
	}
}

// analyzeStateFile analyzes state transformation patterns
func analyzeStateFile(path string, resourceName string, analysis *MigrationAnalysis) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Look for the specific state transformation function
	funcName := fmt.Sprintf("transform%sStateJSON", toCamelCase(resourceName))
	
	if strings.Contains(string(content), funcName) {
		// Extract patterns from the state transformation
		// This is simplified - in reality we'd parse the Go code more thoroughly
		
		// Common patterns in state transformations
		if strings.Contains(string(content), "set(json, attrPath+\".allowed_idps\"") {
			analysis.State.TypeConversions = append(analysis.State.TypeConversions, TypeConversion{
				Attribute: "allowed_idps",
				FromType:  "set",
				ToType:    "list",
			})
		}
		
		if strings.Contains(string(content), "set(json, attrPath+\".custom_pages\"") {
			analysis.State.TypeConversions = append(analysis.State.TypeConversions, TypeConversion{
				Attribute: "custom_pages",
				FromType:  "set",
				ToType:    "list",
			})
		}
		
		// Array to object conversions
		for _, field := range []string{"cors_headers", "landing_page_design", "saas_app", "scim_config"} {
			if strings.Contains(string(content), fmt.Sprintf("\".%s\"", field)) &&
			   strings.Contains(string(content), "IsArray()") {
				analysis.State.ArrayToObject = append(analysis.State.ArrayToObject, field)
			}
		}
		
		// Policies transformation
		if strings.Contains(string(content), "transformedPolicies") {
			analysis.State.StructuralChanges = append(analysis.State.StructuralChanges, StructuralChange{
				Type:   "string_array_to_object_array",
				Source: "policies",
				Parameters: map[string]interface{}{
					"object_key": "id",
				},
			})
		}
	}
	
	return nil
}

// toCamelCase converts snake_case to CamelCase
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: v1_analyzer <resource_name> [config_path] [state_path]")
		os.Exit(1)
	}

	resourceName := os.Args[1]
	configPath := ""
	statePath := ""
	
	if len(os.Args) > 2 {
		configPath = os.Args[2]
	}
	if len(os.Args) > 3 {
		statePath = os.Args[3]
	}

	// Default paths if not provided
	if configPath == "" {
		configPath = fmt.Sprintf("../../cmd/migrate/%s.go", resourceName)
	}
	if statePath == "" {
		statePath = "../../cmd/migrate/state.go"
	}

	analysis, err := AnalyzeV1Migration(resourceName, configPath, statePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output YAML
	output, err := yaml.Marshal(analysis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}