package config

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/common"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlocksToListConverter(t *testing.T) {
	hclContent := `resource "test" "example" {
  name = "test"
  
  item {
    key = "value1"
    id = 1
  }
  
  item {
    key = "value2"
    id = 2
  }
  
  other_block {
    data = "keep"
  }
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &common.TransformContext{}

	transformer := BlocksToListConverter("item")
	err := transformer(block, ctx)
	require.NoError(t, err)

	body := block.Body()
	
	// Check that item attribute was created
	assert.NotNil(t, body.GetAttribute("item"))
	
	// Check that item blocks were removed
	itemBlockCount := 0
	otherBlockCount := 0
	for _, b := range body.Blocks() {
		if b.Type() == "item" {
			itemBlockCount++
		}
		if b.Type() == "other_block" {
			otherBlockCount++
		}
	}
	assert.Equal(t, 0, itemBlockCount, "Item blocks should be removed")
	assert.Equal(t, 1, otherBlockCount, "Other blocks should be preserved")
}