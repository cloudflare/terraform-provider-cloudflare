package transformations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// HCLTransformer handles the transformation of HCL files
type HCLTransformer struct {
	config *TransformationConfig
	parser *hclparse.Parser
	file   *hclwrite.File
}

// NewHCLTransformer creates a new HCL transformer
func NewHCLTransformer(configPath string) (*HCLTransformer, error) {
	if configPath == "" {
		return nil, fmt.Errorf("configuration file path is required")
	}

	config, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &HCLTransformer{
		config: config,
		parser: hclparse.NewParser(),
	}, nil
}

// TransformFile transforms an HCL file according to the configuration
func (ht *HCLTransformer) TransformFile(inputPath, outputPath string) error {
	// Read the input file
	src, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse the HCL file for writing
	var diags hcl.Diagnostics
	ht.file, diags = hclwrite.ParseConfig(src, inputPath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return fmt.Errorf("failed to parse HCL file: %w", diags)
	}
	if ht.file == nil {
		return fmt.Errorf("failed to parse HCL file: parser returned nil")
	}

	// Process all resource blocks
	for _, block := range ht.file.Body().Blocks() {
		if block.Type() != "resource" {
			continue
		}

		labels := block.Labels()
		if len(labels) < 1 {
			continue
		}

		resourceType := strings.Trim(labels[0], "\"")

		// Apply block-to-attribute transformations
		if err := TransformResourceBlock(ht.config, block, resourceType); err != nil {
			return fmt.Errorf("failed to transform resource %s: %w", resourceType, err)
		}

		// Apply attribute renames
		if err := ApplyAttributeRenames(ht.config, block, resourceType); err != nil {
			return fmt.Errorf("failed to apply attribute renames for %s: %w", resourceType, err)
		}

		// Apply attribute removals
		if err := ApplyAttributeRemovals(ht.config, block, resourceType); err != nil {
			return fmt.Errorf("failed to apply attribute removals for %s: %w", resourceType, err)
		}
	}

	// Write the transformed content with formatting
	output := hclwrite.Format(ht.file.Bytes())
	if outputPath == "" || outputPath == inputPath {
		outputPath = inputPath
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	// Post-process to fix various issues
	if err := postProcessFile(outputPath); err != nil {
		fmt.Printf("Warning: Failed to post-process %s: %v\n", outputPath, err)
		// Don't fail the whole transformation for this
	}

	return nil
}

// TransformDirectory transforms all .tf files in a directory
func (ht *HCLTransformer) TransformDirectory(dirPath string, recursive bool) error {
	fmt.Printf("      DEBUG: TransformDirectory called on %s (recursive=%v)\n", dirPath, recursive)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())

		if entry.IsDir() {
			if recursive {
				fmt.Printf("      DEBUG: Recursing into subdirectory %s\n", fullPath)
				// Recursively process subdirectories
				if err := ht.TransformDirectory(fullPath, recursive); err != nil {
					return err
				}
			} else {
				fmt.Printf("      DEBUG: Skipping subdirectory %s (recursive=false)\n", fullPath)
			}
			continue
		}

		// Process .tf files
		if strings.HasSuffix(entry.Name(), ".tf") {
			fmt.Printf("      Transforming: %s\n", fullPath)
			if err := ht.TransformFile(fullPath, fullPath); err != nil {
				fmt.Printf("      ERROR transforming %s: %v\n", fullPath, err)
				return fmt.Errorf("failed to transform %s: %w", fullPath, err)
			}
			fmt.Printf("      SUCCESS: Transformed %s\n", fullPath)
		} else if !entry.IsDir() {
			fmt.Printf("      DEBUG: Skipping non-.tf file: %s\n", fullPath)
		}
	}

	return nil
}

// GetTransformationType is a helper that delegates to the block_to_attribute module
func (ht *HCLTransformer) GetTransformationType(resourceType, blockName string) string {
	return GetTransformationType(ht.config, resourceType, blockName)
}

// HasAttributeRename is a helper that delegates to the attribute_renames module
func (ht *HCLTransformer) HasAttributeRename(resourceType, attrName string) (string, bool) {
	return HasAttributeRename(ht.config, resourceType, attrName)
}

// ShouldRemoveAttribute is a helper that delegates to the attribute_renames module
func (ht *HCLTransformer) ShouldRemoveAttribute(resourceType, attrName string) bool {
	return ShouldRemoveAttribute(ht.config, resourceType, attrName)
}
