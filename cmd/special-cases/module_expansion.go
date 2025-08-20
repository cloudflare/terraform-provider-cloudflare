package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type ModuleInfo struct {
	Name   string
	Source string
	Path   string
}

type ModuleCall struct {
	Name       string
	Source     string
	Arguments  map[string]*hclwrite.Attribute
	Block      *hclwrite.Block
}

func findZoneSettingsModules(rootDir string) ([]ModuleInfo, error) {
	var modules []ModuleInfo
	
	// Walk through all directories to find modules with zone_settings_override
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if !d.IsDir() || d.Name() == ".terraform" {
			return nil
		}
		
		// Check if this directory contains a main.tf with zone_settings_override
		mainTfPath := filepath.Join(path, "main.tf")
		if _, err := os.Stat(mainTfPath); err != nil {
			return nil // No main.tf, skip
		}
		
		hasZoneSettings, err := fileContainsZoneSettingsOverride(mainTfPath)
		if err != nil {
			return nil // Error reading file, skip
		}
		
		if hasZoneSettings {
			// This looks like a zone settings module
			relPath, _ := filepath.Rel(rootDir, path)
			modules = append(modules, ModuleInfo{
				Name:   filepath.Base(path),
				Source: relPath,
				Path:   path,
			})
		}
		
		return nil
	})
	
	return modules, err
}

func fileContainsZoneSettingsOverride(filename string) (bool, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	
	file, diags := hclwrite.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return false, nil // Can't parse, assume no
	}
	
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 {
			if block.Labels()[0] == "cloudflare_zone_settings_override" {
				return true, nil
			}
		}
	}
	
	return false, nil
}

func fileUsesZoneSettingsModules(filename string, modules []ModuleInfo) (bool, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	
	file, diags := hclwrite.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return false, nil
	}
	
	for _, block := range file.Body().Blocks() {
		if block.Type() == "module" && len(block.Labels()) >= 1 {
			sourceAttr := block.Body().GetAttribute("source")
			if sourceAttr == nil {
				continue
			}
			
			// Get source value - this is a simplified extraction
			sourceTokens := sourceAttr.Expr().BuildTokens(nil)
			sourceValue := strings.Trim(string(sourceTokens.Bytes()), `"`)
			
			// Check if this source matches any of our zone settings modules
			for _, module := range modules {
				if strings.Contains(sourceValue, module.Source) || sourceValue == "./"+module.Source {
					return true, nil
				}
			}
		}
	}
	
	return false, nil
}

func transformFileWithModuleExpansion(filename string, modules []ModuleInfo, dryRun bool) error {
	originalContent, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	
	file, diags := hclwrite.ParseConfig(originalContent, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return fmt.Errorf("failed to parse HCL: %v", diags)
	}
	
	body := file.Body()
	var blocksToRemove []*hclwrite.Block
	var newBlocks []*hclwrite.Block
	
	// Process each module block
	for _, block := range body.Blocks() {
		if block.Type() != "module" || len(block.Labels()) < 1 {
			continue
		}
		
		moduleCall, moduleInfo, err := matchModuleToZoneSettings(block, modules, filename)
		if err != nil {
			continue // Skip if we can't match or process
		}
		if moduleCall == nil {
			continue // This module is not a zone settings module
		}
		
		// Expand this module
		expandedResources, importBlocks, err := expandZoneSettingsModule(moduleCall, moduleInfo)
		if err != nil {
			fmt.Printf("  Warning: Failed to expand module %s: %v\n", moduleCall.Name, err)
			continue
		}
		
		// Mark original module block for removal
		blocksToRemove = append(blocksToRemove, block)
		
		// Add expanded resources and imports
		newBlocks = append(newBlocks, expandedResources...)
		newBlocks = append(newBlocks, importBlocks...)
		
		fmt.Printf("  Expanded module %s into %d resources with %d imports\n", 
			moduleCall.Name, len(expandedResources), len(importBlocks))
	}
	
	if len(blocksToRemove) == 0 {
		return nil // No changes needed
	}
	
	// Remove old module blocks
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}
	
	// Add new blocks
	for _, block := range newBlocks {
		body.AppendBlock(block)
	}
	
	// Format the result
	newContent := hclwrite.Format(file.Bytes())
	
	if dryRun {
		fmt.Printf("  ✗ Would update %s\n", filename)
		return nil
	}
	
	// Write back to file
	if err := os.WriteFile(filename, newContent, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	
	fmt.Printf("  ✓ Updated %s\n", filename)
	return nil
}

func matchModuleToZoneSettings(block *hclwrite.Block, modules []ModuleInfo, currentFile string) (*ModuleCall, *ModuleInfo, error) {
	sourceAttr := block.Body().GetAttribute("source")
	if sourceAttr == nil {
		return nil, nil, fmt.Errorf("module has no source")
	}
	
	// Extract source value
	sourceTokens := sourceAttr.Expr().BuildTokens(nil)
	sourceValue := strings.Trim(string(sourceTokens.Bytes()), `"`)
	
	// Find matching module
	var matchedModule *ModuleInfo
	for _, module := range modules {
		if strings.Contains(sourceValue, module.Source) || sourceValue == "./"+module.Source {
			matchedModule = &module
			break
		}
	}
	
	if matchedModule == nil {
		return nil, nil, nil // Not a zone settings module
	}
	
	// Build module call info
	moduleCall := &ModuleCall{
		Name:      block.Labels()[0],
		Source:    sourceValue,
		Arguments: make(map[string]*hclwrite.Attribute),
		Block:     block,
	}
	
	// Collect all module arguments
	for name, attr := range block.Body().Attributes() {
		if name != "source" {
			moduleCall.Arguments[name] = attr
		}
	}
	
	// Debug: print collected arguments (comment out for production)
	// fmt.Printf("    Module %s has %d arguments:\n", moduleCall.Name, len(moduleCall.Arguments))
	// for name, attr := range moduleCall.Arguments {
	//	tokens := attr.Expr().BuildTokens(nil)
	//	fmt.Printf("      %s = %s\n", name, string(tokens.Bytes()))
	// }
	
	return moduleCall, matchedModule, nil
}

func expandZoneSettingsModule(moduleCall *ModuleCall, moduleInfo *ModuleInfo) ([]*hclwrite.Block, []*hclwrite.Block, error) {
	// Read the module's main.tf file
	mainTfPath := filepath.Join(moduleInfo.Path, "main.tf")
	content, err := os.ReadFile(mainTfPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read module main.tf: %v", err)
	}
	
	file, diags := hclwrite.ParseConfig(content, mainTfPath, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, nil, fmt.Errorf("failed to parse module main.tf: %v", diags)
	}
	
	var resourceBlocks []*hclwrite.Block
	var importBlocks []*hclwrite.Block
	
	// Process each block in the module
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 && 
		   block.Labels()[0] == "cloudflare_zone_settings_override" {
			
			// Transform this zone_settings_override resource
			transformedResources, transformedImports := transformZoneSettingsOverrideInModule(block, moduleCall)
			resourceBlocks = append(resourceBlocks, transformedResources...)
			importBlocks = append(importBlocks, transformedImports...)
		}
	}
	
	return resourceBlocks, importBlocks, nil
}

func transformZoneSettingsOverrideInModule(block *hclwrite.Block, moduleCall *ModuleCall) ([]*hclwrite.Block, []*hclwrite.Block) {
	var resourceBlocks []*hclwrite.Block
	var importBlocks []*hclwrite.Block
	
	originalResourceName := block.Labels()[1]
	
	// Get zone_id attribute from the original block
	var zoneIDAttr *hclwrite.Attribute
	if attr := block.Body().GetAttribute("zone_id"); attr != nil {
		zoneIDAttr = attr
	}
	
	// Find the settings block
	for _, settingsBlock := range block.Body().Blocks() {
		if settingsBlock.Type() == "settings" {
			// Process regular attributes - but ONLY if they have corresponding module arguments
			for name, attr := range settingsBlock.Body().Attributes() {
				// Skip this setting if it's not explicitly set in the module call
				// The only exception is zone_id which we always need to substitute
				if name != "zone_id" {
					if _, hasModuleArg := moduleCall.Arguments[name]; !hasModuleArg {
						fmt.Printf("      Skipping %s (not in module call)\n", name)
						continue
					}
				}
				
				// Map the v4 setting name to the correct v5 setting name
				mappedSettingName := mapSettingName(name)
				resourceName := fmt.Sprintf("%s_%s_%s", moduleCall.Name, originalResourceName, name)
				
				// Create the zone setting resource with variable substitution
				newBlock := createZoneSettingResourceWithSubstitution(
					resourceName,
					mappedSettingName,
					zoneIDAttr,
					attr,
					moduleCall,
				)
				resourceBlocks = append(resourceBlocks, newBlock)
				
				// Create import block with variable substitution
				importBlock := createImportBlockWithSubstitution(resourceName, mappedSettingName, zoneIDAttr, moduleCall)
				importBlocks = append(importBlocks, importBlock)
			}
			
			// Process nested blocks (security_header, nel) - only if they have relevant module arguments
			for _, nestedBlock := range settingsBlock.Body().Blocks() {
				if nestedBlock.Type() == "security_header" {
					// Check if any security_header_* variables are set in module call
					hasSecurityHeaderArgs := false
					for argName := range moduleCall.Arguments {
						if strings.HasPrefix(argName, "security_header_") {
							hasSecurityHeaderArgs = true
							break
						}
					}
					if !hasSecurityHeaderArgs {
						fmt.Printf("      Skipping security_header (no security_header_* args in module call)\n")
						continue
					}
					
					resourceName := fmt.Sprintf("%s_%s_security_header", moduleCall.Name, originalResourceName)
					newBlock := transformSecurityHeaderBlockWithSubstitution(resourceName, zoneIDAttr, nestedBlock, moduleCall)
					resourceBlocks = append(resourceBlocks, newBlock)
					
					importBlock := createImportBlockWithSubstitution(resourceName, "security_header", zoneIDAttr, moduleCall)
					importBlocks = append(importBlocks, importBlock)
				} else if nestedBlock.Type() == "nel" {
					// Check if enable_network_error_logging is set in module call
					if _, hasNELArg := moduleCall.Arguments["enable_network_error_logging"]; !hasNELArg {
						fmt.Printf("      Skipping nel (enable_network_error_logging not in module call)\n")
						continue
					}
					
					resourceName := fmt.Sprintf("%s_%s_nel", moduleCall.Name, originalResourceName)
					newBlock := transformNELBlockWithSubstitution(resourceName, zoneIDAttr, nestedBlock, moduleCall)
					resourceBlocks = append(resourceBlocks, newBlock)
					
					importBlock := createImportBlockWithSubstitution(resourceName, "nel", zoneIDAttr, moduleCall)
					importBlocks = append(importBlocks, importBlock)
				}
			}
		}
	}
	
	return resourceBlocks, importBlocks
}