package page_rule

import (
	"encoding/json"
	"errors"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/pagerules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m PageRuleModel) marshalCustom() (data []byte, err error) {
	if data, err = apijson.MarshalRoot(m); err != nil {
		return
	}
	if data, err = m.marshalTargets(data); err != nil {
		return
	}
	if data, err = m.marshalActions(data); err != nil {
		return
	}
	return
}

func (m PageRuleModel) marshalTargets(b []byte) (data []byte, err error) {
	var T struct {
		ID         string `json:"id,omitempty"`
		ZoneID     string `json:"zone_id,omitempty"`
		Priority   int64  `json:"priority,omitempty"`
		Status     string `json:"status,omitempty"`
		CreatedOn  string `json:"created_on,omitempty"`
		ModifiedOn string `json:"modified_on,omitempty"`
		Target     string `json:"target,omitempty"`
		Targets    []any  `json:"targets,omitempty"`
	}
	if err = json.Unmarshal(b, &T); err != nil {
		return nil, err
	}
	T.Targets = []any{
		map[string]any{
			"target": "url",
			"constraint": map[string]any{
				"operator": "matches",
				"value":    T.Target,
			},
		},
	}
	T.Target = ""
	return json.Marshal(T)
}

func (m PageRuleModel) marshalActions(b []byte) (data []byte, err error) {
	var T struct {
		ID         string           `json:"id,omitempty"`
		ZoneID     string           `json:"zone_id,omitempty"`
		Priority   int64            `json:"priority,omitempty"`
		Status     string           `json:"status,omitempty"`
		CreatedOn  string           `json:"created_on,omitempty"`
		ModifiedOn string           `json:"modified_on,omitempty"`
		Targets    []any            `json:"targets,omitempty"`
		Actions    []map[string]any `json:"actions,omitempty"`
	}
	if err = json.Unmarshal(b, &T); err != nil {
		return nil, err
	}

	T.Actions, err = m.Actions.Encode()
    if err != nil {
        return nil, err
    }

	return json.Marshal(T)
}

func (m PageRuleModel) PageruleNewParams() pagerules.PageruleNewParams {
	return pagerules.PageruleNewParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
		Targets: cloudflare.F([]pagerules.TargetParam{
			{
				Constraint: cloudflare.F(pagerules.TargetConstraintParam{
					Operator: cloudflare.F(pagerules.TargetConstraintOperatorMatches),
					Value:    cloudflare.F(m.Target.String()),
				}),
			},
		}),
		Actions: cloudflare.F([]pagerules.PageruleNewParamsActionUnion{}),
	}
}

type PageRuleActionsModel struct {
	AutomaticHTTPSRewrites types.String `tfsdk:"automatic_https_rewrites" json:"automatic_https_rewrites,optional"`
	DisableApps            types.Bool   `tfsdk:"disable_apps" json:"disable_apps,optional"`
}

func (m *PageRuleActionsModel) Encode() (encoded []map[string]any, err error) {

	switch {
	case !m.AutomaticHTTPSRewrites.IsNull():
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDAutomaticHTTPSRewrites, "value": m.AutomaticHTTPSRewrites.String()})
	case !m.DisableApps.IsNull():
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDDisableApps, "value": m.DisableApps.ValueBool()})
	default:
    // TODO: Throw error for unknown page rule
        return nil, errors.New("missing or unknown page rule")
	}

	return
}
