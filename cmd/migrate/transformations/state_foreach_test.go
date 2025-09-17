package transformations

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPreservesForEachIndexKey tests that index_key is preserved for for_each resources
func TestPreservesForEachIndexKey(t *testing.T) {
	// Create a test state file with for_each resources
	inputState := &TerraformState{
		Version:          4,
		TerraformVersion: "1.5.0",
		Serial:           1,
		Lineage:          "test",
		Resources: []TerraformStateResource{
			{
				Mode:     "managed",
				Type:     "cloudflare_load_balancer_pool",
				Name:     "test_pools",
				Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
				Instances: []TerraformStateInstance{
					{
						IndexKey:      "pool1",
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":         "id-1",
							"name":       "pool1",
							"account_id": "test-account",
						},
						Private: "bnVsbA==",
					},
					{
						IndexKey:      "pool2",
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":         "id-2",
							"name":       "pool2",
							"account_id": "test-account",
						},
						Private: "bnVsbA==",
					},
				},
			},
		},
	}

	// Write the input state to a temp file
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "input.tfstate")
	outputPath := filepath.Join(tmpDir, "output.tfstate")

	inputBytes, err := json.MarshalIndent(inputState, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(inputPath, inputBytes, 0644))

	// Create a minimal transformation config
	config := &StateTransformationConfig{
		Version:     "1.0",
		Description: "Test transformation",
		SchemaVersionReset: SchemaVersionReset{
			AllCloudflareResources: true,
		},
		StateAttributeRenames: map[string]map[string]interface{}{
			"cloudflare_load_balancer_pool": {
				"old_field": "new_field",
			},
		},
	}

	// Run the transformation
	err = TransformStateFile(config, inputPath, outputPath)
	require.NoError(t, err)

	// Read the output state
	outputBytes, err := os.ReadFile(outputPath)
	require.NoError(t, err)

	var outputState TerraformState
	err = json.Unmarshal(outputBytes, &outputState)
	require.NoError(t, err)

	// Verify the resources still exist
	require.Len(t, outputState.Resources, 1)
	resource := outputState.Resources[0]
	
	// Verify instances are preserved
	require.Len(t, resource.Instances, 2)
	
	// Verify index_key fields are preserved
	assert.Equal(t, "pool1", resource.Instances[0].IndexKey)
	assert.Equal(t, "pool2", resource.Instances[1].IndexKey)
	
	// Verify other attributes are preserved
	assert.Equal(t, "id-1", resource.Instances[0].Attributes["id"])
	assert.Equal(t, "id-2", resource.Instances[1].Attributes["id"])
}

// TestPreservesCountIndex tests that index_key is preserved for count resources
func TestPreservesCountIndex(t *testing.T) {
	// Create a test state file with count resources
	inputState := &TerraformState{
		Version:          4,
		TerraformVersion: "1.5.0",
		Serial:           1,
		Lineage:          "test",
		Resources: []TerraformStateResource{
			{
				Mode:     "managed",
				Type:     "cloudflare_dns_record",
				Name:     "test_records",
				Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
				Instances: []TerraformStateInstance{
					{
						IndexKey:      0,
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":      "id-1",
							"name":    "record-0",
							"zone_id": "test-zone",
						},
					},
					{
						IndexKey:      1,
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":      "id-2",
							"name":    "record-1",
							"zone_id": "test-zone",
						},
					},
					{
						IndexKey:      2,
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":      "id-3",
							"name":    "record-2",
							"zone_id": "test-zone",
						},
					},
				},
			},
		},
	}

	// Write the input state to a temp file
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "input.tfstate")
	outputPath := filepath.Join(tmpDir, "output.tfstate")

	inputBytes, err := json.MarshalIndent(inputState, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(inputPath, inputBytes, 0644))

	// Create a minimal transformation config
	config := &StateTransformationConfig{
		Version:     "1.0",
		Description: "Test transformation",
		SchemaVersionReset: SchemaVersionReset{
			AllCloudflareResources: true,
		},
	}

	// Run the transformation
	err = TransformStateFile(config, inputPath, outputPath)
	require.NoError(t, err)

	// Read the output state
	outputBytes, err := os.ReadFile(outputPath)
	require.NoError(t, err)

	var outputState TerraformState
	err = json.Unmarshal(outputBytes, &outputState)
	require.NoError(t, err)

	// Verify the resources still exist
	require.Len(t, outputState.Resources, 1)
	resource := outputState.Resources[0]
	
	// Verify instances are preserved
	require.Len(t, resource.Instances, 3)
	
	// Verify index_key fields are preserved for count (numeric indices)
	assert.Equal(t, float64(0), resource.Instances[0].IndexKey) // JSON unmarshals numbers as float64
	assert.Equal(t, float64(1), resource.Instances[1].IndexKey)
	assert.Equal(t, float64(2), resource.Instances[2].IndexKey)
}

// TestHandlesMixedInstanceTypes tests resources with both regular and indexed instances
func TestHandlesMixedInstanceTypes(t *testing.T) {
	// Create a test state with a regular resource (no for_each/count)
	inputState := &TerraformState{
		Version:          4,
		TerraformVersion: "1.5.0",
		Serial:           1,
		Lineage:          "test",
		Resources: []TerraformStateResource{
			{
				Mode:     "managed",
				Type:     "cloudflare_zone",
				Name:     "test_zone",
				Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
				Instances: []TerraformStateInstance{
					{
						// No IndexKey for regular resources
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":   "zone-id",
							"name": "example.com",
						},
					},
				},
			},
			{
				Mode:     "managed",
				Type:     "cloudflare_record",
				Name:     "test_records",
				Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
				Instances: []TerraformStateInstance{
					{
						IndexKey:      "www",
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":   "record-1",
							"name": "www",
						},
					},
					{
						IndexKey:      "api",
						SchemaVersion: 0,
						Attributes: map[string]interface{}{
							"id":   "record-2",
							"name": "api",
						},
					},
				},
			},
		},
	}

	// Write the input state to a temp file
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "input.tfstate")
	outputPath := filepath.Join(tmpDir, "output.tfstate")

	inputBytes, err := json.MarshalIndent(inputState, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(inputPath, inputBytes, 0644))

	// Create a minimal transformation config
	config := &StateTransformationConfig{
		Version:     "1.0",
		Description: "Test transformation",
		SchemaVersionReset: SchemaVersionReset{
			AllCloudflareResources: true,
		},
	}

	// Run the transformation
	err = TransformStateFile(config, inputPath, outputPath)
	require.NoError(t, err)

	// Read the output state
	outputBytes, err := os.ReadFile(outputPath)
	require.NoError(t, err)

	var outputState TerraformState
	err = json.Unmarshal(outputBytes, &outputState)
	require.NoError(t, err)

	// Verify both resources exist
	require.Len(t, outputState.Resources, 2)
	
	// Check the regular resource (no index_key)
	zoneResource := outputState.Resources[0]
	require.Len(t, zoneResource.Instances, 1)
	assert.Nil(t, zoneResource.Instances[0].IndexKey)
	
	// Check the for_each resource (with index_key)
	recordResource := outputState.Resources[1]
	require.Len(t, recordResource.Instances, 2)
	assert.Equal(t, "www", recordResource.Instances[0].IndexKey)
	assert.Equal(t, "api", recordResource.Instances[1].IndexKey)
}