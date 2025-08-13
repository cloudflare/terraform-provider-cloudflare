package api_token

import (
	"encoding/json"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m APITokenModel) marshalCustom() (data []byte, err error) {
	if data, err = apijson.MarshalRoot(m); err != nil {
		return
	}
	return m.marshalPolicies(data)
}

func (m APITokenModel) marshalCustomForUpdate(state APITokenModel) (data []byte, err error) {
	if data, err = apijson.MarshalForUpdate(m, state); err != nil {
		return
	}
	return m.marshalPolicies(data)
}

func (m APITokenModel) marshalPolicies(b []byte) ([]byte, error) {
	var base map[string]json.RawMessage
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}
	if _, ok := base["policies"]; !ok {
		return b, nil
	}

	type onlyPolicies struct {
		Policies types.Dynamic `json:"policies,required"`
	}
	pb, err := apijson.MarshalRoot(onlyPolicies{Policies: m.Policies})
	if err != nil {
		return nil, err
	}
	var pm map[string]json.RawMessage
	if err := json.Unmarshal(pb, &pm); err != nil {
		return nil, err
	}
	if raw, ok := pm["policies"]; ok {
		base["policies"] = raw
	}
	return json.Marshal(base)
}
