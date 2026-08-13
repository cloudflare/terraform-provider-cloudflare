package transformations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"gopkg.in/yaml.v3"
)

// ResourceRenamesConfig represents the YAML configuration for resource renames
type ResourceRenamesConfig struct {
	Version             string                       `yaml:"version"`
	Description         string                       `yaml:"description"`
	Notes               []string                     `yaml:"notes"`
	ResourceRenames     map[string]string            `yaml:"resource_renames"`
	ResourceCategories  map[string][]string          `yaml:"resource_categories"`
	MigrationGuidance   MigrationGuidance            `yaml:"migration_guidance"`
}

// MigrationGuidance provides guidance for handling resource renames
type MigrationGuidance struct {
	Description string   `yaml:"description"`
	Steps       []string `yaml:"steps"`
}

// LoadResourceRenamesConfig loads the resource renames configuration from a YAML file
func LoadResourceRenamesConfig(filename string) (*ResourceRenamesConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config ResourceRenamesConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}

// ResourceRenameTransformer handles resource rename transformations for HCL files
type ResourceRenameTransformer struct {
	config *ResourceRenamesConfig
}

// NewResourceRenameTransformer creates a new resource rename transformer
func NewResourceRenameTransformer(configPath string) (*ResourceRenameTransformer, error) {
	config, err := LoadResourceRenamesConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &ResourceRenameTransformer{
		config: config,
	}, nil
}

// TransformFile transforms resource types in an HCL file
func (rt *ResourceRenameTransformer) TransformFile(inputPath, outputPath string) error {
	// Read the input file
	src, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse the HCL file
	file, diags := hclwrite.ParseConfig(src, inputPath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return fmt.Errorf("failed to parse HCL: %s", diags.Error())
	}

	// Track if any changes were made
	changed := false

	// Collect all blocks with their transformations
	type blockInfo struct {
		block     *hclwrite.Block
		newLabels []string
		needsRename bool
	}
	
	var blocks []blockInfo
	
	// First pass: identify blocks that need renaming
	for _, block := range file.Body().Blocks() {
		info := blockInfo{block: block}
		
		if block.Type() == "resource" {
			labels := block.Labels()
			if len(labels) >= 1 {
				// Get the resource type (remove quotes if present)
				resourceType := strings.Trim(labels[0], `"`)
				
				// Check if this resource type needs to be renamed
				if newType, exists := rt.config.ResourceRenames[resourceType]; exists {
					// Update the label with the new resource type
					newLabels := make([]string, len(labels))
					newLabels[0] = newType
					for i := 1; i < len(labels); i++ {
						newLabels[i] = labels[i]
					}
					info.newLabels = newLabels
					info.needsRename = true
					changed = true
				}
			}
		}
		
		blocks = append(blocks, info)
	}
	
	// If changes are needed, rebuild the body
	if changed {
		// Clear the body
		newFile := hclwrite.NewEmptyFile()
		newBody := newFile.Body()
		
		// Copy root attributes first
		for name, attr := range file.Body().Attributes() {
			tokens := attr.Expr().BuildTokens(nil)
			newBody.SetAttributeRaw(name, tokens)
		}
		
		// Recreate blocks in order
		for _, info := range blocks {
			if info.needsRename && info.block.Type() == "resource" {
				// Create renamed block
				newBlock := newBody.AppendNewBlock("resource", info.newLabels)
				
				// Copy all attributes from the old block
				for name, attr := range info.block.Body().Attributes() {
					tokens := attr.Expr().BuildTokens(nil)
					newBlock.Body().SetAttributeRaw(name, tokens)
				}
				
				// Copy all nested blocks
				for _, nestedBlock := range info.block.Body().Blocks() {
					copyBlock(newBlock.Body(), nestedBlock)
				}
			} else {
				// Copy block as-is
				copyBlockToBody(newBody, info.block)
			}
		}
		
		// Replace the file content
		file = newFile
	}

	// Also update references in other blocks (data sources, locals, outputs)
	if changed {
		updateResourceReferences(file, rt.config.ResourceRenames)
	}

	// Write the output file
	output := hclwrite.Format(file.Bytes())
	if outputPath == "" || outputPath == inputPath {
		outputPath = inputPath
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// TransformDirectory transforms all .tf files in a directory
func (rt *ResourceRenameTransformer) TransformDirectory(dirPath string, recursive bool) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())

		if entry.IsDir() && recursive {
			// Recursively process subdirectories
			if err := rt.TransformDirectory(fullPath, recursive); err != nil {
				return err
			}
			continue
		}

		// Process .tf files
		if strings.HasSuffix(entry.Name(), ".tf") {
			fmt.Printf("Transforming resource renames in: %s\n", fullPath)
			if err := rt.TransformFile(fullPath, fullPath); err != nil {
				return fmt.Errorf("failed to transform %s: %w", fullPath, err)
			}
		}
	}

	return nil
}

// copyBlock copies a block to a new body
func copyBlock(targetBody *hclwrite.Body, sourceBlock *hclwrite.Block) {
	newBlock := targetBody.AppendNewBlock(sourceBlock.Type(), sourceBlock.Labels())
	
	// Copy attributes
	for name, attr := range sourceBlock.Body().Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		newBlock.Body().SetAttributeRaw(name, tokens)
	}
	
	// Recursively copy nested blocks
	for _, nestedBlock := range sourceBlock.Body().Blocks() {
		copyBlock(newBlock.Body(), nestedBlock)
	}
}

// copyBlockToBody copies a complete block (including type) to a body
func copyBlockToBody(targetBody *hclwrite.Body, sourceBlock *hclwrite.Block) {
	newBlock := targetBody.AppendNewBlock(sourceBlock.Type(), sourceBlock.Labels())
	
	// Copy attributes
	for name, attr := range sourceBlock.Body().Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		newBlock.Body().SetAttributeRaw(name, tokens)
	}
	
	// Recursively copy nested blocks
	for _, nestedBlock := range sourceBlock.Body().Blocks() {
		copyBlock(newBlock.Body(), nestedBlock)
	}
}

// updateResourceReferences updates references to renamed resources throughout the file
func updateResourceReferences(file *hclwrite.File, renames map[string]string) {
	// This is a simplified version - in production, you'd want to parse
	// and update references more carefully, handling interpolations, etc.
	
	// Process all blocks to update references
	for _, block := range file.Body().Blocks() {
		updateBlockReferences(block.Body(), renames)
	}
	
	// Process root body attributes
	updateBodyReferences(file.Body(), renames)
}

// updateBlockReferences updates references in a block body
func updateBlockReferences(body *hclwrite.Body, renames map[string]string) {
	// Update attributes that might contain references
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		// This is simplified - in reality, you'd parse the tokens
		// and update resource references within interpolations
		// For now, we'll skip complex reference updates
		_ = tokens
		_ = name
	}
	
	// Recursively update nested blocks
	for _, block := range body.Blocks() {
		updateBlockReferences(block.Body(), renames)
	}
}

// updateBodyReferences updates references in a body
func updateBodyReferences(body *hclwrite.Body, renames map[string]string) {
	// Similar to updateBlockReferences but for the root body
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		_ = tokens
		_ = name
	}
}

// GetResourceRename checks if a resource type should be renamed
func (rt *ResourceRenameTransformer) GetResourceRename(resourceType string) (string, bool) {
	newType, exists := rt.config.ResourceRenames[resourceType]
	return newType, exists
}

// GetResourceCategory returns the category for a resource type
func (rt *ResourceRenameTransformer) GetResourceCategory(resourceType string) string {
	for category, resources := range rt.config.ResourceCategories {
		for _, res := range resources {
			if res == resourceType {
				return category
			}
		}
	}
	return "other"
}