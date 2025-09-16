package transformations

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// postProcessFile runs all post-processing steps on a transformed file
func postProcessFile(filePath string) error {
	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	original := string(content)
	processed := original

	// Only apply double dollar fix to specific resources
	processed = fixVariableInterpolationInTargetResources(processed)

	// Fix heredoc issues (indentation)
	processed = fixHeredocIssues(processed)

	// Add missing commas between object attributes
	processed = fixObjectCommas(processed)

	// Normalize for expression spacing
	processed = normalizeForExpressionSpacing(processed)

	// Only write if changes were made
	if processed != original {
		if err := os.WriteFile(filePath, []byte(processed), 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}

	return nil
}

// fixHeredocIssues fixes heredoc-related problems including indentation and commas
func fixHeredocIssues(content string) string {
	lines := strings.Split(content, "\n")
	inHeredoc := false
	isIndentedHeredoc := false

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Check if we're entering a heredoc
		if strings.Contains(line, "<<-EOF") {
			inHeredoc = true
			isIndentedHeredoc = true
			continue
		} else if strings.Contains(line, "<<EOF") {
			inHeredoc = true
			isIndentedHeredoc = false
			continue
		}

		// Check if we're exiting a heredoc
		if inHeredoc && trimmed == "EOF" {
			inHeredoc = false
			if isIndentedHeredoc {
				// For <<-EOF, EOF must be at column 0 (no indentation)
				lines[i] = "EOF"
			}
			isIndentedHeredoc = false
		}
	}

	// Remove any blank lines after EOF that break the structure
	result := []string{}
	for i := 0; i < len(lines); i++ {
		result = append(result, lines[i])

		// If this is an EOF line followed by a blank line and then an attribute,
		// skip the blank line
		if strings.TrimSpace(lines[i]) == "EOF" &&
			i+1 < len(lines) && strings.TrimSpace(lines[i+1]) == "" &&
			i+2 < len(lines) && strings.Contains(lines[i+2], "=") {
			i++ // Skip the blank line
		}
	}

	return strings.Join(result, "\n")
}

// fixObjectCommas adds missing commas between object attributes and removes extra commas
func fixObjectCommas(content string) string {
	// First pass: remove {, patterns
	content = strings.ReplaceAll(content, "{,", "{")

	// Only apply comma fixes to specific resource types
	targetResources := map[string]bool{
		"cloudflare_load_balancer":      true,
		"cloudflare_load_balancer_pool": true,
		"cloudflare_ruleset":            true,
	}

	// Check if this file contains any target resources
	hasTargetResource := false
	for resourceType := range targetResources {
		if strings.Contains(content, `resource "`+resourceType+`"`) {
			hasTargetResource = true
			break
		}
	}

	// If no target resources, return content with just {, fixed
	if !hasTargetResource {
		return content
	}

	lines := strings.Split(content, "\n")
	inTargetResource := false
	resourceDepth := 0
	inList := false
	listDepth := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Check if we're entering a target resource block
		if strings.HasPrefix(trimmed, "resource ") {
			inTargetResource = false
			for resourceType := range targetResources {
				if strings.Contains(line, `"`+resourceType+`"`) {
					inTargetResource = true
					resourceDepth = 0
					break
				}
			}
		}

		// Track braces to know when we exit the resource block
		if inTargetResource {
			openBraces := strings.Count(line, "{")
			closeBraces := strings.Count(line, "}")
			resourceDepth += openBraces - closeBraces

			// Check if we've exited the resource block
			if resourceDepth <= 0 && strings.Contains(trimmed, "}") {
				inTargetResource = false
			}

			// Track list context within the resource
			if strings.Contains(line, "[") {
				inList = true
				listDepth += strings.Count(line, "[")
			}
			if strings.Contains(line, "]") {
				listDepth -= strings.Count(line, "]")
				if listDepth <= 0 {
					inList = false
					listDepth = 0
				}
			}
		}

		// Remove comma after EOF
		if trimmed == "EOF," {
			lines[i] = strings.ReplaceAll(lines[i], "EOF,", "EOF")
		}

		// Remove trailing comma before newline after EOF
		if strings.HasSuffix(trimmed, "EOF") && i+1 < len(lines) {
			nextLine := lines[i+1]
			if strings.TrimSpace(nextLine) == "," {
				// Remove the line with just a comma
				lines = append(lines[:i+1], lines[i+2:]...)
			}
		}

		// Skip empty lines, comments, and closing braces
		if trimmed == "" || strings.HasPrefix(trimmed, "#") ||
			strings.HasPrefix(trimmed, "}") || strings.HasPrefix(trimmed, "]") {
			continue
		}

		// Only process if we're inside a target resource AND inside a list
		if !inTargetResource || !inList {
			continue
		}

		// Skip lines that are inside objects/lists (check indentation)
		// Only add commas for attributes inside object/list literals, not at root level
		leadingSpaces := len(line) - len(strings.TrimLeft(line, " \t"))

		// Check if this line is an attribute (contains = but not := or ==)
		// AND it's indented (not at root level)
		// AND it doesn't end with { (which means it's starting an object)
		if leadingSpaces >= 4 && strings.Contains(trimmed, "=") && !strings.Contains(trimmed, ":=") &&
			!strings.Contains(trimmed, "==") && !strings.Contains(trimmed, "!=") &&
			!strings.HasSuffix(trimmed, ",") && !strings.HasSuffix(trimmed, "{") {

			// Check if next non-empty line is another attribute or closing brace
			for j := i + 1; j < len(lines); j++ {
				nextTrimmed := strings.TrimSpace(lines[j])
				if nextTrimmed == "" {
					continue
				}

				// If next line is another attribute, add comma
				if strings.Contains(nextTrimmed, "=") &&
					!strings.HasPrefix(nextTrimmed, "}") &&
					!strings.HasPrefix(nextTrimmed, "]") {
					// Special case: don't add comma after EOF
					if !strings.HasSuffix(trimmed, "EOF") {
						lines[i] = lines[i] + ","
					}
				}
				break
			}
		}
	}

	return strings.Join(lines, "\n")
}

// normalizeForExpressionSpacing removes extra spaces in for expressions
func normalizeForExpressionSpacing(content string) string {
	// Replace { for with {for (remove space after opening brace in for expressions)
	content = strings.ReplaceAll(content, "{ for ", "{for ")
	// Replace } at end of for expressions (remove space before closing brace)
	content = strings.ReplaceAll(content, " }", "}")
	return content
}

// fixVariableInterpolationInTargetResources only fixes double dollars in specific resource types
func fixVariableInterpolationInTargetResources(content string) string {
	// Parse the HCL to identify resource blocks
	file, diags := hclwrite.ParseConfig([]byte(content), "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		// If we can't parse, return original content
		return content
	}

	// Target resource types
	targetResources := map[string]bool{
		"cloudflare_load_balancer":      true,
		"cloudflare_load_balancer_pool": true,
		"cloudflare_ruleset":            true,
	}

	modified := false
	for _, block := range file.Body().Blocks() {
		if block.Type() != "resource" {
			continue
		}

		labels := block.Labels()
		if len(labels) < 1 {
			continue
		}

		resourceType := strings.Trim(labels[0], "\"")
		if !targetResources[resourceType] {
			continue
		}

		// Get the block content as string
		blockBytes := hclwrite.Format(block.BuildTokens(nil).Bytes())
		blockContent := string(blockBytes)

		// Fix double dollars in this block
		fixedBlock := fixDoubleDollarInContent(blockContent)

		if fixedBlock != blockContent {
			modified = true
			// We need to replace the block content in the original string
			// This is a bit tricky with HCL, so we'll use a different approach
			// We'll fix the entire file content but only for lines within these resource blocks
		}
	}

	// If we found target resources that need fixing, apply fixes to the whole content
	// This is simpler than trying to surgically replace blocks
	if modified {
		return fixDoubleDollarInContentForResources(content, targetResources)
	}

	return content
}

// fixDoubleDollarInContentForResources fixes double dollars only within specific resource blocks
func fixDoubleDollarInContentForResources(content string, targetResources map[string]bool) string {
	lines := strings.Split(content, "\n")
	inTargetResource := false
	resourceDepth := 0
	result := []string{}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check if we're entering a resource block
		if strings.HasPrefix(trimmed, "resource ") {
			// Check if it's a target resource
			for resourceType := range targetResources {
				if strings.Contains(line, "\""+resourceType+"\"") {
					inTargetResource = true
					resourceDepth = 0
					break
				}
			}
			if !inTargetResource {
				// Check if any other resource type is mentioned (to mark we're in a non-target resource)
				if strings.Contains(line, "resource ") {
					inTargetResource = false
				}
			}
		}

		// Track braces to know when we exit the resource block
		if inTargetResource {
			// Count braces to track nesting
			openBraces := strings.Count(line, "{")
			closeBraces := strings.Count(line, "}")
			resourceDepth += openBraces - closeBraces

			// Apply fixes to this line
			line = fixDoubleDollarInContent(line)

			// Check if we've exited the resource block
			if resourceDepth <= 0 && strings.Contains(trimmed, "}") {
				inTargetResource = false
			}
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func fixDoubleDollarInContent(content string) string {
	// Common patterns where $$ should be $ in Terraform:

	// 1. Fix $${var. patterns (variable interpolations)
	// These should be ${var.
	content = regexp.MustCompile(`\$\$\{var\.`).ReplaceAllString(content, `${var.`)

	// 2. Fix $${local. patterns (local references)
	// These should be ${local.
	content = regexp.MustCompile(`\$\$\{local\.`).ReplaceAllString(content, `${local.`)

	// 3. Fix $${data. patterns (data source references)
	// These should be ${data.
	content = regexp.MustCompile(`\$\$\{data\.`).ReplaceAllString(content, `${data.`)

	// 4. Fix $${module. patterns (module references)
	// These should be ${module.
	content = regexp.MustCompile(`\$\$\{module\.`).ReplaceAllString(content, `${module.`)

	// 5. Fix $${resource_type. patterns (resource references)
	// These should be ${resource_type.
	// Common Cloudflare resource types
	resourceTypes := []string{
		"cloudflare_ruleset",
		"cloudflare_load_balancer",
		"cloudflare_load_balancer_pool",
		"cloudflare_access_application",
		"cloudflare_access_group",
		"cloudflare_access_policy",
		"cloudflare_teams_account",
		"cloudflare_teams_location",
		"cloudflare_teams_rule",
		"cloudflare_zone",
		"cloudflare_record",
		"cloudflare_page_rule",
		"cloudflare_rate_limit",
		"cloudflare_filter",
		"cloudflare_firewall_rule",
		"cloudflare_tunnel",
		"cloudflare_tunnel_config",
		"cloudflare_certificate_pack",
		"cloudflare_custom_hostname",
		"cloudflare_custom_ssl",
		"cloudflare_healthcheck",
		"cloudflare_logpush_job",
		"cloudflare_notification_policy",
		"cloudflare_pages_project",
		"cloudflare_spectrum_application",
		"cloudflare_waiting_room",
		"cloudflare_worker_route",
		"cloudflare_worker_script",
		"cloudflare_zone_lockdown",
	}

	for _, resourceType := range resourceTypes {
		pattern := regexp.MustCompile(`\$\$\{` + regexp.QuoteMeta(resourceType) + `\.`)
		content = pattern.ReplaceAllString(content, `${`+resourceType+`.`)
	}

	// 6. Fix $${count. patterns
	// These should be ${count.
	content = regexp.MustCompile(`\$\$\{count\.`).ReplaceAllString(content, `${count.`)

	// 7. Fix $${each. patterns (for for_each)
	// These should be ${each.
	content = regexp.MustCompile(`\$\$\{each\.`).ReplaceAllString(content, `${each.`)

	// 8. Fix $${self. patterns
	// These should be ${self.
	content = regexp.MustCompile(`\$\$\{self\.`).ReplaceAllString(content, `${self.`)

	// 9. Fix $${path. patterns
	// These should be ${path.
	content = regexp.MustCompile(`\$\$\{path\.`).ReplaceAllString(content, `${path.`)

	// 10. Fix $${terraform. patterns
	// These should be ${terraform.
	content = regexp.MustCompile(`\$\$\{terraform\.`).ReplaceAllString(content, `${terraform.`)

	// 11. Fix template string interpolations in expressions like regex_replace
	// Pattern: $${1}, $${2}, etc. in regex replacements should stay as $${1}
	// But $${var.something} should be ${var.something}
	// So we need to be more careful here

	// 12. Fix function calls that were incorrectly escaped
	// $${func( should be ${func(
	functions := []string{
		"abs", "ceil", "floor", "log", "max", "min", "parseint", "pow", "signum",
		"chomp", "format", "formatlist", "indent", "join", "lower", "regex", "regexall",
		"replace", "regex_replace", "split", "strrev", "substr", "title", "trim", "trimprefix",
		"trimsuffix", "trimspace", "upper", "chunklist", "coalesce", "coalescelist",
		"compact", "concat", "contains", "distinct", "element", "flatten", "index",
		"keys", "length", "list", "lookup", "map", "matchkeys", "merge", "range",
		"reverse", "setintersection", "setproduct", "setsubtract", "setunion", "slice",
		"sort", "transpose", "values", "zipmap", "base64decode", "base64encode",
		"base64gzip", "csvdecode", "jsondecode", "jsonencode", "textdecodebase64",
		"textencodebase64", "urlencode", "yamldecode", "yamlencode", "abspath",
		"dirname", "pathexpand", "basename", "file", "fileexists", "fileset",
		"filebase64", "templatefile", "formatdate", "timeadd", "timestamp", "base64sha256",
		"base64sha512", "bcrypt", "filebase64sha256", "filebase64sha512", "filemd5",
		"filesha1", "filesha256", "filesha512", "md5", "rsadecrypt", "sha1", "sha256",
		"sha512", "uuid", "uuidv5", "cidrhost", "cidrnetmask", "cidrsubnet", "cidrsubnets",
		"tobool", "tolist", "tomap", "tonumber", "toset", "tostring", "try", "can",
		"defaults", "nonsensitive", "sensitive", "templatestring", "one", "sum",
		"alltrue", "anytrue",
	}

	for _, fn := range functions {
		pattern := regexp.MustCompile(`\$\$\{` + regexp.QuoteMeta(fn) + `\(`)
		content = pattern.ReplaceAllString(content, `${`+fn+`(`)
	}

	// 13. General pattern: Any $${word. where word starts with a letter
	// should likely be ${word. unless it's a number like $${1}
	// This catches any remaining terraform interpolation patterns
	content = regexp.MustCompile(`\$\$\{([a-zA-Z][a-zA-Z0-9_]*\.)`).ReplaceAllString(content, `${$1`)

	// 14. Fix any $${" patterns (string interpolations)
	// These should be ${"
	content = regexp.MustCompile(`\$\$\{"`).ReplaceAllString(content, `${"`)

	// 15. Fix for expressions: $${[for
	// These should be ${[for
	content = regexp.MustCompile(`\$\$\{\[for`).ReplaceAllString(content, `${[for`)

	// 16. Fix conditional expressions with ?
	// $${condition ? should be ${condition ?
	// This catches any remaining conditional patterns not already caught
	content = regexp.MustCompile(`\$\$\{([^}]*\?[^}]*)\}`).ReplaceAllString(content, `${$1}`)

	// 17. Special case: In regex_replace or similar functions, $${1}, $${2} etc. should remain as is
	// They are correctly escaped for regex backreferences
	// So we don't touch patterns like $${number}

	return content
}
