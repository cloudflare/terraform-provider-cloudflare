package structural

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrayToObjectArrayConverter(t *testing.T) {
	hclContent := `resource "test" "example" {
  name = "test"
  policies = ["id1", "id2", cloudflare_access_policy.policy1.id]
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &basic.TransformContext{}

	transformer := ArrayToObjectArrayConverter("policies", "id")
	err := transformer(block, ctx)
	require.NoError(t, err)

	body := block.Body()
	attr := body.GetAttribute("policies")
	assert.NotNil(t, attr)
	
	// Check that the expression was transformed
	tokens := attr.Expr().BuildTokens(nil)
	result := string(tokens.Bytes())
	
	assert.Contains(t, result, "{ id = \"id1\" }")
	assert.Contains(t, result, "{ id = \"id2\" }")
	assert.Contains(t, result, "{ id = cloudflare_access_policy.policy1.id }")
}

func TestStringArrayToObjectArrayWithRename(t *testing.T) {
	hclContent := `resource "test" "example" {
  name = "test"
  policies = [cloudflare_access_policy.old.id, "static_id"]
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &basic.TransformContext{}

	transformer := StringArrayToObjectArrayWithRename(
		"policies", 
		"id",
		"cloudflare_access_policy.",
		"cloudflare_zero_trust_access_policy.",
	)
	err := transformer(block, ctx)
	require.NoError(t, err)

	body := block.Body()
	attr := body.GetAttribute("policies")
	assert.NotNil(t, attr)
	
	// Check that the expression was transformed and renamed
	tokens := attr.Expr().BuildTokens(nil)
	result := string(tokens.Bytes())
	
	assert.Contains(t, result, "{ id = cloudflare_zero_trust_access_policy.old.id }")
	assert.Contains(t, result, "{ id = \"static_id\" }")
	assert.NotContains(t, result, "cloudflare_access_policy")
}

func TestArrayToObjectArrayConverter_EmptyArray(t *testing.T) {
	hclContent := `resource "test" "example" {
  name = "test"
  policies = []
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &basic.TransformContext{}

	transformer := ArrayToObjectArrayConverter("policies", "id")
	err := transformer(block, ctx)
	require.NoError(t, err)

	body := block.Body()
	attr := body.GetAttribute("policies")
	assert.NotNil(t, attr)
	
	// Check that empty array remains empty
	tokens := attr.Expr().BuildTokens(nil)
	result := string(tokens.Bytes())
	assert.Equal(t, "[]", result)
}

func TestArrayToObjectArrayConverter_AlreadyTransformed(t *testing.T) {
	hclContent := `resource "test" "example" {
  name = "test"
  policies = [{ id = "already_transformed" }, { id = "another_one" }]
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &basic.TransformContext{}

	transformer := ArrayToObjectArrayConverter("policies", "id")
	err := transformer(block, ctx)
	require.NoError(t, err)

	body := block.Body()
	attr := body.GetAttribute("policies")
	assert.NotNil(t, attr)
	
	// Check that already transformed objects are preserved
	tokens := attr.Expr().BuildTokens(nil)
	result := string(tokens.Bytes())
	assert.Contains(t, result, "{ id = \"already_transformed\" }")
	assert.Contains(t, result, "{ id = \"another_one\" }")
}