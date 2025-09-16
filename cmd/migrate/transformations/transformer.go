package transformations

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
		return fmt.Errorf("failed to parse HCL file %s: %s", inputPath, diags.Error())
	}
	if ht.file == nil {
		return fmt.Errorf("failed to parse HCL file %s: file is nil", inputPath)
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
	
	// Post-process to fix double dollar signs and heredoc issues in cloudflare_ruleset expressions
	// Only apply if we're dealing with heredoc expressions that need it
	if strings.Contains(string(output), "<<-EOF") || strings.Contains(string(output), "<<-EOT") {
		output = fixCloudflareRulesetDoubleDollarSigns(output)
	}
	
	if outputPath == "" || outputPath == inputPath {
		outputPath = inputPath
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	// Skip terraform fmt during transformation to avoid working directory issues
	// Run 'terraform fmt -recursive .' after migration to format all files

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
// fixCloudflareRulesetDoubleDollarSigns fixes double dollar sign escaping and heredoc issues ONLY within cloudflare_ruleset resources
func fixCloudflareRulesetDoubleDollarSigns(content []byte) []byte {
	result := string(content)
	
	// Only process content within cloudflare_ruleset resources to avoid affecting other resources
	rulesetPattern := regexp.MustCompile(`(?s)(resource\s+"cloudflare_ruleset"\s+"[^"]+"\s*\{)(.*?)(\n\})`)
	result = rulesetPattern.ReplaceAllStringFunc(result, func(match string) string {
		// Extract the ruleset resource content
		parts := rulesetPattern.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match // Return original if parsing fails
		}
		
		resourceStart := parts[1]
		resourceContent := parts[2]
		resourceEnd := parts[3]
		
		// Fix double dollar signs in variable references within this cloudflare_ruleset resource
		// This handles cases like $${var.example} -> ${var.example}
		resourceContent = regexp.MustCompile(`\$\$\{([^}]+)\}`).ReplaceAllString(resourceContent, `${$1}`)
		
		// Fix escaped heredoc expressions that were converted to string literals
		// Only process expressions that are definitely escaped heredocs
		// Must contain both <<- and literal \n to avoid corrupting simple expressions
		heredocPattern := regexp.MustCompile(`expression\s*=\s*"(<<-[A-Z]+)\\n([^"]+)\\n\s*(EOF|EOT)"`)
		resourceContent = heredocPattern.ReplaceAllStringFunc(resourceContent, func(match string) string {
			// Only process if this contains both heredoc markers and escaped newlines
			if !strings.Contains(match, "<<-") || !strings.Contains(match, "\\n") {
				return match // Skip anything that's not clearly an escaped heredoc
			}
			
			// Extract the heredoc marker and content using the same pattern
			parts := heredocPattern.FindStringSubmatch(match)
			if len(parts) < 4 {
				return match // Return original if parsing fails
			}
			
			startMarker := parts[1]
			content := parts[2]
			endMarker := parts[3]
			
			// Clean up the content: convert \n to actual newlines and unescape quotes
			content = strings.ReplaceAll(content, "\\n", "\n")  // Convert literal \n to actual newlines
			content = strings.ReplaceAll(content, `\"`, `"`)     // Unescape quotes
			content = strings.TrimPrefix(content, "\n")          // Remove leading newline
			content = strings.TrimSuffix(content, "\n")          // Remove trailing newline before end marker
			
			// Keep the original content structure - terraform fmt will handle proper indentation
			// Just ensure we have the basic structure right
			formattedContent := strings.TrimSpace(content)
			
			return fmt.Sprintf(`expression = %s
%s
    %s`, startMarker, formattedContent, endMarker)
		})
		
		return resourceStart + resourceContent + resourceEnd
	})
	
	return []byte(result)
}

// runTerraformFmt runs terraform fmt on the specified file to apply canonical formatting
func (ht *HCLTransformer) runTerraformFmt(filePath string) error {
	// Convert to absolute path to avoid working directory issues
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	
	// Check if file exists before trying to format it
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", absPath)
	}
	
	cmd := exec.Command("terraform", "fmt", absPath)
	
	// Set working directory to the file's directory
	cmd.Dir = filepath.Dir(absPath)
	
	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("terraform fmt failed: %w\nOutput: %s", err, string(output))
	}
	
	return nil
}
