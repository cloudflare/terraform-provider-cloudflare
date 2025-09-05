package main

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/zclconf/go-cty/cty"
)

// WorkersSecretInfo holds extracted information from a workers_secret resource
type WorkersSecretInfo struct {
	ScriptName  string // From script_name attribute
	SecretName  string // From name attribute
	SecretText  string // From secret_text attribute
	AccountID   string // From account_id attribute
	ResourceID  string // Resource identifier (for removing from config)
}

// isWorkersSecretResource checks if a block is a workers_secret resource
func isWorkersSecretResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_secret" || block.Labels()[0] == "cloudflare_worker_secret")
}

// migrateWorkersSecretsToBindings performs cross-resource migration of workers_secret to workers_script bindings
// This function should be called after all individual block transformations are complete
func migrateWorkersSecretsToBindings(file *hclwrite.File, diags ast.Diagnostics) {
	// Phase 1: Collect all workers_secret resources
	var secretsToMigrate []WorkersSecretInfo
	var secretBlocksToRemove []*hclwrite.Block

	for _, block := range file.Body().Blocks() {
		if isWorkersSecretResource(block) {
			// Extract secret information
			secretInfo := extractWorkersSecretInfo(block, diags)
			if secretInfo != nil {
				secretsToMigrate = append(secretsToMigrate, *secretInfo)
				secretBlocksToRemove = append(secretBlocksToRemove, block)
			}
		}
	}

	// Phase 2: Find workers_script resources and add secret_text bindings
	for _, block := range file.Body().Blocks() {
		if isWorkersScriptResource(block) {
			// Find secrets that belong to this script
			secretsForScript := findSecretsForScript(block, secretsToMigrate, diags)
			if len(secretsForScript) > 0 {
				// Add secret_text bindings to this script
				addSecretBindingsToScript(block, secretsForScript, diags)
			}
		}
	}

	// Phase 3: Remove workers_secret blocks and add migration warnings
	for _, secretBlock := range secretBlocksToRemove {
		// Add warning comment before removing
		warningTokens := []*hclwrite.Token{
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenComment, Bytes: []byte(`# MIGRATION: cloudflare_workers_secret has been migrated to secret_text binding in cloudflare_workers_script`)},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		}
		file.Body().AppendUnstructuredTokens(warningTokens)

		// Remove the workers_secret block
		file.Body().RemoveBlock(secretBlock)
	}
}

// extractWorkersSecretInfo extracts key information from a workers_secret block
func extractWorkersSecretInfo(block *hclwrite.Block, diags ast.Diagnostics) *WorkersSecretInfo {
	if len(block.Labels()) < 2 {
		// Invalid block structure - add warning
		addManualMigrationWarning(block, "Invalid workers_secret block structure - please migrate manually", diags)
		return nil
	}

	attributes := block.Body().Attributes()

	// Extract script_name
	scriptName, err := extractAttributeString(attributes, "script_name", diags)
	if err != nil {
		addManualMigrationWarning(block, "Unable to extract script_name - please migrate manually", diags)
		return nil
	}

	// Extract secret name
	secretName, err := extractAttributeString(attributes, "name", diags)
	if err != nil {
		addManualMigrationWarning(block, "Unable to extract secret name - please migrate manually", diags)
		return nil
	}

	// Extract secret text
	secretText, err := extractAttributeString(attributes, "secret_text", diags)
	if err != nil {
		addManualMigrationWarning(block, "Unable to extract secret_text - please migrate manually", diags)
		return nil
	}

	// Extract account_id
	accountID, err := extractAttributeString(attributes, "account_id", diags)
	if err != nil {
		addManualMigrationWarning(block, "Unable to extract account_id - please migrate manually", diags)
		return nil
	}

	return &WorkersSecretInfo{
		ScriptName:  scriptName,
		SecretName:  secretName,
		SecretText:  secretText,
		AccountID:   accountID,
		ResourceID:  block.Labels()[1], // Resource instance name
	}
}

// extractAttributeString extracts a string value from an attribute, handling various formats
func extractAttributeString(attributes map[string]*hclwrite.Attribute, attrName string, diags ast.Diagnostics) (string, error) {
	attr, exists := attributes[attrName]
	if !exists {
		return "", fmt.Errorf("%s attribute not found", attrName)
	}

	// Convert to string representation for simple parsing
	attrStr := ast.Expr2S(ast.WriteExpr2Expr(attr.Expr(), diags), diags)
	
	// Handle quoted strings
	if strings.HasPrefix(attrStr, `"`) && strings.HasSuffix(attrStr, `"`) {
		return strings.Trim(attrStr, `"`), nil
	}
	
	// Handle variable references - return as-is for now (will need manual migration)
	if strings.Contains(attrStr, "var.") || strings.Contains(attrStr, "local.") {
		return "", fmt.Errorf("variable reference detected: %s", attrStr)
	}
	
	// Handle resource references (e.g., cloudflare_workers_script.foo.script_name)
	if strings.Contains(attrStr, "cloudflare_workers_script.") || strings.Contains(attrStr, "cloudflare_worker_script.") {
		return "", fmt.Errorf("resource reference detected: %s", attrStr)
	}
	
	// Return literal value
	return attrStr, nil
}

// findSecretsForScript finds all secrets that belong to a specific workers_script
func findSecretsForScript(scriptBlock *hclwrite.Block, secrets []WorkersSecretInfo, diags ast.Diagnostics) []WorkersSecretInfo {
	var matchingSecrets []WorkersSecretInfo
	
	// Get the script name from the workers_script block
	attributes := scriptBlock.Body().Attributes()
	scriptName, err := extractAttributeString(attributes, "script_name", diags)
	if err != nil {
		// Try "name" attribute (v4 format)
		scriptName, err = extractAttributeString(attributes, "name", diags)
		if err != nil {
			return matchingSecrets // No script name found
		}
	}
	
	accountID, err := extractAttributeString(attributes, "account_id", diags)
	if err != nil {
		return matchingSecrets // No account ID found
	}
	
	// Find secrets that match this script
	for _, secret := range secrets {
		if secret.ScriptName == scriptName && secret.AccountID == accountID {
			matchingSecrets = append(matchingSecrets, secret)
		}
	}
	
	return matchingSecrets
}

// addSecretBindingsToScript adds secret_text bindings to a workers_script resource
func addSecretBindingsToScript(block *hclwrite.Block, secrets []WorkersSecretInfo, diags ast.Diagnostics) {
	// Create secret_text bindings for each secret
	var newBindings []hclsyntax.Expression
	
	for _, secret := range secrets {
		// Create a secret_text binding object
		binding := &hclsyntax.ObjectConsExpr{
			Items: []hclsyntax.ObjectConsItem{
				{
					KeyExpr:   ast.NewKeyExpr("type"),
					ValueExpr: &hclsyntax.LiteralValueExpr{Val: cty.StringVal("secret_text")},
				},
				{
					KeyExpr:   ast.NewKeyExpr("name"),
					ValueExpr: &hclsyntax.LiteralValueExpr{Val: cty.StringVal(secret.SecretName)},
				},
				{
					KeyExpr:   ast.NewKeyExpr("text"),
					ValueExpr: &hclsyntax.LiteralValueExpr{Val: cty.StringVal(secret.SecretText)},
				},
			},
		}
		newBindings = append(newBindings, binding)
	}
	
	if len(newBindings) == 0 {
		return
	}
	
	// Check if the block already has bindings attribute
	existingBindingsAttr := block.Body().GetAttribute("bindings")
	var allBindings []hclsyntax.Expression
	
	if existingBindingsAttr != nil {
		// Parse existing bindings
		existingExpr := ast.WriteExpr2Expr(existingBindingsAttr.Expr(), diags)
		if tuple, ok := existingExpr.(*hclsyntax.TupleConsExpr); ok {
			allBindings = append(allBindings, tuple.Exprs...)
		} else {
			// Single binding or other format - preserve it
			allBindings = append(allBindings, existingExpr)
		}
	}
	
	// Add new secret bindings
	allBindings = append(allBindings, newBindings...)
	
	// Create the new bindings expression
	bindingsExpr := &hclsyntax.TupleConsExpr{
		Exprs: allBindings,
	}
	
	// Set the bindings attribute
	transforms := map[string]ast.ExprTransformer{
		"bindings": func(expr *hclsyntax.Expression, diags ast.Diagnostics) {
			*expr = bindingsExpr
		},
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

// addManualMigrationWarning adds a warning comment for manual migration
func addManualMigrationWarning(block *hclwrite.Block, message string, diags ast.Diagnostics) {
	warningTokens := []*hclwrite.Token{
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenComment, Bytes: []byte("# MIGRATION WARNING: " + message)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
	}
	block.Body().AppendUnstructuredTokens(warningTokens)
}

// transformWorkersSecretStateJSON migrates workers_secret state to workers_script bindings
// This function collects workers_secret state information for cross-resource migration
func transformWorkersSecretStateJSON(jsonStr string, path string) string {
	// Extract workers_secret information from state
	scriptNamePath := path + ".attributes.script_name"
	secretNamePath := path + ".attributes.name"
	secretTextPath := path + ".attributes.secret_text"
	accountIDPath := path + ".attributes.account_id"

	scriptName := gjson.Get(jsonStr, scriptNamePath)
	secretName := gjson.Get(jsonStr, secretNamePath)
	secretText := gjson.Get(jsonStr, secretTextPath)
	accountID := gjson.Get(jsonStr, accountIDPath)

	if scriptName.Exists() && secretName.Exists() && secretText.Exists() && accountID.Exists() {
		// Store this secret info for later cross-resource migration
		// For now, we return the original state - the removal will be handled by the main migration pipeline
		secretInfo := WorkersSecretStateInfo{
			ScriptName: scriptName.String(),
			SecretName: secretName.String(),
			SecretText: secretText.String(),
			AccountID:  accountID.String(),
			StatePath:  path,
		}
		
		// Add to global collection for cross-resource processing
		addWorkersSecretForStateMigration(secretInfo)
	}
	
	// Return original state for now - actual removal happens in migrateWorkersSecretsInState
	return jsonStr
}

// WorkersSecretStateInfo holds state migration information
type WorkersSecretStateInfo struct {
	ScriptName string
	SecretName string
	SecretText string
	AccountID  string
	StatePath  string
}

// Global collection for state migration (will be moved to proper context later)
var workersSecretsForStateMigration []WorkersSecretStateInfo

// addWorkersSecretForStateMigration adds a secret to the migration collection
func addWorkersSecretForStateMigration(secret WorkersSecretStateInfo) {
	workersSecretsForStateMigration = append(workersSecretsForStateMigration, secret)
}

// migrateWorkersSecretsInState performs cross-resource state migration
// This should be called after all individual resource state transformations
func migrateWorkersSecretsInState(jsonStr string) string {
	result := jsonStr
	
	// Find all workers_script resources in state and add secret bindings
	scriptsPath := "resources"
	scriptsValue := gjson.Get(jsonStr, scriptsPath)
	
	if scriptsValue.Exists() && scriptsValue.IsArray() {
		for i, resource := range scriptsValue.Array() {
			resourceType := resource.Get("type").String()
			if resourceType == "cloudflare_workers_script" || resourceType == "cloudflare_worker_script" {
				// Find secrets that belong to this script
				scriptPath := fmt.Sprintf("resources.%d", i)
				result = addSecretsToScriptState(result, scriptPath, workersSecretsForStateMigration)
			}
		}
	}
	
	// Remove all workers_secret resources from state
	result = removeWorkersSecretResourcesFromState(result)
	
	// Clear the collection for next migration
	workersSecretsForStateMigration = nil
	
	return result
}

// addSecretsToScriptState adds secret_text bindings to a workers_script resource in state
func addSecretsToScriptState(jsonStr string, scriptResourcePath string, secrets []WorkersSecretStateInfo) string {
	result := jsonStr
	
	// Get script information
	scriptNamePath := scriptResourcePath + ".instances.0.attributes.script_name"
	scriptAccountPath := scriptResourcePath + ".instances.0.attributes.account_id"
	
	// Try both script_name and name (for v4 compatibility)
	scriptName := gjson.Get(jsonStr, scriptNamePath)
	if !scriptName.Exists() {
		scriptNamePath = scriptResourcePath + ".instances.0.attributes.name"
		scriptName = gjson.Get(jsonStr, scriptNamePath)
	}
	
	scriptAccount := gjson.Get(jsonStr, scriptAccountPath)
	
	if !scriptName.Exists() || !scriptAccount.Exists() {
		return result // Cannot match script
	}
	
	// Find matching secrets
	var matchingSecrets []WorkersSecretStateInfo
	for _, secret := range secrets {
		if secret.ScriptName == scriptName.String() && secret.AccountID == scriptAccount.String() {
			matchingSecrets = append(matchingSecrets, secret)
		}
	}
	
	if len(matchingSecrets) == 0 {
		return result // No secrets for this script
	}
	
	// Get existing bindings
	bindingsPath := scriptResourcePath + ".instances.0.attributes.bindings"
	existingBindings := gjson.Get(jsonStr, bindingsPath)
	
	var allBindings []interface{}
	
	// Parse existing bindings
	if existingBindings.Exists() && existingBindings.IsArray() {
		for _, binding := range existingBindings.Array() {
			allBindings = append(allBindings, binding.Value())
		}
	}
	
	// Add secret bindings
	for _, secret := range matchingSecrets {
		secretBinding := map[string]interface{}{
			"type": "secret_text",
			"name": secret.SecretName,
			"text": secret.SecretText,
		}
		allBindings = append(allBindings, secretBinding)
	}
	
	// Update state with new bindings
	result, _ = sjson.Set(result, bindingsPath, allBindings)
	
	return result
}

// removeWorkersSecretResourcesFromState removes all workers_secret resources from state
func removeWorkersSecretResourcesFromState(jsonStr string) string {
	result := jsonStr
	
	// Get all resources
	resourcesPath := "resources"
	resourcesValue := gjson.Get(jsonStr, resourcesPath)
	
	if !resourcesValue.Exists() || !resourcesValue.IsArray() {
		return result
	}
	
	// Filter out workers_secret resources
	var remainingResources []interface{}
	for _, resource := range resourcesValue.Array() {
		resourceType := resource.Get("type").String()
		if resourceType != "cloudflare_workers_secret" && resourceType != "cloudflare_worker_secret" {
			remainingResources = append(remainingResources, resource.Value())
		}
	}
	
	// Update state with filtered resources
	result, _ = sjson.Set(result, resourcesPath, remainingResources)
	
	return result
}
