// This file is intentionally NOT code-generated.

package snippet_rules

import "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"

// unmarshalRulesResult reads the rules of a snippet rules response.
//
// The endpoint answers with the rules as the result itself, a list
// (`{"result":[{"snippet_name":...}]}`), not an object that holds them, so
// SnippetRulesResultEnvelope never fills Rules. With `no_refresh` on Rules, Read
// and import then kept whatever the state held, and a rule changed outside
// Terraform never showed in a plan.
func unmarshalRulesResult(raw []byte) (*[]*SnippetRulesRulesModel, error) {
	var env struct {
		Result *[]*SnippetRulesRulesModel `json:"result"`
	}
	if err := apijson.Unmarshal(raw, &env); err != nil {
		return nil, err
	}
	if env.Result == nil {
		empty := []*SnippetRulesRulesModel{}
		return &empty, nil
	}
	return env.Result, nil
}
