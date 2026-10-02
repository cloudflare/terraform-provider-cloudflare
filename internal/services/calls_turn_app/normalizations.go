package calls_turn_app

// normalizeCallsTURNKeyID copies uid → key_id when key_id is not set.
//
// The Calls TURN API returns the resource identifier as uid in the JSON
// response body (e.g. POST /calls/turn_keys response). The provider uses
// key_id as the URL path parameter for Get/Update/Delete:
//
//	GET  /accounts/{account_id}/calls/turn_keys/{key_id}
//	PUT  /accounts/{account_id}/calls/turn_keys/{key_id}
//	DELETE /accounts/{account_id}/calls/turn_keys/{key_id}
//
// Because key_id is a path parameter (not a JSON field), the apijson decoder
// never populates it from the API response. Without this normalization, every
// operation after Create fails with "missing required key_id parameter".
//
// The uid and key_id values are identical — they both represent the same
// Cloudflare-generated identifier. This function bridges the naming gap until
// the upstream OpenAPI spec (cffs.git) exposes key_id as a readable top-level
// JSON field in the POST/GET/PUT response body.
func normalizeCallsTURNKeyID(data *CallsTURNAppModel) {
	if data == nil {
		return
	}
	if !data.KeyID.IsNull() && !data.KeyID.IsUnknown() && data.KeyID.ValueString() != "" {
		return
	}
	if !data.UID.IsNull() && !data.UID.IsUnknown() {
		data.KeyID = data.UID
	}
}
