package zero_trust_access_ai_controls_mcp_portal

import (
	"bytes"
	"encoding/json"
)

// injectServerIDs rewrites a raw portal API response so that each
// servers[].id is also present under the key servers[].server_id.
//
// The portal API returns each attached server's identity under the JSON key
// "id", but the resource schema and the Create/Update request body use
// "server_id" for the same value. apijson maps a struct field to a single JSON
// key for both read and write, so a plain unmarshal looking for "server_id"
// finds nothing in a response that uses "id", leaving server_id null in state.
//
// Populating server_id BEFORE the unmarshal is required because servers is
// modeled as a set: set elements are identified by their full object value, so
// elements whose server_id is null would collapse together and could not be
// matched against config by identity. Injecting the key lets apijson populate
// server_id natively during (Computed)Unmarshal for Read, Import, Create, and
// Update alike.
//
// Best effort: returns the input unchanged on any parse error, and never
// overwrites a server_id that is already present.
func injectServerIDs(body []byte) []byte {
	var env map[string]json.RawMessage
	if json.Unmarshal(body, &env) != nil {
		return body
	}
	resRaw, ok := env["result"]
	if !ok {
		return body
	}
	var portal map[string]json.RawMessage
	if json.Unmarshal(resRaw, &portal) != nil {
		return body
	}
	serversRaw, ok := portal["servers"]
	if !ok {
		return body
	}
	var servers []map[string]json.RawMessage
	if json.Unmarshal(serversRaw, &servers) != nil {
		return body
	}

	changed := false
	for _, s := range servers {
		if _, has := s["server_id"]; has {
			continue
		}
		id, ok := s["id"]
		if !ok {
			continue
		}
		// Skip an explicit JSON null id. server_id is the required element
		// identity; copying null would only yield a null server_id (drift/churn),
		// so treat a null id the same as an absent one.
		if bytes.Equal(bytes.TrimSpace(id), []byte("null")) {
			continue
		}
		s["server_id"] = id
		changed = true
	}
	if !changed {
		return body
	}

	ns, err := json.Marshal(servers)
	if err != nil {
		return body
	}
	portal["servers"] = ns
	np, err := json.Marshal(portal)
	if err != nil {
		return body
	}
	env["result"] = np
	out, err := json.Marshal(env)
	if err != nil {
		return body
	}
	return out
}
