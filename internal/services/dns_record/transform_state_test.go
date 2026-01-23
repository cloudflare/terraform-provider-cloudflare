package dns_record

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/tf-migrate/pkg/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestTransformDNSRecordState_SimpleARecord(t *testing.T) {
	// v4 state with 'value' field (should become 'content' in v5)
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1",
		"ttl": 300,
		"proxied": true
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Check value was renamed to content
	assert.Equal(t, "192.168.1.1", output["content"])
	assert.Nil(t, output["value"], "value field should be removed")

	// Check other fields preserved
	assert.Equal(t, "abc123", output["zone_id"])
	assert.Equal(t, "example.com", output["name"])
	assert.Equal(t, "A", output["type"])
	assert.Equal(t, float64(300), output["ttl"])
	assert.Equal(t, true, output["proxied"])

	// Check timestamps were added
	assert.NotNil(t, output["created_on"])
	assert.NotNil(t, output["modified_on"])

	// Check data field was removed for simple types
	assert.Nil(t, output["data"])
}

func TestTransformDNSRecordState_SimpleARecordWithContent(t *testing.T) {
	// v4 state that already has 'content' field
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"content": "192.168.1.1",
		"value": "192.168.1.1",
		"ttl": 300
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Content should be preserved, value removed
	assert.Equal(t, "192.168.1.1", output["content"])
	assert.Nil(t, output["value"])
}

func TestTransformDNSRecordState_MissingTTL(t *testing.T) {
	// v4 state missing TTL (required in v5)
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1"
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// TTL should be set to default value 1
	assert.Equal(t, float64(1), output["ttl"])
}

func TestTransformDNSRecordState_DeprecatedFields(t *testing.T) {
	// v4 state with deprecated fields
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1",
		"ttl": 300,
		"hostname": "example.com",
		"allow_overwrite": true,
		"proxiable": true
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Deprecated fields should be removed
	assert.Nil(t, output["hostname"])
	assert.Nil(t, output["allow_overwrite"])
	assert.Nil(t, output["proxiable"])
}

func TestTransformDNSRecordState_MetadataToMeta(t *testing.T) {
	// v4 state with metadata (renamed to meta in v5)
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1",
		"ttl": 300,
		"metadata": {"auto_added": "true"}
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// metadata should be renamed to meta
	assert.Nil(t, output["metadata"])
	assert.NotNil(t, output["meta"])

	// meta should be a JSON string, not an object
	// (jsontypes.NormalizedType expects a string containing JSON)
	metaValue, ok := output["meta"].(string)
	assert.True(t, ok, "meta should be a string, got %T", output["meta"])
	assert.Contains(t, metaValue, "auto_added")
}

func TestTransformDNSRecordState_ExistingMetaObject(t *testing.T) {
	// v4 state with existing meta object (from API response)
	// This should be converted to a JSON string for v5
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1",
		"ttl": 300,
		"meta": {"auto_added": true, "managed_by_apps": false}
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// meta should be a JSON string, not an object
	metaValue, ok := output["meta"].(string)
	assert.True(t, ok, "meta should be a string, got %T", output["meta"])
	assert.Contains(t, metaValue, "auto_added")
	assert.Contains(t, metaValue, "managed_by_apps")
}

func TestTransformDNSRecordState_CAARecord(t *testing.T) {
	// v4 CAA record with data array
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "CAA",
		"ttl": 300,
		"data": [{
			"flags": 0,
			"tag": "issue",
			"content": "letsencrypt.org"
		}]
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Data should be converted from array to object
	data, ok := output["data"].(map[string]interface{})
	require.True(t, ok, "data should be an object")

	// content should be renamed to value in data
	assert.Nil(t, data["content"])
	assert.Equal(t, "letsencrypt.org", data["value"])
	assert.Equal(t, "issue", data["tag"])

	// flags should be wrapped in NormalizedDynamicValue format
	flags, ok := data["flags"].(map[string]interface{})
	require.True(t, ok, "flags should be wrapped")
	assert.Equal(t, "number", flags["type"])
}

func TestTransformDNSRecordState_SRVRecord(t *testing.T) {
	// v4 SRV record with data array
	input := `{
		"zone_id": "abc123",
		"name": "_sip._tcp.example.com",
		"type": "SRV",
		"ttl": 300,
		"data": [{
			"priority": 10,
			"weight": 5,
			"port": 5060,
			"target": "sipserver.example.com"
		}]
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Data should be converted from array to object
	data, ok := output["data"].(map[string]interface{})
	require.True(t, ok, "data should be an object")

	assert.Equal(t, float64(10), data["priority"])
	assert.Equal(t, float64(5), data["weight"])
	assert.Equal(t, float64(5060), data["port"])
	assert.Equal(t, "sipserver.example.com", data["target"])

	// Priority should also be at root level for SRV records
	assert.Equal(t, float64(10), output["priority"])
}

func TestTransformDNSRecordState_EmptyDataArray(t *testing.T) {
	// v4 state with empty data array (should be removed)
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1",
		"ttl": 300,
		"data": []
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Empty data array should be removed for simple types
	assert.Nil(t, output["data"])
}

func TestTransformDNSRecordState_SettingsAllNull(t *testing.T) {
	// v4 state with settings where all values are null
	input := `{
		"zone_id": "abc123",
		"name": "example.com",
		"type": "A",
		"value": "192.168.1.1",
		"ttl": 300,
		"settings": {
			"flatten_cname": null,
			"ipv4_only": null,
			"ipv6_only": null
		}
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	// Settings with all null values should be removed
	assert.Nil(t, output["settings"])
}

func TestTransformDNSRecordState_WithAttributesWrapper(t *testing.T) {
	// State in terraform state file format with "attributes" wrapper
	input := `{
		"attributes": {
			"zone_id": "abc123",
			"name": "example.com",
			"type": "A",
			"value": "192.168.1.1",
			"ttl": 300,
			"hostname": "example.com"
		}
	}`

	result, err := state.TransformDNSRecordState(gjson.Parse(input))
	require.NoError(t, err)

	var output map[string]interface{}
	err = json.Unmarshal([]byte(result), &output)
	require.NoError(t, err)

	attrs := output["attributes"].(map[string]interface{})

	// Check transformations were applied
	assert.Equal(t, "192.168.1.1", attrs["content"])
	assert.Nil(t, attrs["value"])
	assert.Nil(t, attrs["hostname"])
}
