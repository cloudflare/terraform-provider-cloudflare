package account_token

import (
	"encoding/json"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func MarshalCustom(m AccountTokenModel) (data []byte, err error) {
	if data, err = apijson.MarshalRoot(m); err != nil {
		return
	}
	var base map[string]json.RawMessage
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	// for each policy, marshal the resources string as raw json
	policyJsons := make([]json.RawMessage, len(*m.Policies))
	for i, policy := range *m.Policies {
		policyData, err := apijson.MarshalRoot(policy)
		if err != nil {
			return nil, err
		}
		var policyBase map[string]json.RawMessage
		if err := json.Unmarshal(policyData, &policyBase); err != nil {
			return nil, err
		}
		resources := json.RawMessage(policy.Resources.ValueString())
		policyBase["resources"] = resources
		policyJsons[i], err = json.Marshal(policyBase)
		if err != nil {
			return nil, err
		}
	}
	base["policies"], err = json.Marshal(policyJsons)
	if err != nil {
		return nil, err
	}

	return json.Marshal(base)
}

func UnmarshalCustom(data []byte, model *AccountTokenResultEnvelope) (err error) {
	if err = apijson.Unmarshal(data, model); err != nil {
		return
	}

	// pull out the raw JSON values for each policy resource and map to the model
	var base map[string]json.RawMessage
	if err := json.Unmarshal(data, &base); err != nil {
		return err
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(base["result"], &result); err != nil {
		return err
	}
	var policyJsons []json.RawMessage
	if err := json.Unmarshal(result["policies"], &policyJsons); err != nil {
		return err
	}
	for i, policyJson := range policyJsons {
		var policy map[string]json.RawMessage
		if err := json.Unmarshal(policyJson, &policy); err != nil {
			return err
		}
		(*model.Result.Policies)[i].Resources = types.StringValue(string(policy["resources"]))
	}
	return
}

func UnmarshalComputedCustom(data []byte, model *AccountTokenResultEnvelope) (err error) {
	if err = apijson.UnmarshalComputed(data, model); err != nil {
		return
	}
	forPoliciesOnly := AccountTokenResultEnvelope{}
	err = UnmarshalCustom(data, &forPoliciesOnly)
	if err != nil {
		return err
	}
	model.Result.Policies = forPoliciesOnly.Result.Policies
	return nil
}
