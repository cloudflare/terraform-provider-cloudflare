package core

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ChangeType represents the type of change in a migration
type ChangeType string

const (
	ChangeTypeAdd      ChangeType = "ADD"
	ChangeTypeRemove   ChangeType = "REMOVE"
	ChangeTypeModify   ChangeType = "MODIFY"
	ChangeTypeRename   ChangeType = "RENAME"
	ChangeTypeRestructure ChangeType = "RESTRUCTURE"
)

// ImpactLevel represents the impact of a migration
type ImpactLevel string

const (
	ImpactLow      ImpactLevel = "LOW"      // Only renames, no functional change
	ImpactMedium   ImpactLevel = "MEDIUM"   // Structural changes, defaults added
	ImpactHigh     ImpactLevel = "HIGH"     // Logic changes, may affect behavior
	ImpactBreaking ImpactLevel = "BREAKING" // Requires manual intervention
)

// Change represents a single change in a migration
type Change struct {
	Type        ChangeType  `json:"type"`
	Path        string      `json:"path"`
	OldValue    string      `json:"old_value,omitempty"`
	NewValue    string      `json:"new_value,omitempty"`
	Description string      `json:"description"`
}

// ResourcePreview represents the preview for a single resource
type ResourcePreview struct {
	Type          string       `json:"type"`
	Name          string       `json:"name"`
	NewType       string       `json:"new_type,omitempty"`
	ConfigChanges []Change     `json:"config_changes"`
	StateChanges  []Change     `json:"state_changes,omitempty"`
	Impact        ImpactLevel  `json:"impact"`
	Warnings      []string     `json:"warnings,omitempty"`
}

// PreviewSummary provides a summary of all changes
type PreviewSummary struct {
	TotalResources   int            `json:"total_resources"`
	ChangedResources int            `json:"changed_resources"`
	TotalChanges     int            `json:"total_changes"`
	ImpactCounts     map[string]int `json:"impact_counts"`
	TypeCounts       map[string]int `json:"type_counts"`
}

// MigrationPreview represents the complete preview of a migration
type MigrationPreview struct {
	FromVersion string            `json:"from_version"`
	ToVersion   string            `json:"to_version"`
	Resources   []ResourcePreview `json:"resources"`
	Summary     PreviewSummary    `json:"summary"`
}

// PreviewGenerator generates previews for migrations
type PreviewGenerator struct {
	registry MigrationRegistry
}

// NewPreviewGenerator creates a new preview generator
func NewPreviewGenerator(registry MigrationRegistry) *PreviewGenerator {
	return &PreviewGenerator{
		registry: registry,
	}
}

// GeneratePreview generates a preview for the given resources
func (pg *PreviewGenerator) GeneratePreview(
	blocks []*hclwrite.Block,
	sourceVersion string,
	targetVersion string,
) (*MigrationPreview, error) {
	preview := &MigrationPreview{
		FromVersion: sourceVersion,
		ToVersion:   targetVersion,
		Resources:   []ResourcePreview{},
	}

	for _, block := range blocks {
		// Skip non-resource blocks
		if block.Type() != "resource" {
			continue
		}
		
		labels := block.Labels()
		if len(labels) < 2 {
			continue // Need at least resource type and name
		}
		
		resourceType := labels[0]
		resourceName := labels[1]

		// Get the migration for this resource
		migration, err := pg.registry.Get(resourceType, sourceVersion, targetVersion)
		if err != nil {
			// No migration available, skip
			continue
		}

		// Generate preview for this resource
		resourcePreview := pg.previewResource(block, migration, resourceType, resourceName)
		preview.Resources = append(preview.Resources, resourcePreview)
	}

	// Calculate summary
	preview.Summary = pg.calculateSummary(preview.Resources)

	return preview, nil
}

// previewResource generates a preview for a single resource
func (pg *PreviewGenerator) previewResource(
	block *hclwrite.Block,
	migration Migration,
	resourceType string,
	resourceName string,
) ResourcePreview {
	preview := ResourcePreview{
		Type:          resourceType,
		Name:          resourceName,
		ConfigChanges: []Change{},
		StateChanges:  []Change{},
		Warnings:      []string{},
	}

	// Clone the block for dry-run
	clonedBlock := pg.cloneBlock(block)
	
	// Create a preview context to capture changes
	ctx := NewMigrationContext()
	ctx.DryRun = true

	// Run migration in dry-run mode to detect changes
	if err := migration.MigrateConfig(clonedBlock, ctx); err == nil {
		// Compare original and migrated blocks
		preview.ConfigChanges = pg.detectConfigChanges(block, clonedBlock)
	}

	// Determine new resource type if it changes
	if newType := pg.getNewResourceType(migration); newType != resourceType {
		preview.NewType = newType
		preview.ConfigChanges = append(preview.ConfigChanges, Change{
			Type:        ChangeTypeRename,
			Path:        "resource_type",
			OldValue:    resourceType,
			NewValue:    newType,
			Description: fmt.Sprintf("Resource type changed from %s to %s", resourceType, newType),
		})
	}

	// Collect diagnostics as warnings
	for _, diag := range ctx.Diagnostics {
		if diag.Severity == DiagnosticSeverityWarning {
			preview.Warnings = append(preview.Warnings, diag.Summary)
		}
	}

	// Determine impact level
	preview.Impact = pg.calculateImpact(preview.ConfigChanges)

	return preview
}

// detectConfigChanges compares two blocks and returns the changes
func (pg *PreviewGenerator) detectConfigChanges(original, modified *hclwrite.Block) []Change {
	changes := []Change{}

	origBody := original.Body()
	modBody := modified.Body()

	// Check for removed attributes
	for name, attr := range origBody.Attributes() {
		if modBody.GetAttribute(name) == nil {
			changes = append(changes, Change{
				Type:        ChangeTypeRemove,
				Path:        name,
				OldValue:    pg.getAttributeValue(attr),
				Description: fmt.Sprintf("Attribute '%s' will be removed", name),
			})
		}
	}

	// Check for added or modified attributes
	for name, modAttr := range modBody.Attributes() {
		origAttr := origBody.GetAttribute(name)
		if origAttr == nil {
			// Added
			changes = append(changes, Change{
				Type:        ChangeTypeAdd,
				Path:        name,
				NewValue:    pg.getAttributeValue(modAttr),
				Description: fmt.Sprintf("Attribute '%s' will be added", name),
			})
		} else {
			// Check if modified
			origValue := pg.getAttributeValue(origAttr)
			modValue := pg.getAttributeValue(modAttr)
			if origValue != modValue {
				changes = append(changes, Change{
					Type:        ChangeTypeModify,
					Path:        name,
					OldValue:    origValue,
					NewValue:    modValue,
					Description: fmt.Sprintf("Attribute '%s' will be modified", name),
				})
			}
		}
	}

	// Check for block changes
	origBlocks := pg.getBlocksByType(origBody)
	modBlocks := pg.getBlocksByType(modBody)

	for blockType, count := range origBlocks {
		modCount := modBlocks[blockType]
		if modCount != count {
			if modCount == 0 {
				changes = append(changes, Change{
					Type:        ChangeTypeRemove,
					Path:        blockType,
					OldValue:    fmt.Sprintf("%d blocks", count),
					Description: fmt.Sprintf("Block '%s' will be removed", blockType),
				})
			} else {
				changes = append(changes, Change{
					Type:        ChangeTypeRestructure,
					Path:        blockType,
					OldValue:    fmt.Sprintf("%d blocks", count),
					NewValue:    fmt.Sprintf("%d blocks", modCount),
					Description: fmt.Sprintf("Block '%s' structure will change", blockType),
				})
			}
		}
	}

	// Check for new blocks
	for blockType, count := range modBlocks {
		if origBlocks[blockType] == 0 && count > 0 {
			changes = append(changes, Change{
				Type:        ChangeTypeAdd,
				Path:        blockType,
				NewValue:    fmt.Sprintf("%d blocks", count),
				Description: fmt.Sprintf("Block '%s' will be added", blockType),
			})
		}
	}

	return changes
}

// calculateImpact determines the impact level based on changes
func (pg *PreviewGenerator) calculateImpact(changes []Change) ImpactLevel {
	if len(changes) == 0 {
		return ImpactLow
	}

	hasStructuralChanges := false
	hasRemovals := false
	hasAdditions := false

	for _, change := range changes {
		switch change.Type {
		case ChangeTypeRestructure:
			hasStructuralChanges = true
		case ChangeTypeRemove:
			hasRemovals = true
		case ChangeTypeAdd:
			hasAdditions = true
		}
	}

	if hasStructuralChanges {
		return ImpactHigh
	}
	if hasRemovals {
		return ImpactMedium
	}
	if hasAdditions {
		return ImpactMedium
	}

	return ImpactLow
}

// calculateSummary generates a summary of all changes
func (pg *PreviewGenerator) calculateSummary(resources []ResourcePreview) PreviewSummary {
	summary := PreviewSummary{
		TotalResources:   len(resources),
		ChangedResources: 0,
		TotalChanges:     0,
		ImpactCounts:     make(map[string]int),
		TypeCounts:       make(map[string]int),
	}

	for _, resource := range resources {
		if len(resource.ConfigChanges) > 0 || len(resource.StateChanges) > 0 {
			summary.ChangedResources++
		}

		summary.TotalChanges += len(resource.ConfigChanges) + len(resource.StateChanges)
		
		// Count impact levels
		summary.ImpactCounts[string(resource.Impact)]++

		// Count change types
		for _, change := range resource.ConfigChanges {
			summary.TypeCounts[string(change.Type)]++
		}
	}

	return summary
}

// RenderDiff renders the preview as a colored diff for terminal output
func (preview *MigrationPreview) RenderDiff() string {
	var output strings.Builder

	// Header
	output.WriteString(fmt.Sprintf("\n=== Migration Preview: %s → %s ===\n\n", 
		preview.FromVersion, preview.ToVersion))

	// Summary
	output.WriteString("Summary:\n")
	output.WriteString(fmt.Sprintf("  • Resources to migrate: %d\n", preview.Summary.TotalResources))
	output.WriteString(fmt.Sprintf("  • Resources with changes: %d\n", preview.Summary.ChangedResources))
	output.WriteString(fmt.Sprintf("  • Total changes: %d\n\n", preview.Summary.TotalChanges))

	// Resource changes
	for _, resource := range preview.Resources {
		if len(resource.ConfigChanges) == 0 && len(resource.StateChanges) == 0 {
			continue
		}

		output.WriteString(fmt.Sprintf("Resource: %s.%s\n", resource.Type, resource.Name))
		if resource.NewType != "" {
			output.WriteString(fmt.Sprintf("  → Will be renamed to: %s.%s\n", resource.NewType, resource.Name))
		}
		output.WriteString(fmt.Sprintf("  Impact: %s\n", resource.Impact))

		if len(resource.ConfigChanges) > 0 {
			output.WriteString("  Configuration changes:\n")
			for _, change := range resource.ConfigChanges {
				output.WriteString(formatChange(change))
			}
		}

		if len(resource.Warnings) > 0 {
			output.WriteString("  Warnings:\n")
			for _, warning := range resource.Warnings {
				output.WriteString(fmt.Sprintf("    ⚠ %s\n", warning))
			}
		}

		output.WriteString("\n")
	}

	// Impact breakdown
	if len(preview.Summary.ImpactCounts) > 0 {
		output.WriteString("Impact Analysis:\n")
		for level, count := range preview.Summary.ImpactCounts {
			output.WriteString(fmt.Sprintf("  • %s: %d resources\n", level, count))
		}
	}

	return output.String()
}

// formatChange formats a single change for display
func formatChange(change Change) string {
	symbol := ""
	switch change.Type {
	case ChangeTypeAdd:
		symbol = "+"
	case ChangeTypeRemove:
		symbol = "-"
	case ChangeTypeModify:
		symbol = "~"
	case ChangeTypeRename:
		symbol = "→"
	case ChangeTypeRestructure:
		symbol = "※"
	}

	if change.OldValue != "" && change.NewValue != "" {
		return fmt.Sprintf("    %s %s: %s → %s\n", symbol, change.Path, 
			truncate(change.OldValue, 30), truncate(change.NewValue, 30))
	} else if change.OldValue != "" {
		return fmt.Sprintf("    %s %s: %s\n", symbol, change.Path, truncate(change.OldValue, 30))
	} else if change.NewValue != "" {
		return fmt.Sprintf("    %s %s: %s\n", symbol, change.Path, truncate(change.NewValue, 30))
	} else {
		return fmt.Sprintf("    %s %s: %s\n", symbol, change.Path, change.Description)
	}
}

// Helper methods

func (pg *PreviewGenerator) cloneBlock(block *hclwrite.Block) *hclwrite.Block {
	// This is a simplified clone - in production, we'd need a deep copy
	tokens := block.BuildTokens(nil)
	content := string(tokens.Bytes())
	
	file, _ := hclwrite.ParseConfig([]byte(content), "clone.tf", hcl.InitialPos)
	if len(file.Body().Blocks()) > 0 {
		return file.Body().Blocks()[0]
	}
	return block
}

func (pg *PreviewGenerator) getAttributeValue(attr *hclwrite.Attribute) string {
	tokens := attr.Expr().BuildTokens(nil)
	return strings.TrimSpace(string(tokens.Bytes()))
}

func (pg *PreviewGenerator) getBlocksByType(body *hclwrite.Body) map[string]int {
	counts := make(map[string]int)
	for _, block := range body.Blocks() {
		counts[block.Type()]++
	}
	return counts
}

func (pg *PreviewGenerator) getNewResourceType(migration Migration) string {
	// This would need to be extended to support resource type changes
	// For now, return empty string indicating no change
	return ""
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}