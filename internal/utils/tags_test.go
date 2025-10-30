package utils

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestConvertTerraformTagsToAPI(t *testing.T) {
	t.Run("converts tags correctly", func(t *testing.T) {
		input := map[string]types.String{
			"env":  types.StringValue("production"),
			"team": types.StringValue("platform"),
		}
		result := ConvertTerraformTagsToAPI(&input)

		assert.Equal(t, "production", result["env"])
		assert.Equal(t, "platform", result["team"])
		assert.Len(t, result, 2)
	})

	t.Run("returns empty map for nil input", func(t *testing.T) {
		result := ConvertTerraformTagsToAPI(nil)
		assert.NotNil(t, result)
		assert.Len(t, result, 0)
	})

	t.Run("returns empty map for empty input", func(t *testing.T) {
		input := map[string]types.String{}
		result := ConvertTerraformTagsToAPI(&input)
		assert.NotNil(t, result)
		assert.Len(t, result, 0)
	})
}

func TestConvertAPITagsToTerraform(t *testing.T) {
	t.Run("converts tags correctly", func(t *testing.T) {
		input := map[string]string{
			"env":  "production",
			"team": "platform",
		}
		result := ConvertAPITagsToTerraform(input)

		assert.NotNil(t, result)
		assert.Equal(t, "production", (*result)["env"].ValueString())
		assert.Equal(t, "platform", (*result)["team"].ValueString())
		assert.Len(t, *result, 2)
	})

	t.Run("returns nil for empty input", func(t *testing.T) {
		input := map[string]string{}
		result := ConvertAPITagsToTerraform(input)
		assert.Nil(t, result)
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		result := ConvertAPITagsToTerraform(nil)
		assert.Nil(t, result)
	})
}

func TestTagsChanged(t *testing.T) {
	t.Run("returns false when both nil", func(t *testing.T) {
		assert.False(t, TagsChanged(nil, nil))
	})

	t.Run("returns true when planned is nil but state has tags", func(t *testing.T) {
		state := map[string]types.String{
			"env": types.StringValue("production"),
		}
		assert.True(t, TagsChanged(nil, &state))
	})

	t.Run("returns true when state is nil but planned has tags", func(t *testing.T) {
		planned := map[string]types.String{
			"env": types.StringValue("production"),
		}
		assert.True(t, TagsChanged(&planned, nil))
	})

	t.Run("returns false when tags are identical", func(t *testing.T) {
		planned := map[string]types.String{
			"env":  types.StringValue("production"),
			"team": types.StringValue("platform"),
		}
		state := map[string]types.String{
			"env":  types.StringValue("production"),
			"team": types.StringValue("platform"),
		}
		assert.False(t, TagsChanged(&planned, &state))
	})

	t.Run("returns true when tag value changes", func(t *testing.T) {
		planned := map[string]types.String{
			"env": types.StringValue("staging"),
		}
		state := map[string]types.String{
			"env": types.StringValue("production"),
		}
		assert.True(t, TagsChanged(&planned, &state))
	})

	t.Run("returns true when tag is added", func(t *testing.T) {
		planned := map[string]types.String{
			"env":  types.StringValue("production"),
			"team": types.StringValue("platform"),
		}
		state := map[string]types.String{
			"env": types.StringValue("production"),
		}
		assert.True(t, TagsChanged(&planned, &state))
	})

	t.Run("returns true when tag is removed", func(t *testing.T) {
		planned := map[string]types.String{
			"env": types.StringValue("production"),
		}
		state := map[string]types.String{
			"env":  types.StringValue("production"),
			"team": types.StringValue("platform"),
		}
		assert.True(t, TagsChanged(&planned, &state))
	})

	t.Run("returns false when both are empty maps", func(t *testing.T) {
		planned := map[string]types.String{}
		state := map[string]types.String{}
		assert.False(t, TagsChanged(&planned, &state))
	})
}
