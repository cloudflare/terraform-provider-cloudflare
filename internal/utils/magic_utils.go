package utils

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func UnmarshalMagicModel(bytes []byte, env any, wrapperField string, unmarshalComputedOnly bool) (err error) {
	if wrapperField != "" {
		// remove extra wrapper field
		bytes, err = stripWrapper(bytes, wrapperField)
		if err != nil {
			return err
		}
	}
	if unmarshalComputedOnly {
		return apijson.UnmarshalComputed(bytes, env)
	}
	return apijson.Unmarshal(bytes, env)
}

// stripWrapper removes a wrapper key from the json response
// and moves the value object to root level
// example
//   - input: {"result": {"ipsec_tunnel": {"name":"tunnel", ...}}}
//   - output: {"result": {"name":"tunnel", ...}}
func stripWrapper(input []byte, key string) ([]byte, error) {
	jsonStr := string(input)
	path := fmt.Sprintf("result.%s", key)

	obj := gjson.Get(jsonStr, path)
	if !obj.Exists() {
		return nil, fmt.Errorf("key %s not found in response", key)
	}

	var err error
	for k, v := range obj.Map() {
		jsonStr, err = sjson.Set(jsonStr, fmt.Sprintf("result.%s", k), v.Value())
		if err != nil {
			return nil, fmt.Errorf("failed to set json value %s: %s", k, v.String())
		}
	}

	jsonStr, err = sjson.Delete(jsonStr, path)
	if err != nil {
		return nil, fmt.Errorf("failed to remove key %s in json", path)
	}

	return []byte(jsonStr), nil
}

// LookupMagicWanCfIP returns a usable Anycast IP for a specific account
// used for test only
func LookupMagicWanCfIP(t *testing.T, accountID string) string {
	var cfIP string
	// CF anycast IPs are allocated per account
	switch accountID {
	case "5e51c34a8fffd8b96136e87be3af7110":
		cfIP = "162.159.68.213"
	case "f037e56e89293a057740de681ac9abbe":
		cfIP = "162.159.73.109"
	default:
		t.Fatalf("need to specify an anycast IP for account %s", accountID)
	}
	return cfIP
}
