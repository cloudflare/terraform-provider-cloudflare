package workers_script

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContainersMetadataSerialization(t *testing.T) {
	containers := []*WorkersScriptMetadataContainersModel{
		{ClassName: types.StringValue("MyContainer")},
	}

	metadata := WorkersScriptMetadataModel{
		MainModule: types.StringValue("index.js"),
		Containers: &containers,
	}

	payload, err := apijson.Marshal(metadata)
	require.NoError(t, err)

	var raw map[string]interface{}
	err = json.Unmarshal(payload, &raw)
	require.NoError(t, err)

	// Verify containers array is present
	containersRaw, ok := raw["containers"]
	require.True(t, ok, "metadata should contain 'containers' key")

	containersList, ok := containersRaw.([]interface{})
	require.True(t, ok, "containers should be an array")
	require.Len(t, containersList, 1)

	first := containersList[0].(map[string]interface{})
	assert.Equal(t, "MyContainer", first["class_name"])
}

func TestContainersMetadataMultipleClasses(t *testing.T) {
	containers := []*WorkersScriptMetadataContainersModel{
		{ClassName: types.StringValue("ContainerA")},
		{ClassName: types.StringValue("ContainerB")},
	}

	metadata := WorkersScriptMetadataModel{
		MainModule: types.StringValue("index.js"),
		Containers: &containers,
	}

	payload, err := apijson.Marshal(metadata)
	require.NoError(t, err)

	var raw map[string]interface{}
	err = json.Unmarshal(payload, &raw)
	require.NoError(t, err)

	containersList := raw["containers"].([]interface{})
	require.Len(t, containersList, 2)

	assert.Equal(t, "ContainerA", containersList[0].(map[string]interface{})["class_name"])
	assert.Equal(t, "ContainerB", containersList[1].(map[string]interface{})["class_name"])
}

func TestContainersMetadataOmittedWhenNil(t *testing.T) {
	metadata := WorkersScriptMetadataModel{
		MainModule: types.StringValue("index.js"),
		Containers: nil,
	}

	payload, err := apijson.Marshal(metadata)
	require.NoError(t, err)

	var raw map[string]interface{}
	err = json.Unmarshal(payload, &raw)
	require.NoError(t, err)

	_, ok := raw["containers"]
	assert.False(t, ok, "containers should not be present when nil")
}

func TestContainersSchemaExists(t *testing.T) {
	ctx := context.Background()
	s := ResourceSchema(ctx)

	attr, ok := s.Attributes["containers"]
	require.True(t, ok, "schema should have 'containers' attribute")
	assert.True(t, attr.IsOptional(), "containers should be optional")
}
