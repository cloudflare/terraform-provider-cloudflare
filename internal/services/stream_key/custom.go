package stream_key

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
)

// streamKeyListEnvelope decodes GET /accounts/{account_id}/stream/keys, which is
// a list endpoint: result is an array of every signing key on the account. The
// API has no per-key GET, so Read has to pick the element belonging to this
// resource itself.
//
// Result is a pointer so that an absent array is distinguishable from an empty
// one: the former means the response was not the list this code expects, the
// latter that the account really has no signing keys left.
type streamKeyListEnvelope struct {
	Result *[]StreamKeyModel `json:"result"`
}

// resolveStreamKeyID returns the id of the signing key this resource is bound
// to, or an empty string when it cannot be determined.
//
// Refreshes before this fix decoded the array response into a single object,
// which left id null in state. Such a state is only recoverable through the
// retained jwk, whose kid is the same value the list endpoint reports as id, so
// it is used as the fallback to heal states written by those versions.
func resolveStreamKeyID(data *StreamKeyModel) string {
	if id := data.ID.ValueString(); id != "" {
		return id
	}
	return keyIDFromJWK(data.Jwk.ValueString())
}

// keyIDFromJWK returns the kid of a signing key in JWK format. The API returns
// the jwk as base64 of a JSON object; the plain JSON form is accepted as well so
// that a change of encoding degrades into a no-op rather than a wrong id.
func keyIDFromJWK(jwk string) string {
	if jwk == "" {
		return ""
	}

	raw := []byte(jwk)
	if decoded, err := base64.StdEncoding.DecodeString(jwk); err == nil {
		raw = decoded
	} else if decoded, err := base64.RawStdEncoding.DecodeString(jwk); err == nil {
		raw = decoded
	}

	var key struct {
		Kid string `json:"kid"`
	}
	if err := json.Unmarshal(raw, &key); err != nil {
		return ""
	}
	return key.Kid
}

// refreshStreamKeyFromList applies the list element whose id is keyID onto data.
// It reports whether such an element was present; false means the key is gone
// from the account and the resource should be removed from state.
func refreshStreamKeyFromList(data *StreamKeyModel, keyID string, body []byte) (bool, error) {
	// A body that cannot be read as the expected list is reported as an error
	// rather than as "no such key": the caller removes the resource from state
	// when no element matches, and an unreadable response is no reason to do so.
	if !json.Valid(body) {
		return false, errors.New("response body is not valid JSON")
	}

	env := streamKeyListEnvelope{}
	if err := apijson.Unmarshal(body, &env); err != nil {
		return false, err
	}
	if env.Result == nil {
		return false, errors.New("response body has no result array")
	}

	for _, key := range *env.Result {
		if key.ID.ValueString() != keyID {
			continue
		}

		// The list endpoint reports id and created only. Writing a null over a
		// field the response does not carry is what wiped state in the first
		// place, so each field is taken from the response only when it is there.
		if !key.ID.IsNull() {
			data.ID = key.ID
		}
		if !key.KeyID.IsNull() {
			data.KeyID = key.KeyID
		}
		if !key.Created.IsNull() {
			data.Created = key.Created
		}
		return true, nil
	}

	return false, nil
}
