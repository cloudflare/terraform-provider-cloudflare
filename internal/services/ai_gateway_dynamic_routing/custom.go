package ai_gateway_dynamic_routing

import (
	"bytes"
	"encoding/json"
)

// normalizeDynamicRoutingResponse adapts the alternate API response shape to the
// provider model while preserving populated documented elements responses.
func normalizeDynamicRoutingResponse(raw []byte) ([]byte, error) {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return raw, nil
	}

	resultRaw, ok := envelope["result"]
	if !ok {
		return raw, nil
	}

	var result map[string]json.RawMessage
	if err := json.Unmarshal(resultRaw, &result); err != nil {
		return raw, nil
	}
	elementsRaw, hasElements := result["elements"]
	if hasElements {
		elementsRaw = bytes.TrimSpace(elementsRaw)
		if bytes.Equal(elementsRaw, []byte("null")) {
			hasElements = false
		} else {
			var elements []json.RawMessage
			if err := json.Unmarshal(elementsRaw, &elements); err != nil || elements == nil {
				return raw, nil
			}
			if len(elements) == 0 {
				hasElements = false
			} else if !isDynamicRoutingGraph(elements) {
				return raw, nil
			}
		}
	}

	versionRaw, ok := result["version"]
	if !ok {
		return raw, nil
	}

	var version map[string]json.RawMessage
	if err := json.Unmarshal(versionRaw, &version); err != nil {
		return raw, nil
	}

	dataRaw, ok := version["data"]
	if !ok {
		return raw, nil
	}

	var graph []json.RawMessage
	if err := json.Unmarshal(dataRaw, &graph); err != nil || graph == nil {
		return raw, nil
	}
	if !isDynamicRoutingGraph(graph) {
		return raw, nil
	}

	var compact bytes.Buffer
	if err := json.Compact(&compact, dataRaw); err != nil {
		return raw, nil
	}

	if !hasElements {
		result["elements"] = json.RawMessage(append([]byte(nil), compact.Bytes()...))
	}
	// The generated model exposes version.data as a string, while this API
	// variant returns the graph as an array.
	versionData, err := json.Marshal(compact.String())
	if err != nil {
		return nil, err
	}
	version["data"] = versionData

	versionRaw, err = json.Marshal(version)
	if err != nil {
		return nil, err
	}
	result["version"] = versionRaw
	resultRaw, err = json.Marshal(result)
	if err != nil {
		return nil, err
	}
	envelope["result"] = resultRaw

	return json.Marshal(envelope)
}

func isDynamicRoutingGraph(graph []json.RawMessage) bool {
	for _, elementRaw := range graph {
		var element map[string]json.RawMessage
		if err := json.Unmarshal(elementRaw, &element); err != nil || element == nil {
			return false
		}

		var id, elementType string
		if err := json.Unmarshal(element["id"], &id); err != nil || id == "" {
			return false
		}
		if err := json.Unmarshal(element["type"], &elementType); err != nil || elementType == "" {
			return false
		}
		var outputs map[string]json.RawMessage
		if err := json.Unmarshal(element["outputs"], &outputs); err != nil || outputs == nil {
			return false
		}
	}

	return true
}
