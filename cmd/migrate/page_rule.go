package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// transformPageRuleConfig does string-level transformations on page_rule resources
// This is called AFTER Grit has done its transformations
func transformPageRuleConfig(content string) string {
	// Remove minify completely - it's not supported in v5
	content = removeMinifyFromActions(content)

	// Consolidate multiple cache_ttl_by_status entries into a single map
	content = consolidateCacheTTLByStatus(content)

	return content
}

func removeMinifyFromActions(content string) string {
	// Pattern to match minify blocks after Grit has converted them to attribute syntax
	// Match from "minify = {" to its closing "}" including any blank lines after
	minifyPattern := regexp.MustCompile(`(?ms)\s*minify\s*=\s*\{[^{}]*\}\s*`)
	content = minifyPattern.ReplaceAllString(content, "\n    ")
	return content
}

func consolidateCacheTTLByStatus(content string) string {
	// Find page_rule resources and process their actions
	pageRulePattern := regexp.MustCompile(`(?ms)resource\s+"cloudflare_page_rule"[^{]+\{`)

	// Check if we have any page rules to process
	if !pageRulePattern.MatchString(content) {
		return content
	}

	// First, find all cache_ttl_by_status blocks and collect the data
	// Pattern matches the multiline structure after Grit:
	// cache_ttl_by_status = {
	//   codes = "XXX"
	//   ttl   = YYY
	// }
	ttlBlockPattern := regexp.MustCompile(`(?ms)cache_ttl_by_status\s*=\s*\{\s*codes\s*=\s*"([^"]+)"\s*ttl\s*=\s*(\d+)\s*\}`)

	matches := ttlBlockPattern.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return content
	}

	// Collect all the code-ttl pairs
	pairs := make(map[string]string)
	for _, match := range matches {
		if len(match) >= 3 {
			code := match[1]
			ttl := match[2]
			pairs[code] = ttl
		}
	}

	if len(pairs) == 0 {
		return content
	}

	// Sort codes for consistent output
	var codes []string
	for code := range pairs {
		codes = append(codes, code)
	}
	sort.Strings(codes)

	// Build the consolidated map
	var mapEntries []string
	for _, code := range codes {
		mapEntries = append(mapEntries, fmt.Sprintf(`"%s" = %s`, code, pairs[code]))
	}

	newTTL := "cache_ttl_by_status = { " + strings.Join(mapEntries, ", ") + " }"

	// Remove all old cache_ttl_by_status blocks including any trailing whitespace
	// This pattern includes any whitespace/newlines before and after the block
	ttlRemovePattern := regexp.MustCompile(`(?ms)\s*cache_ttl_by_status\s*=\s*\{[^{}]*\}\s*`)
	content = ttlRemovePattern.ReplaceAllString(content, "\n    ")

	// Now insert the consolidated one after cache_level
	// Find the cache_level line and add our new TTL after it
	cacheLevelPattern := regexp.MustCompile(`(cache_level\s*=\s*"[^"]+")`)
	replaced := false
	content = cacheLevelPattern.ReplaceAllStringFunc(content, func(match string) string {
		if !replaced {
			replaced = true
			return match + "\n    " + newTTL
		}
		return match
	})

	return content
}